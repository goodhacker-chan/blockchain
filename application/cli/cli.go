package cli

import (
	"flag"
	"fmt"
	"log"

	"os"
)

// CLI 负责处理命令参数
type CLI struct{}

// 显示使用方法
func (cli *CLI) printUsage() {
	fmt.Println("使用方法:")
	fmt.Println("  createwallet - 创建一个钱包并将秘钥存放到钱包文件, 返回钱包地址")
	fmt.Println("  getbalance -address ADDRESS - 获取钱包地址余额")
	fmt.Println("  createblockchain -address ADDRESS - 创建一个区块")
	fmt.Println("  listaddresses - 显示钱包文件中的所有钱包地址")
	fmt.Println("  printchain - 打印出区块链结构")
	fmt.Println("  reindexutxo - 重新构建UTXO集")
	//fmt.Println("  send -from FROM -to TO -amount AMOUNT - 交易发送给一个地址金币")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT -mine - 将一定数量的硬币从一个地址发送到另一个地址,在同一个节点上时设置 -mine(标志表示该块将被同一节点立即开采)")
	fmt.Println("  startnode -miner ADDRESS - 启动具有在NODE_ID env中指定的ID的节点, -miner实现采矿")
}

// 验证参数
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// 执行命令和解析参数
func (cli *CLI) Run() {
	cli.validateArgs()

	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Printf("NODE_ID env 没有设置(终端中执行: export NODE_ID=3000(节点))!")
		os.Exit(1)
	}

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)
	reindexUTXOCmd := flag.NewFlagSet("reindexutxo", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	startNodeCmd := flag.NewFlagSet("startnode", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "获取余额地址")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "发送区块链地址")
	sendFrom := sendCmd.String("from", "", "源的钱包地址")
	sendTo := sendCmd.String("to", "", "目的钱包地址")
	sendAmount := sendCmd.Int("amount", 0, "交易发送数量")
	sendMine := sendCmd.Bool("mine", false, "立即在同一节点上进行开采")
	startNodeMiner := startNodeCmd.String("miner", "", "启用挖矿模式并将奖励发送给ADDRESS")

	switch os.Args[1] {
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "listaddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "reindexutxo":
		err := reindexUTXOCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "startnode":
		err := startNodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		cli.printUsage()
		os.Exit(1)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress, nodeID)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress, nodeID)
	}

	if createWalletCmd.Parsed() {
		cli.createWallet(nodeID)
	}

	if listAddressesCmd.Parsed() {
		cli.listAddresses(nodeID)
	}

	if printChainCmd.Parsed() {
		cli.printChain(nodeID)
	}

	if reindexUTXOCmd.Parsed() {
		cli.reindexUTXO(nodeID)
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFrom, *sendTo, *sendAmount, nodeID, *sendMine)
	}

	if startNodeCmd.Parsed() {
		nodeID := os.Getenv("NODE_ID")
		if nodeID == "" {
			startNodeCmd.Usage()
			os.Exit(1)
		}
		cli.startNode(nodeID, *startNodeMiner)
	}
}
