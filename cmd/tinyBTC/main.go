package main

import (
	"log/slog"

	"github.com/lat1992/tiny-btc/internal/chain"
	"github.com/lat1992/tiny-btc/internal/handlers"
	"github.com/lat1992/tiny-btc/internal/services"
	"github.com/lat1992/tiny-btc/internal/txPool"
)

func main() {
	// init
	tp := txPool.NewTxPool()
	ts := services.NewTransactionService(tp)
	c := chain.NewChain(tp)
	cs := services.NewChainService(c)
	router := handlers.GetRouter(cs, ts)

	// start
	go c.Start()
	if err := router.Run(":8080"); err != nil {
		slog.Error("error when router start", err)
	}
}
