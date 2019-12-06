package cli

import (
	"blockchain/base"
	"blockchain/tools"
	"fmt"
	"log"
)

// 创建钱包地址
func (cli *CLI) getBalance(address, nodeID string) {
	if !base.ValidateAddress(address) {
		log.Panic("错误: 钱包地址无效")
	}
	bc := base.NewBlockchain(nodeID)
	UTXOSet := base.UTXOSet{bc}
	defer bc.DB.Close()

	balance := 0
	pubKeyHash := tools.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("钱包地址余额 '%s': %d\n", address, balance)
}
