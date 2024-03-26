package blockchain

type Block struct {
	Hash         []byte `json:"hash"`
	Data         []byte `json:"data"`
	PreviousHash []byte `json:"previousHash"`
	Nonce        int    `json:"nonce"`
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash, 0}
	pow := NewProof(block)

	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

func BlockGenesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
