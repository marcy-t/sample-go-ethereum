package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	/*
		秘密鍵の生成
		go-ethereumでアカウントアドレスを使用するには、
		common.Addressタイプに変換する必要がある
	*/
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("crypto.GenerateKey : %v \n", privateKey)
	/*
		crypto.GenerateKey : &{{0x140001242d0 85531770078635377878588618836923944938376432816497610975064824155376987625769 73381096783975819817347464479463283975019231468181598306572331410529970762723} 108599437771372318051953569921321951891799511147820265422186698700866808497918}
	*/

	/*
		golang crypto/ecdsaパッケージを
		インポートしてFromECDSAメソッドを使うことで、バイトに変換
	*/
	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("privateKeyBytes : %v \n", privateKeyBytes)
	/*
		privateKeyBytes : [225 4 16 6 17 249 172 133 49 68 138 207 240 88 49 94 167 118 85 244 216 196 88 162 210 251 128 98 141 101 92 90]
	*/

	/*
		byte sliceを受けと るEncodeメソッドを提供する、go-ethereum hexutilパッケージを使用して、
		これを16進文字列に変換することができる
		そして16進数エンコードされた後に「0x」を削除します
	*/
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("hexutil encode : %v \n", hexutil.Encode(privateKeyBytes)[2:])

	/*
		例）同じものは生成されない
		hexutil encode : f541403d08cc26a9a3cf9a0d6d4a71a82ac13e12b2268c7c58f3c0300712bfb1
		これは署名に使用される秘密鍵で、パスワードのように扱われ、決して共有はしてはならない
	*/

	/*
		公開鍵
		公開鍵は秘密鍵から派生する
		公開鍵を返すPUblicメソッドがある
	*/
	publicKey := privateKey.Public()
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("publicKey : %v \n", publicKey)
	/*
		例）
		publicKey : &{0x1400009c2d0 52102182653970875267639900307390818886038331764372604145470117703195129749270 27736048816714796860018845946691315842503537206458021737106851536258101483271}
	*/

	/*
		16進数への変換は、秘密鍵の場合と同様の手順で行う
		0xと最初の2文字04は常に接頭辞であり不要なので削除
		publicKeyにアサーションしてチェック（*ecdsa.PublicKey）
	*/
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("publicKeyBytes : %v \n", hexutil.Encode(publicKeyBytes))
	fmt.Printf("publicKeyBytes : %v \n", hexutil.Encode(publicKeyBytes)[4:])
	/*
		publicKeyBytes : 0x0400b9527e167c488b3f5c86df66ef93ebe64c1c24f087437d7089c9dfc8470a72e1069fd5d92608b7611823b957f21ab414bdba58a0cbd44838af83a99f4107dd
		publicKeyBytes : 00b9527e167c488b3f5c86df66ef93ebe64c1c24f087437d7089c9dfc8470a72e1069fd5d92608b7611823b957f21ab414bdba58a0cbd44838af83a99f4107dd
	*/

	/*
		上記で公開鍵が手に入るので、見慣れた公開アドレスを生成できる
		「go-ethereum crypto」パッケージには、PubKeyToAddressメソッドがあり
		ECDSA公開鍵を受け取って公開鍵アドレスを返す
	*/
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("address : %v \n", address)
	/*
		address : 0xb7192a0229E46Ad4Af30ADB59080E08cA21c6A0D
	*/

	/*
		公開アドレスは単純に公開鍵の「Keccak-256」で、最初の40文字（20バイト）をとって
		0xを前置しています
		crypto/sha3 keccak256関数を使って手動で行う方法は以下
	*/
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("hash : %v \n", hash)
	fmt.Printf("hash hexutil.Encode : %v \n", hexutil.Encode(hash.Sum(nil)[12:]))
}
