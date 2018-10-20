package main

import (
	"fmt"
    "net/http"
    "encoding/json"
)

type NewBlockRequest struct {
    Data string `json:"data"`
}

type NewTransactionRequest struct {
    From string `json:"from"`
    To string `json:"to"`
    Amount int `json:"amount"`
}


func index(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "Blockchain test") // send data to client side
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {

    resp, _ := json.Marshal(Blockchain)

    fmt.Fprintf(w, string(resp))
}

func postBlock(w http.ResponseWriter, r *http.Request) {

    decoder := json.NewDecoder(r.Body)

    var body NewBlockRequest
    err := decoder.Decode(&body)

	if err != nil {
		panic(err)
	}

	//fmt.Println(body.Data)

    newBlock := mineBlock(body.Data)
    if addBlockToChain(newBlock) {

        resp, _ := json.Marshal(newBlock)

        fmt.Fprintf(w, string(resp)) // send data to client side
    } else {
        fmt.Fprintf(w, "Invalid block")
    }
}

func newTransaction(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var body NewTransactionRequest
    err := decoder.Decode(&body)

	if err != nil {
		panic(err)
    }
    

}

func node() {

     // Set routes
    http.HandleFunc("/", index)
    http.HandleFunc("/blockchain", getBlockchain)
    http.HandleFunc("/mineblock", postBlock)
    http.HandleFunc("/newTransaction", newTransaction)

    // Start server
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        fmt.Println(err)
    }
}