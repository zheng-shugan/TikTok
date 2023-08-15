package logic

import (
	"context"

	"github.com/sunflower10086/TikTok/interaction/internal/svc"
	___interaction"github.com/sunflower10086/TikTok/interaction/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *___interaction.FavoriteActionReq) (*___interaction.Empty, error) {
	// todo: add your logic here and delete this line

	return &___interaction.Empty{}, nil
}
