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
