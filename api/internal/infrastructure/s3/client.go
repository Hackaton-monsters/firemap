package s3

import (
	"context"
	"io"

	cnfgs "firemap/internal/infrastructure/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryClient interface {
	UploadImage(ctx context.Context, file io.Reader, publicID string) (string, error)
	DeleteImage(ctx context.Context, publicID string) error
}

type CloudinaryClientImpl struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryClient(cnfgs *cnfgs.Config) (*CloudinaryClientImpl, error) {
	cld, err := cloudinary.NewFromURL(cnfgs.S3.URL)
	if err != nil {
		return nil, err
	}

	return &CloudinaryClientImpl{
		cld: cld,
	}, nil
}

func (c *CloudinaryClientImpl) UploadImage(ctx context.Context, file io.Reader, publicID string) (string, error) {
	uploadResult, err := c.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: publicID,
		Folder:   "firemap",
	})
	if err != nil {
		return "", err
	}
	return uploadResult.SecureURL, nil
}

func (c *CloudinaryClientImpl) DeleteImage(ctx context.Context, publicID string) error {
	_, err := c.cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	return err
}

func ProvideS3Client(cnfgs *cnfgs.Config) CloudinaryClient {
	client, err := NewCloudinaryClient(cnfgs)
	if err != nil {
		return nil
	}
	return client
}
