package cli

import (
	"blockchain/base"
	"fmt"
	"log"
)

// 地址列表
func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := base.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
