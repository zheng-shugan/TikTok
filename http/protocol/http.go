package protocol

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/restful-api-demo/apps"
	"github.com/sunflower10086/restful-api-demo/conf"
)

func NewHTTPService() *HTTPService {
	g := gin.Default()
	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HTTPAddr(),
		Handler:           g,
	}

	return &HTTPService{
		Service: server,
		L:       log.New(os.Stderr, "[Host] ", log.Ldate|log.Ltime|log.Lshortfile),
		Conf:    conf.C(),
		engin:   g,
	}
}

type HTTPService struct {
	Service *http.Server
	L       *log.Logger
	Conf    *conf.Config
	engin   *gin.Engine
}

func (h *HTTPService) Start() error {
	if err := apps.InitGinHandler(h.engin); err != nil {
		fmt.Println(err)
		return err
	}

	// 已加载的app的日志信息打印
	handlers := apps.LoadedGinHandler()
	h.L.Println("loaded handler has", handlers)

	// 监听端口
	h.L.Printf("%s running in %s \n", h.Conf.App.Name, h.Conf.App.HTTPAddr())
	if err := h.Service.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			h.L.Fatalf("listen: %s\n", err)
			return nil
		} else {
			return fmt.Errorf("start service err: %s", err.Error())
		}
	}
	return nil
}

func (h *HTTPService) Stop() {
	h.L.Println("start graceful shutdown")
	// 开始关闭
	timeOut := 2
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := h.Service.Shutdown(ctx); err != nil {
		h.L.Fatalf("%s Shutdown err: %v\n", h.Conf.App.HTTPAddr(), err)
	}
	// 关闭超时
	select {
	case <-ctx.Done():
		h.L.Printf("timeout of %d seconds.", timeOut)
	}

	// 正常关闭
	h.L.Printf("%s exiting", h.Conf.App.Name)
}
