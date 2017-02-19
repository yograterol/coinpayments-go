package coinpayments

import (
	"net/http"

	"github.com/dghubble/sling"
)

type AccountInfoService struct {
	sling        *sling.Sling
	ApiPublicKey string
	Params       APIParams
}

type AccountInfo struct {
	Username   string `json:"username"`
	MerchantID string `json:"merchant_id"`
	Email      string `json:"email"`
	PublicName string `json:"public_name"`
	TimeJoined int64  `json:"time_joined"`
}

type AccountInfoResponse struct {
	Error  string       `json:"error"`
	Result *AccountInfo `json:"result"`
}

func newAccountInfoService(sling *sling.Sling, apiPublicKey string) *AccountInfoService {
	accountInfo := &AccountInfoService{
		sling:        sling.Path("api.php"),
		ApiPublicKey: apiPublicKey,
	}
	// In each request the params are the same.
	accountInfo.getParams()
	return accountInfo
}

func (s *AccountInfoService) getHMAC() string {
	return getHMAC(getPayload(s.Params))
}

func (s *AccountInfoService) Show() (AccountInfoResponse, *http.Response, error) {
	accountInfoResponse := new(AccountInfoResponse)
	resp, err := s.sling.New().Set("HMAC", s.getHMAC()).Post(
		"api.php").BodyForm(s.Params).ReceiveSuccess(accountInfoResponse)
	return *accountInfoResponse, resp, err
}

func (s *AccountInfoService) getParams() {
	s.Params.Command = "get_basic_info"
	s.Params.Key = s.ApiPublicKey
	s.Params.Version = "1"
}
