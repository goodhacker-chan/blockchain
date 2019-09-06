package easyBlockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// 块结构
type Block struct {
	Timestamp     int64	// 创建时间戳
	Data          []byte	// 区块中包含的实际信息
	PrevBlockHash []byte	// 前一个块的哈希
	Hash          []byte	// 当前块的哈希
}

// 区块链原型
type Blockchain struct {
	blocks []*Block
}

/**
 * 设置哈希
 */
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

/**
 * 创建块
 * @param data 数据
 */
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

/**
 * 添加块
 */
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

/**
 * 成因地块
 * @des
 */
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

/**
 * 创建一个带城因地块的块链
 */
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func MyBlockchain() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	fmt.Println("111")
	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}