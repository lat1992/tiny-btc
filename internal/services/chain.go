package services

import (
	"github.com/lat1992/tiny-btc/internal"
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

func (s *ChainService) GetBlocks() []internal.Block {
	blocks := s.chain.GetBlocks()
	var bs []internal.Block
	for _, b := range blocks {
		bs = append(bs, internal.Block{
			Hash:   b.Hash(),
			Number: b.Number(),
			Txs:    b.Txs(),
		})
	}
	return bs
}

func (s *ChainService) GetBlock(blockNumber uint) internal.Block {
	b := s.chain.GetBlock(blockNumber)
	return internal.Block{
		Hash:   b.Hash(),
		Number: b.Number(),
		Txs:    b.Txs(),
	}
}
