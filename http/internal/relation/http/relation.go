package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/relation"
	"github.com/sunflower10086/TikTok/http/internal/relation/impl"
	"net/http"
)

func RelationAction(ctx *gin.Context) {
	var actionParam relation.ActionRequest

	if err := ctx.ShouldBind(&actionParam); err != nil {
		msg := err.Error()
		ctx.JSON(http.StatusOK, relation.ActionResponse{
			Response: result.Response{
				StatusCode: result.ParamErrCode,
				StatusMsg:  &msg,
			},
		})
		return
	}

	_, err := impl.Action(ctx, &actionParam)
	if err != nil {
		msg := err.Error()
		ctx.JSON(http.StatusOK, relation.ActionResponse{
			Response: result.Response{
				StatusCode: result.ServerErrCode,
				StatusMsg:  &msg,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, relation.ActionResponse{
		Response: result.Response{
			StatusCode: result.SuccessCode,
		},
	})
}

func FollowList(ctx *gin.Context) {
	var followListParam relation.FollowListRequest

	if err := ctx.ShouldBind(&followListParam); err != nil {
		msg := err.Error()
		ctx.JSON(http.StatusOK, relation.FollowListResponse{
			Response: result.Response{
				StatusCode: result.ParamErrCode,
				StatusMsg:  &msg,
			},
		})
		return
	}

	rep, err := impl.FollowerList(ctx, &followListParam)
	if err != nil {
		msg := err.Error()
		ctx.JSON(http.StatusOK, relation.FollowListResponse{
			Response: result.Response{
				StatusCode: result.ServerErrCode,
				StatusMsg:  &msg,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, relation.FollowListResponse{
		Response: result.Response{
			StatusCode: result.SuccessCode,
		},
		UserList: rep.UserList,
	})
}

func FollowerList(ctx *gin.Context) {
	var followerListParam relation.FollowerListRequest

	if err := ctx.ShouldBind(&followerListParam); err != nil {
		msg := err.Error()
		ctx.JSON(http.StatusOK, relation.FollowerListResponse{
			Response: result.Response{
				StatusCode: result.ParamErrCode,
				StatusMsg:  &msg,
			},
		})
		return
	}

	rep, err := impl.FollowerList(ctx, &followerListParam)
	if err != nil {
		msg := err.Error()
		ctx.JSON(http.StatusOK, relation.FollowerListResponse{
			Response: result.Response{
				StatusCode: result.ServerErrCode,
				StatusMsg:  &msg,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, relation.FollowerListResponse{
		Response: result.Response{
			StatusCode: result.SuccessCode,
		},
		UserList: rep.UserList,
	})
}

func FriendList(ctx *gin.Context) {
	var friendListParam relation.FriendListRequest

	if err := ctx.ShouldBind(&friendListParam); err != nil {
		msg := err.Error()
		ctx.JSON(http.StatusOK, relation.FriendListResponse{
			Response: result.Response{
				StatusCode: result.ParamErrCode,
				StatusMsg:  &msg,
			},
		})
		return
	}

	rep, err := impl.FollowerList(ctx, &friendListParam)
	if err != nil {
		msg := err.Error()
		ctx.JSON(http.StatusOK, relation.FriendListResponse{
			Response: result.Response{
				StatusCode: result.ServerErrCode,
				StatusMsg:  &msg,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, relation.FriendListResponse{
		Response: result.Response{
			StatusCode: result.SuccessCode,
		},
		UserList: rep.UserList,
	})
}
