package contract

import (
	"mime/multipart"
)

type ImageUploader interface {
	UploadImage(token string, file multipart.File, filename string) (int64, error)
}
