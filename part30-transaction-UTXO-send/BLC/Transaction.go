package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// UTXO
type Transaction struct {
	//1. 交易hash
	TxHash []byte

	//2. 输入
	Vins []*TXInput

	//3. 输出
	Vouts []*TXOutput
}

//1.Transaction 创建分两种情况
//2.创世区块创建时的Transaction
func NewCoinBaseTransaction(address string) *Transaction {

	//输入 代表消费
	txInput := &TXInput{[]byte{}, -1, "Genesis Data"}

	//输出
	txOutput := &TXOutput{10, address}

	txCoinBase := &Transaction{TxHash: []byte{}, Vins: []*TXInput{txInput}, Vouts: []*TXOutput{txOutput}}

	//设置hash值
	txCoinBase.HashTransaction()

	return txCoinBase
}

//将区块对象序列化 成字节数组
func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())

	tx.TxHash = hash[:]
}

//2.转账时的Transaction
