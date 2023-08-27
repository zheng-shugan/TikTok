package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/video"
	"github.com/sunflower10086/TikTok/http/internal/video/impl"
)

func GetFeedVideo(ctx *gin.Context) {
	var latest_time int64
	var token string

	// 参数校验并绑定
	latestTimeParam := ctx.Query("latest_time")
	if len(latestTimeParam) != 0 {
		if latestTime, err := strconv.ParseInt(latestTimeParam, 10, 64); err != nil {
			ctx.JSON(http.StatusOK, video.GetFeedVideoResp{
				StatusCode: result.ParamErrCode,
				StatusMsg:  err.Error(),
				VideoList:  nil,
				NextTime:   nil,
			})
			return
		} else {
			latest_time = latestTime
		}
	} else {
		latest_time = 0
	}

	token = ctx.Query("token")

	var req = video.GetFeedVideoReq{
		LatestTime: latest_time,
		Token:      token,
	}

	// 调用服务
	resp, err := impl.GetFeedVideo(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, video.GetFeedVideoResp{
			StatusCode: result.ServerErrCode,
			StatusMsg:  err.Error(),
			VideoList:  nil,
			NextTime:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func PublishAction(ctx *gin.Context) {
	//	SaveUploadedFile上传表单文件到指定的路径
	// ctx.SaveUploadedFile(f, "./"+f.Filename)
	var publishParam video.PublishRequest
	fmt.Println(ctx.Query("token"))
	// 绑定参数
	if err := ctx.ShouldBind(&publishParam); err != nil {
		// msg := result.ParamErrMsg
		msg := err.Error()
		ctx.JSON(http.StatusOK, video.PublishResponse{
			Response: result.Response{
				StatusCode: result.ParamErrCode,
				StatusMsg:  &msg,
			},
		})
		return
	}

	_, err := impl.PublishAction(ctx, &publishParam)
	if err != nil {
		// msg := result.ParamErrMsg
		msg := err.Error()
		ctx.JSON(http.StatusOK, video.PublishResponse{
			Response: result.Response{
				StatusCode: result.ServerErrCode,
				StatusMsg:  &msg,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, video.PublishResponse{
		Response: result.Response{StatusCode: result.SuccessCode},
	})
}

func GetPublishList(ctx *gin.Context) {
	var userID int64
	var token string

	tokenParam := ctx.Query("token")
	userIDParam := ctx.Query("user_id")

	// 参数校验并绑定
	if userIDParam == "" {
		ctx.JSON(http.StatusOK, video.GetPublishListResp{
			StatusCode: result.ParamErrCode,
			StatusMsg:  result.ParamErrMsg,
			VideoList:  nil,
		})
		return
	}

	if ID, err := strconv.ParseInt(userIDParam, 10, 64); err != nil {
		ctx.JSON(http.StatusOK, video.GetPublishListResp{
			StatusCode: result.ParamErrCode,
			StatusMsg:  err.Error(),
			VideoList:  nil,
		})
		return
	} else {
		userID = ID
	}
	token = tokenParam

	var req = video.GetPublishListReq{
		UserID: userID,
		Token:  token,
	}

	// 调用服务
	resp, err := impl.GetPublishList(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusOK, video.GetPublishListResp{
			StatusCode: result.ParamErrCode,
			StatusMsg:  err.Error(),
			VideoList:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)

}
