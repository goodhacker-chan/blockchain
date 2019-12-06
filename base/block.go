package base

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// 块结构
type Block struct {
	Timestamp     int64  // 创建时间戳
	Transactions  []*Transaction	// 交易事务
	PrevBlockHash []byte // 前一个块的哈希
	Hash          []byte // 当前块的哈希
	Nonce         int    // 块计数器
	Height        int	// 块长度
}

// 创建块
func NewBlock(transactions []*Transaction, prevBlockHash []byte, height int) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0, height}
	//block.SetHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// 起源块初始化链
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}, 0)
}

// 返回块中事务的hash
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := NewMerkleTree(transactions)

	return mTree.RootNode.Data
}

// 序列化块
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// 反序列化一个块
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}