package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/middleware"
	relation "github.com/sunflower10086/TikTok/http/internal/relation/http"
	user "github.com/sunflower10086/TikTok/http/internal/user/http"
	video "github.com/sunflower10086/TikTok/http/internal/video/http"
)

func Init(r *gin.Engine) {
	// noAuth
	noAuthRouter := r.Group("/douyin")
	{
		noAuthRouter.POST("/user/register/", user.Register)
		noAuthRouter.POST("/user/login/", user.Login)
		noAuthRouter.POST("/publish/action/", video.PublishAction)
		noAuthRouter.GET("/feed/", video.GetFeedVideo)
	}

	// 加的middleware.JWTAuthMiddlewareQuery()是一个登录验证的中间件
	withAuthRouter := r.Group("/douyin", middleware.JWTAuthMiddlewareQuery())
	{
		withAuthRouter.GET("/test/", user.Test) // 获取用户信息
		withAuthRouter.GET("/publish/list/", video.GetPublishList)
		withAuthRouter.POST("/relation/action", relation.RelationAction)
		withAuthRouter.POST("/relation/follow/list/", relation.FollowList)
		withAuthRouter.POST("/relation/follower/list/", relation.FollowerList)
		withAuthRouter.POST("/relation/friend/list/", relation.FriendList)
	}
}
