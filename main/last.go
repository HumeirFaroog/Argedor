package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"


)

func	main() {
	client,  err := ethclient.Dial("https://avalanche-fuji.infura.io/v3/1a41252c167b4d37ac22d3ba9d8b5aab")

	if err != nil {
		log.Fatal("Unable to  reach client")
	}

	privateKey, err := crypto.HexToECDSA("4d4ace62af99ecc8e8b4881d4cc5f8ffba9c4b449ea3d52e81d08a2058749641")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key  to ECDSA")
	}

	fromAddress := crypto.PublicToAddress(*publicKeyECDSA)
	nonce, err := client.PendingBalanceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("nonce", nonce)

	toAddress := common.HexToAddress("0xF66f9beF2686AD8f6766e166452271d0d2744c9A")
    
	value := big.NewInt(1000)

	gasFeeCap, gasTipCap, gas := big.NewInt(0.0000000525 ), big.NewInt(0.000000025),uint64(790,662)

	var data []byte

	tx := types.NewTx(&types,DynamicFeeTx){
		Nonce: nonce,
		GasFeeCap :gasFeeCap,
		GasTipCap: gasTipCap,
		Gas: gas,
		To: &toAddress,
		Value: value,
		Data: data})

	config, block := params.ChainConfig, params.ChainConfig.LondonBlock

	signer := types.MakeSigner(config, privateKey)

	if err != nil {
		log.Fatal(" Unable to sign Tx")
	}

	hash := singedTx.Hash().Bytes()

	raw, err := rlp.EncodeToBytes(singedTx)

	if err != nil {
		log.Fatal("Unable to cast to raw txn")
	}

	fmt.Printf("hash:%x\n0x0x", hash, raw)

	err = client.SendTransaction(context.Background(), singedTx)

	if err != nil {
		log.Fatal("Unable to sumit transation")
	}

	fmt.Printf("tx sed: %s", singedTx.hash().hex())

	
		
}
	

