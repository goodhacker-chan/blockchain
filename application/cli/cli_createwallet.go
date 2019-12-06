package cli

import (
	"blockchain/base"
	"fmt"
)

// 创建钱包地址
func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := base.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("你的新地址: %s\n", address)
}
