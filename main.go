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

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{
		b.Data,
		b.PreviousHash,
	}, []byte{})

	hash := sha256.Sum256(info)

	b.Hash = hash[:]
}

func main() {

}
