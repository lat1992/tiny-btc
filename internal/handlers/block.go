package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lat1992/tiny-btc/internal/services"
)

func GetBlockHeight(s services.IChainService) func(c *gin.Context) {
	return func(c *gin.Context) {
		bh := s.GetBlockHeight()
		c.JSON(http.StatusOK, gin.H{
			"blockHeight": bh,
		})
	}
}

func GetBlock(s services.IChainService) func(c *gin.Context) {
	return func(c *gin.Context) {
		bnString := c.Param("number")
		bn, err := strconv.Atoi(bnString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "input no valid"})
			return
		}
		b := s.GetBlock(uint(bn))
		c.JSON(http.StatusOK, b)
	}
}

func GetBlocks(s services.IChainService) func(c *gin.Context) {
	return func(c *gin.Context) {
		b := s.GetBlocks()
		c.JSON(http.StatusOK, b)
	}
}
