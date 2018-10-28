package main

import (
	"fmt"
	"crypto/sha256"
	"crypto"
	"crypto/rsa"
	"crypto/rand"
	"encoding/hex"
	"./utils"
)

//https://play.golang.org/p/dx37I-1g0ga
//https://gist.github.com/miguelmota/3ea9286bd1d3c2a985b67cac4ba2130a

type TransactionOut struct {
	Id string // Id of transaction
	Index string // Used to make the hashes unique
	ToAddress string // Address that to coins belong (Unspent) or are sent (When spending)
	Amount int
	Unspent bool // True means that the coins belong to ToAddress
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

var UnspentTransactionsOut []TransactionOut
var PendingTransactions []Transaction
var PrivateKey *rsa.PrivateKey
var PublicKey *rsa.PublicKey

func transactionTest() {
	PrivateKey = createNewPrivateKey()
	PublicKey = &PrivateKey.PublicKey
	fmt.Println(string(utils.PublicKeyToBytes(PublicKey)))
	fmt.Println(string(utils.PrivateKeyToBytes(PrivateKey)))

	//UnspentTransactionsOut = []TransactionOut{TransactionOut{Id: "id", Index: "index", ToAddress: "123", Amount: 100}}

	//createNewTransaction("abc", "123", 100)
}

func createNewPrivateKey() *rsa.PrivateKey {
	PrivateKey, _ := rsa.GenerateKey(rand.Reader, 512)
	return PrivateKey
}

func createNewTransaction(to string, from string, amount int) (bool, string) {
	unspentTransactions, leftOverAmount := findUnspentTransactionsFor(from, amount)
	
	if len(unspentTransactions) == 0 {
		fmt.Println("Couldn't find enough unspent transactions")
		return false, "Couldn't find enough unspent transactions"
	} else {
		var unSignedTransactionsIn []*TransactionIn
		for _, txOut := range unspentTransactions {
			unSignedTransactionsIn = append(unSignedTransactionsIn, createUnsignedTransactionIn(txOut))
		}

		transactionsOut := createNewTransactionsOut(from, to, amount, leftOverAmount)
		var transaction Transaction
		transaction.Outputs = transactionsOut
		transaction.Id = GetTransactionHash(transactionsOut, unSignedTransactionsIn)

		transaction.Inputs = signTransactionsIn(transaction, unSignedTransactionsIn)

		if ValidateTransaction(transaction) && ValidTransactionToPool(transaction) {
			PendingTransactions = append(PendingTransactions, transaction)
			fmt.Println(transaction)
			return true, ""
		} else {
			return false, "Invalid transaction"
		}
	}
}

func findUnspentTransactionsFor(fromAddr string, amount int) ([]TransactionOut, int) {

	var found []TransactionOut
	amountFound := 0

	for _, txOut := range UnspentTransactionsOut {
		if txOut.ToAddress == fromAddr {
			found = append(found, txOut)
			amountFound += txOut.Amount
		}
	}

	leftOver := 0
	if amountFound >= amount {
		leftOver = amountFound - amount
	} else {
		found = []TransactionOut{}
	}

	return found, leftOver
}	

func createUnsignedTransactionIn(unspentTransaction TransactionOut) *TransactionIn {

	return &TransactionIn{
		TransactionOutId: unspentTransaction.Id,
		TransactionOutIndex: unspentTransaction.Index,
	}
}

func createNewTransactionsOut(from string, to string, amount int, leftOverAmount int) []TransactionOut {

	txsOut := []TransactionOut{
		TransactionOut{
			ToAddress: to,
			Amount: amount,
		},
	}

	if leftOverAmount > 0 {
		leftOverTxOut := TransactionOut{
			ToAddress: from,
			Amount: leftOverAmount,
			Unspent: true,
		}
		txsOut = append(txsOut, leftOverTxOut)
	} 
	
	return txsOut
}

func signTransactionsIn(transaction Transaction, unsigned []*TransactionIn) []TransactionIn {
	values := []TransactionIn{}
	for _, txIn := range unsigned {
		signature := signTransactionIn(PrivateKey, transaction)
		fmt.Println("Signature: ", signature)
		txIn.Signature = signature
		values = append(values, *txIn)
	}

	return values
}

func signTransactionIn(PrivateKey *rsa.PrivateKey, transaction Transaction) string {
	stringAsBytes := []byte(transaction.Id)
	hash := sha256.Sum256(stringAsBytes)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, PrivateKey, crypto.SHA256, hash[:])
	fmt.Println("Transaction ID:", transaction.Id, "\nHash: ", hash, "\nSignature:", hex.EncodeToString(signature))
	return hex.EncodeToString(signature)
}

func GetTransactionHash(transactionsOut []TransactionOut, transactionsIn []*TransactionIn) string {
	combinedString := ""
	
	for _, txOut := range transactionsOut {
		combinedString += (txOut.ToAddress + string(txOut.Amount))
	}

	for _, txIn := range transactionsIn {
		combinedString += (txIn.TransactionOutId + string(txIn.TransactionOutIndex))
	}

	return utils.CalculateHash(combinedString)
}


// Validation
func ValidTransactionToPool(transaction Transaction) bool {
	for _, txIn := range transaction.Inputs {
		if isAlreadyPending(txIn) {
			return false
		}
	}

	return true
}

func ValidateTransactions(transactions []Transaction) bool {
	
	for _, transaction := range transactions[1:] {
		if !ValidateTransaction(transaction) {
			return false
		}
	}

	return true
}

func ValidateTransaction(transaction Transaction) bool {

	txInAmount := 0
	txOutAmount := 0

	for _, txIn := range transaction.Inputs {
		if !validateTxIn(transaction, txIn) {
			return false
		}
		txOut,_ := findReferencedTxOut(txIn)
		txInAmount += txOut.Amount
	}

	for _, txOut := range transaction.Outputs {
		txOutAmount += txOut.Amount
	}

	if txInAmount != txOutAmount {
		return false
	}

	return true
}

func validateTxIn(transaction Transaction, txIn TransactionIn) bool {

	referencedTxOut, foundReferencedTxOut := findReferencedTxOut(txIn)

	// Verify the signature, i.e. the coins belong to the spender
	if foundReferencedTxOut {
		publicKey := utils.BytesToPublicKey([]byte(referencedTxOut.ToAddress))

		stringAsBytes := []byte(transaction.Id)
		txIdhash := sha256.Sum256(stringAsBytes)

		decodedSig, _ := hex.DecodeString(txIn.Signature)

		invalidSignature := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, txIdhash[:], decodedSig)

		if invalidSignature != nil {
			fmt.Println("Invalid Signature", transaction.Id, txIn)
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func findReferencedTxOut(txIn TransactionIn) (TransactionOut, bool) {
	// Find the referenced unspent transaction out
	for _, unSpentTxOut := range UnspentTransactionsOut {
		if unSpentTxOut.Id == txIn.TransactionOutId && unSpentTxOut.Index == txIn.TransactionOutIndex {
			return unSpentTxOut, true
		}
	}

	return TransactionOut{}, false
}

func isAlreadyPending(txIn TransactionIn) bool {

	for _, tx := range PendingTransactions {
		for _, pendingTxIn := range tx.Inputs {
			if (txIn.TransactionOutId == pendingTxIn.TransactionOutId && 
				txIn.TransactionOutIndex == pendingTxIn.TransactionOutIndex) {
					return true
			}
		}
	}

	return false
}
