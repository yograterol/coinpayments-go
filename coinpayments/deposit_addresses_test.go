package coinpayments

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepositAddresses(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"error":"ok","result":{"address":"154FN6HFr9FjUQmodrR2SbtgLgGqDXuswP"}}`)
	})

	client := NewClient("", "", httpClient)
	depositAddress, _, err := client.DepositAddresses.GetNewAddreess(&DepositAddressParams{"BTC"})

	expected := DepositAddressResponse{
		Error: "ok",
		Result: &DepositAddress{
			Address: "154FN6HFr9FjUQmodrR2SbtgLgGqDXuswP",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, depositAddress)
}
