package main

import (
	"flag"
	"fmt"
)

func main() {
	flagString := flag.String("printChain", "", "输出所有的区块信息...")

	flagInt := flag.Int("number", 6, "输出一个整数...")

	flagBool := flag.Bool("open", false, "判断真假...")

	flag.Parse()

	fmt.Printf("%s\n", *flagString)
	fmt.Printf("%d\n", *flagInt)
	fmt.Printf("%v\n", *flagBool)
}

//bc
//./bc addBlock -data "feige"
//./bc printChain
//即将输出所有区块
