package main

import (
    "fmt"
    "bytes"
    "net/http"
    "encoding/json"
    "os"
    "./utils"
)

type NewBlockRequest struct {
    Data string `json:"data"`
}

type NewTransactionRequest struct {
    From string `json:"from"`
    To string `json:"to"`
    Amount int `json:"amount"`
}

type Peer struct {
    Address string `json:"address"`
    Port string `json:"port"`
}

type BlockBroadcast struct {
    Block Block `json:"block"`
    Peer Peer `json:"peer"`
}

type TransactionBroadcast struct {
    Transaction Transaction `json:"transaction"`
    Peer Peer `json:"peer"`
}

func index(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "Blockchain test") // send data to client side
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {

    resp, _ := json.Marshal(Blockchain)

    fmt.Fprintf(w, string(resp))
}

func mineBlockRequest(w http.ResponseWriter, r *http.Request) {

    decoder := json.NewDecoder(r.Body)

    var body NewBlockRequest
    err := decoder.Decode(&body)

	if err != nil {
		panic(err)
	}

	//fmt.Println(body.Data)

    newBlock := mineBlock(body.Data)
    if addBlockToChain(newBlock) {

        broadcastNewBlock(newBlock)

        resp, _ := json.Marshal(newBlock)

        fmt.Fprintf(w, string(resp)) // send data to client side
    } else {
        fmt.Fprintf(w, "Invalid block")
    }
}

func newTransaction(w http.ResponseWriter, r *http.Request) {

    if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

        fmt.Fprintf(w, "OK")
    } else {
        decoder := json.NewDecoder(r.Body)

        var body NewTransactionRequest
        err := decoder.Decode(&body)

        if err != nil {
            fmt.Println(err)
            panic(err)
        }
        
        resp, errText, transaction := createNewTransaction(body.To, body.From, body.Amount)
        if resp {
            broadcastNewTransaction(transaction)
            fmt.Fprintf(w, "Added to transaction pool.")
        } else {
            fmt.Fprintf(w, errText)
        }
    }
}

func getTransactionPool(w http.ResponseWriter, r *http.Request) {
    resp, _ := json.Marshal(PendingTransactions)

    fmt.Fprintf(w, string(resp))
}

func getBalances(w http.ResponseWriter, r *http.Request) {
    balances := make(map[string]int)

    for _, txOut := range UnspentTransactionsOut {
        if _, keyExists := balances[txOut.ToAddress]; keyExists == true {
            balances[txOut.ToAddress] += txOut.Amount
        } else {
            balances[txOut.ToAddress] = txOut.Amount
        }
    }

    resp, _ := json.Marshal(balances)

    fmt.Fprintf(w, string(resp))
}

func getPeers(w http.ResponseWriter, r *http.Request) {
    resp, _ := json.Marshal(Peers)

    fmt.Fprintf(w, string(resp))
}

func addPeer(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var newPeer Peer
    err := decoder.Decode(&newPeer)

	if err != nil {
        fmt.Println(err)
		panic(err)
    }
    
    Peers = append(Peers, newPeer)

    queryBlockchainFromPeer(newPeer)

    resp, _ := json.Marshal(Peers)

    fmt.Fprintf(w, string(resp))
}

func getOwnAddress(w http.ResponseWriter, r *http.Request) {
    type AddressRequest struct {
        Address string `json:"address"`
    }

    resp, _ := json.Marshal(AddressRequest{Address: string(utils.PublicKeyToBytes(PublicKey))})

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Content-Type", "application/json")

    fmt.Fprintf(w, string(resp))
}

var Peers []Peer
var ThisPeer Peer

func node() {

    Peers = []Peer{}

     // Set routes
    http.HandleFunc("/", index)
    http.HandleFunc("/blockchain", getBlockchain)
    http.HandleFunc("/mineblock", mineBlockRequest)
    http.HandleFunc("/newTransaction", newTransaction)
    http.HandleFunc("/transactions", getTransactionPool)
    http.HandleFunc("/balances", getBalances)

    http.HandleFunc("/peers", getPeers)

    http.HandleFunc("/peer", addPeer)
    http.HandleFunc("/peer/block", newBlockFromPeer)
    http.HandleFunc("/peer/transaction", newTransactionFromPeer)

    http.HandleFunc("/utils/getOwnAddress", getOwnAddress)

    port := "9090"
    if len(os.Args) > 1 {
        port = os.Args[1]
    }

    ThisPeer = Peer{Address: "http://localhost", Port: port}

    // Start server
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        fmt.Println(err)
    }
}


// Broadcast
func broadcast(data interface{}, endpoint string) {
    for _, peer := range Peers {
        jsonPeer, _ := json.Marshal(data)
        url := peer.Address + ":" + peer.Port + endpoint
        req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPeer))
        req.Header.Set("Content-Type", "application/json")
        
        fmt.Println("Broadcasted ", endpoint , " to ", peer)

        client := &http.Client{}
        _, err := client.Do(req)
        if err != nil {
            fmt.Println(err)
        }
    }
}

func broadcastNewBlock(block Block) {
    data := BlockBroadcast{
        Block: block, 
        Peer: ThisPeer,
    }
    broadcast(data, "/peer/block")
}

func broadcastNewTransaction(transaction Transaction) {
    data := TransactionBroadcast{
        Transaction: transaction, 
        Peer: ThisPeer,
    }
    broadcast(data, "/peer/transaction")
}

func newBlockFromPeer(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var newBlock BlockBroadcast
    err := decoder.Decode(&newBlock)

    if err != nil {
		panic(err)
    }

    added := addBlockToChain(newBlock.Block)
    if !added {
        queryBlockchainFromPeer(newBlock.Peer)
    }

    fmt.Fprintf(w, "ok")
}

func queryBlockchainFromPeer(peer Peer) {

    data := getBlockchainFromPeer(peer)

    peer_difficulty := CalculateCumulativeDifficulty(data)
    current_difficulty := CalculateCumulativeDifficulty(Blockchain)

    if peer_difficulty > current_difficulty {
        Blockchain = data

        UnspentTransactionsOut = []TransactionOut{}
        for _, block := range Blockchain {
            updateTransactions(block)
        }
    }
}

func getBlockchainFromPeer(peer Peer) []Block {
    fmt.Println("Trying to fetch blockchain from peer: ", peer)
    peerUrl := peer.Address + ":" + peer.Port + "/blockchain"

    client := &http.Client{}
    req, _ := http.NewRequest("GET", peerUrl, nil)
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)

    if err != nil {
        fmt.Println(err)
        panic(err)
    }

    defer resp.Body.Close()

    decoder := json.NewDecoder(resp.Body)

    var data []Block
    decode_err := decoder.Decode(&data)

    if decode_err != nil {
        fmt.Println(decode_err)
        panic(decode_err)
    }

    return data
}

func newTransactionFromPeer(w http.ResponseWriter, r *http.Request) {

    decoder := json.NewDecoder(r.Body)

    var data TransactionBroadcast
    err := decoder.Decode(&data)

    if err != nil {
		panic(err)
    }

    if ValidateTransaction(data.Transaction) {
        PendingTransactions = append(PendingTransactions, data.Transaction)
        fmt.Println("Added transaction, ", data.Transaction)
    }

    fmt.Fprintf(w, "ok")
}

