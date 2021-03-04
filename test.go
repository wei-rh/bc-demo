package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	//b := blockchain.NewBlock("","Gensis Block.")
	//fmt.Println(b)
	//fmt.Println(b)

	//dbpath := "data"
	//db,err := leveldb.OpenFile(dbpath,nil)
	//if err!=nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	//区块链测试
	//bc := blockchain.NewBlockChain(db)
	////创建创世区块
	//bc.AddGensisBlock()
	////添加新区块
	////bc.AddBlock("first block").
	////	AddBlock("second block")
	//bc.Iterate()

	bits := 24
	target := big.NewInt(1)
	//前8为0
	target.Lsh(target,uint(256-bits+1))
	fmt.Println(target.String())
	fmt.Println("------------------------------")
	nonce := 0
	serviceStr := "block data"
	var hashInt big.Int
	for  {
		data := serviceStr + strconv.Itoa(nonce)
		hash := sha256.Sum256([]byte(data))
		hashInt.SetBytes(hash[:])
		//fmt.Printf("%x\n",hash)
		fmt.Println(hashInt.String(),nonce)
		if hashInt.Cmp(target)==-1 {
			fmt.Println("本机挖矿成功")
			return
		}
		nonce++
	}


}