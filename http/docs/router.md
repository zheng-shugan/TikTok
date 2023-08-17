# 路由说明

[router.go](..%2Finternal%2Frouter%2Frouter.go)


路由的末尾必须加入斜杠，否则"抖声"会报错307

[gin相关issue](https://github.com/gin-gonic/gin/issues/1004)

`
[GIN-debug] redirecting request 307: /douyin/user/login --> /douyin/user/login?username=testuser%40qq.com&password=123456
`

```go
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
}
```