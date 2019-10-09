package cli

import (
	"bolckchain/baseBlockchain"
	"bolckchain/tools"
	"fmt"
	"log"
)

// 创建钱包地址
func (cli *CLI) getBalance(address string) {
	if !baseBlockchain.ValidateAddress(address) {
		log.Panic("错误: 地址无效")
	}
	bc := baseBlockchain.NewBlockchain(address)
	defer bc.DB.Close()

	balance := 0
	pubKeyHash := tools.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := bc.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("地址余额 '%s': %d\n", address, balance)
}