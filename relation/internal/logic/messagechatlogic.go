package logic

import (
	"context"

	"github.com/sunflower10086/TikTok/relation/internal/svc"
	___relation "github.com/sunflower10086/TikTok/relation/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageChatLogic) MessageChat(in *___relation.MessageChatRequest) (*___relation.MessageChatResponse, error) {
	// todo: add your logic here and delete this line

	return &___relation.MessageChatResponse{}, nil
}
