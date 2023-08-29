package logic

import (
	"context"
	"github.com/sunflower10086/TikTok/interaction/internal/dao"
	"github.com/sunflower10086/TikTok/interaction/internal/svc"
	___interaction "github.com/sunflower10086/TikTok/interaction/pb"
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
	//点赞 or 取消点赞
	var err error
	if in.ActionType == 1 {
		err = dao.AddFavorite(l.ctx, in.UserId, in.VideoId)
	} else {
		err = dao.DelFavorite(l.ctx, in.UserId, in.VideoId)
	}

	return &___interaction.Empty{}, err
}
