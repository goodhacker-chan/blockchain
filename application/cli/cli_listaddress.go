package cli

import (
	"blockchain/baseBlockchain"
	"fmt"
	"log"
)

// 地址列表
func (cli *CLI) listAddresses() {
	wallets, err := baseBlockchain.NewWallets()
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}