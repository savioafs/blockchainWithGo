package main

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash         []byte `json:"hash"`
	Data         []byte `json:"data"`
	PreviousHash []byte `json:"previousHash"`
}

type Blockchain struct {
	Blocks []*Block `json:"blocks"`
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{
		b.Data,
		b.PreviousHash,
	}, []byte{})

	hash := sha256.Sum256(info)

	b.Hash = hash[:]
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash}
	block.DeriveHash()

	return block
}

func (chain *Blockchain) AddBlock(data string) {
	previousBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, previousBlock.PreviousHash)

	chain.Blocks = append(chain.Blocks, newBlock)
}

func main() {

}
