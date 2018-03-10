package main

import (
	"fmt"
	//"math/big"

	"github.com/onrik/ethrpc"
)

func main() {
	fmt.Println("Start application")
	defer fmt.Println("Stop application")

	tun, err := NewTuner("./config.toml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connect address: ", tun.Ethereum.Address)

	client := ethrpc.New(tun.Ethereum.Address)
	if client == nil {
		fmt.Println("Not create new client")
		return
	}
	fmt.Println("New client created")

	// fmt.Println(client.NetVersion())

	version, err := client.Web3ClientVersion()
	if err != nil {
		fmt.Println("Error get version from Geth:", err)
		return
	}
	fmt.Printf("Geth version: %s\n", version)

	// Send eth
	txid, err := client.EthSendTransaction(ethrpc.T{
		From:  tun.Test.From,
		To:    tun.Test.To,
		Value: client.EthSum(tun.Test.Amount),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Transactions ID: %s\n", txid)
}
