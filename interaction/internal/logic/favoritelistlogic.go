package logic

import (
	"context"

	"github.com/sunflower10086/TikTok/interaction/internal/svc"
	___interaction "github.com/sunflower10086/TikTok/interaction/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *___interaction.FavoriteListReq) (*___interaction.FavoriteListResp, error) {
	// todo: add your logic here and delete this line

	return &___interaction.FavoriteListResp{}, nil
}
