package cli

import (
	"blockchain/base"
	"fmt"
	"log"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("起始节点 %s\n", nodeID)
	if len(minerAddress) > 0 {
		if base.ValidateAddress(minerAddress) {
			fmt.Println("采矿进行中... 收益转入到地址: ", minerAddress)
		} else {
			log.Panic("矿工地址错误!")
		}
	}
	base.StartServer(nodeID, minerAddress)
}
