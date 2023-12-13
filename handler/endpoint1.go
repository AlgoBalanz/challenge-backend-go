package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Endpoint1 struct {
	Ticker     string    `json:"ticker"`
	Code       int32     `json:"code" binding:"gte=0"`
	Total      float64   `json:"total" binding:"gte=0"`
	Exceeded   bool      `json:"exceeded"`
	LastChange time.Time `json:"last_change"`
}

func HandleEndpoint1(c *gin.Context) {
	// TODO: Implementar logica para el endpoint 1

	c.JSON(http.StatusOK, gin.H{"message": "Endpoint 1"})
}
