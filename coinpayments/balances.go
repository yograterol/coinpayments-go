package coinpayments

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

type BalanceService struct {
	sling        *sling.Sling
	ApiPublicKey string
	Params       BalanceBodyParams
}

// Balances struct for JSON Response
// `Balance` field is commented because the
// response has a mix of type for the field "balance" (satoshi)
type Balance struct {
	// Balance    uint64 `json:"balance,string"`
	BalanceF   string `json:"balancef"`
	CoinStatus string `json:"coin_status"`
	Status     string `json:"status"`
}

func (b *Balance) GetSatoshi() uint64 {
	balance, err := strconv.ParseFloat(b.BalanceF, 64)
	if err != nil {
		log.Fatal(err)
	}
	return uint64(balance * 100000000.00)
}

type BalanceResponse struct {
	Error  string             `json:"error"`
	Result map[string]Balance `json:"result"`
}

type BalanceParams struct {
	All uint8 `url:"all"`
}

type BalanceBodyParams struct {
	APIParams
	BalanceParams
}

func newBalanceService(sling *sling.Sling, apiPublicKey string) *BalanceService {
	balanceService := &BalanceService{
		sling:        sling.Path("api.php"),
		ApiPublicKey: apiPublicKey,
	}
	// In each request the params are the same.
	balanceService.getParams()
	return balanceService
}

func (s *BalanceService) getHMAC() string {
	return getHMAC(getPayload(s.Params))
}

func (s *BalanceService) Show(balanceParams *BalanceParams) (BalanceResponse, *http.Response, error) {
	balanceResponse := new(BalanceResponse)
	s.Params.All = balanceParams.All
	resp, err := s.sling.New().Set("HMAC", s.getHMAC()).Post(
		"api.php").BodyForm(s.Params).ReceiveSuccess(balanceResponse)
	return *balanceResponse, resp, err
}

func (s *BalanceService) getParams() {
	s.Params.Command = "balances"
	s.Params.Key = s.ApiPublicKey
	s.Params.Version = "1"
}
