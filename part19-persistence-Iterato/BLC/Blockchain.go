package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
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

type BlockChainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

func (blockChainIterator *BlockChainIterator) Next() *Block {

	var block *Block

	err := blockChainIterator.DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			currentBlockBytes := b.Get(blockChainIterator.CurrentHash)
			//获取到 当前迭代器里面的currentHash所对应的区块
			block = DeserializeBlock(currentBlockBytes)

			//更新迭代器里面的CurrentHash
			blockChainIterator.CurrentHash = block.PrevBlockHash

		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return block
}

//遍历输出所有区块的信息
func (blc *BlockChain) PrintChain() {

	blockChainIterator := blc.Iterator()

	for {
		block := blockChainIterator.Next()

		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)

		fmt.Println()

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
func (blc *BlockChain) AddBlockToBlockChain(data string) {

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
			newBlock := NewBlock(data, block.Height+1, block.Hash)
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
func CreateBlockChainWithGenesisBlock() *BlockChain {

	//创建或打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {

		//获取表
		b := tx.Bucket([]byte(blockTableName))

		if b == nil {
			//创建数据库表
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panic(err)
			}
		}

		if b != nil {
			//创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Data.....")
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
			blockHash = genesisBlock.Hash
		}
		return nil
	})
	//defer db.Close()

	//返回区块链对象
	//return &BlockChain{Blocks: []*Block{genesisBlock}}
	return &BlockChain{blockHash, db}
}
