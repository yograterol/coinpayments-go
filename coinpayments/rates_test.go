package coinpayments

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRates(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
		"error":"ok",
		"result":{"BTC":{"is_fiat":0,"rate_btc":"1.000000000000000000000000","last_update":"1375473661",
		"tx_fee":"0.00010000","status":"online","name":"Bitcoin","confirms":"2","can_convert":1,
		"capabilities":["wallet","transfers","convert"]},"LTC":{"is_fiat":0,
		"rate_btc":"0.003598757777777800000000","last_update":"1487464863","tx_fee":"0.00100000",
		"status":"online","name":"Litecoin","confirms":"3","can_convert":1,
		"capabilities":["wallet","transfers","convert"]}}
		}`)
	})

	client := NewClient("", "", httpClient)
	rates, _, err := client.Rates.Show(&RateParams{Short: 1, Accepted: 0})
	expectedMap := map[string]RateInfo{}
	expectedMap["BTC"] = RateInfo{
		IsFiat:         0,
		RateBTC:        "1.000000000000000000000000",
		LastUpdate:     "1375473661",
		TransactionFee: "0.00010000",
		Name:           "Bitcoin",
		Capabilities:   []string{"wallet", "transfers", "convert"},
		Confirms:       "2",
		CanConvert:     1,
		Status:         "online",
	}
	expectedMap["LTC"] = RateInfo{
		IsFiat:         0,
		RateBTC:        "0.003598757777777800000000",
		LastUpdate:     "1487464863",
		TransactionFee: "0.00100000",
		Name:           "Litecoin",
		Capabilities:   []string{"wallet", "transfers", "convert"},
		Confirms:       "3",
		CanConvert:     1,
		Status:         "online",
	}
	expected := RateResponse{
		Error:  "ok",
		Result: expectedMap,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, rates)
}

func TestRatesShort(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
		"error":"ok",
		"result":{
		  "BTC":{"is_fiat":0,"rate_btc":"1.000000000000000000000000","last_update":"1375473661",
		  "tx_fee":"0.00010000","status":"online"},"LTC":{"is_fiat":0,"rate_btc":"0.003593641111111100000000",
		  "last_update":"1487463962","tx_fee":"0.00100000","status":"online"}
		  }
		}`)
	})

	client := NewClient("", "", httpClient)
	rates, _, err := client.Rates.Show(&RateParams{Short: 1, Accepted: 0})
	expectedMap := map[string]RateInfo{}
	expectedMap["BTC"] = RateInfo{
		IsFiat:         0,
		RateBTC:        "1.000000000000000000000000",
		LastUpdate:     "1375473661",
		TransactionFee: "0.00010000",
		Status:         "online",
	}
	expectedMap["LTC"] = RateInfo{
		IsFiat:         0,
		RateBTC:        "0.003593641111111100000000",
		LastUpdate:     "1487463962",
		TransactionFee: "0.00100000",
		Status:         "online",
	}
	expected := RateResponse{
		Error:  "ok",
		Result: expectedMap,
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, rates)
}
