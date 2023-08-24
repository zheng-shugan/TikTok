package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/TikTok/http/internal/relation"
)

// Controller

func Action(ctx *gin.Context, req *relation.ActionRequest) (*relation.ActionResponse, error) {
	//TODO 关注操作
	return nil, nil
}

func FollowList(ctx *gin.Context, req *relation.FollowListRequest) (*relation.FollowListResponse, error) {
	//TODO 关注列表
	return nil, nil
}

func FollowerList(ctx *gin.Context, req *relation.FollowerListRequest) (*relation.FollowerListResponse, error) {
	//TODO 粉丝列表
	return nil, nil
}

func FriendList(ctx *gin.Context, req relation.FriendListRequest) (*relation.FriendListResponse, error) {
	//TODO 好友列表
	return nil, nil
}
