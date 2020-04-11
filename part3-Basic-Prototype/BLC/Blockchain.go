package BLC

type BlockChain struct {
	Blocks []*Block //存储有序得区块
}

//1.创建带有创世区块得区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	//创建创世区块
	genesisBlock := CreateGenesisBlock("Genesis Data.....")
	//返回区块链对象
	return &BlockChain{Blocks: []*Block{genesisBlock}}
}
