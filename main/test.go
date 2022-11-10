package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

// func showBalance(w, http.ResponseWriter, r *http.Request) {
// 	address := mux.Vars(r)["0xF66f9beF2686AD8f6766e166452271d0d2744c9"]

// }

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Welcome home!")
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/transfers", transferTonkens).Methods("POST")
	// router.HandleFunc("/balance", showBalance).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}
