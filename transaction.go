package fpjs_coding_task

import (
	"math"
)

type Transaction struct {
	// a UUID of transaction
	ID string
	// in USD, typically a value betwen 0.01 and 1000 USD.
	Amount float32
	// a 2-letter country code of where the bank is located
	BankCountryCode string
}

type simpleTransaction struct {
	Amount    float32
	LatencyMs int
}

type Transactions interface {
	prioritize(transactions []Transaction, totalTimeMs int) []Transaction
}
type TransactionsService struct {
	latencyService LatencyService
}

func (ts *TransactionsService) bestProfit(transactions []Transaction, totalTimeMs int) float32 {
	bestTransactions := ts.prioritize(transactions, totalTimeMs)
	result := float32(0)
	for _, transaction := range bestTransactions {
		result += transaction.Amount
	}

	return result
}

func (ts *TransactionsService) prioritize(transactions []Transaction, totalTimeMs int) []Transaction {
	simpleTransactions := ts.simplify(transactions)
	resultIds := ts.maxProfit(simpleTransactions, totalTimeMs)

	result := make([]Transaction, len(resultIds))
	for i := 0; i < len(resultIds); i++ {
		result[i] = transactions[resultIds[i]-1]
	}

	return result
}

func (ts *TransactionsService) maxProfit(transactions []simpleTransaction, totalTimeMs int) []int {
	matrix := buildProfitMatrix(transactions, totalTimeMs)
	i, t := len(transactions)-1, totalTimeMs
	maxProfit := matrix[i+1][t]
	result := []int{}
	profit := maxProfit
	for profit > 0 || i > 0 {
		//if profit changes that transaction was a must
		if matrix[i][t] != matrix[i+1][t] {
			profit -= int(transactions[i].Amount)
			t -= transactions[i].LatencyMs
			result = append(result, i)
		}
		i--
	}
	return result
}

func buildProfitMatrix(transactions []simpleTransaction, totalTimeMs int) [][]int {
	// create the empty matrix
	memo := make([][]int, len(transactions)+1)
	for i := range memo {
		memo[i] = make([]int, totalTimeMs+1)
	}

	for i := 1; i <= len(transactions); i++ {
		for time := 1; time <= totalTimeMs; time++ {
			if transactions[i-1].LatencyMs <= time {
				val1 := float64(memo[i-1][time])
				val2 := float64(transactions[i-1].Amount)
				if (time - transactions[i-1].LatencyMs) >= 0 {
					val2 += float64(memo[i-1][time-transactions[i-1].LatencyMs])
				}
				memo[i][time] = int(math.Max(val1, val2))

			} else {
				memo[i][time] = memo[i-1][time]
			}
		}
	}
	return memo
}
func (ts *TransactionsService) simplify(transactions []Transaction) []simpleTransaction {
	result := make([]simpleTransaction, len(transactions)+1)
	//start with 0 to avoid null checks in the matrix
	result[0] = simpleTransaction{Amount: 0, LatencyMs: 0}

	for i := 1; i <= len(transactions); i++ {
		result[i] = simpleTransaction{
			Amount:    transactions[i-1].Amount,
			LatencyMs: ts.latencyService.GetTransactionLatency(transactions[i-1]),
		}
	}
	return result
}