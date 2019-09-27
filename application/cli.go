package application

import (
	"blockchain/baseBlockchain"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// CLI 负责处理命令参数
type CLI struct {}

// 创建区块链
func (cli *CLI) createBlockchain(address string) {
	bc := baseBlockchain.CreateBlockchain(address)
	bc.DB.Close()
	fmt.Println("创建完成!")
}

// 获取余额
func (cli *CLI) getBalance(address string) {
	bc := baseBlockchain.NewBlockchain(address)
	defer bc.DB.Close()

	balance := 0
	UTXOs := bc.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("地址 '%s' 余额: %d\n", address, balance)
}

// 验证参数
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// 在终端打印块
func (cli *CLI) printChain() {
	// TODO: Fix this
	bc := baseBlockchain.NewBlockchain("")
	defer bc.DB.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := baseBlockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// 给对方交易块
func (cli *CLI) send(from, to string, amount int) {
	bc := baseBlockchain.NewBlockchain(from)
	defer bc.DB.Close()

	tx := baseBlockchain.NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*baseBlockchain.Transaction{tx})
	fmt.Println("交易成功!")
}

// 显示使用方法
func (cli *CLI) printUsage() {
	fmt.Println("使用方法:")
	fmt.Println("  getbalance -address ADDRESS - 获取地址钱包余额")
	fmt.Println("  createblockchain -address ADDRESS - 创建一个区块")
	fmt.Println("  printchain - 打印出区块链结构")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT - 交易发送给一个地址金币")
}

// 执行命令和解析参数
func (cli *CLI) Run() {
	cli.validateArgs()

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "获取余额地址")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "发送区块链地址")
	sendFrom := sendCmd.String("from", "", "源的钱包地址")
	sendTo := sendCmd.String("to", "", "目的钱包地址")
	sendAmount := sendCmd.Int("amount", 0, "交易发送数量")

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
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
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
		cli.getBalance(*getBalanceAddress)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFrom, *sendTo, *sendAmount)
	}
}