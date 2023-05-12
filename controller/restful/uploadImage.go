package restful

import (
	"context"
	"github.com/gin-gonic/gin"
	"upload-image-s3/env"
	"upload-image-s3/helper"
	"upload-image-s3/helper/config"
	"upload-image-s3/service"
)

func UploadImg(gc *gin.Context) {
	req := new(service.ReqOfUploadImage)
	if err := gc.Bind(req); err != nil {
		helper.ResJSON(gc, nil, err)
		return
	}

	c, cancel := context.WithTimeout(gc, config.ConForge().GetDuration(env.TimeOfAPI))
	defer cancel()

	if err := service.UploadImage(c, req); err != nil {
		helper.ResJSON(gc, nil, err)
	}

	helper.ResJSON(gc, nil, nil)
}
