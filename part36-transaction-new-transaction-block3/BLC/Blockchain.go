package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
)

const dbName = "blockChain.db"
const blockTableName = "blocks"

type BlockChain struct {
	//Blocks []*Block //存储有序得区块
	Tip []byte //最新得区块得Hash
	DB  *bolt.DB
}

//迭代器
func (blockChain *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{
		CurrentHash: blockChain.Tip,
		DB:          blockChain.DB,
	}
}

//判断数据库是否存在
func DBExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

//遍历输出所有区块的信息
func (blc *BlockChain) PrintChain() {

	blockChainIterator := blc.Iterator()

	for {
		block := blockChainIterator.Next()

		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)

		fmt.Println("Txs:")

		for _, tx := range block.Txs {
			fmt.Printf("%x\n", tx.TxHash)

			fmt.Println("Vins:-----------------------")
			for _, in := range tx.Vins {
				fmt.Printf("%x\n", in.TxHash)
				fmt.Printf("%d\n", in.Vout)
				fmt.Printf("%s\n", in.ScriptSig)
			}

			fmt.Println("Vouts:----------------------")
			for _, out := range tx.Vouts {
				fmt.Printf("%d\n", out.Value)
				fmt.Printf("%s\n", out.ScriptPubKey)
			}

		}

		fmt.Println("----------------------------")

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}

	}

	//var block *Block
	//
	//var currentHash []byte = blc.Tip
	//
	//for {
	//	err := blc.DB.View(func(tx *bolt.Tx) error {
	//		//1.表
	//		b := tx.Bucket([]byte(blockTableName))
	//		if b != nil {
	//			//获取当前区块的字节数组
	//			blockBytes := b.Get(currentHash)
	//			//反序列化
	//			block = DeserializeBlock(blockBytes)
	//
	//			fmt.Printf("Height: %d\n", block.Height)
	//			fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
	//			fmt.Printf("Data: %s\n", block.Data)
	//			fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
	//			fmt.Printf("Hash: %x\n", block.Hash)
	//			fmt.Printf("Nonce: %d\n", block.Nonce)
	//
	//		}
	//
	//		return nil
	//	})
	//
	//	fmt.Println()
	//
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//
	//	var hashInt big.Int
	//	hashInt.SetBytes(block.PrevBlockHash)
	//
	//	if big.NewInt(0).Cmp(&hashInt) == 0 {
	//		break
	//	}
	//
	//	currentHash = block.PrevBlockHash

}

//2.增加区块到区块链里面
func (blc *BlockChain) AddBlockToBlockChain(txs []*Transaction) {

	//往链里面添加区块
	//blc.Blocks = append(blc.Blocks, newBlock)
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//1.获取表
		b := tx.Bucket([]byte(blockTableName))
		//2.创建新区块
		if b != nil {

			//先获取最新区块
			blockBytes := b.Get(blc.Tip)
			//反序列化
			block := DeserializeBlock(blockBytes)

			//3.将区块序列化并且存储到数据库中
			newBlock := NewBlock(txs, block.Height+1, block.Hash)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//4.更新数据库里面"l"对应得hash
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//5.更新blockChain的Tip
			blc.Tip = newBlock.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

//1.创建带有创世区块得区块链
func CreateBlockChainWithGenesisBlock(address string) *BlockChain {

	// 判断数据库是否存在

	if DBExists() {

		fmt.Println("数据库已存在.....")
		os.Exit(1)
	}

	fmt.Println("正在创建创世区块......")

	//创建或打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var genesisHash []byte

	err = db.Update(func(tx *bolt.Tx) error {

		//创建数据库表
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic(err)
		}

		if b != nil {
			//创建创世区块

			//创建了一个coinBase Transaction
			txCoinBase := NewCoinBaseTransaction(address)

			genesisBlock := CreateGenesisBlock([]*Transaction{txCoinBase})
			//将创世区块存储到表当中
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//存储最新区块得Hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			genesisHash = genesisBlock.Hash
		}
		return nil
	})

	return &BlockChain{
		Tip: genesisHash,
		DB:  db,
	}

}

func GetBlockChainObject() *BlockChain {
	//创建或打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var tip []byte

	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))

		if b != nil {

			//读取最新区块的hash
			tip = b.Get([]byte("l"))

		}

		return nil
	})

	return &BlockChain{
		Tip: tip,
		DB:  db,
	}
}

//如果一个地址对应的Txoutput未花费 那么这个Transaction就应该添加到数组中返回
func UnSpentTransactionsWithAddress(address string) []*Transaction {

	return nil
}

//挖到新的区块
func (blockchain *BlockChain) MineNewBlock(from []string, to []string, amount []string) {

	fmt.Println(from)
	fmt.Println(to)
	fmt.Println(amount)

	//./bc send -from '["d"]' -to '["e"]' -amount '["2"]'
	//[d]
	//[e]
	//[2]
	//1.建立一笔交易
	value, _ := strconv.Atoi(amount[0])
	tx := NewSimpleTransaction(from[0], to[0], value)

	//1.通过相关算法 建立Transaction数组
	var txs []*Transaction
	txs = append(txs, tx)

	var block *Block
	blockchain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			hash := b.Get([]byte("l"))

			blockBytes := b.Get(hash)

			block = DeserializeBlock(blockBytes)
		}

		return nil
	})

	//2.建新的区块
	block = NewBlock(txs, block.Height+1, block.Hash)

	//1.将新区块存储到数据库
	blockchain.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			b.Put(block.Hash, block.Serialize())

			b.Put([]byte("l"), block.Hash)

			blockchain.Tip = block.Hash
		}

		return nil
	})
}
