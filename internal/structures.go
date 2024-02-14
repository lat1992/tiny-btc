package internal

type Transaction struct {
	Hash        string
	RawTx       string
	BlockNumber uint
	Status      string
}

type Block struct {
	Number uint
	Hash   string
	Txs    []Transaction
}
