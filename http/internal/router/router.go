package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/middleware"
	user "github.com/sunflower10086/TikTok/http/internal/user/http"
	video "github.com/sunflower10086/TikTok/http/internal/video/http"
)

func Init(r *gin.Engine) {
	// noAuth
	noAuthRouter := r.Group("/douyin")
	{
		noAuthRouter.POST("/user/register/", user.Register)
		noAuthRouter.POST("/user/login/", user.Login)
		noAuthRouter.GET("/feed/", video.GetFeedVideo)
	}

	// 加的middleware.JWTAuthMiddlewareQuery()是一个登录验证的中间件
	withAuthRouter := r.Group("/douyin", middleware.JWTAuthMiddlewareQuery())
	{
		withAuthRouter.GET("/test/", user.Test) // 获取用户信息
		withAuthRouter.POST("/publish/action/", video.PublishAction)
		withAuthRouter.GET("/publish/list/", video.GetPublishList)
		withAuthRouter.GET("/user/", user.GetUserInfo)
	}
}
