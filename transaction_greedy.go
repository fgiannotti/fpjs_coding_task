package fpjs_coding_task

import (
	"math"
	"sort"
)

type TransactionsGreedyService struct {
	latencyService LatencyService
}

//Function to tests profits
func (ts *TransactionsGreedyService) bestProfit(transactions []Transaction, totalTimeMs int) float64 {
	simpleTransactions := ts.simplify(transactions)
	return ts.maxProfit(simpleTransactions,totalTimeMs)
}

func (ts *TransactionsGreedyService) prioritize(transactions []Transaction, totalTimeMs int) []Transaction {
	//simpleTransactions := ts.simplify(transactions)
	//TODO: Make profit functions to save used transactions
	return []Transaction{}
}

func (ts *TransactionsGreedyService) maxProfit(transactions []simpleTransaction, totalTimeMs int) float64 {
	transactionsByAmount := make([]simpleTransaction,len(transactions))
	transactionsByTime := make([]simpleTransaction,len(transactions))
	copy(transactionsByAmount,transactions)
	copy(transactionsByTime,transactions)

	sort.SliceStable(transactionsByAmount, func(i, j int) bool {
		return transactions[i].Amount > transactions[j].Amount
	})
	sort.SliceStable(transactionsByTime, func(i, j int) bool {
		return transactions[i].LatencyMs < transactions[j].LatencyMs
	})
	profitFromAmountOrder := getGreedyProfit(transactionsByAmount, totalTimeMs)
	profitFromLatencyOrder := getGreedyProfit(transactionsByTime, totalTimeMs)

	return math.Max(float64(profitFromAmountOrder),float64(profitFromLatencyOrder))
}

func getGreedyProfit(transactions []simpleTransaction, totalTimeMs int) float32 {
	time := totalTimeMs
	i := 0
	profit := float32(0)
	for time > 0 && i < len(transactions) {
		if transactions[i].LatencyMs <= time {
			profit += transactions[i].Amount
			time -= transactions[i].LatencyMs
		}
		i++
	}
	return profit
}


func (ts *TransactionsGreedyService) simplify(transactions []Transaction) []simpleTransaction {
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