package logic

import (
	"context"

	"github.com/sunflower10086/TikTok/relation/internal/svc"
	___relation "github.com/sunflower10086/TikTok/relation/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowerListLogic) FollowerList(in *___relation.FollowerListRequest) (*___relation.FollowerListResponse, error) {
	// todo: add your logic here and delete this line

	return &___relation.FollowerListResponse{}, nil
}
