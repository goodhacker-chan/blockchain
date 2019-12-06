package cli

import (
	"blockchain/base"
	"fmt"
	"log"
)

// 交易金币
func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !base.ValidateAddress(from) {
		log.Panic("错误: 发送人地址无效")
	}
	if !base.ValidateAddress(to) {
		log.Panic("错误: 接收人地址无效")
	}

	bc := base.NewBlockchain(nodeID)
	UTXOSet := base.UTXOSet{bc}
	defer bc.DB.Close()

	wallets, err := base.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := base.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := base.NewCoinbaseTX(from, "")
		//txs := []*base.Transaction{tx}
		txs := []*base.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		base.SendTx(base.KnownNodes[0], tx)
	}

	fmt.Println("交易成功!")
}
