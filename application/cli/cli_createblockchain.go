package cli

import (
	"blockchain/baseBlockchain"
	"fmt"
	"log"
)

// 创建区块链
func (cli *CLI) createBlockchain(address string) {
	if !baseBlockchain.ValidateAddress(address) {
		log.Panic("错误: 地址无效")
	}

	bc := baseBlockchain.CreateBlockchain(address)
	defer bc.DB.Close()

	UTXOSet := baseBlockchain.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("成功!")
}