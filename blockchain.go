package main

import (
	"blockchain/application/cli"
)

func main() {
	BaseBlockchain()
}

// 基础区块实例
func BaseBlockchain() {
	cli := cli.CLI{}
	cli.Run()
}
