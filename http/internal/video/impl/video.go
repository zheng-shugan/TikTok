package impl

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sunflower10086/TikTok/http/internal/dao"
	"github.com/sunflower10086/TikTok/http/internal/pkg/jwt"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
	"github.com/sunflower10086/TikTok/http/internal/video"
)

const (
	B  = 1 << 3
	KB = B << 10
	MB = KB << 10
)

const (
	LIMIT     = 30 //返回的视频数
	MAX_VIDEO = 10 * MB
)

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

func PublishAction(ctx *gin.Context, req *video.PublishRequest) (*video.PublishResponse, error) {
	// 保证唯一的 videoName
	videoName := uuid.New().String()

	// ossConf := config.C().Oss

	// var uploader oss.Uploader

	// uploader, err := aliyun.NewAliOssStore(ossConf)
	// if err != nil {
	// 	return nil, err
	// }
	// 阿里云oss存储，域名没有备案，换个方式
	// err = uploader.Upload(ossConf.BucketName, ossConf.OssVideoDir+videoName+".mp4", req.Data)
	// if err != nil {
	// 	return nil, err
	// }

	// 存储视频在本地
	dir, _ := os.Getwd()
	_ = os.Mkdir("video", os.ModeDir)
	f, _ := os.Create(dir + "\\video\\" + videoName + ".mp4")
	defer f.Close()
	f2, _ := req.Data.Open()
	data := make([]byte, MAX_VIDEO)
	n, _ := f2.Read(data)
	f.Write(data[:n])

	// 从cookie中取user_id，(生成token的时候会把user_id存入cookie)
	CookieGetId, ok := ctx.Get("user_id")
	if ok == false {
		return nil, errors.New("从cookie中获取userId失败")
	}
	userIdStr := fmt.Sprintf("%v", CookieGetId)
	userId, err := strconv.Atoi(userIdStr)

	err = dao.SaveVideo(ctx, ossConf.OssVideoDir+videoName+".mp4", int64(userId), req.Title)
	if err != nil {
		log.Println("视频存入数据库失败！")
		return nil, err
	}

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
