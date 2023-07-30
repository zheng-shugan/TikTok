package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	config "github.com/sunflower10086/TikTok/http/config"
	"github.com/sunflower10086/TikTok/http/internal/dao/db"
	"github.com/sunflower10086/TikTok/http/internal/pkg"

	_ "github.com/sunflower10086/TikTok/http/internal/all"
)

var (
	configFile string
)

var StartCmd = &cobra.Command{
	Use:     "start",
	Long:    "tiktok API后端",
	Short:   "tiktok API后端",
	Example: "tiktok API后端 commands",
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

		return run()
	},
}

func run() error {
	g := gin.Default()

	if err := pkg.InitGinHandler(g); err != nil {
		return err
	}

	handlers := pkg.LoadedGinApp()
	fmt.Println("loaded handler has", handlers)

	srv := &http.Server{
		Addr:    config.C().Apps.HTTP.GetAddr(),
		Handler: g,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
		return err
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 2 seconds.")
	}
	log.Println("Server exiting")

	return nil
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&configFile, "config", "f", "./etc/demo.toml", "demo config file")
	RootCmd.AddCommand(StartCmd)
}
