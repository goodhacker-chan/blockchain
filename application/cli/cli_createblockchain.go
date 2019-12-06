package cli

import (
	"blockchain/base"
	"fmt"
	"log"
)

// 创建区块链
func (cli *CLI) createBlockchain(address, nodeID string) {
	if !base.ValidateAddress(address) {
		log.Panic("错误: 地址无效")
	}

	bc := base.CreateBlockchain(address, nodeID)
	defer bc.DB.Close()

	UTXOSet := base.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("成功!")
}
