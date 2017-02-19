package coinpayments

import (
	"net/http"

	"github.com/dghubble/sling"
)

type BalanceService struct {
	sling        *sling.Sling
	ApiPublicKey string
	Params       APIParams
}

type Balance struct {
	Balance  string `json:"balance"`
	BalanceF string `json:"balancef"`
}

type BalanceResponse struct {
	Error  string             `json:"error"`
	Result map[string]Balance `json:"result"`
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

func (s *BalanceService) Show() (BalanceResponse, *http.Response, error) {
	balanceResponse := new(BalanceResponse)
	resp, err := s.sling.New().Set("HMAC", s.getHMAC()).Post(
		"api.php").BodyForm(s.Params).ReceiveSuccess(balanceResponse)
	return *balanceResponse, resp, err
}

func (s *BalanceService) getParams() {
	s.Params.Command = "balances"
	s.Params.Key = s.ApiPublicKey
	s.Params.Version = "1"
}
