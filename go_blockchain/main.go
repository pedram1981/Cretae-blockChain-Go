package main

import (
	"fmt"
	"go_blockchain/blockChain"
)

func main() {
	var chain = blockChain.InitBlockChain()
	chain.AddBlock("I want to say hello to all people")
	chain.AddBlock("I am new generation of AI")
	fmt.Printf("\n\n+++++++++++++++++++++++++++++++++++++++\n")
	for i := 0; i < len(chain.Blocks); i++ {
		fmt.Printf("pervious Hash:%x\n", chain.Blocks[i].PrevHash)
		fmt.Printf("Data in Block:%s\n", chain.Blocks[i].Data)
		fmt.Printf("Hash:%x\n", chain.Blocks[i].Hash)
		fmt.Printf("--------------------------------\n")
	}
}
