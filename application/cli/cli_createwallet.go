package cli

import (
	"bolckchain/baseBlockchain"
	"fmt"
)

// 创建钱包地址
func (cli *CLI) createWallet() {
	wallets, _ := baseBlockchain.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("你的新地址: %s\n", address)
}