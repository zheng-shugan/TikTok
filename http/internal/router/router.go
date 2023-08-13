package router

import (
	"github.com/gin-gonic/gin"
	feed "github.com/sunflower10086/TikTok/http/internal/feed/http"
	user "github.com/sunflower10086/TikTok/http/internal/user/http"
	videostream "github.com/sunflower10086/TikTok/http/internal/videostream/http"
)

func Init(r *gin.Engine) {
	// noAuth
	noAuthRouter := r.Group("/douyin")
	{
		noAuthRouter.POST("/user/register", user.Register)
		noAuthRouter.POST("/user/login", user.Login)
		noAuthRouter.POST("/publish/action/", feed.Publish)
		noAuthRouter.GET("/videostream", videostream.GetFeedVideo)
	}

	// 加的auth.JwtAuthMiddleware()是一个登录验证的中间件
	// withAuthRouter := g.Group("/douyin", auth.JwtAuthMiddleware())
	// {
	// 	withAuthRouter.GET("/user/", u.GetInfoHandler) // 获取用户信息
	//  withAuthRouter.GET("/publishlist", publishlist.GetPublishList())
	// }
}
