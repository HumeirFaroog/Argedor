package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://avalanche-fuji.infura.io/v3/1a41252c167b4d37ac22d3ba9d8b5aab")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("4d4ace62af99ecc8e8b4881d4cc5f8ffba9c4b449ea3d52e81d08a2058749641")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

        value := big.NewInt(10000000000000000) // in wei (0.01 eth)
	gasLimit := uint64(21000)               // in units
	tip := big.NewInt(2000000000)           // maxPriorityFeePerGas = 2 Gwei
	feeCap := big.NewInt(20000000000)       // maxFeePerGas = 20 Gwei
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x515f03aB6D698F634EDa46159f094954b41A8F7e")
	var data []byte

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: feeCap,
		GasTipCap: tip,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     value,
		Data:      data,
	})
        
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)

	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %s", signedTx.Hash().Hex())

}