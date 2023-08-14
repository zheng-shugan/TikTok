package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/publishlist"
	"github.com/sunflower10086/TikTok/http/internal/publishlist/impl"
	"net/http"
	"strconv"
)

func GetPublishList(ctx *gin.Context) {
	var userID int64
	var token string

	tokenParam := ctx.Query("token")
	userIDParam := ctx.Query("user_id")

	// 参数校验
	if tokenParam == "" || userIDParam == "" {
		ctx.JSON(http.StatusOK, publishlist.GetPublishListResp{
			StatusCode: result.ParamErrCode,
			StatusMsg:  result.ParamErrMsg,
			VideoList:  nil,
		})
		return
	}

	if ID, err := strconv.ParseInt(userIDParam, 10, 64); err != nil {
		ctx.JSON(http.StatusOK, publishlist.GetPublishListResp{
			StatusCode: result.ParamErrCode,
			StatusMsg:  err.Error(),
			VideoList:  nil,
		})
		return
	} else {
		userID = ID
	}
	token = tokenParam

	var req = publishlist.GetPublishListReq{
		UserID: userID,
		Token:  token,
	}

	// 调用服务
	resp, err := impl.GetPublishList(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusOK, publishlist.GetPublishListResp{
			StatusCode: result.ParamErrCode,
			StatusMsg:  err.Error(),
			VideoList:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)

}
