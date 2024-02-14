package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lat1992/tiny-btc/internal/services"
)

func StartMine(s services.IChainService) func(c *gin.Context) {
	return func(c *gin.Context) {
		s.StartMine()
		c.JSON(http.StatusOK, gin.H{
			"message": "start mine",
		})
	}
}

func StopMine(s services.IChainService) func(c *gin.Context) {
	return func(c *gin.Context) {
		s.StopMine()
		c.JSON(http.StatusOK, gin.H{
			"message": "stop mine",
		})
	}
}

type SetDifficultyRequest struct {
	Value uint `json:"difficulty"`
}

func SetDifficulty(s services.IChainService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request SetDifficultyRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "json format error",
			})
			return
		}
		s.SetDifficulty(request.Value)
		c.JSON(http.StatusOK, gin.H{
			"message": "difficulty set",
		})
	}
}
