package cli

import (
	"blockchain/baseBlockchain"
	"fmt"
)

// 重置未使用的块
func (cli *CLI) reindexUTXO() {
	bc := baseBlockchain.NewBlockchain()
	UTXOSet := baseBlockchain.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("完成! UTXO集中有 %d 事务.\n", count)
}