package services

import (
	"github.com/lat1992/tiny-btc/internal"
	"github.com/lat1992/tiny-btc/internal/txPool"
)

type TransactionService struct {
	txPool *txPool.TxPool
}

func NewTransactionService(tp *txPool.TxPool) *TransactionService {
	return &TransactionService{
		txPool: tp,
	}
}

func (s *TransactionService) SendTransaction(hash, tx string) error {
	return s.txPool.AddTransaction(hash, tx)
}

func (s *TransactionService) GetTransaction(hash string) *internal.Transaction {
	return s.txPool.GetTx(hash)
}
