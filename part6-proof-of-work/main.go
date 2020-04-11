package main

import (
	"fmt"
	"publicChain/part6-proof-of-work/BLC"
)

func main() {
	//创世区块
	blockChain := BLC.CreateBlockChainWithGenesisBlock()

	//添加新区块
	blockChain.AddBlockToBlockChain("Send 100RMB To A", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 200RMB To B", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 300RMB To C", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 400RMB To D", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlockToBlockChain("Send 500RMB To E", blockChain.Blocks[len(blockChain.Blocks)-1].Height+1, blockChain.Blocks[len(blockChain.Blocks)-1].Hash)

	fmt.Println(blockChain)
	fmt.Println(blockChain.Blocks)
}
