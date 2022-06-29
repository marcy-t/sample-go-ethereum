package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
	AddressCheck
	アドレスを検証し、検証するアドレスがスマートコントラクトの
	アドレスであるかどうか判断する内容
*/
func main() {
	/*
		Check if Address is Valid
		簡単な正規表現を使って、Ethereumのアドレスが有効かどうかチェック
	*/
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("%v \n", "------------------------------------------")
	// is valid: true
	fmt.Printf("is valid 0x323b5d4c32345ced77393b3530b1eed0f346429d : %v \n",
		re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d"))
	// is valid: false
	fmt.Printf("is valid 0xZYXb5d4c32345ced77393b3530b1eed0f346429d : %v \n",
		re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d"))

	/*
		Check if Address is an Account or a Smart Contract
		・アドレスがアカウントorスマートコントラクトか確認する
		・アドレスにバイトコードが保存されていれば、そのアドレスがスマートコントラクトであるかどうか判断する
	*/
	// トークンのスマートコントラクトコード取得し、長さをチェック
	// スマートコントラクトであるか確認
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498") // true
	// address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498aa") // false

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	// bytecode, err := ethclient.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	fmt.Printf("%v \n", "------------------------------------------")
	// is valid: true
	fmt.Printf("is contract : %v \n", isContract)
}
