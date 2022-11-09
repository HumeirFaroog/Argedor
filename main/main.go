package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var avalancheURL = "https://avalanche-mainnet.infura.io/v3/1a41252c167b4d37ac22d3ba9d8b5aab"

func main() {
	client, err := ethclient.DiaContex(context.Background(), avalancheURL)
	if err != nil {
		log.Fatalf("Erro to  create token :%v", err)
	}
	defer client.close()

	block, err := client.blockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Erro to  find  token :%v", err)
	}

	fmt.PrintLn(block.Number())

}
