package main

import (
	"fmt"
	"crypto/sha256"
	"time"
	"encoding/hex"
	"strings"
)

type Block struct {
	Timestamp int64 `json:"timestamp"`
	Index int `json:"index"`
	Data string `json:"data"`
	Hash string `json:"hash"`
	PrevHash string `json:"prevHash"`
	Difficulty int 
	Nonce int 
}

type TransactionOut struct {
	Id string
	Index string
	Address string
	Amount int
} 

type TransactionIn struct {
	TransactionOutId string
	TransactionOutIndex string
	Signature string
}

type Transaction struct {
	Id string
	Inputs []TransactionIn
	Outputs []TransactionOut
}

const BLOCK_GENERATION_INTERVAL = 10
const DIFFICULTY_ADJUSTMENT_INTERVAL = 10

var Blockchain []Block

func main() {

	Blockchain = append(Blockchain, generateGenesisBlock())
	fmt.Println(Blockchain)
	//Blockchain = append(Blockchain, generateNewBlock("Block number 2"))
	//fmt.Println(Blockchain)

	node()
}

func calculateHash(blockString string) string {
	h := sha256.New()
	h.Write([]byte(blockString))
	sum := h.Sum(nil)
	//fmt.Printf("%x", sum)
	return hex.EncodeToString(sum)
}

func getDifficulty() int {
	lastBlock := getLatestBlock()

	if lastBlock.Index % DIFFICULTY_ADJUSTMENT_INTERVAL == 0 && lastBlock.Index != 0 {
		return adjustedDifficulty(lastBlock)
	} else {
		return lastBlock.Difficulty
	}
}

func adjustedDifficulty(lastBlock Block) int {
	lastAdjustedBlock := Blockchain[len(Blockchain) - DIFFICULTY_ADJUSTMENT_INTERVAL]
	timeTakenBetween := int(lastBlock.Timestamp - lastAdjustedBlock.Timestamp) / 1000
	timeExpected := DIFFICULTY_ADJUSTMENT_INTERVAL * BLOCK_GENERATION_INTERVAL

	if timeTakenBetween > timeExpected {
		return lastBlock.Difficulty - 1
	} else if timeTakenBetween < timeExpected {
		return lastBlock.Difficulty + 1
	} else {
		return lastBlock.Difficulty
	}
}

func HexToBin(hexString string) string {
	byteArr, _ := hex.DecodeString(hexString)
	s := ""
	for _, n := range(byteArr) {
		s += fmt.Sprintf("%08b", n)
	}
	return s
}

func hashMatchesDifficulty(hash string, difficulty int) bool {
	binaryHash := HexToBin(hash)
	//fmt.Println("Binary hash: ", binaryHash)
	//fmt.Printf("Binary hash: %v\n", binaryHash)
	difficultyPrefix := strings.Repeat("0", difficulty)
	if strings.HasPrefix(binaryHash, difficultyPrefix) {
		return true
	} else {
		return false
	}
}

func getTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func buildBlockString(index int, timestamp int64, prevHash string, data string, nonce int, difficulty int) string {
	return fmt.Sprintf("%d, %d, %s, %s, %d, %d", index, timestamp, prevHash, data, nonce, difficulty)
}

func getLatestBlock() Block {
	return Blockchain[len(Blockchain)-1]
}

func generateGenesisBlock() Block {
	prevHash := "0"
	index := 0
	timestamp := getTimestamp()
	data := "Genesis block"
	nonce := 0
	difficulty := 1
	fmt.Println(timestamp)

	blockString := buildBlockString(index, timestamp, prevHash, data, nonce, difficulty)
	hash := calculateHash(blockString)

	hashMatchesDifficulty(hash, difficulty)

	return Block{
		PrevHash: prevHash, 
		Index: index, 
		Hash: hash, 
		Timestamp: timestamp, 
		Data: data, 
		Difficulty: difficulty, 
		Nonce: nonce,
	}
}

func isValidNewBlock(prevBlock Block, block Block) bool {
	calculatedHash := calculateHash(buildBlockString(block.Index, block.Timestamp, block.PrevHash, block.Data, block.Nonce, block.Difficulty))

	if (block.Index != prevBlock.Index + 1 || block.PrevHash != prevBlock.Hash || calculatedHash != block.Hash) {
		return false
	}

	return true
}

func isValidBlockchain(blockchain []Block) bool {
	for i := range blockchain {
		if (!isValidNewBlock(blockchain[i], blockchain[i+1])) {
			return false
		}
	}

	return true
}

func addBlockToChain(newBlock Block) bool {
	if isValidNewBlock(getLatestBlock(), newBlock) {
		Blockchain = append(Blockchain, newBlock)
		fmt.Println("Found new block! Difficulty: ", newBlock.Difficulty, " , Hash: ", newBlock.Hash)
		return true
	}

	return false
}

/*func generateNewBlock(data string, nonce int) Block {
	prevBlock := getLatestBlock()

	index := prevBlock.Index + 1
	prevHash := prevBlock.Hash
	timestamp := getTimestamp()

	blockString := buildBlockString(index, timestamp, prevHash, data)
	hash := calculateHash(blockString)

	return Block{Index: index, Timestamp: timestamp, PrevHash: prevHash, Hash: hash, Data: data}
}*/


func mineBlock(data string) Block {
	prevBlock := getLatestBlock()

	index := prevBlock.Index + 1
	prevHash := prevBlock.Hash
	timestamp := getTimestamp()
	difficulty := getDifficulty()
	nonce := 0

	hash := calculateHash(buildBlockString(index, timestamp, prevHash, data, nonce, difficulty))

	for !hashMatchesDifficulty(hash, difficulty) {
		nonce += 1
		hash = calculateHash(buildBlockString(index, timestamp, prevHash, data, nonce, difficulty))
	}

	return Block{
		Index: index, 
		Timestamp: timestamp, 
		PrevHash: prevHash, 
		Hash: hash, 
		Data: data,
		Nonce: nonce,
		Difficulty: difficulty,
	}
}
