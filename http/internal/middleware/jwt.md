## Gin中间件 - JWT

### 1. 配置保护路由

[router.go](..%2Frouter%2Frouter.go)

所有在withAuthRouter中的路由都会被保护，需要登录才能访问

```go
// 加的middleware.JwtAuthMiddleware()是一个登录验证的中间件
withAuthRouter := r.Group("/douyin", middleware.JWTAuthMiddleware())
{
    withAuthRouter.GET("/test/", user.Test) // 获取用户信息
}
```

如果用户没有登录，访问这些路由，会返回401错误

登录状态需要在请求头中携带token，token的key是Authorization，value是`Bearer + 空格 + token`

![img.png](..%2F..%2F..%2Fimages%2Fimg.png)

### 2. 通过token获取用户信息

```go
func Test(ctx *gin.Context) {
	// 通过token获取用户信息
	userId, username := token.GetUserIDAndUsernameFromCtx(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"userId":   userId,
		"username": username,
	})
}
```