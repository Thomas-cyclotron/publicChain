package BLC

import "fmt"

//查询余额
func (cli *CLI) getBalance(address string) {

	fmt.Println("测试测试" + address)

	blockchain := GetBlockChainObject()
	defer blockchain.DB.Close()

	amount := blockchain.GetBalance(address)

	fmt.Printf("%s一共有%d个Token\n", address, amount)

}
