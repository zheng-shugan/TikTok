package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/interaction/rpc"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	___interaction "github.com/sunflower10086/TikTok/interaction/pb"
	"net/http"
	"strconv"
)

func FavoriteAction(ctx *gin.Context) {
	var videoID, actionType int64
	var msg string

	videoIDParam := ctx.Query("video_id")
	actionTypeParam := ctx.Query("action_type")

	// 绑定参数
	if videoIDParam == "" || actionTypeParam == "" {
		msg = "参数为空!"
		ctx.JSON(http.StatusOK, result.Response{
			StatusCode: result.ParamErrCode,
			StatusMsg:  &msg,
		})
		return
	}

	ID, err := strconv.ParseInt(videoIDParam, 10, 64)
	Type, err2 := strconv.ParseInt(actionTypeParam, 10, 64)

	if err != nil || err2 != nil {
		msg = err.Error()
		ctx.JSON(http.StatusOK, result.Response{
			StatusCode: result.ParamErrCode,
			StatusMsg:  &msg,
		})
		return
	}

	videoID = ID
	actionType = Type

	// 调用服务
	err = rpc.FavoriteAction(ctx, &___interaction.FavoriteActionReq{
		VideoId:    videoID,
		ActionType: actionType,
	})

	if err != nil {
		msg = err.Error()
		ctx.JSON(http.StatusOK, result.Response{
			StatusCode: result.ParamErrCode,
			StatusMsg:  &msg,
		})
	} else {
		msg = "success"
		ctx.JSON(http.StatusOK, result.Response{
			StatusCode: result.SuccessCode,
			StatusMsg:  &msg,
		})
	}
}

func FavoriteList(ctx *gin.Context) {

}
