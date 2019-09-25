package main

import (
	"blockchain/baseBlockchain"
	"fmt"
	"strconv"
)

func main() {
	BaseBlockchain()
}

// 基础区块实例
func BaseBlockchain() {
	bc := baseBlockchain.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		// 计算散列
		pow := baseBlockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
