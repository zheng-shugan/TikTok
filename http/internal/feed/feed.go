package feed

import (
	"context"
	"mime/multipart"

	"github.com/sunflower10086/TikTok/http/internal/pkg/result"
)

type Service interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
}

type PublishRequest struct {
	Data  *multipart.FileHeader `json:"data" binding:"required" form:"data"`
	Title string                `json:"title" binding:"required" form:"title"`
}

type PublishResponse struct {
	result.Response
}
