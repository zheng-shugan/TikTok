package logic

import (
	"context"
	"github.com/sunflower10086/TikTok/interaction/internal/dao"
	"github.com/sunflower10086/TikTok/interaction/internal/svc"
	___interaction "github.com/sunflower10086/TikTok/interaction/pb"
	"log"

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
	// 查询点赞列表
	videos, err := dao.GetFavoriteList(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	// 判断每个视频作者是否被关注
	for _, v := range videos {
		v.IsFavorite = true

		// 默认自己不能关注自己
		if v.Author.ID == in.UserId {
			v.Author.IsFollow = false
			continue
		}

		check, err := dao.CheckIsFollow(l.ctx, v.Author.ID, in.UserId)
		if err != nil {
			log.Println("判断用户是否关注视频作者失败:", err)
			return nil, err
		}

		v.Author.IsFollow = check
	}

	return &___interaction.FavoriteListResp{VideoList: videos}, nil
}
