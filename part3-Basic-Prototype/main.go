package main

import (
	"fmt"
	"publicChain/part3-Basic-Prototype/BLC"
)

func main() {
	genesisBlockChain := BLC.CreateBlockChainWithGenesisBlock()
	fmt.Println(genesisBlockChain)
	fmt.Println(genesisBlockChain.Blocks)
	fmt.Println(genesisBlockChain.Blocks[0])
}
