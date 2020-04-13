package BLC

//创建创世区块
func (cli *CLI) createGenesisBlockChain(address string) {

	//1.创建coinbase交易
	blockchain := CreateBlockChainWithGenesisBlock(address)
	defer blockchain.DB.Close()
}
