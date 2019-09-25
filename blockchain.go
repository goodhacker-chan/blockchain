package main

import (
	"blockchain/application"
	"blockchain/baseBlockchain"
)

func main() {
	BaseBlockchain()
}

// 基础区块实例
func BaseBlockchain() {
	bc := baseBlockchain.NewBlockchain()
	defer bc.DB.Close()

	cli := application.CLI{bc}
	cli.Run()
}
