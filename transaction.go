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

type Transactions interface {
	prioritize(transactions []Transaction, totalTimeMs int) []Transaction
}
type TransactionsService struct {
	latencyService LatencyService
}

func (ts *TransactionsService) bestProfit(transactions []Transaction, totalTimeMs int) float32 {
	bestTransactions := ts.prioritize(transactions,totalTimeMs)
	result := float32(0)
	for _,transaction := range bestTransactions {
		result += transaction.Amount
	}

	return result
}

func (ts *TransactionsService) prioritize(transactions []Transaction, totalTimeMs int) []Transaction {
	amounts, times := ts.mapToAmountsAndTimes(transactions)
	resultIds := maxProfit(amounts, times, totalTimeMs)

	result := make([]Transaction, len(resultIds))
	for i := 0; i < len(resultIds); i++ {
		result[i] = transactions[resultIds[i]-1]
	}

	return result
}

func maxProfit(amounts []float32, times []int, totalTimeMs int) []int {
	matrix := buildProfitMatrix(amounts, times, totalTimeMs)
	i, t := len(amounts)-1, totalTimeMs
	maxProfit := matrix[i+1][t]
	result := []int{}
	profit := maxProfit
	for profit > 0 || i < 0 {
		//if profit changes that transaction was a must
		if matrix[i][t] != matrix[i+1][t] {
			profit -= int(amounts[i])
			t -= times[i]
			result = append(result,i)
		}
		i--
	}
	return result
}

func buildProfitMatrix(amounts []float32, times []int, totalTimeMs int) [][]int {
	// create the empty matrix
	memo := make([][]int, len(amounts)+1)
	for i := range memo {
		memo[i] = make([]int, totalTimeMs+1)
	}

	for i := 1; i <= len(amounts); i++ {
		for t := 1; t <= totalTimeMs; t++ {
			if times[i-1] <= t {
				val1 := float64(memo[i-1][t])
				val2 := float64(amounts[i-1])
				if (t - times[i-1]) >= 0 {
					val2 += float64(memo[i-1][t-times[i-1]])
				}
				memo[i][t] = int(math.Max(val1, val2))

			} else {
				memo[i][t] = memo[i-1][t]
			}
		}
	}
	return memo
}
func (ts *TransactionsService) mapToAmountsAndTimes(transactions []Transaction) ([]float32, []int) {
	amounts := make([]float32, len(transactions)+1)
	times := make([]int, len(transactions)+1)
	amounts[0], times[0] = 0, 0
	for i := 1; i <= len(transactions); i++ {
		amounts[i] = transactions[i-1].Amount
		times[i] = ts.latencyService.GetTransactionLatency(transactions[i-1])
	}
	return amounts, times
}

//caso 100 USD 1 seg 0 , 10 veces 20 USD en 0,01. Gana la segunda a pesar de que es mas chico
