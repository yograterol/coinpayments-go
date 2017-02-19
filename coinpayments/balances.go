package coinpayments

import (
	"net/http"

	"github.com/dghubble/sling"
)

type BalanceService struct {
	sling        *sling.Sling
	ApiPublicKey string
	Params       BalanceBodyParams
}

type Balance struct {
	Balance    string `json:"balance"`
	BalanceF   string `json:"balancef"`
	CoinStatus string `json:"coin_status"`
	Status     string `json:"status"`
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
