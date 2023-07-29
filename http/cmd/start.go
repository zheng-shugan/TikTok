package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	config "github.com/sunflower10086/TikTok/http/config"
	"github.com/sunflower10086/TikTok/http/internal/dao/db"
	"github.com/sunflower10086/TikTok/http/internal/pkg"
	"github.com/sunflower10086/TikTok/http/protocol"
)

var (
	configFile string
)

var StartCmd = &cobra.Command{
	Use:     "start",
	Long:    "demo API后端",
	Short:   "demo API后端",
	Example: "demo API后端 commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载配置文件
		if err := config.LoadConfigFromYaml(configFile); err != nil {
			return err
		}

		// 启动与mysql的连接
		if err := db.Init(); err != nil {
			return err
		}

		// 注册服务到IOC
		pkg.InitImpl()

		master := newMaster()

		// 优雅启停
		// 相当于监听一下 kill -2 和 kill -9
		quit := make(chan os.Signal)
		// kill (no param) default send syscanll.SIGTERM
		// kill -2 is syscall.SIGINT (Ctrl + C)
		// kill -9 is syscall.SIGKILL
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		go master.WaitStop(quit)

		return master.Start()
	},
}

func newMaster() *master {
	return &master{
		http: protocol.NewHTTPService(),
	}
}

type master struct {
	http *protocol.HTTPService
}

func (m *master) Start() error {
	if err := m.http.Start(); err != nil {
		return err
	}
	return nil
}

func (m *master) Stop() {
	log.Printf("Shutdown %s ...\n", m.http.Conf.App.Name)
}

func (m *master) WaitStop(quit <-chan os.Signal) {
	for v := range quit {
		switch v {
		default:
			m.http.L.Printf("received signal: %s", v)
			m.http.Stop()
		}
	}
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&configFile, "config", "f", "./etc/demo.toml", "demo config file")
	RootCmd.AddCommand(StartCmd)
}
