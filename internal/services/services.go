package services

import (
	"github.com/lat1992/tiny-btc/internal"
	"github.com/lat1992/tiny-btc/internal/block"
)

type ITransactionService interface {
	SendTransaction(hash, tx string)
	GetTransaction(hash string) *internal.Transaction
}

type IBlockService interface {
}

type IChainService interface {
	StopMine()
	StartMine()
	SetDifficulty(v uint)
	GetBlockHeight() uint
	GetBlocks() []*block.Block
	GetBlock(blockNumber uint) *block.Block
}
