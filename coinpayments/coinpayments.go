package coinpayments

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"
	"net/http"

	"github.com/dghubble/sling"
	goquery "github.com/google/go-querystring/query"
)

const apiBase = "https://www.coinpayments.net/"

var apiPrivateKey string

type Client struct {
	sling       *sling.Sling
	AccountInfo *AccountInfoService
	Rates       *RateService
	Balances    *BalanceService
}

type APIParams struct {
	Version string `url:"version"`
	Command string `url:"cmd"`
	Key     string `url:"key"`
}

func NewClient(publicKey string, privateKey string, httpClient *http.Client) *Client {
	baseClient := sling.New().Client(httpClient).Base(apiBase)
	apiPrivateKey = privateKey
	return &Client{
		sling:       baseClient,
		AccountInfo: newAccountInfoService(baseClient.New(), publicKey),
		Rates:       newRateService(baseClient.New(), publicKey),
		Balances:    newBalanceService(baseClient.New(), publicKey),
	}
}

func getHMAC(payload string) string {
	mac := hmac.New(sha512.New, []byte(apiPrivateKey))
	mac.Write([]byte(payload))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func getPayload(payload interface{}) string {
	bodyForm, err := goquery.Values(payload)
	if err != nil {
		log.Fatal(err)
	}
	return bodyForm.Encode()
}
