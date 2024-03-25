package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) HandleEndpoint1(c *gin.Context) {
	// TODO: Implementar logica para el endpoint 1

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint 1"})
}
