package aliyun

import (
	"fmt"
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sunflower10086/TikTok/http/config"
	myOss "github.com/sunflower10086/TikTok/http/internal/pkg/oss"
)

var (
	_ myOss.Uploader = &AliOssStore{}
)

type AliOssStore struct {
	client      *oss.Client
	aliOssStore *config.Oss
	listener    oss.ProgressListener
}

func NewAliOssStore(conf *config.Oss) (*AliOssStore, error) {
	client, err := oss.New(conf.OssEndpoint, conf.AccessKeyId, conf.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client:      client,
		aliOssStore: conf,
		listener:    &ProgressListener{},
	}, nil
}

func (a *AliOssStore) Upload(bucketName, objectKey string, file *multipart.FileHeader) error {
	// 2.获得我们的bucket对象
	bucket, err := a.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	// 3.上传文件
	if err := bucket.PutObject(objectKey, f, oss.Progress(a.listener)); err != nil {
		return err
	}

	// 4. 打印下载链接
	downloadUrl, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Println("上传云商: 阿里云[oss-cn-beijing.aliyuncs.com]")
	fmt.Printf("上传用户: [%s]\n", a.aliOssStore.OssEndpoint)
	fmt.Printf("文件下载url: [%s]\n\n", downloadUrl)
	fmt.Println("注意:下载链接有效期一天，请在一天内下载")
	return nil
}
