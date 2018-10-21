package main

import (
	"fmt"
	"crypto/sha256"
	"crypto"
	"crypto/rsa"
	"crypto/rand"
	"encoding/hex"
	"../utils"
)

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

var UnspentTransactionsOut []TransactionOut
var privateKey  *rsa.PrivateKey

func main() {
	privateKey = createNewPrivateKey()
	fmt.Println(string(utils.PublicKeyToBytes(&privateKey.PublicKey)))
	fmt.Println(string(utils.PrivateKeyToBytes(privateKey)))

	UnspentTransactionsOut = []TransactionOut{TransactionOut{Id: "id", Index: "index", Address: "123", Amount: 100}}

	createNewTransaction("abc", "123", 100)
}

func createNewPrivateKey() *rsa.PrivateKey {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 512)
	return privateKey
}

/*func PublicKeyToBytes(pub *rsa.PublicKey) []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		fmt.Println("Error while converting public key to bytes")
	}
	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes
}

func PrivateKeyToBytes(priv *rsa.PrivateKey) []byte {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return privBytes
}

func calculateHash(blockString string) string {
	strAsBytes := []byte(blockString)
	sum := sha256.Sum256(strAsBytes)

	return hex.EncodeToString(sum[:])
}*/

func createNewTransaction(to string, from string, amount int) {
	unspentTransactions, leftOverAmount := findUnspentTransactionsFor(from, amount)
	
	if len(unspentTransactions) == 0 {
		fmt.Println("Not enough unspent transactions")
	} else {
		var unSignedTransactionsIn []*TransactionIn
		for _, txOut := range unspentTransactions {
			unSignedTransactionsIn = append(unSignedTransactionsIn, createUnsignedTransactionIn(txOut))
		}

		transactionsOut := createNewTransactionsOut(from, to, amount, leftOverAmount)
		var transaction Transaction
		transaction.Outputs = transactionsOut
		transaction.Id = getTransactionHash(transactionsOut, unSignedTransactionsIn)

		transaction.Inputs = signTransactionsIn(transaction, unSignedTransactionsIn)

		fmt.Println(transaction)
	}

}

func findUnspentTransactionsFor(fromAddr string, amount int) ([]TransactionOut, int) {

	var found []TransactionOut
	amountFound := 0

	for _, txOut := range UnspentTransactionsOut {
		if txOut.Address == fromAddr {
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
			Address: to,
			Amount: amount,
		},
	}

	if leftOverAmount > 0 {
		leftOverTxOut := TransactionOut{
			Address: from,
			Amount: leftOverAmount,
		}
		txsOut = append(txsOut, leftOverTxOut)
	} 
	
	return txsOut
}

func signTransactionsIn(transaction Transaction, unsigned []*TransactionIn) []TransactionIn {
	values := []TransactionIn{}
	for _, txIn := range unsigned {
		signature := signTransactionIn(privateKey, transaction)
		fmt.Println("Signature: ", signature)
		txIn.Signature = signature
		values = append(values, *txIn)
	}

	return values
}

func signTransactionIn(privateKey *rsa.PrivateKey, transaction Transaction) string {
	stringAsBytes := []byte(transaction.Id)
	hash := sha256.Sum256(stringAsBytes)
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	return hex.EncodeToString(signature)
}

func getTransactionHash(transactionsOut []TransactionOut, transactionsIn []*TransactionIn) string {
	combinedString := ""
	
	for _, txOut := range transactionsOut {
		combinedString += (txOut.Address + string(txOut.Amount))
	}

	for _, txIn := range transactionsIn {
		combinedString += (txIn.TransactionOutId + string(txIn.TransactionOutIndex))
	}

	return utils.CalculateHash(combinedString)
}
