package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

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
