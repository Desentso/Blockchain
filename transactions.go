package main

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

var PendingTransactions []Transaction
var UnspentTransactionsOut []TransactionOut

func createNewTransaction(to string, from string, amount int) {
	unspentTransactions, leftOverAmount := findUnspentTransactionsFor(from, amount)
	
	if len(unspentTransactions) == 0 {
		
	} else {
		var unSignedTransactionsIn []TransactionIn
		for _, txOut := range unspentTransactions {
			unSignedTransactionsIn = append(unSignedTransactionsIn, createUnsignedTransactionIn(txOut))
		}

		transactionsOut := createNewTransactionsOut(from, to, amount, leftOverAmount)
		var transaction Transaction
		transaction.Id = getTransactionHash(transactionsOut, unSignedTransactionsIn)
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

func createUnsignedTransactionIn(unspentTransaction TransactionOut) TransactionIn {

	return TransactionIn{
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


func getTransactionHash(transactionsOut []TransactionOut, transactionsIn []TransactionIn) string {
	combinedString := ""
	
	for _, txOut := range transactionsOut {
		combinedString += (txOut.Address + string(txOut.Amount))
	}

	for _, txIn := range transactionsIn {
		combinedString += (txIn.TransactionOutId + string(txIn.TransactionOutIndex))
	}

	return calculateHash(combinedString)
}
