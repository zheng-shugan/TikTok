package oss

import "mime/multipart"

type Uploader interface {
	Upload(bucketName, objectKey string, file *multipart.FileHeader) error
}
