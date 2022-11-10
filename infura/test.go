package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"context"
    "crypto/ecdsa"
    "math/big"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/sha3"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

type Transfers []Transfer
type Transfer struct {
	From     string `json:oxooo`
	To       string `json:oxooooo`
	balances string `json:tokens`
}

func transferTonkens(w http.ResponseWriter, r *http.Request) {
	transfers := Transfers{
		Transfer{From: "", To: "", balances: ""},
	}

	fmt.Println("Transfer: all the addresses")
	json.NewEncoder(w).Encode(transfers)
}

func showBalance(w, http.ResponseWriter, r *http.Request) {

	account := common.HexToAddress("0xF66f9beF2686AD8f6766e166452271d0d2744c9A")
    balance, err := client.BalanceAt(context.Background(), account, nil)
     if err != nil {
    log.Fatal(err)

   }

      fmt.Println(balance) //  show in console 

}

func main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Welcome home!")
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/transfers", transferTonkens).Methods("POST")
	router.HandleFunc("/balance", showBalance).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}
