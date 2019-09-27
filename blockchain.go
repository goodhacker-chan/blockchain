package main

import (
	"blockchain/application"
)

func main() {
	BaseBlockchain()
}

// 基础区块实例
func BaseBlockchain() {
	cli := application.CLI{}
	cli.Run()
}
