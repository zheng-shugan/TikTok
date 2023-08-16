package router

import (
	"github.com/gin-gonic/gin"
	feed "github.com/sunflower10086/TikTok/http/internal/feed/http"
	"github.com/sunflower10086/TikTok/http/internal/middleware"
	user "github.com/sunflower10086/TikTok/http/internal/user/http"
)

func Init(r *gin.Engine) {
	// noAuth
	noAuthRouter := r.Group("/douyin")
	{
		noAuthRouter.POST("/user/register", user.Register)
		noAuthRouter.POST("/user/login", user.Login)
		noAuthRouter.POST("/publish/action/", feed.Publish)
	}

	// 加的middleware.JwtAuthMiddleware()是一个登录验证的中间件
	withAuthRouter := r.Group("/douyin", middleware.JWTAuthMiddleware())
	{
		withAuthRouter.GET("/test/", user.Test) // 获取用户信息
	}
}
