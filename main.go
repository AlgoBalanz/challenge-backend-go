package main

import (
	"challenge/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/endpoint1", handler.HandleEndpoint1)
	r.POST("/endpoint2", handler.HandleEndpoint2)

	r.Run()
}
