package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/user"
	"github.com/sunflower10086/TikTok/http/internal/user/impl"
)

// Controller

// Login 用户登录
func Login(ctx *gin.Context) {
	var loginParam user.LoginRequest

	// 参数校验
	if err := ctx.ShouldBind(&loginParam); err != nil {
		msg := result.ParamErrMsg
		ctx.JSON(http.StatusOK, user.LoginResponse{
			Response: result.Response{
				StatusCode: result.ParamErrCode,
				StatusMsg:  &msg,
			},
		})
	}

	// 调用服务的接口
	resp, err := impl.Login(ctx, &loginParam)
	if err != nil {
		errMsg := err.Error()
		ctx.JSON(http.StatusOK, user.LoginResponse{
			Response: result.Response{
				StatusCode: result.ServerErrCode,
				StatusMsg:  &errMsg,
			},
		})
	}

	ctx.JSON(http.StatusOK, user.LoginResponse{
		Response: result.Response{StatusCode: result.SuccessCode},
		Token:    resp.Token,
		UserID:   resp.UserID,
	})
}

func Register(ctx *gin.Context) {
	var registerParam user.RegisterRequest

	// 参数校验
	if err := ctx.ShouldBind(&registerParam); err != nil {
		msg := result.ParamErrMsg
		ctx.JSON(http.StatusOK, user.RegisterResponse{
			Response: result.Response{
				StatusCode: result.ParamErrCode,
				StatusMsg:  &msg,
			},
		})
	}

	// 调用服务的接口
	resp, err := impl.Register(ctx, &registerParam)
	if err != nil {
		errMsg := err.Error()
		ctx.JSON(http.StatusOK, user.RegisterResponse{
			Response: result.Response{
				StatusCode: result.ServerErrCode,
				StatusMsg:  &errMsg,
			},
		})
	}

	ctx.JSON(http.StatusOK, user.RegisterResponse{
		Response: result.Response{StatusCode: result.SuccessCode},
		Token:    resp.Token,
		UserID:   resp.UserID,
	})
}

func GetUserInfo(ctx *gin.Context) {
	var getUserInfo user.GetInfoRequest

	// 参数校验
	if err := ctx.ShouldBind(&getUserInfo); err != nil {
		msg := result.ParamErrMsg
		ctx.JSON(http.StatusOK, user.LoginResponse{
			Response: result.Response{
				StatusCode: result.ParamErrCode,
				StatusMsg:  &msg,
			},
		})
	}

	// 调用服务的接口
	resp, err := impl.GetInfo(ctx, &getUserInfo)
	if err != nil {
		errMsg := err.Error()
		ctx.JSON(http.StatusOK, user.GetInfoResponse{
			Response: result.Response{
				StatusCode: result.ServerErrCode,
				StatusMsg:  &errMsg,
			},
		})
	}

	ctx.JSON(http.StatusOK, user.GetInfoResponse{
		Response: result.Response{StatusCode: result.SuccessCode},
		User:     resp.User,
	})
}
