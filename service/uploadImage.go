package service

import (
	"context"
	"mime/multipart"
)

type (
	taskOfUploadImage struct {
		req *ReqOfUploadImage
	}

	ReqOfUploadImage struct {
		Image *multipart.File
	}

	storageOfUpload struct {
		ImgFingerprint string
	}
)

func UploadImage(ctx context.Context, in *ReqOfUploadImage) error {
	return nil
}

func newTaskOfUploadImage(in *ReqOfUploadImage) *taskOfUploadImage {
	return &taskOfUploadImage{
		req: in,
	}
}

func (t *taskOfUploadImage) Exec(ctx context.Context) error {
	return nil
}

func (t *taskOfUploadImage) getImageFingerprint(ctx context.Context) error {
	return nil
}

func (t *taskOfUploadImage) uploadToS3(ctx context.Context) error {
	return nil
}
