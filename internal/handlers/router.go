package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lat1992/tiny-btc/internal/services"
)

func GetRouter(cs services.IChainService, ts services.ITransactionService) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.POST("/sendTransaction", SendTransaction(ts))
	router.GET("/transaction/:hash", GetTransaction(ts))

	router.GET("/blocks", GetBlocks(cs))
	router.GET("/block/:number", GetBlock(cs))
	router.GET("/blockHeight", GetBlockHeight(cs))

	router.POST("/setDifficulty", SetDifficulty(cs))
	router.POST("/mine/start", StartMine(cs))
	router.POST("/mine/stop", StopMine(cs))

	return router
}
