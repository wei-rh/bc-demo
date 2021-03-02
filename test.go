package main

import (
	"fmt"
	"github.com/wei-rh/bc-demo/blockchain"
)

func main() {
	b := blockchain.NewBlock("","Gensis Block.")
	fmt.Println(b)
	fmt.Println(b)
}