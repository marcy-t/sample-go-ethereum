package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

/*
	QueryingBlocks
	ブロック情報を照会する方法
	2通り
*/
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	// client, err := ethclient.Dial("https://cloudflare-eth.com") // 今現時点でアクセスエラー
	client, err := ethclient.Dial(os.Getenv("HOST"))
	if err != nil {
		log.Fatal(err)
	}
	/*
		Block Header
		クライアントのHeaderByNumberを呼び出すと、ブロックヘッダー情報を返すことができる
		nilを渡すと最新のブロックヘッダーが返される
	*/
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("header : %v \n", header.Number.String())
	// header : 7174802

	/*
		Full block
		ブロック番号を指摘し取得
		クライアントのBlockByNumberメソッドを呼び出して以下を取得
		- ブロック番号
		- ブロックタイムスタンプ
		- ブロックハッシュ
		- ブロック難易度
		- ブロックの全てのコンテンツとメタデータ
		- トランザクションリスト
	*/
	currentBlockHeight := header.Number.Int64()
	blockNumber := big.NewInt(currentBlockHeight)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("block.Number : %v \n", block.Number().Uint64()) // block.Number : 7174942
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("block.Time : %v \n", block.Time()) // block.Time : 1657042552
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("block.Difficulty : %v \n", block.Difficulty().Uint64()) // block.Difficulty : 2
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("block.Hash : %v \n", block.Hash().Hex()) // block.Hash : 0x31991f9a0909ff9b19002076bc86584139430a2dbd9aba7fa8377f316e94be89
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("block.Transactions list : %v \n", block.Transactions()) // block.Transactions list : [0x1400042e060 0x1400042e120 0x1400042e1e0 0x1400042e2a0 0x1400042e300 0x1400042e360 0x1400042e3c0 0x1400042e420 0x1400042e480 0x1400042e4e0 0x1400042e540 0x1400042e5a0 0x1400042e600 0x1400042e660 0x1400042e6c0 0x1400042e720 0x1400042e780 0x1400042e7e0]
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("block.Transactions len : %v \n", len(block.Transactions())) // block.Transactions len : 18

	/*
		TransactionCountを呼び出すと、ブロック内のトランザクションの数だけ返す
	*/
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("count : %v \n", count) // count : 18

}
