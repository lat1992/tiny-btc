package txPool

import (
	"slices"

	"github.com/lat1992/tiny-btc/internal"
)

type TxPool struct {
	pendingTxs []*internal.Transaction
	txMap      map[string]*internal.Transaction
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
	// duplicate the pending transaction list.
	var pendingTxs []*internal.Transaction
	copy(pendingTxs, tp.pendingTxs)
	return pendingTxs
}

func (tp *TxPool) Validate(txs []*internal.Transaction, blockNumber uint) {
	for _, tx := range txs {
		tp.txMap[tx.Hash].Status = "validate"
		tp.txMap[tx.Hash].BlockNumber = blockNumber
	}
	slices.Delete(tp.pendingTxs, 0, len(txs)-1)
}

func (tp *TxPool) GetTx(hash string) *internal.Transaction {
	return tp.txMap[hash]
}