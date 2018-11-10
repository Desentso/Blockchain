package main

import (
	"fmt"
	"time"
	"math"
	"encoding/hex"
	"strings"
	"./utils"
)

type Block struct {
	Timestamp int64 `json:"timestamp"`
	Index int `json:"index"`
	Data string `json:"data"`
	Transactions []Transaction `json:"transactions"`
	Hash string `json:"hash"`
	PrevHash string `json:"prevHash"`
	Difficulty int 
	Nonce int 
}

const BLOCK_GENERATION_INTERVAL = 10
const DIFFICULTY_ADJUSTMENT_INTERVAL = 10
const BLOCK_REWARD_AMOUNT = 50

var Blockchain []Block

func main() {

	Blockchain = append(Blockchain, generateGenesisBlock())
	fmt.Println(Blockchain)
	//Blockchain = append(Blockchain, generateNewBlock("Block number 2"))
	//fmt.Println(Blockchain)
	
	transactionTest()
	node()
}

/*func CalculateHash(blockString string) string {
	/*h := sha256.New()
	h.Write([]byte(blockString))
	sum := h.Sum(nil)*//*
	strAsBytes := []byte(blockString)
	sum := sha256.Sum256(strAsBytes)
	//fmt.Printf("%x", sum)
	return hex.EncodeToString(sum[:])
}*/

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

func buildBlockString(index int, timestamp int64, prevHash string, data string, nonce int, difficulty int, transactions []Transaction) string {
	return fmt.Sprintf("%d, %d, %s, %s, %d, %d, %v\n", index, timestamp, prevHash, data, nonce, difficulty, transactions)
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

	blockString := buildBlockString(index, timestamp, prevHash, data, nonce, difficulty, []Transaction{})
	hash := utils.CalculateHash(blockString)

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

func CalculateCumulativeDifficulty(blockchain []Block) int {
	difficulty := 0
	for _, block := range blockchain {
		difficulty += int(math.Pow(2, float64(block.Difficulty)))
	}

	return difficulty
}

func isValidNewBlock(prevBlock Block, block Block) bool {
	calculatedHash := utils.CalculateHash(buildBlockString(block.Index, block.Timestamp, block.PrevHash, block.Data, block.Nonce, block.Difficulty, block.Transactions))

	if (block.Index != prevBlock.Index + 1 || block.PrevHash != prevBlock.Hash || calculatedHash != block.Hash) {
		return false
	}

	if !ValidateTransactions(block.Transactions) {
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
	if isValidNewBlock(getLatestBlock(), newBlock) && ValidateTransactions(newBlock.Transactions) {
		Blockchain = append(Blockchain, newBlock)
		//UnspentTransactionsOut = append(UnspentTransactionsOut, reward)
		updateTransactions(newBlock)

		fmt.Println("Found new block! Difficulty: ", newBlock.Difficulty, " , Hash: ", newBlock.Hash)
		return true
	}

	return false
}

func minerReward(address string) Transaction {
	reward := TransactionOut{
		Amount: BLOCK_REWARD_AMOUNT, 
		ToAddress: address, 
		Index: string(len(Blockchain)), 
		Unspent: true,
	}
	transactionsOut := []TransactionOut{reward}

	var transaction Transaction
	transaction.Outputs = transactionsOut
	transaction.Inputs = []TransactionIn{}
	transaction.Id = GetTransactionHash(transactionsOut, []*TransactionIn{})

	return transaction
}

func updateTransactions(block Block) {
	consumedTxOuts := []TransactionOut{}

	for _, transaction := range block.Transactions {
		// Add the new unspent transactions
		for i, txOut := range transaction.Outputs {
			newTxOut := TransactionOut{Id: transaction.Id, Index: string(i), ToAddress: txOut.ToAddress, Amount: txOut.Amount}
			UnspentTransactionsOut = append(UnspentTransactionsOut, newTxOut)
		}
		
		// Remove spent txOuts from unspent
		for _, txIn := range transaction.Inputs {
			consumedTxOuts = append(consumedTxOuts, )
			for index, unspentTxOut := range UnspentTransactionsOut {
				if unspentTxOut.Id == txIn.TransactionOutId && unspentTxOut.Index == txIn.TransactionOutIndex {
					UnspentTransactionsOut = append(UnspentTransactionsOut[:index], UnspentTransactionsOut[index+1:]...)
				}
			}
		}

		// Remove transaction from pending transactions
		for index, pendingTx := range PendingTransactions {
			if pendingTx.Id == transaction.Id {
				PendingTransactions = append(PendingTransactions[:index], PendingTransactions[index+1:]...)
			}
		}
	}
}

func mineBlock(data string) Block {
	prevBlock := getLatestBlock()

	index := prevBlock.Index + 1
	prevHash := prevBlock.Hash
	timestamp := getTimestamp()
	difficulty := getDifficulty()
	nonce := 0
	transactions := []Transaction{minerReward(string(utils.PublicKeyToBytes(PublicKey)))}
	transactions = append(transactions, PendingTransactions...)

	hash := utils.CalculateHash(buildBlockString(index, timestamp, prevHash, data, nonce, difficulty, transactions))

	for !hashMatchesDifficulty(hash, difficulty) {
		nonce += 1
		hash = utils.CalculateHash(buildBlockString(index, timestamp, prevHash, data, nonce, difficulty, transactions))
	}

	return Block{
		Index: index, 
		Timestamp: timestamp, 
		PrevHash: prevHash, 
		Hash: hash, 
		Data: data,
		Nonce: nonce,
		Difficulty: difficulty,
		Transactions: transactions,
	}
}
