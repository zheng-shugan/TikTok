package impl

import (
	"context"
	"fmt"
	"github.com/sunflower10086/TikTok/http/internal/dao"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sunflower10086/TikTok/http/config"
	"github.com/sunflower10086/TikTok/http/internal/pkg/oss"
	"github.com/sunflower10086/TikTok/http/internal/pkg/oss/aliyun"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/pkg/token"
	"github.com/sunflower10086/TikTok/http/internal/video"
	"github.com/sunflower10086/TikTok/http/pkg/jwt"
)

const (
	LIMIT = 30 //返回的视频数
)

func GetFeedVideo(ctx context.Context, req *video.GetFeedVideoReq) (*video.GetFeedVideoResp, error) {
	// latest_time默认为当前时间，若请求参数不为空则更新
	latestTime := time.Now().Unix()
	if req.LatestTime != 0 && req.LatestTime < 253402300799 { // 253402300799对应10000-01-01 07:59:59 +0800 CST
		latestTime = req.LatestTime
	}
	log.Println(time.Unix(latestTime, 0))

	// 获取视频流
	videos, err := dao.QueryFeedVideo(ctx, LIMIT, latestTime)
	if err != nil {
		log.Println("视频流获取失败:", err)
		return nil, err
	}

	// 确定登录用户的视频点赞和关注信息
	token := req.Token
	if token != "" {
		userID, _, err := jwt.GetUserIDAndUsername(token)
		if err != nil {
			log.Println("token验证失败:", err)
			return nil, err
		}

		for _, v := range videos {
			check, err := dao.CheckIsFavorite(ctx, v.ID, userID)
			if err != nil {
				log.Println("判断用户是否给视频点赞失败:", err)
				return nil, err
			}

			v.IsFavorite = check

			check, err = dao.CheckIsFollow(ctx, v.Author.ID, userID)
			if err != nil {
				log.Println("判断用户是否关注视频作者失败:", err)
				return nil, err
			}

			v.Author.IsFollow = check
		}
	}

	var nextTime *int64 = nil
	if len(videos) > 0 {
		nextTime = &(videos[len(videos)-1]).PublishTime
	}

	return &video.GetFeedVideoResp{
		StatusCode: result.SuccessCode,
		StatusMsg:  result.SuccessMsg,
		VideoList:  videos,
		NextTime:   nextTime,
	}, nil
}

func PublishAction(ctx *gin.Context, req *video.PublishRequest) (*video.PublishResponse, error) {
	// 保证唯一的 videoName
	videoName := uuid.New().String()

	ossConf := config.C().Oss

	var uploader oss.Uploader

	uploader, err := aliyun.NewAliOssStore(ossConf)
	if err != nil {
		return nil, err
	}

	downURL, err := uploader.Upload(ossConf.BucketName, ossConf.PlayUrlPrefix+videoName+".mp4", req.Data)
	if err != nil {
		return nil, err
	}

	//userId, _ := token.GetUserIDAndUsernameFromCtx(ctx)

	//nowUserId, _ := token.GetUserIDAndUsernameFromCtx(ctx)

	userId, username := token.GetUserIDAndUsernameFromCtx(ctx)
	fmt.Println(userId, username)

	err = dao.SaveVideo(ctx, downURL, req.Title, userId)
	if err != nil {
		log.Println("视频存入数据库失败:", err)
		return nil, err
	}

	return &video.PublishResponse{
		Response: result.Response{StatusCode: result.SuccessCode},
	}, nil
}

func GetPublishList(ctx context.Context, req *video.GetPublishListReq) (*video.GetPublishListResp, error) {
	userID := req.UserID
	token := req.Token

	// token不为空则鉴权
	if token != "" {
		_, err := jwt.ParseToken(token)
		if err != nil {
			return &video.GetPublishListResp{
				StatusCode: result.ParamErrCode,
				StatusMsg:  result.ParamErrMsg,
				VideoList:  nil,
			}, err
		}
	}

	// 查询发布列表
	videos, err := dao.QueryPublishList(ctx, userID)
	if err != nil {
		log.Println("查询用户发布列表失败:", err)
		return nil, err
	}

	// 默认自己不能关注自己
	// 特判自己是否给自己的视频点赞
	for _, v := range videos {
		check, err := dao.CheckIsFavorite(ctx, v.ID, userID)
		if err != nil {
			log.Println("判断用户是否给视频点赞失败:", err)
			return nil, err
		}

		v.IsFavorite = check
	}

	return &video.GetPublishListResp{
		StatusCode: result.SuccessCode,
		StatusMsg:  result.SuccessMsg,
		VideoList:  videos,
	}, nil
}
