package cli

import (
	"blockchain/baseBlockchain"
	"fmt"
	"log"
)

// 打印链
func (cli *CLI) send(from, to string, amount int) {
	if !baseBlockchain.ValidateAddress(from) {
		log.Panic("错误: 发送人地址无效")
	}
	if !baseBlockchain.ValidateAddress(to) {
		log.Panic("错误: 接收人地址无效")
	}

	bc := baseBlockchain.NewBlockchain()
	UTXOSet := baseBlockchain.UTXOSet{bc}
	defer bc.DB.Close()

	tx := baseBlockchain.NewUTXOTransaction(from, to, amount, &UTXOSet)
	//cbTx := baseBlockchain.NewCoinbaseTX(from, "")
	//txs := []*baseBlockchain.Transaction{cbTx, tx}
	txs := []*baseBlockchain.Transaction{tx}
	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)

	fmt.Println("交易成功!")
}