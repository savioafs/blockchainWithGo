package blockchain

type Block struct {
	Hash         []byte `json:"hash"`
	Data         []byte `json:"data"`
	PreviousHash []byte `json:"previousHash"`
	Nonce        int    `json:"nonce"`
}

type Blockchain struct {
	Blocks []*Block `json:"blocks"`
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash, 0}
	pow := NewProof(block)

	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

func (chain *Blockchain) AddBlock(data string) {
	previousBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)

	chain.Blocks = append(chain.Blocks, newBlock)
}

func BlockGenesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{BlockGenesis()}}
}
