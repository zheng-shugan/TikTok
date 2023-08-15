package logic

import (
	"context"

	"github.com/sunflower10086/TikTok/interaction/internal/svc"
	___interaction "github.com/sunflower10086/TikTok/interaction/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *___interaction.CommentListReq) (*___interaction.CommentListResp, error) {
	// todo: add your logic here and delete this line

	return &___interaction.CommentListResp{}, nil
}
