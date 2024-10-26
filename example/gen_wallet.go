package main

import (
	"fmt"
	"github.com/cuckoo-wallet/hdwallet"
)

func main() {

	mnemonic, _ := hdwallet.NewMnemonic(12, "")
	fmt.Println("mnemonic: ", mnemonic)
	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("助记词：", mnemonic)
	fmt.Println()
	wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.ETH))
	address, _ := wallet.GetAddress()
	fmt.Println("ETH私钥：", wallet.GetKey().PrivateHex())
	fmt.Println("ETH: ", address)

}
