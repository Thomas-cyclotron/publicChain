package BLC

type BlockChain struct {
	Blocks []*Block //存储有序得区块
}

//2.增加区块到区块链里面
func (blc *BlockChain) AddBlockToBlockChain(data string, height int64, prevHash []byte) {
	//创建新区块
	newBlock := NewBlock(data, height, prevHash)
	//往链里面添加区块
	blc.Blocks = append(blc.Blocks, newBlock)
}

//1.创建带有创世区块得区块链
func CreateBlockChainWithGenesisBlock() *BlockChain {
	//创建创世区块
	genesisBlock := CreateGenesisBlock("Genesis Data.....")
	//返回区块链对象
	return &BlockChain{Blocks: []*Block{genesisBlock}}
}
