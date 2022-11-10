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
		Transfer{From: "here", To: "here", balances: " here"},
	}

	fmt.Println("Transfer: all the addresses")
	json.NewEncoder(w).Encode(transfers)
}

func showToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "  show my current balance")

}
func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Welcome home!")
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/transfers", transferTonkens).Methods("GET")
	router.HandleFunc("/balance", showToken)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}
