package blockchain

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"time"
)

type BlockChain struct {
	lastHash Hash  //最后一个区块哈希
	//blocks map[Hash]*Block //全部区块信息，由区块哈希作为key检验
	db *leveldb.DB  //leveldb 的连接
}

//构造方法
func NewBlockChain(db *leveldb.DB) *BlockChain {
	//实例化 blockchain
	bc:= &BlockChain{
		db: db,
	}
	//初始化lastHash
	data, err := bc.db.Get([]byte("lastHash"),nil)
	if err==nil {
		bc.lastHash = Hash(data)
	}
	return bc
}
//添加创世区块
func (bc *BlockChain) AddGensisBlock() *BlockChain {
	//只有txs时特殊的
	if bc.lastHash!="" {
		//已经存在区块，不需要再添加创世区块
		return bc
	}
	return bc.AddBlock("The Gensis Block")
}

func (bc *BlockChain) AddBlock(txs string) *BlockChain {
	//构建区块
	b := NewBlock(bc.lastHash,txs)
	// 将区块加入到链的存储结构中
	if bs,err := BlockSerialize(*b);err!=nil{
		log.Fatal("block can not be serialize")
	}else if err = bc.db.Put([]byte("b_"+b.hashCurr),bs,nil);err!=nil{
		log.Fatal("block can not be saved")
	}

	//将最后的区块哈希设置为当前区块
	bc.lastHash = b.hashCurr
	//将最后的区块hash存到数据库中
	if err := bc.db.Put([]byte("lastHash"),[]byte(b.hashCurr),nil);err!=nil{
		log.Fatal("lastHash not be saved")
	}
	return bc
}

//通过hash获取区块
func (bc *BlockChain) GetBlock(hash Hash) (*Block,error) {
	data,err := bc.db.Get([]byte("b_"+hash),nil)
	if err != nil{
		return nil, err
	}
	//反序列化
	b,err := BlockUnSerialize(data)
	if err!=nil {
		return nil, err
	}
	return &b,err
}


//迭代展示区块的方法
func (bc *BlockChain) Iterate()  {
	//最后的哈希
	for hash := bc.lastHash;hash!="";{
		b,err := bc.GetBlock(hash)
		if err!=nil {
			log.Fatalf("Block <%s> is not exits",hash)
		}
		fmt.Println("HashCurr:",b.hashCurr)
		fmt.Println("Txs:",b.txs)
		fmt.Println("Time",b.header.time.Format(time.UnixDate))
		fmt.Println("HashPerv:",b.header.hashPrevBlock)
		fmt.Println()
		hash = b.header.hashPrevBlock
	}
}

//清空
func (bc *BlockChain) Clear()  {
	//数据库中全部的key全部删除
	bc.db.Delete([]byte("lastHash"),nil)
	iter := bc.db.NewIterator(util.BytesPrefix([]byte("b_")),nil)
	for iter.Next() {
		bc.db.Delete(iter.Key(),nil)
	}
	iter.Release()
	//清空bc对象
	bc.lastHash=""
}