package BLC

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

//代表256位hash 里面前面至少要有16个0
const targetBit = 16

type ProofOfWork struct {
	Block  *Block   //当前要验证得区块
	target *big.Int //大数据存储
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{},
	)
	return data
}

func (proofofWork *ProofOfWork) IsValid() bool {
	//1.proofofWork.Block.Hash
	//2.proofofWork.Target
	var hashInt big.Int
	hashInt.SetBytes(proofofWork.Block.Hash)
	if proofofWork.target.Cmp(&hashInt) == 1 {
		return true
	}

	return false
}

func (proofofWork *ProofOfWork) Run() ([]byte, int64) {
	//1.将Block得属性拼接成字节数序
	//2.生成Hash
	//3.判断hash有效性 如果满足条件跳出循环
	nonce := 0

	var hashInt big.Int //存储新生成得哈希值
	var hash [32]byte

	for {
		//准备数据
		dataBytes := proofofWork.prepareData(nonce)

		//生成hash
		hash = sha256.Sum256(dataBytes)
		//fmt.Printf("\r%x", hash)
		//将hash存储到hashInt
		hashInt.SetBytes(hash[:])

		//判断hashInt 是否小于区块Block里面得target
		if proofofWork.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce = nonce + 1

	}
	return hash[:], int64(nonce)
}

//创建新得工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	//1.big,Int对象

	//1.创建一个初始值为1 得target
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBit)
	//2.左移256 - targetBit
	return &ProofOfWork{Block: block, target: target}
}
