package coinpayments

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalances(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
		"error":"ok",
		"result":{"BTC":{"balance":"0","balancef":"0.00000000","status":"available","coin_status":"online"}}
		}`)
	})

	client := NewClient("", "", httpClient)
	balances, _, err := client.Balances.Show(&BalanceParams{All: 0})
	expectedMap := map[string]Balance{}
	expectedMap["BTC"] = Balance{
		Balance:    "0",
		BalanceF:   "0.00000000",
		Status:     "available",
		CoinStatus: "online",
	}

	expected := BalanceResponse{
		Error:  "ok",
		Result: expectedMap,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, balances)
}
