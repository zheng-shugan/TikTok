package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/interaction/rpc"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/pkg/jwt"
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

func GetFavoriteList(ctx *gin.Context) {
	var userID int64
	var msg, token string

	userIDParam := ctx.Query("user_id")
	token = ctx.Query("token")

	// 参数绑定
	if userIDParam == "" {
		msg = "参数为空"
		ctx.JSON(http.StatusOK, map[string]any{
			"status_code": result.ParamErrCode,
			"status_msg":  &msg,
			"video_list":  nil,
		})
		return
	}

	ID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		msg = err.Error()
		ctx.JSON(http.StatusOK, map[string]any{
			"status_code": result.ParamErrCode,
			"status_msg":  &msg,
			"video_list":  nil,
		})
		return
	}

	userID = ID

	// 若token不为空则鉴权
	if token != "" {
		_, err = jwt.ParseToken(token)
		if err != nil {
			msg = err.Error()
			ctx.JSON(http.StatusOK, map[string]any{
				"status_code": result.ParamErrCode,
				"status_msg":  &msg,
				"video_list":  nil,
			})
			return
		}
	}

	// 调用服务
	list, err := rpc.FavoriteList(ctx, &___interaction.FavoriteListReq{
		UserId: userID,
	})

	if err != nil {
		msg = err.Error()
		ctx.JSON(http.StatusOK, map[string]any{
			"status_code": result.ParamErrCode,
			"status_msg":  &msg,
			"video_list":  nil,
		})
		return
	}

	msg = "success"
	ctx.JSON(http.StatusOK, map[string]any{
		"status_code": result.ParamErrCode,
		"status_msg":  &msg,
		"video_list":  list,
	})
}
