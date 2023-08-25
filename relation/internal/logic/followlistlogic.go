package logic

import (
	"context"

	"github.com/sunflower10086/TikTok/relation/internal/svc"
	___relation "github.com/sunflower10086/TikTok/relation/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowListLogic) FollowList(in *___relation.FollowListRequest) (*___relation.FollowListResponse, error) {
	// todo: add your logic here and delete this line

	return &___relation.FollowListResponse{}, nil
}
