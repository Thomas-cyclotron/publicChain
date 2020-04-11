package main

import "fmt"
import "publicChain/part1-Basic-Prototype/BLC"

func main() {
	fmt.Println("测试")
	block := BLC.NewBlock("Genenis Block",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Println(block)
}
