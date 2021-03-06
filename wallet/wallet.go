package wallet

import (
	"crypto/sha256"
	"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ripemd160"
	"log"
)

const KeyBitSize = 256

type Address = string

type Wallet struct {
	//私钥为 *bip32.Key类型
	privateKey *bip32.Key
	//公钥由私钥计算推到，使用下面的掉用
	//publicKey ecdsa.PublicKey
	//助记词
	mnemonic string
	Address Address
}

//构造函数
func NewWallet(pass string) *Wallet {
	w := &Wallet{}
	//生成key
	w.GenKey(pass)
	//生成地址
	w.GenAddress()

	return w
}

//生成秘钥
func (w *Wallet) GenKey(pass string) *Wallet {

	//使用bip39
	//熵（随机）
	entropy,err := bip39.NewEntropy(KeyBitSize)
	if err!=nil {
		log.Fatal(err)
	}
	//助记词
	mnemonic,err := bip39.NewMnemonic(entropy)
	if err!=nil {
		log.Fatal(err)
	}
	w.mnemonic=mnemonic
	//key种子
	seed := bip39.NewSeed(mnemonic,pass)
	//生成秘钥
	privateKey,err := bip32.NewMasterKey(seed)
	w.privateKey = privateKey
	return w


	//elliptic.P256() 生成椭圆
	//rand.Reader，生成随机数
	//privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//if err != nil {
	//	panic(err)
	//}
	//w.privateKey = privateKey
	//return w
}

//生成Address
func (w *Wallet) GenAddress() *Wallet {
	//利用私钥形成公钥
	pubKey := w.privateKey.PublicKey().String()

	//hash pubkey
	//pubHash := ripemd160(sha256(pubkey))
	hashPubKey := HashPubKey([]byte(pubKey))

	//计算checkSum 校验值
	h1 := sha256.Sum256(hashPubKey)
	checkSum := sha256.Sum256(h1[:])

	//组合继续base64
	data := append(append([]byte{0},hashPubKey...,),checkSum[:4]...)
	w.Address = base58.Encode(data)
	return w
}

//生成当前公钥hashkey
func HashPubKey(pubKey []byte) []byte {
	shaHash := sha256.Sum256(pubKey)
	rpmd := ripemd160.New()
	rpmd.Write(shaHash[:])
	pubHash := rpmd.Sum(nil)
	return pubHash
}

//get Mnemonic
func (w *Wallet) GetMnemonic() string {
	return w.mnemonic
}