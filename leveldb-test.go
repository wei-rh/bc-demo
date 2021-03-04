package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	//打开
	dbpath := "testdb"
	db,err := leveldb.OpenFile(dbpath,nil)
	if err!=nil {
		log.Fatal(err)
	}
	//读取key
	key := "joke Han"

	//if err := db.Put([]byte(key),[]byte("BlockChain demo"),nil);err!=nil {
	//	log.Fatal(err)
	//}
	//log.Fatal("put success")
	data,err := db.Get([]byte(key),nil)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
