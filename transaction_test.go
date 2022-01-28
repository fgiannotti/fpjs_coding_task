package fpjs_coding_task

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestTransactionsService_mapToAmountsAndTimes(t *testing.T) {
	mockTransactionAR300 := Transaction{"id", float32(300), "ar"}
	mockTransactionAR200 := Transaction{"id", float32(200), "ar"}
	ls := NewLatencyService()
	type fields struct {
		latencyService LatencyService
	}
	type args struct {
		transactions []Transaction
		totalTimeMs  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []simpleTransaction
	}{
		{
			"Test 2 transaction and take the best",
			fields{ls},
			args{[]Transaction{mockTransactionAR200, mockTransactionAR300}, 100},
			[]simpleTransaction{
				{mockTransactionAR200.Amount,ls.GetTransactionLatency(mockTransactionAR200)},
				{mockTransactionAR300.Amount, ls.GetTransactionLatency(mockTransactionAR300)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TransactionsService{
				latencyService: tt.fields.latencyService,
			}
			got := ts.simplify(tt.args.transactions)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapToAmountsAndTimes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionsService_prioritize(t *testing.T) {
	mockTransactionAR300 := Transaction{"id", float32(300), "ar"}
	mockTransactionAR200 := Transaction{"id", float32(200), "ar"}

	type fields struct {
		latencyService LatencyService
	}
	type args struct {
		transactions []Transaction
		totalTimeMs  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Transaction
	}{
		{
			"Test 2 transaction and take the best",
			fields{NewLatencyService()},
			args{[]Transaction{mockTransactionAR200, mockTransactionAR300}, 100},
			[]Transaction{mockTransactionAR300},
		},
		{
			"Test 2 transaction and take both",
			fields{NewLatencyService()},
			args{[]Transaction{mockTransactionAR200, mockTransactionAR300}, 1000},
			[]Transaction{mockTransactionAR300, mockTransactionAR200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TransactionsService{
				latencyService: tt.fields.latencyService,
			}
			if got := ts.prioritize(tt.args.transactions, tt.args.totalTimeMs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prioritize() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestTransactionsService_bestProfit(t *testing.T) {
	csv ,err := os.ReadFile("test_transactions.csv")
	if err != nil { panic(err) }
	parsedTransactions := parseTransactions(csv)
	bestProfitExpected1000ms := float32(35471.812)
	bestProfitExpected50ms := float32(4139.43)
	bestProfitExpected60ms := float32(4675.71)
	bestProfitExpected90ms := float32(6972.29)

	type fields struct {
		latencyService LatencyService
	}
	type args struct {
		transactions []Transaction
		totalTimeMs  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		{
			"Test max profit in transactions with 1000ms",
			fields{NewLatencyService()},
			args{parsedTransactions, 1000},
			bestProfitExpected1000ms,
		},
		{
			"Test max profit in transactions with 50",
			fields{NewLatencyService()},
			args{parsedTransactions, 50},
			bestProfitExpected50ms,
		},
		{
			"Test max profit in transactions with 60",
			fields{NewLatencyService()},
			args{parsedTransactions, 60},
			bestProfitExpected60ms,
		},
		{
			"Test 2 transaction and take both",
			fields{NewLatencyService()},
			args{parsedTransactions, 90},
			bestProfitExpected90ms,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TransactionsService{
				latencyService: tt.fields.latencyService,
			}
			if got := ts.bestProfit(tt.args.transactions, tt.args.totalTimeMs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prioritize() = %v, want %v", got, tt.want)
			}
		})
	}
}
func parseTransactions(csv []byte) []Transaction {
	lines := strings.Split(string(csv), "\n")

	transactions := make([]Transaction, 0)

	for i, v := range lines {
		if i == 0 {
			continue
		}
		s := strings.Split(v, ",")
		newItemWorth, _ := strconv.ParseFloat(s[1],32)
		newTransaction := Transaction{ID: s[0], Amount: float32(newItemWorth), BankCountryCode: s[2]}
		transactions = append(transactions, newTransaction)
	}
	return transactions
}
