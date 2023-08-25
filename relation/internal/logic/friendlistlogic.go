package logic

import (
	"context"

	"github.com/sunflower10086/TikTok/relation/internal/svc"
	___relation "github.com/sunflower10086/TikTok/relation/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *___relation.FriendListRequest) (*___relation.FriendListResponse, error) {
	// todo: add your logic here and delete this line

	return &___relation.FriendListResponse{}, nil
}
