package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/pkg/jwt"
	"net/http"
	"strings"
)

type jwtResponse struct {
	result.Response
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			msg := "请求头中auth为空"
			ctx.JSON(http.StatusOK, jwtResponse{
				Response: result.Response{
					StatusCode: result.AuthErrCode,
					StatusMsg:  &msg,
				},
			})
			ctx.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			msg := "请求头中auth格式有误，必须是Bearer {token}"
			ctx.JSON(http.StatusOK, jwtResponse{
				Response: result.Response{
					StatusCode: result.AuthErrCode,
					StatusMsg:  &msg,
				},
			})
			ctx.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			msg := "无效的Token"
			ctx.JSON(http.StatusOK, jwtResponse{
				Response: result.Response{
					StatusCode: result.AuthErrCode,
					StatusMsg:  &msg,
				},
			})

			ctx.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文ctx上
		ctx.Set("userID", mc.UserID)
		ctx.Set("username", mc.Username)

		ctx.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
