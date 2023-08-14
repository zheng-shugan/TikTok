package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/videostream"
	"github.com/sunflower10086/TikTok/http/internal/videostream/impl"
	"net/http"
	"strconv"
)

func GetFeedVideo(ctx *gin.Context) {
	var latest_time int64
	var token string

	// 参数绑定
	latestTimeParam := ctx.Query("latest_time")
	if len(latestTimeParam) != 0 {
		if latestTime, err := strconv.ParseInt(latestTimeParam, 10, 64); err != nil {
			ctx.JSON(http.StatusOK, videostream.GetFeedVideoResp{
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

	var req = videostream.GetFeedVideoReq{
		LatestTime: latest_time,
		Token:      token,
	}

	// 调用服务
	resp, err := impl.GetFeedVideo(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, videostream.GetFeedVideoResp{
			StatusCode: result.ServerErrCode,
			StatusMsg:  err.Error(),
			VideoList:  nil,
			NextTime:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
