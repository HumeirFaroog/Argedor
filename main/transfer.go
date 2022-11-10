package main

import (
	"github.com/ethereum/go-ethereum"
	"fmt"
	"golang.org/x/crypto/sha3"
	"math/big"
	"context"
	"crypto/ecdsa"
	"log"
     
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	//  my transation address 
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

	value := big.NewInt(0)
	getPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}


     // my account address
	toAddress := common.HexToAddress("0xF66f9beF2686AD8f6766e166452271d0d2744c9A")
	// my smart contract address
	tokenAddress := common.HexToAddress("0x515f03aB6D698F634EDa46159f094954b41A8F7e")

	// using the signature function for transfer 

	transferFnSignature := []byte("transfer(address, uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Printf("Method ID: %s\n", hexutil.Encode(methodID))

	paddedtrAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Printf("To address: %s\n", hexutil.Encode(paddedAddress))

	//sendinf amount of token 

	balance :=(big.Int)
	balance.setString("1000000000000000000000", 10)
	paddedBalance := common.LeftPadBytes(balance.Bytes(), 32)
	fmt.Printf("Token balance: %s", hexutil.Encode(paddedBalance))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedBalance...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To: &toAddress,
		Data: data
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Gas limit: %d", gasLimit)

	tx := types.NewTransation(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.signedTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransation(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tokens sent at TX: %s", signedTx.Hash().Hex())


}
