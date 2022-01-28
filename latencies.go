package fpjs_coding_task

import (
	"encoding/json"
	"os"
	"strings"
)

type LatencyService interface {
	GetTransactionLatency(transaction Transaction) int
}

type LocalLatencyService struct {
	Latencies map[string]int
}

func NewLatencyService() LatencyService {
	file, err := os.ReadFile("api_latencies.json")
	if err != nil {
		panic(err)
	}
	latencies := map[string]int{}

	err = json.Unmarshal(file, &latencies)
	if err != nil {
		panic(err)
	}

	return &LocalLatencyService{Latencies: latencies}
}

func (ls *LocalLatencyService) GetTransactionLatency(transaction Transaction) int {
	return ls.Latencies[strings.ToLower(transaction.BankCountryCode)]
}
