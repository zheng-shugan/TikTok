package impl

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sunflower10086/TikTok/http/config"
	"github.com/sunflower10086/TikTok/http/internal/dao"
	"github.com/sunflower10086/TikTok/http/internal/pkg/jwt"
	"github.com/sunflower10086/TikTok/http/internal/pkg/oss"
	"github.com/sunflower10086/TikTok/http/internal/pkg/oss/aliyun"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/video"
)

const LIMIT = 30 //返回的视频数

func GetFeedVideo(ctx context.Context, req *video.GetFeedVideoReq) (*video.GetFeedVideoResp, error) {
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

	return &video.GetFeedVideoResp{
		StatusCode: result.SuccessCode,
		StatusMsg:  result.SuccessMsg,
		VideoList:  videos,
		NextTime:   nextTime,
	}, nil
}

func PublishAction(ctx context.Context, req *video.PublishRequest) (*video.PublishResponse, error) {
	// 保证唯一的 videoName
	videoName := uuid.New().String()

	ossConf := config.C().Oss

	var uploader oss.Uploader

	uploader, err := aliyun.NewAliOssStore(ossConf)
	if err != nil {
		return nil, err
	}

	err = uploader.Upload(ossConf.BucketName, ossConf.PlayUrlPrefix+videoName+".mp4", req.Data)
	if err != nil {
		return nil, err
	}

	// err = dao.UploadVideo(videoName, userId, title)
	// if err != nil {
	// 	log.Println("视频存入数据库失败！")
	// 	return nil, err
	// }

	return &video.PublishResponse{
		Response: result.Response{StatusCode: result.SuccessCode},
	}, nil
}

func GetPublishList(ctx context.Context, req *video.GetPublishListReq) (*video.GetPublishListResp, error) {
	userID := req.UserID

	videos, err := dao.QueryPublishList(ctx, userID)
	if err != nil {
		log.Println("查询用户发布列表错误!")
		return nil, err
	}

	return &video.GetPublishListResp{
		StatusCode: result.SuccessCode,
		StatusMsg:  result.SuccessMsg,
		VideoList:  videos,
	}, nil
}
