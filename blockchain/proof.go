package blockchain

import "math/big"

const Difficulty = 12 // create algorithm for chose to difficulty

type ProofOfWork struct {
	Block  *Block   `json:"block"`
	Target *big.Int `json:"target"`
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{
		b,
		target,
	}

	return pow
}
