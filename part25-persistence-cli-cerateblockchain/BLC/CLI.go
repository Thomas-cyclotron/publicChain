package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	BC *BlockChain
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateBlockChain -data DATA -- 交易数据")
	fmt.Println("\taddBlock -data DATA -- 交易数据")
	fmt.Println("\tprintChain -- 输出区块信息")
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.BC.AddBlockToBlockChain(data)
}

func (cli *CLI) printChain() {
	cli.BC.PrintChain()
}

func (cli *CLI) createGenesisBlockChain(data string) {
	fmt.Println(data)
}

func (cli *CLI) Run() {

	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createBlockChain", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "http://www.baidu.com", "交易数据")

	flagCreateBlockChainWithData := createBlockChainCmd.String("data", "Genesis block data......", "创世区块交易数据")

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
		cli.addBlock(*flagAddBlockData)

	}

	if printChainCmd.Parsed() {

		fmt.Println("输出所有区块的数据....")
		cli.printChain()
	}

	if createBlockChainCmd.Parsed() {

		if *flagCreateBlockChainWithData == "" {
			fmt.Println("交易数据不能为空......")
			printUsage()
			os.Exit(1)
		}

		cli.createGenesisBlockChain(*flagCreateBlockChainWithData)
	}

}
