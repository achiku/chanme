package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
)

var iniBlock = genesis()
var blockChain = []*Block{iniBlock}

// Block block
type Block struct {
	Index        int64
	PreviousHash string
	Timestamp    int64
	Data         string
	Hash         string
}

// GenerateNextBlock generate next block
func GenerateNextBlock(prev *Block, data string) *Block {
	b := &Block{
		Index:        prev.Index + 1,
		PreviousHash: prev.Hash,
		Timestamp:    time.Now().UnixNano(),
		Data:         data,
	}
	b.Hash = Hash(b)
	return b
}

// genesis get the first block
func genesis() *Block {
	genesis := &Block{
		Index:        0,
		PreviousHash: "0",
		Data:         "I'm the begining",
		Timestamp:    time.Date(1985, 8, 18, 0, 0, 0, 0, time.UTC).UnixNano(),
	}
	genesis.Hash = Hash(genesis)
	return genesis
}

// Hash get block hash
func Hash(b *Block) string {
	str := fmt.Sprintf("%d%s%d%s", b.Index, b.PreviousHash, b.Timestamp, b.Data)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

// ValidateNewBlock validate paired blocks
func ValidateNewBlock(prev, n *Block) error {
	switch {
	case prev.Index+1 != n.Index:
		return errors.New("invalid index")
	case prev.Hash != n.PreviousHash:
		return errors.New("invalid previous hash")
	case n.Hash != Hash(n):
		return errors.New("invalid hash")
	default:
		return nil
	}
}

// ValidateBlockChain validate block chain
func ValidateBlockChain(bc []*Block) error {
	ib := bc[0]
	if ib.Hash != iniBlock.Hash {
		return errors.New("invalid genesis block hash")
	}
	tmpChain := []*Block{ib}
	for i := 1; i < len(bc); i++ {
		if err := ValidateNewBlock(tmpChain[i-1], bc[i]); err != nil {
			return errors.New("invalid chain")
		}
		tmpChain = append(tmpChain, bc[i])
	}
	return nil
}

// ReplaceChain replace chain
func ReplaceChain(n []*Block) (bool, error) {
	if err := ValidateBlockChain(n); err != nil {
		return false, errors.New("invalid block chain")
	}
	if len(n) > len(blockChain) {
		blockChain = n
		return true, nil
	}
	return false, nil
}
