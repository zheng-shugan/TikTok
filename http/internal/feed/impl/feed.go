package impl

import (
	"context"

	"github.com/google/uuid"
	"github.com/sunflower10086/TikTok/http/config"
	"github.com/sunflower10086/TikTok/http/internal/feed"
	"github.com/sunflower10086/TikTok/http/internal/pkg/oss"
	"github.com/sunflower10086/TikTok/http/internal/pkg/oss/aliyun"
	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
)

func Publish(ctx context.Context, req *feed.PublishRequest) (*feed.PublishResponse, error) {
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

	return &feed.PublishResponse{
		Response: result.Response{StatusCode: result.SuccessCode},
	}, nil
}
