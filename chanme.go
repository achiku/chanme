package chanme

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Chain chain
type Chain struct {
	Blocks []*Block `json:"blocks"`
	mux    sync.Mutex
}

// NewBlockChain new chain
func NewBlockChain() *Chain {
	return &Chain{
		Blocks: []*Block{genesis()},
	}
}

// Block block
type Block struct {
	Index        int64  `json:"index"`
	PreviousHash string `json:"previous_hash"`
	Timestamp    int64  `json:"timestamp"`
	Data         string `json:"data"`
	Hash         string `json:"hash"`
}

// GenerateNextBlock generate next block
func (cn *Chain) GenerateNextBlock(data string) (*Block, error) {
	if len(cn.Blocks) == 0 {
		return nil, errors.New("block lenght = 0")
	}
	cn.mux.Lock()
	defer cn.mux.Unlock()

	prev := cn.Blocks[len(cn.Blocks)-1]
	b := &Block{
		Index:        prev.Index + 1,
		PreviousHash: prev.Hash,
		Timestamp:    time.Now().UnixNano(),
		Data:         data,
	}
	b.Hash = Hash(b)
	cn.Blocks = append(cn.Blocks, b)
	return b, nil
}

// genesis the first block
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

// ValidateNewBlock validate new blocks
func ValidateNewBlock(prev, next *Block) error {
	switch {
	case prev.Index+1 != next.Index:
		return errors.New("invalid index")
	case prev.Hash != next.PreviousHash:
		return errors.New("invalid previous hash")
	case next.Hash != Hash(next):
		return errors.New("invalid hash")
	default:
		return nil
	}
}

// ValidateBlockChain validate block chain
func ValidateBlockChain(chn *Chain, genesis *Block) error {
	chn.mux.Lock()
	defer chn.mux.Unlock()

	ib := chn.Blocks[0]
	if ib.Hash != genesis.Hash {
		return errors.New("invalid genesis block hash")
	}
	tmpChain := []*Block{ib}
	for i := 1; i < len(chn.Blocks); i++ {
		if err := ValidateNewBlock(tmpChain[i-1], chn.Blocks[i]); err != nil {
			return errors.New("invalid chain")
		}
		tmpChain = append(tmpChain, chn.Blocks[i])
	}
	return nil
}

// RefreshChain refresh chain
func RefreshChain(curr, received *Chain, genesis *Block) (*Chain, error) {
	if err := ValidateBlockChain(received, genesis); err != nil {
		return nil, errors.New("invalid block chain")
	}
	if len(received.Blocks) > len(curr.Blocks) {
		return received, nil
	}
	return curr, nil
}
