package BLC

import (
	"fmt"
	"time"
)

//定义区块
type Block struct {
	//1.区块高度
	Height int64
	//2.上一个区块得hash
	PrevBlockHash []byte
	//3.交易数据
	Data []byte
	//4.时间戳
	Timestamp int64
	//5.hash
	Hash []byte
	//6.Nonce
	Nonce int64
}

//func (block *Block) SetHash() {
//	//1.height转换为字节数组[]byte
//	heightBytes := IntToHed(block.Height)
//
//	//2.时间戳转换为字节数组[]byte
//	timeString := strconv.FormatInt(block.Timestamp, 2)
//	timeBytes := []byte(timeString)
//	//3.拼接所有得属性
//	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})
//	//4.生成Hash
//	//将二维的切片数组连接起来返回一个一维的切片
//	hash := sha256.Sum256(blockBytes)
//	//sha256 返回得是一个32字节得数组 需要转换一下
//	block.Hash = hash[:]
//}

//1.创建新得区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {

	//创建区块
	block := &Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
		Nonce:         0,
	}

	//设置Hash
	//block.SetHash()
	//调用工作量证明得方法并且返回有效得Hash和Nonce值
	pow := NewProofOfWork(block)

	//挖矿验证
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	fmt.Println()
	return block
}

//2.单独写一个方法 生成创世区块
func CreateGenesisBlock(data string) *Block {

	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
