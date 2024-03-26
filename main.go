package main

import (
	"fmt"
	"strconv"

	"github.com/savioafs/blockchainWithGo/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {
		fmt.Println("---------------------------")
		fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
