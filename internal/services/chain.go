package services

import (
	"github.com/lat1992/tiny-btc/internal/block"
	"github.com/lat1992/tiny-btc/internal/chain"
)

type ChainService struct {
	chain *chain.Chain
}

func NewChainService(c *chain.Chain) *ChainService {
	return &ChainService{
		chain: c,
	}
}

func (s *ChainService) StopMine() {
	s.chain.SetIsStop(true)
}

func (s *ChainService) StartMine() {
	s.chain.SetIsStop(false)
}

func (s *ChainService) SetDifficulty(v uint) {
	s.chain.SetDifficulty(v)
}

func (s *ChainService) GetBlockHeight() uint {
	return s.chain.GetBlockHeight()
}

func (s *ChainService) GetBlocks() []*block.Block {
	return s.GetBlocks()
}

func (s *ChainService) GetBlock(blockNumber uint) *block.Block {
	return s.GetBlock(blockNumber)
}
