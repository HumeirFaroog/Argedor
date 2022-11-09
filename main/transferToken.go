package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	transfer, err := http.Get("https://avalanche-fuji.infura.io/v3/1a41252c167b4d37ac22d3ba9d8b5aab")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	transferData, err := ioutil.ReadAll(transfer.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(transferData))

}
