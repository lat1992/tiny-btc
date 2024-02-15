package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lat1992/tiny-btc/internal/services"
)

type SendTransactionRequest struct {
	Hash  string `json:"hash"`
	RawTx string `json:"rawTx"`
}

func SendTransaction(s services.ITransactionService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request SendTransactionRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "json format error",
			})
			return
		}
		if err := s.SendTransaction(request.Hash, request.RawTx); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "transaction sent",
		})
	}
}

func GetTransaction(s services.ITransactionService) func(c *gin.Context) {
	return func(c *gin.Context) {
		txHash := c.Param("hash")
		tx := s.GetTransaction(txHash)
		c.JSON(http.StatusOK, tx)
	}
}
