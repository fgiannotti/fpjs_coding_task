package fpjs_coding_task

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalLatencyService_GetTransactionLatency(t *testing.T) {
	expectedBrLatency := 37

	ls := NewLatencyService()
	lat := ls.GetTransactionLatency(Transaction{BankCountryCode: "br"})

	assert.NotZerof(t, lat, "error: latency is zero")
	assert.Equal(t, expectedBrLatency, lat)
}
