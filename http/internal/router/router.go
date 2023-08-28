package router

import (
	"github.com/gin-gonic/gin"
	interaction "github.com/sunflower10086/TikTok/http/internal/interaction/http"
	"github.com/sunflower10086/TikTok/http/internal/middleware"
	user "github.com/sunflower10086/TikTok/http/internal/user/http"
	video "github.com/sunflower10086/TikTok/http/internal/video/http"
)

func Init(r *gin.Engine) {
	// noAuth
	noAuthRouter := r.Group("/douyin")
	{
		noAuthRouter.POST("/user/register/", user.Register)              // 注册
		noAuthRouter.POST("/user/login/", user.Login)                    // 登录
		noAuthRouter.GET("/feed/", video.GetFeedVideo)                   // 获取视频流
		noAuthRouter.GET("/publish/list/", video.GetPublishList)         // 获取发布列表
		noAuthRouter.GET("/favorite/list/", interaction.GetFavoriteList) // 获取喜欢列表
	}

	// 加的middleware.JWTAuthMiddlewareQuery()是一个登录验证的中间件
	withAuthRouter := r.Group("/douyin", middleware.JWTAuthMiddlewareQuery())
	{
		//withAuthRouter.GET("/test/", user.Test)                      // Test: 获取用户信息
		withAuthRouter.POST("/publish/action/", video.PublishAction)         // 投稿
		withAuthRouter.GET("/user/", user.GetUserInfo)                       // 获取用户信息
		withAuthRouter.POST("/favorite/action/", interaction.FavoriteAction) // 点赞
	}
}
