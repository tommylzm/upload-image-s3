package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"upload-image-s3/helper/log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	g.Use(gin.Recovery())

	Routing(g)

	fmt.Println("server is online")

	server := &http.Server{
		Addr:    ":3000",
		Handler: g,
	}

	log.Error(fmt.Sprintf("Server in fatal. %v", server.ListenAndServe()))
}
