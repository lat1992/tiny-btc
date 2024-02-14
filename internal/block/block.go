package block

import "github.com/lat1992/tiny-btc/internal"

type Block struct {
	number uint
	hash   string
	txs    []*internal.Transaction
}

func NewBlock(hash string, txs []*internal.Transaction) *Block {
	newBlock := Block{
		hash: hash,
		txs:  txs,
	}
	return &newBlock
}

func (b *Block) Hash() string {
	return b.hash
}

func (b *Block) Number() uint {
	return b.number
}

func (b *Block) Txs() []internal.Transaction {
	var txs []internal.Transaction
	for _, tx := range b.txs {
		txs = append(txs, *tx)
	}
	return txs
}
