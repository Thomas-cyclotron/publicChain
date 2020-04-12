package main

import (
	"publicChain/part24-persistence-cli/BLC"
)

func main() {
	//创世区块
	blockChain := BLC.CreateBlockChainWithGenesisBlock()
	defer blockChain.DB.Close()
	////添加新区块
	//blockChain.AddBlockToBlockChain("Send 100RMB To A")
	//blockChain.AddBlockToBlockChain("Send 200RMB To B")
	//blockChain.AddBlockToBlockChain("Send 300RMB To C")
	//blockChain.AddBlockToBlockChain("Send 400RMB To D")
	//blockChain.AddBlockToBlockChain("Send 500RMB To E")
	//blockChain.PrintChain()

	cli := BLC.CLI{blockChain}

	cli.Run()

}
