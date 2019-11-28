package cli

import (
	"blockchain/baseBlockchain"
	"fmt"
	"strconv"
)

// 打印链
func (cli *CLI) printChain() {
	bc := baseBlockchain.NewBlockchain()
	defer bc.DB.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ 块 %x ============\n", block.Hash)
		fmt.Printf("上一个. 块: %x\n", block.PrevBlockHash)
		pow := baseBlockchain.NewProofOfWork(block)
		fmt.Printf("散列: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}