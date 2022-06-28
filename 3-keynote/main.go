package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/k0kubun/pp/v3"
)

func createKs() accounts.Account {
	/*
		・KeyStoreは、暗号化されたウォレットプライベートキーを含むファイル
		・「go-ethereum」のキーストアには、1ファイルにつき1つのウォレットキーペアしか格納できない
		・キーストアの生成はNewKeyStoreを呼び出し、キーストアを保存するパスを指定
		・その後、NewAccountを呼び出し暗号化用のパスワードを渡すことで新しいウォレットを生成できます
		・NewAccountを呼び出すたびに、新しいキーストアファイルがディスク上に生成されます
	*/
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	log.Println("########")
	log.Println("ks : ", ks)
	log.Println("########")

	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("account Address: %v \n", account.Address.Hex())
	fmt.Printf("account : %v \n", account.URL)
	/*
		account : {0xdb510F95Dc75b292E7E3F12B57dF3840f6e6A358 keystore:///Users/makito/B-LAND/sample-go-ethereum/3-keynote/tmp/UTC--2022-06-28T12-06-12.353488000Z--db510f95dc75b292e7e3f12b57df3840f6e6a358}
	*/
	return account
}

func importKs() {
	file := "./wallets/UTC--2022-06-28T18-07-03.968841000Z--497ce35065662c265311e48248af55b939e069cb"

	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("keystore.StandardScryptN : %v \n", keystore.StandardScryptN)

	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("keystore.StandardScryptP : %v \n", keystore.StandardScryptP)

	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	/*
		・鍵をインポートするには、基本的に通常通りNewKeyStoreを再起動し、
		鍵ストアのJsonデータをバイトとして受け取るImportメソッドを呼び出す必要がある
	*/
	passPhrase := "secret"
	// newPassphrase := "hoge"
	/*
		・第二引数は暗号化したパスワードで、これを複合化するために使用
		・第三者引数は新しい暗号化用パスワードを指定するが、例では同じものを使用
		・以下はキーストアをインポートしてアカウントにアクセス
		・鍵ストアをインポートしてアカウントにアクセスする例
	*/
	/*
		https://pkg.go.dev/github.com/0xsequence/ethkit/go-ethereum/accounts/keystore#KeyStore.Import
		func (ks *KeyStore) Import(keyJSON []byte, passphrase, newPassphrase string) (accounts.Account, error)
	*/
	account, err := ks.Import(jsonBytes, passPhrase, passPhrase)
	if err != nil {
		// could not decrypt key with given password
		// account already exists
		log.Fatal(err)
	}

	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("account : %v \n", account.Address.Hex())

	// if err := os.Remove(file); err != nil {
	// 	log.Fatal(err)
	// }
}

func updateKs() {
	file := "./wallets/UTC--2022-06-28T16-43-32.884125000Z--77e080a900be5f0f5b9f76c87e80394382117f20"

	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	account, err := ks.Import(jsonBytes, "secret", "secret")
	if err != nil {
		// could not decrypt key with given password
		// account already exists
		pp.Print(err)
		log.Fatal(err)
	}

	passPhrase := "secret"
	_ = ks.Update(account, passPhrase, passPhrase)

	fmt.Printf("%v \n", "------------------------------------------")
	fmt.Printf("importKs Address : %v \n", account.Address.Hex())

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	/*

	 */
	// account := createKs()
	// fmt.Printf("%v \n", "------------------------------------------")
	// fmt.Printf("account createKs() :  %v \n", account)
	importKs()
	// updateKs()
}
