package main

import (
	//"crypto/sha256"
	//"fmt"
	//"math/big"
	//"strconv"
	"fmt"
	//"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip32"
)

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)

	//反推
	userMnemonic := mnemonic
	//生成殇
	//userEntropy,_ := bip39.EntropyFromMnemonic(userMnemonic)
	//生成种子
	userSeed := bip39.NewSeed(userMnemonic,"Secret Passphrase")
	usermasterKey, _ := bip32.NewMasterKey(userSeed)
	userpublicKey := usermasterKey.PublicKey()
	fmt.Println("Mnemonic: ", userMnemonic)
	fmt.Println("Master private key: ", usermasterKey)
	fmt.Println("Master public key: ", userpublicKey)



}
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

	//bits := 24
	//target := big.NewInt(1)
	////前8为0
	//target.Lsh(target,uint(256-bits+1))
	//fmt.Println(target.String())
	//fmt.Println("------------------------------")
	//nonce := 0
	//serviceStr := "block data"
	//var hashInt big.Int
	//for  {
	//	data := serviceStr + strconv.Itoa(nonce)
	//	hash := sha256.Sum256([]byte(data))
	//	hashInt.SetBytes(hash[:])
	//	//fmt.Printf("%x\n",hash)
	//	fmt.Println(hashInt.String(),nonce)
	//	if hashInt.Cmp(target)==-1 {
	//		fmt.Println("本机挖矿成功")
	//		return
	//	}
	//	nonce++
	//}
	//encoded := "1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq"
	//num, err := base58.Decode(encoded)
	//if err != nil {
	//	fmt.Printf("Demo %v, got error %s\n", encoded, err)
	//}
	//chk := base58.Encode(num)
	//if encoded == string(chk) {
	//	fmt.Printf ( "Successfully decoded then re-encoded %s\n", encoded )
	//}

//}