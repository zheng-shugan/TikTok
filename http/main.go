package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/config"
	"github.com/sunflower10086/TikTok/http/internal/dao/db"
	"github.com/sunflower10086/TikTok/http/internal/router"
)

var configFile = flag.String("f", "etc/", "the config file")

func init() {
	config.LoadConfigFromYaml(*configFile)

	db.Init()
}

func main() {
	r := gin.Default()

	router.Init(r)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello")
	})

	httpConf := config.C().Apps.HTTP
	addr := fmt.Sprintf("%s:%s", httpConf.Host, httpConf.Port)
	r.Run(addr)
}
