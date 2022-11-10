package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Transfer struct {
	From     string `oxooo`
	To       string `oxooooo`
	balances string `tokens`
}

type Transfers []Transfer

func transferTonkens(w http.ResponseWriter, r *http.Request) {
	transfers := Transfers{
		Transfer{From: " sender address", To: "reciever  wallet", balances: "myCurrentBalance"},
	}

	fmt.Println("Transfer: all the addresses")
	json.NewEncoder(w).Encode(transfers)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here we'll show the current balances")
}

func handleRequests() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/transfers", transferTonkens)
	// http.HandleFunc("/balance", transferTonkens)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
