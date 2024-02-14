package chain

import (
	"log/slog"
	"time"

	"github.com/lat1992/tiny-btc/internal"
	"github.com/lat1992/tiny-btc/internal/block"
	"github.com/lat1992/tiny-btc/internal/pow"
	"github.com/lat1992/tiny-btc/internal/txPool"
)

type Chain struct {
	blocks            []*block.Block
	blockHeight       uint
	currentDifficulty uint
	txPool            *txPool.TxPool
	isStop            bool
}

func NewChain(tp *txPool.TxPool) *Chain {
	hash := pow.Proof(0, "genesis", "")
	b := block.NewBlock(hash, []*internal.Transaction{})
	return &Chain{
		blocks:            []*block.Block{b},
		currentDifficulty: 1,
		txPool:            tp,
		isStop:            false,
	}
}

func (c *Chain) Mine(pendingTxs []*internal.Transaction) {
	c.blockHeight = uint(len(c.blocks))
	lastBlock := c.blocks[c.blockHeight-1]
	var hashString string
	for _, pendingTx := range pendingTxs {
		hashString += pendingTx.Hash
	}
	hash := pow.Proof(c.currentDifficulty, lastBlock.Hash(), hashString)
	b := block.NewBlock(hash, pendingTxs)
	c.blocks = append(c.blocks, b)
	c.txPool.Validate(pendingTxs, c.blockHeight)
}

func (c *Chain) Start() {
	for {
		if !c.isStop {
			pendingTxs := c.txPool.GetPendingTransactions()
			if len(pendingTxs) > 3 {
				c.Mine(pendingTxs)
				slog.Info("block mined")
			}
			time.Sleep(time.Second)
		}
	}
}

func (c *Chain) SetIsStop(v bool) {
	c.isStop = v
}

func (c *Chain) SetDifficulty(v uint) {
	c.currentDifficulty = v
}

func (c *Chain) GetBlockHeight() uint {
	return c.blockHeight
}

func (c *Chain) GetBlocks() []*block.Block {
	return c.blocks
}

func (c *Chain) GetBlock(blockNumber uint) *block.Block {
	return c.blocks[blockNumber]
}
