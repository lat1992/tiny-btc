package txPool

import (
	"slices"
	"sync"

	"github.com/lat1992/tiny-btc/internal"
)

type TxPool struct {
	pendingTxs []*internal.Transaction
	txMap      map[string]*internal.Transaction
	mapLock    sync.RWMutex
}

func NewTxPool() *TxPool {
	return &TxPool{
		pendingTxs: []*internal.Transaction{},
		txMap:      make(map[string]*internal.Transaction),
	}
}

func (tp *TxPool) AddTransaction(hash, rawTx string) {
	tx := internal.Transaction{
		Hash:   hash,
		RawTx:  rawTx,
		Status: "pending",
	}
	tp.txMap[hash] = &tx
	tp.pendingTxs = append(tp.pendingTxs, &tx)
}

func (tp *TxPool) GetPendingTransactions() []*internal.Transaction {
	return tp.pendingTxs
}

func (tp *TxPool) ValidateAndDeletePending(txs []*internal.Transaction, blockNumber uint) {
	tp.mapLock.Lock()
	for _, tx := range txs {
		tp.txMap[tx.Hash].Status = "validate"
		tp.txMap[tx.Hash].BlockNumber = blockNumber
	}
	tp.mapLock.Unlock()
	tp.pendingTxs = slices.Delete(tp.pendingTxs, 0, len(txs))
}

func (tp *TxPool) GetTx(hash string) *internal.Transaction {
	return tp.txMap[hash]
}
