package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/endpoint1", handleEndpoint1)
	r.POST("/endpoint2", handleEndpoint2)

	r.Run()
}

func handleEndpoint1(c *gin.Context) {
	// TODO: Implementar logica para el endpoint 1

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint 1"})
}

func handleEndpoint2(c *gin.Context) {
	// TODO: Implementar logica para el endpoint 2

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint 2"})
}
