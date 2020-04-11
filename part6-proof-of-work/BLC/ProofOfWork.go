package BLC

import "math/big"

//代表256位hash 里面前面至少要有16个0
const targetBit = 16

type ProofOfWork struct {
	block  *Block   //当前要验证得区块
	target *big.Int //大数据存储
}

func (proofofWork *ProofOfWork) Run() ([]byte, int64) {
	return nil, 0
}

//创建新得工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	//1.big,Int对象

	//1.创建一个初始值为1 得target
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBit)
	//2.左移256 - targetBit
	return &ProofOfWork{block: block, target: target}
}
