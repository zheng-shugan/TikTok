package impl

import (
	"context"
	"github.com/sunflower10086/TikTok/http/internal/dao"
	"github.com/sunflower10086/TikTok/http/internal/pkg/jwt"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/videostream"
	"log"
	"time"
)

const LIMIT = 30 //返回的视频数

func GetFeedVideo(ctx context.Context, req *videostream.GetFeedVideoReq) (*videostream.GetFeedVideoResp, error) {
	// latest_time默认为当前时间，若请求参数不为空则更新
	latestTime := time.Now().Unix()
	if req.LatestTime != 0 {
		latestTime = req.LatestTime
	}

	// 获取视频流
	videos, err := dao.QueryFeedVideo(ctx, LIMIT, latestTime)
	if err != nil {
		log.Println("视频流获取失败!")
		return nil, err
	}

	// 确定登录用户的视频点赞和关注信息
	token := req.Token
	if token != "" {
		userID, err := jwt.VerifyJWT(token) // TODO: jwt.VerifyJWT接口待实现
		if err != nil {
			log.Println("token验证失败!")
			return nil, err
		}

		err = dao.CheckIsFavorite(ctx, videos, userID)
		if err != nil {
			log.Println("判断视频点赞失败!")
			return nil, err
		}

		err = dao.CheckIsFollow(ctx, videos, userID)
		if err != nil {
			log.Println("判断用户是否关注视频作者失败!")
			return nil, err
		}
	}

	var nextTime *int64 = nil
	if len(videos) > 0 {
		nextTime = &(*videos[len(videos)-1]).PublishTime
	}

	return &videostream.GetFeedVideoResp{
		StatusCode: result.SuccessCode,
		StatusMsg:  result.SuccessMsg,
		VideoList:  videos,
		NextTime:   nextTime,
	}, nil
}
