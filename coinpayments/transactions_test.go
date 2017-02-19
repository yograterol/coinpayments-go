package coinpayments

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactions(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"error":"ok","result":{"amount":"0.01000000","txn_id":"CPBB0FWI7BW9NUN2REJ3MS56KL",
		"address":"3BHLa87gpzKSemapUtY6gHhJMEE6UWqPuf","confirms_needed":"2","timeout":55800,
		"status_url":"https:\/\/www.coinpayments.net\/index.php?cmd=status&id=CPBB0FWI7BW9NUN2REJ3MS56KL&key=862f8acbd764c59d25de1ffb098d7c73",
		"qrcode_url":"https:\/\/www.coinpayments.net\/qrgen.php?id=CPBB0FWI7BW9NUN2REJ3MS56KL&key=862f8acbd764c59d25de1ffb098d7c73"}}`)
	})

	client := NewClient("", "", httpClient)
	transaction, _, err := client.Transactions.NewTransaction(&TransactionParams{
		Amount:    0.01,
		Currency1: "BTC",
		Currency2: "BTC",
	})

	expected := TransactionResponse{
		Error: "ok",
		Result: &Transaction{
			Amount:         "0.01000000",
			Address:        "3BHLa87gpzKSemapUtY6gHhJMEE6UWqPuf",
			TXNId:          "CPBB0FWI7BW9NUN2REJ3MS56KL",
			ConfirmsNeeded: "2",
			Timeout:        55800,
			StatusUrl: "https://www.coinpayments.net/index" +
				".php?cmd=status&id=CPBB0FWI7BW9NUN2REJ3MS56KL&key=862f8acbd764c59d25de1ffb098d7c73",
			QRCodeUrl: "https://www.coinpayments.net/qrgen" +
				".php?id=CPBB0FWI7BW9NUN2REJ3MS56KL&key=862f8acbd764c59d25de1ffb098d7c73",
		},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, transaction)
}
