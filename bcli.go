package main

import (
	"flag"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/wei-rh/bc-demo/blockchain"
	"log"
	"os"
	"strings"
)

//命令行工具
func main() {
	//初始化数据库
	//数据库连接
	dbpath := "data"
	db,err := leveldb.OpenFile(dbpath,nil)
	if err!=nil {
		log.Fatal(err)
	}
	defer db.Close()
	//区块链测试
	bc := blockchain.NewBlockChain(db)
	//创建创世区块
	bc.AddGensisBlock()

	arg1 := ""
	if len(os.Args)>=2 {
		arg1 = os.Args[1]
	}
	switch strings.ToLower(arg1) {
	case "create:block":
		fs := flag.NewFlagSet("create:block",flag.ExitOnError)
		txs := fs.String("txs","","")
		fs.Parse(os.Args[2:])
		//完成区块的创建
		bc.AddBlock(*txs)
	case "show":
		bc.Iterate()
	case "init":
		//清空
		bc.Clear()
		bc.AddGensisBlock()
	case "help":
		fallthrough
	default:
		Usage()
	}
}

func Usage()  {
	fmt.Println("bcli is a tool for Blockchain. ")
	fmt.Println()
	fmt.Println("Usage: ")
	fmt.Printf("\t%s\t\t%s\n","bcli createblock <txs>","create block on blockchain")
	fmt.Printf("\t%s\t\t\t%s\n", "bcli init", "initial blockchain")
	fmt.Printf("\t%s \t\t\t%s\n", "bcli he1p","help info for bcli")
	fmt.Printf("\t%s \t\t\t%s\n", "bc1i show", "show blocks in chain.")

}