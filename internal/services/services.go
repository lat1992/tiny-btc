package services

import (
	"github.com/lat1992/tiny-btc/internal"
)

type ITransactionService interface {
	SendTransaction(hash, tx string)
	GetTransaction(hash string) *internal.Transaction
}

type IChainService interface {
	StopMine()
	StartMine()
	SetDifficulty(v uint)
	GetBlockHeight() uint
	GetBlocks() []internal.Block
	GetBlock(blockNumber uint) internal.Block
}
