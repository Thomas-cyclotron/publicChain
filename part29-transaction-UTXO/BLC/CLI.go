package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateBlockChain -address -- 交易数据")
	fmt.Println("\taddBlock -data DATA -- 交易数据")
	fmt.Println("\tprintChain -- 输出区块信息")
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(txs []*Transaction) {

	if DBExists() == false {
		fmt.Println("数据库不存在....")
		os.Exit(1)
	}

	BlockChain := GetBlockChainObject()

	defer BlockChain.DB.Close()

	BlockChain.AddBlockToBlockChain(txs)
}

func (cli *CLI) printChain() {

	if DBExists() == false {
		fmt.Println("数据库不存在....")
		os.Exit(1)
	}

	BlockChain := GetBlockChainObject()

	defer BlockChain.DB.Close()

	BlockChain.PrintChain()

}

func (cli *CLI) createGenesisBlockChain(address string) {

	//1.创建coinbase交易

	CreateBlockChainWithGenesisBlock(address)
}

func (cli *CLI) Run() {

	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createBlockChain", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "http://www.baidu.com", "交易数据")

	flagCreateBlockChainWithAddress := createBlockChainCmd.String("address", "", "创建创世区块的地址")

	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createBlockChain":
		err := createBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:

		printUsage()
		os.Exit(1)

	}

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}

		//fmt.Println(*flagAddBlockData)
		cli.addBlock([]*Transaction{})

	}

	if printChainCmd.Parsed() {

		//fmt.Println("输出所有区块的数据....")
		cli.printChain()
	}

	if createBlockChainCmd.Parsed() {

		if *flagCreateBlockChainWithAddress == "" {
			fmt.Println("地址不能为空......")
			printUsage()
			os.Exit(1)
		}

		cli.createGenesisBlockChain(*flagCreateBlockChainWithAddress)
	}

}
