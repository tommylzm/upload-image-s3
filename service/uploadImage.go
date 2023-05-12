package service

import (
	"context"
	"github.com/corona10/goimagehash"
	"github.com/disintegration/imaging"
	"image"
	"mime/multipart"
)

type (
	taskOfUploadImage struct {
		req     *ReqOfUploadImage
		storage *storageOfUpload
	}

	ReqOfUploadImage struct {
		Image *multipart.FileHeader
	}

	storageOfUpload struct {
		ImgFingerprint uint64
	}
)

func UploadImage(ctx context.Context, in *ReqOfUploadImage) error {
	task := newTaskOfUploadImage(in)
	return task.Exec(ctx)
}

func newTaskOfUploadImage(in *ReqOfUploadImage) *taskOfUploadImage {
	return &taskOfUploadImage{
		req:     in,
		storage: &storageOfUpload{},
	}
}

func (t *taskOfUploadImage) Exec(ctx context.Context) error {
	if err := t.getImageFingerprint(ctx); err != nil {
		return err
	}
	return nil
}

func (t *taskOfUploadImage) getImageFingerprint(ctx context.Context) error {
	f, _ := t.req.Image.Open()
	defer f.Close()

	im, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	grayImg := imaging.Grayscale(im)
	blurrredImg := imaging.Blur(grayImg, 10)

	fingerprint, err := goimagehash.PerceptionHash(blurrredImg)
	if err != nil {
		return err
	}

	t.storage.ImgFingerprint = fingerprint.GetHash()

	return nil
}

func (t *taskOfUploadImage) uploadToS3(ctx context.Context) error {
	return nil
}
