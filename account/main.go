package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	/*
		go-ethereumでアカウントアドレスを使用するには、
		common.Addressタイプに変換する必要がある
	*/
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	/*
		go-ethereumのメソッドにethereum addressを渡す際はほとんどがこれになる
	*/
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)

	fmt.Printf("Balance : %v \n", balance) // 32625327174001387699

	/*
		ブロックの時点の口座残高を読み取り
		client.BalanceAtのBlockNumberにnilを入れると入れると最新の残高が返される
	*/
	// blockNumber := big.NewInt(5532993) // ETHScanでブロックとか確認
	// balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	balanceAt, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("BalanceAt : %v \n", balanceAt) // 32625327174001387699

	/*
		・Ethereumの数値は固定小数点精度であるため、
		可能な限りは小さな単位で扱われ、ETHの場合は
		「wei」となります
		・ETHの値を読み取るには、wei/10^18の計算を行う必要がある
		・大きい数値はgoのmathとmath.bigを使用
	*/
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	/*
		Balance  : 32625327174001387699
		ethValue : 32.6253271740013877
	*/
	fmt.Printf("ethValue : %v \n", ethValue)

	/*
		Pending balance
		取引送信した後や確認待ちの時に、保留中の口座残高を知りたいことがある
		過程を取得する
	*/
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}

	// pendingBalance : 32625327174001387699
	fmt.Printf("pendingBalance : %v \n", pendingBalance)
}
