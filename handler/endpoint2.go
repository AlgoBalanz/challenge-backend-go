package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleEndpoint2(c *gin.Context) {
	// TODO: Implementar logica para el endpoint 2

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint 2"})
}
