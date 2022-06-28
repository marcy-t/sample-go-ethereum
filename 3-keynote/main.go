package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
	/*
		・KeyStoreは、暗号化されたウォレットプライベートキーを含むファイル
		・「go-ethereum」のキーストアには、1ファイルにつき1つのウォレットキーペアしか格納できない
		・キーストアの生成はNewKeyStoreを呼び出し、キーストアを保存するパスを指定
		・その後、NewAccountを呼び出し暗号化用のパスワードを渡すことで新しいウォレットを生成できます
		・NewAccountを呼び出すたびに、新しいキーストアファイルがディスク上に生成されます
	*/
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("account : %v \n", account)
	/*
		account : {0xdb510F95Dc75b292E7E3F12B57dF3840f6e6A358 keystore:///Users/makito/B-LAND/sample-go-ethereum/3-keynote/tmp/UTC--2022-06-28T12-06-12.353488000Z--db510f95dc75b292e7e3f12b57df3840f6e6a358}
	*/
}

func importKs() {
	file := "./tmp/UTC--2022-06-28T12-44-14.249656000Z--bee2c6283c97a486c899562947da180d053e4086"

	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("keystore.StandardScryptN : %v \n", keystore.StandardScryptN)

	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("keystore.StandardScryptP : %v \n", keystore.StandardScryptP)

	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("####")
		log.Println("err : ", err)
		log.Println("####")
		log.Fatal(err)
	}

	/*
		・鍵をインポートするには、基本的に通常通りNewKeyStoreを再起動し、
		鍵ストアのJsonデータをバイトとして受け取るImportメソッドを呼び出す必要がある
	*/
	password := "secret"
	/*
		・第二引数は暗号化したパスワードで、これを複合化するために使用
		・第三者引数は新しい暗号化用パスワードを指定するが、例では同じものを使用
		・以下はキーストアをインポートしてアカウントにアクセス
		・鍵ストアをインポートしてアカウントにアクセスする例
	*/
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		// could not decrypt key with given password
		log.Fatal(err)
	}

	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("importKs Address : %v \n", account.Address.Hex())

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	/*

	 */
	// createKs()
	importKs()
}
