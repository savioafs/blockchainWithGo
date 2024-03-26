package blockchain

type Blockchain struct {
	Blocks []*Block `json:"blocks"`
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{BlockGenesis()}}
}

func (chain *Blockchain) AddBlock(data string) {
	previousBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)

	chain.Blocks = append(chain.Blocks, newBlock)
}
