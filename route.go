package main

import "github.com/gin-gonic/gin"

func Routing(router *gin.Engine) {
	router.POST("upload/image")
}
