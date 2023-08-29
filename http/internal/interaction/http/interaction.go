package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/interaction"
	"github.com/sunflower10086/TikTok/http/internal/interaction/rpc"
	"github.com/sunflower10086/TikTok/http/internal/interaction/rpcTohttp"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/pkg/token"
	"github.com/sunflower10086/TikTok/http/pkg/jwt"
	___interaction "github.com/sunflower10086/TikTok/interaction/pb"
	"net/http"
	"strconv"
)

func FavoriteAction(ctx *gin.Context) {
	var userID, videoID, actionType int64
	var msg string

	videoIDParam := ctx.Query("video_id")
	actionTypeParam := ctx.Query("action_type")

	// 参数校验以及绑定
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

	if Type != 1 && Type != 2 {
		msg = "action_type参数错误: 请输入1（点赞）或者2（取消点赞）"
		ctx.JSON(http.StatusOK, result.Response{
			StatusCode: result.ParamErrCode,
			StatusMsg:  &msg,
		})
		return
	}

	userID, _ = token.GetUserIDAndUsernameFromCtx(ctx)
	videoID = ID
	actionType = Type

	// 调用服务
	err = rpc.FavoriteAction(ctx, &___interaction.FavoriteActionReq{
		UserId:     userID,
		VideoId:    videoID,
		ActionType: actionType,
	})

	if err != nil {
		msg = err.Error()
		ctx.JSON(http.StatusOK, result.Response{
			StatusCode: result.ServerErrCode,
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
		ctx.JSON(http.StatusOK, interaction.FavoriteListResp{
			StatusCode: result.ParamErrCode,
			StatusMsg:  &msg,
			VideoList:  nil,
		})
		return
	}

	ID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		msg = err.Error()
		ctx.JSON(http.StatusOK, interaction.FavoriteListResp{
			StatusCode: result.ParamErrCode,
			StatusMsg:  &msg,
			VideoList:  nil,
		})
		return
	}

	userID = ID

	// 若token不为空则鉴权
	if token != "" {
		_, err = jwt.ParseToken(token)
		if err != nil {
			msg = err.Error()
			ctx.JSON(http.StatusOK, interaction.FavoriteListResp{
				StatusCode: result.AuthErrCode,
				StatusMsg:  &msg,
				VideoList:  nil,
			})
			return
		}
	}

	// 调用服务
	resp, err := rpc.FavoriteList(ctx, &___interaction.FavoriteListReq{
		UserId: userID,
	})

	if err != nil {
		msg = err.Error()
		ctx.JSON(http.StatusOK, interaction.FavoriteListResp{
			StatusCode: result.ServerErrCode,
			StatusMsg:  &msg,
			VideoList:  nil,
		})
		return
	}

	videos := make([]*interaction.Video, len(resp.VideoList))
	for i, v := range resp.VideoList {
		videos[i] = rpcTohttp.MapHttp(v)
	}

	msg = "success"
	ctx.JSON(http.StatusOK, interaction.FavoriteListResp{
		StatusCode: result.SuccessCode,
		StatusMsg:  &msg,
		VideoList:  videos,
	})
}
