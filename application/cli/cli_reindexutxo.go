package cli

import (
	"blockchain/base"
	"fmt"
)

// 重置未使用的块
func (cli *CLI) reindexUTXO(nodeID string) {
	bc := base.NewBlockchain(nodeID)
	UTXOSet := base.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("完成! UTXO集中有 %d 事务.\n", count)
}
