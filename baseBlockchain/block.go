package baseBlockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 块结构
type Block struct {
	Timestamp     int64  // 创建时间戳
	Data          []byte // 区块中包含的实际信息
	PrevBlockHash []byte // 前一个块的哈希
	Hash          []byte // 当前块的哈希
	Nonce         int    // 块计数器
}

// 设置哈希
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// 创建块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	//block.SetHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// 成因地块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
