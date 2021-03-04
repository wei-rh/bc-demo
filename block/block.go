package block

import (
	"fmt"
	"time"
)

type Hash = string

const BlockBits = 16
const HashLen = 256
const nodeVersion = 0
//区块主体
type Block struct {
	header    BlockHeader
	txs       string //交易列表
	txCounter int    //交易计数器
	hashCurr  Hash   //当前区块哈市值缓存
}
type BlockHeader struct {
	version        int
	hashPrevBlock  Hash      //前一个区块的 Hash
	hashMerkleRoot Hash      //默克尔树的哈希节点
	time           time.Time //区块的创建时间
	bits           int       // 难度相关
	nonce          int       //挖矿相关
}

func NewBlock(prevHash Hash,txs string) *Block {
	//实例化Block
	b := &Block{
		header:      BlockHeader{
			version:       nodeVersion,
			hashPrevBlock: prevHash,  //设置前面的区块哈希
			time:          time.Now(),
			bits:          BlockBits,
		},
		txs:       txs,   //设置数据
		txCounter: 1,   //计算叫姨
	}
	//计算设置当前区块的哈希
	//b.SetHashCurr()
	return b
}

//func (bh *BlockHeader) Stringify() string {
//	return fmt.Sprintf("%d%s%s%d%d%d",
//		bh.version,
//		bh.hashPrevBlock,
//		bh.hashMerkleRoot,
//		bh.time.UnixNano(), //得到时间戳， nano 级别
//		bh.bits,
//		bh.nonce,
//		)
//}



//bits 属性的getter
func (b *Block) GetBits() int {
	return b.header.bits
}
//生成用于 pow 的字符串
func (b *Block) GenServiceStr() string {
	return fmt.Sprintf("%d%s%s%s%d",
		b.header.version,
		b.header.hashPrevBlock,
		b.header.hashMerkleRoot,
		b.header.time.Format("2016-01-02 15:04:05.999999999 -0700 MST"),
		b.header.bits,
	)

}

func (b *Block) SetNonce(nonce int) *Block {
	b.header.nonce=nonce
	return b
}

//设置当前区块 hash
func (b *Block)SetHashCurr(hash Hash) *Block {
	//计算 hash 值
	b.hashCurr =  hash
	return b
}

func (b *Block) GetHashCurr() Hash {
	return b.hashCurr
}
func (b *Block) GetTxs() string {
	return b.txs
}
func (b *Block) GetTime() time.Time {
	return b.header.time
}
func (b *Block) GetHashPrevBlock() Hash {
	return b.header.hashPrevBlock
}
func (b *Block) GetNonce() int {
	return b.header.nonce
}
