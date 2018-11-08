package main

import (
    "fmt"
    "bytes"
    "net/http"
    "encoding/json"
    "os"
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
    decoder := json.NewDecoder(r.Body)

    var body NewTransactionRequest
    err := decoder.Decode(&body)

	if err != nil {
		panic(err)
    }
    
    resp, errText := createNewTransaction(body.To, body.From, body.Amount)
    if resp {
        fmt.Fprintf(w, "Added to transaction pool.")
    } else {
        fmt.Fprintf(w, errText)
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

    var body Peer
    err := decoder.Decode(&body)

	if err != nil {
        fmt.Println(err)
		panic(err)
    }
    
    Peers = append(Peers, body)

    resp, _ := json.Marshal(Peers)

    fmt.Fprintf(w, string(resp))
}

var Peers []Peer

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
    http.HandleFunc("/addPeer", addPeer)
    http.HandleFunc("/newBlock", newBlockFromNode)

    port := "9090"
    if len(os.Args) > 1 {
        port = os.Args[1]
    }
    // Start server
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        fmt.Println(err)
    }
}


// Broadcast
func broadcastNewBlock(block Block) {
    for _, peer := range Peers {
        data := BlockBroadcast{
            Block: block, 
            Peer: peer,
        }
        jsonPeer, _ := json.Marshal(data)
        url := peer.Address + ":" + peer.Port + "/newBlock"
        req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPeer))
        req.Header.Set("Content-Type", "application/json")
        
        fmt.Println("Broadcasted block to ", peer)

        client := &http.Client{}
        _, err := client.Do(req)
        if err != nil {
            fmt.Println(err)
        }
    }
}

func newBlockFromNode(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var newBlock BlockBroadcast
    err := decoder.Decode(&newBlock)

    if err != nil {
		panic(err)
    }

    added := addBlockToChain(newBlock.Block)
    if !added && newBlock.Block.Index > getLatestBlock().Index {
        peerUrl := newBlock.Peer.Address + ":" + newBlock.Peer.Port + "/blockchain"
        resp, err := http.Get(peerUrl)

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

        fmt.Println("Fetched blockchain")
    }

    //fmt.Println("New block: ", newBlock)

    fmt.Fprintf(w, "ok")
}
