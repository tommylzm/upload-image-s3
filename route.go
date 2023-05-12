package main

import (
	"github.com/gin-gonic/gin"
	"upload-image-s3/controller/restful"
)

func Routing(router *gin.Engine) {
	router.POST("upload/image", restful.UploadImg)
}
