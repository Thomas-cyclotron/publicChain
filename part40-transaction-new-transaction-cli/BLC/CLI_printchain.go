package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) printChain() {

	if DBExists() == false {
		fmt.Println("数据库不存在....")
		os.Exit(1)
	}

	BlockChain := GetBlockChainObject()

	defer BlockChain.DB.Close()

	BlockChain.PrintChain()

}
