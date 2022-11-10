package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var avalancheURL = "https://avalanche-fuji.infura.io/v3/1a41252c167b4d37ac22d3ba9d8b5aab"
var ganaheURL = "http://127.0.0.1:8545"

func main() {
	// client, err := ethclient.DiaContex(context.Background(), avalancheURL)
	// if err != nil {
	// 	log.Fatalf("Erro to  create token :%v", err)
	// }
	// defer client.close()

	// block, err := client.blockByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatalf("Erro to  find  token :%v", err)
	// }

	// fmt.PrintLn("The Block Number", block.Number())

	addr := "0xf66f9bef2686ad8f6766e166452271d0d2744c9a"
	address := common.HexToAddress(addr)

	balance, err := BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Erro to  find  token :%v", err)
	}
	fmt.Println("The balance", balance)

}
