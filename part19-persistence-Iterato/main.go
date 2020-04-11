package main

import (
	"publicChain/part19-persistence-Iterato/BLC"
)

func main() {
	//创世区块
	blockChain := BLC.CreateBlockChainWithGenesisBlock()
	defer blockChain.DB.Close()
	////添加新区块
	blockChain.AddBlockToBlockChain("Send 100RMB To A")
	blockChain.AddBlockToBlockChain("Send 200RMB To B")
	blockChain.AddBlockToBlockChain("Send 300RMB To C")
	blockChain.AddBlockToBlockChain("Send 400RMB To D")
	blockChain.AddBlockToBlockChain("Send 500RMB To E")
	blockChain.PrintChain()
	//
	//fmt.Println(blockChain)
	//fmt.Println(blockChain.Blocks)
	//block := BLC.NewBlock("Test", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	//fmt.Printf("%d\n", block.Nonce)
	//fmt.Printf("%x\n", block.Hash)
	//
	//bytes := block.Serialize()
	//fmt.Println(bytes)
	//block = BLC.DeserializeBlock(bytes)
	//
	//fmt.Printf("%d\n", block.Nonce)
	//fmt.Printf("%x\n", block.Hash)
	//
	////创建或打开数据库
	//db, err := bolt.Open("my.db", 0600, nil)
	//if err != nil {
	//	log.Fatal(err)
	//
	//}
	//defer db.Close()

	////更新数据库
	//err = db.Update(func(tx *bolt.Tx) error {
	//	//取表对象
	//	b := tx.Bucket([]byte("blocks"))
	//	if b == nil {
	//		b, err = tx.CreateBucket([]byte("blocks"))
	//		if err != nil {
	//			log.Panic("blocks create failed......")
	//		}
	//	}
	//
	//	err = b.Put([]byte("l"), block.Serialize())
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//
	//	return nil
	//})
	//
	//if err != nil {
	//	log.Panic(err)
	//}

	//err = db.View(func(tx *bolt.Tx) error {
	//	b := tx.Bucket([]byte("blocks"))
	//	if b != nil {
	//		blockData := b.Get([]byte("l"))
	//		//fmt.Println(blockData)
	//		//fmt.Printf("%s\n",blockData)
	//		block := BLC.DeserializeBlock(blockData)
	//		fmt.Printf("%v\n", block)
	//	}
	//	return nil
	//})
	//if err != nil {
	//	log.Panic(err)
	//}
}
