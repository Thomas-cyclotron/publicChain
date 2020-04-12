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
	fmt.Println("\tsend -from From -to TO -amount AMOUNT -- 交易明细")
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

//创建创世区块
func (cli *CLI) createGenesisBlockChain(address string) {

	//1.创建coinbase交易

	CreateBlockChainWithGenesisBlock(address)
}

//转账
func (cli *CLI) send(from []string, to []string, amount []string) {
	MineNewBlock(from, to, amount)
}

func (cli *CLI) Run() {

	isValidArgs()

	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createBlockChain", flag.ExitOnError)

	flagFrom := sendBlockCmd.String("from", "", "转账源地址...")
	flagTo := sendBlockCmd.String("to", "", "转账目的地地址...")
	flagAmount := sendBlockCmd.String("amount", "", "转账金额...")

	flagCreateBlockChainWithAddress := createBlockChainCmd.String("address", "", "创建创世区块的地址")

	switch os.Args[1] {
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
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

	if sendBlockCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			printUsage()
			os.Exit(1)
		}

		//fmt.Println(*flagAddBlockData)
		//cli.addBlock([]*Transaction{})

		//fmt.Println(*flagFrom)
		//fmt.Println(*flagTo)
		//fmt.Println(*flagAmount)

		//fmt.Println(JSONToArray(*flagFrom))
		//fmt.Println(JSONToArray(*flagTo))
		//fmt.Println(JSONToArray(*flagAmount))

		from := JSONToArray(*flagFrom)
		to := JSONToArray(*flagTo)
		amount := JSONToArray(*flagAmount)
		cli.send(from, to, amount)

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
