package cli

import (
	"bolckchain/baseBlockchain"
	"fmt"
	"log"
)

// 创建区块链
func (cli *CLI) createBlockchain(address string) {
	if !baseBlockchain.ValidateAddress(address) {
		log.Panic("错误: 地址无效")
	}
	bc := baseBlockchain.CreateBlockchain(address)
	bc.DB.Close()
	fmt.Println("成功!")
}