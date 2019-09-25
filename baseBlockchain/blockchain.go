package baseBlockchain

import (
	"fmt"
	"log"

	"github.com/boltDB/bolt"
)

// 保存的数据库文件名配置
const DBFile = "blockchain.DB"
const blocksBucket = "blocks"

// 区块链原型
type Blockchain struct {
	Tip []byte
	DB  *bolt.DB
}

// 块迭代器
type BlockchainIterator struct {
	currentHash []byte
	DB          *bolt.DB
}

// 添加块
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.Tip = newBlock.Hash

		return nil
	})
}

// 迭代器
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.DB}

	return bci
}

// 返回下一个块
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodeDBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodeDBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}

// 创建一个带城因地块的新块链
func NewBlockchain() *Blockchain {
	var Tip []byte
	// 打开一个bolt数据库
	DB, err := bolt.Open(DBFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			genesis := NewGenesisBlock()

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			Tip = genesis.Hash
		} else {
			Tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{Tip, DB}

	return &bc
}
