package impl

import (
	"context"
	"github.com/sunflower10086/TikTok/http/internal/dao"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/publishlist"
	"log"
)

func GetPublishList(ctx context.Context, req *publishlist.GetPublishListReq) (*publishlist.GetPublishListResp, error) {
	userID := req.UserID

	videos, err := dao.QueryPublishList(ctx, userID)
	if err != nil {
		log.Println("查询用户发布列表错误!")
		return nil, err
	}

	return &publishlist.GetPublishListResp{
		StatusCode: result.SuccessCode,
		StatusMsg:  result.SuccessMsg,
		VideoList:  videos,
	}, nil
}
