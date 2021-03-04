package pow

import (
	"crypto/sha256"
	"fmt"
	"github.com/wei-rh/bc-demo/block"
	"math"
	"math/big"
	"strconv"
)
type ProofOfWord struct {
	//需要pow工作量区块的区块
	block *block.Block
	//证明参数列表
	target *big.Int
}

//构造方法
func NewPow(b *block.Block) *ProofOfWord {
	p := &ProofOfWord{
		block:  b,
		target: big.NewInt(1),
	}
	//计算 target值
	p.target.Lsh(p.target,uint(block.HashLen-b.GetBits()-1))
	return p
}

//hashcash 证明
// 返回使用的 nonce 和 形成的区块hash
func (p *ProofOfWord) Proof() (int, block.Hash) {

	var hashInt big.Int
	// 基于block 准备 serviceStr
	serviceStr := p.block.GenServiceStr()
	//nonce 计数器
	nonce := 0
	//迭代计算hash，设置防nonce溢出的条件
	fmt.Printf("Target:%d\n",p.target)
	for nonce<math.MaxInt64 {
		//生成 hash
		hash := sha256.Sum256([]byte(serviceStr + strconv.Itoa(nonce)))
		//得到 hash 的big.iny
		hashInt.SetBytes(hash[:])
		fmt.Printf("Hash :%s\t%d\n",hashInt.String(),nonce)
		//判断是否满意难度（数学难题）
		if hashInt.Cmp(p.target)==-1 {
			return nonce, block.Hash(fmt.Sprintf("%x",hash))
		}
		nonce++
	}
	return 0,""
}

//pow验证
func (p *ProofOfWord) Validate() bool {

	serviceStr := p.block.GenServiceStr()
	date := serviceStr + strconv.Itoa(p.block.GetNonce())
	hash := sha256.Sum256([]byte(date))
	if p.block.GetHashCurr() != fmt.Sprintf("%x",hash){
		return false
	}
	//
	target := big.NewInt(1)
	target.Lsh(target,uint(block.HashLen-p.block.GetBits()-1))
	hashInt := new(big.Int)
	hashInt.SetBytes(hash[:])
	if hashInt.Cmp(target)!=-1 {
		return false
	}
	return true
}