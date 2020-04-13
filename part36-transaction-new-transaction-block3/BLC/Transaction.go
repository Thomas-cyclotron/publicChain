package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
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
func NewSimpleTransaction(from string, to string, amount int) *Transaction {

	//1.有一个函数 可以返回from这个人所有的未话费交易输出所对应的Transaction

	//var txIntputs []*TXInput
	//var txOutputs []*TXOutput

	unSpentTx := UnSpentTransactionsWithAddress(from)

	fmt.Println(unSpentTx)

	//Height: 3
	//PrevBlockHash: 0000a4ce2ca9d64c5c068a3c01249cc7ac26c886e9f40caafd7ca1d58461498e
	//Timestamp: 2020-04-13 02:11:14 PM
	//Hash: 0000e0f4250d2c1670ef2726926078182ac47b958df66f29f49a01525519464d
	//Nonce: 77080
	//Txs:
	//	ce8e4ab58680fbe0bba20d8cacd0dfe2ffced2514ab14fa7e33276942df26cd6
	//Vins:-----------------------
	//		e8d2eefd6459cc2c94039e9b22ac035fa25d569333421a6ff846ea23467e1465
	//	0
	//	e
	//Vouts:----------------------
	//	2
	//	f
	//	8
	//	e
	//	----------------------------
	//Height: 2
	//PrevBlockHash: 0000f180439fb29ce2adbdb924fffe0d62a4863a184e3f5e14df7f0dbc9969bd
	//Timestamp: 2020-04-13 02:08:28 PM
	//Hash: 0000a4ce2ca9d64c5c068a3c01249cc7ac26c886e9f40caafd7ca1d58461498e
	//Nonce: 40644
	//Txs:
	//	3ff80d01957610f63ca062987f8df7bddcac1b5e13ad7b9e24b3abc036309aaa
	//Vins:-----------------------
	//		e8d2eefd6459cc2c94039e9b22ac035fa25d569333421a6ff846ea23467e1465
	//	0
	//	d
	//Vouts:----------------------
	//	2
	//	e
	//	8
	//	d
	//	----------------------------
	//Height: 1
	//PrevBlockHash: 0000000000000000000000000000000000000000000000000000000000000000
	//Timestamp: 2020-04-13 02:08:22 PM
	//Hash: 0000f180439fb29ce2adbdb924fffe0d62a4863a184e3f5e14df7f0dbc9969bd
	//Nonce: 7071
	//Txs:
	//	e8d2eefd6459cc2c94039e9b22ac035fa25d569333421a6ff846ea23467e1465
	//Vins:-----------------------
	//
	//		-1
	//	Genesis Data
	//Vouts:----------------------
	//	10
	//	--Test
	//	----------------------------

	//输入 代表消费
	//bytes, _ := hex.DecodeString("e8d2eefd6459cc2c94039e9b22ac035fa25d569333421a6ff846ea23467e1465")
	//
	//txInput := &TXInput{bytes, 0, from}
	//
	////s := hex.EncodeToString(b)
	//
	////fmt.Printf("s: %s\n",s)
	//
	//txIntputs = append(txIntputs, txInput)
	//
	////输出
	//txOutput := &TXOutput{int64(amount), to}
	//txOutputs = append(txOutputs, txOutput)
	//
	////找零
	//txOutput = &TXOutput{4 - int64(amount), from}
	//txOutputs = append(txOutputs, txOutput)
	//
	//tx := &Transaction{TxHash: []byte{}, Vins: txIntputs, Vouts: txOutputs}
	//
	////设置hash值
	//tx.HashTransaction()

	return nil

}
