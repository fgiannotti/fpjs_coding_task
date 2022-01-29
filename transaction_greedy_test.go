package fpjs_coding_task

import (
	"os"
	"reflect"
	"testing"
)

func TestTransactionsGreedyService_bestProfit(t *testing.T) {
	csv ,err := os.ReadFile("test_transactions.csv")
	if err != nil { panic(err) }
	parsedTransactions := parseTransactions(csv)
	bestProfitExpected1000ms := 1849.150009393692
	bestProfitExpected50ms := 33.29999923706055
	bestProfitExpected60ms := 43.58000183105469
	bestProfitExpected90ms := 47.019999265670776

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
		want   float64
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
			"Test max profit in transactions with 90",
			fields{NewLatencyService()},
			args{parsedTransactions, 90},
			bestProfitExpected90ms,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TransactionsGreedyService{
				latencyService: tt.fields.latencyService,
			}
			if got := ts.bestProfit(tt.args.transactions, tt.args.totalTimeMs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bestProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
