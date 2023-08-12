package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/feed"
	"github.com/sunflower10086/TikTok/http/internal/feed/impl"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
)

func Publish(ctx *gin.Context) {

	//SaveUploadedFile上传表单文件到指定的路径
	// ctx.SaveUploadedFile(f, "./"+f.Filename)
	var publishParam feed.PublishRequest

	// 绑定参数
	if err := ctx.ShouldBind(&publishParam); err != nil {
		// msg := result.ParamErrMsg
		msg := err.Error()
		ctx.JSON(http.StatusOK, feed.PublishResponse{
			Response: result.Response{
				StatusCode: result.ParamErrCode,
				StatusMsg:  &msg,
			},
		})
	}

	_, err := impl.Publish(ctx, &publishParam)
	if err != nil {
		msg := result.ParamErrMsg
		ctx.JSON(http.StatusOK, feed.PublishResponse{
			Response: result.Response{
				StatusCode: result.ServerErrCode,
				StatusMsg:  &msg,
			},
		})
	}

	ctx.JSON(http.StatusOK, feed.PublishResponse{
		Response: result.Response{StatusCode: result.SuccessCode},
	})
}
