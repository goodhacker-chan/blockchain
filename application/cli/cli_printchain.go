package cli

import (
	"blockchain/base"
	"fmt"
	"strconv"
)

// 打印链
func (cli *CLI) printChain(nodeID string) {
	bc := base.NewBlockchain(nodeID)
	defer bc.DB.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ 块 %x ============\n", block.Hash)
		fmt.Printf("块高度: %d\n", block.Height)
		fmt.Printf("上一个. 块: %x\n", block.PrevBlockHash)
		pow := base.NewProofOfWork(block)
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
