package coinpayments

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountInfo(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"error":"ok","result":{"username":"yograterol","merchant_id":"2455003b05d6bc669cd80834d3f39149","email":"yohangraterol92@gmail.com","public_name":"","time_joined":1449626152}}`)
	})

	client := NewClient("", "", httpClient)
	accountInfo, _, err := client.AccountInfo.Show()
	expected := AccountInfoResponse{Error: "ok",
		Result: &AccountInfo{
			Username:   "yograterol",
			MerchantID: "2455003b05d6bc669cd80834d3f39149",
			Email:      "yohangraterol92@gmail.com",
			PublicName: "",
			TimeJoined: 1449626152,
		}}
	assert.Nil(t, err)
	assert.Equal(t, expected, accountInfo)
}

func TestAccountInfoError(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"error":"XXXXXX","result": []}`)
	})

	client := NewClient("", "", httpClient)
	_, _, err := client.AccountInfo.Show()

	assert.NotNil(t, err)
}
