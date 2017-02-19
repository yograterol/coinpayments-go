package coinpayments

import (
	"net/http"

	"fmt"

	"github.com/dghubble/sling"
)

type DepositAddressService struct {
	sling        *sling.Sling
	ApiPublicKey string
	Params       DepositAddressBodyParams
}

type DepositAddress struct {
	Address string `json:"address"`
	PubKey  string `json:"pubkey"`
	DestTag string `json:"dest_tag"`
}

type DepositAddressResponse struct {
	Error  string          `json:"error"`
	Result *DepositAddress `json:"result"`
}

type DepositAddressParams struct {
	Currency string `url:"currency"`
}

type DepositAddressBodyParams struct {
	APIParams
	DepositAddressParams
}

func newDepositAddressService(sling *sling.Sling, apiPublicKey string) *DepositAddressService {
	depositAddressService := &DepositAddressService{
		sling:        sling.Path("api.php"),
		ApiPublicKey: apiPublicKey,
	}
	depositAddressService.getParams()
	return depositAddressService
}

func (s *DepositAddressService) getHMAC() string {
	return getHMAC(getPayload(s.Params))
}

func (s *DepositAddressService) Show(depositAddressParams *DepositAddressParams) (DepositAddressResponse, *http.Response,
	error) {
	depositAddressResponse := new(DepositAddressResponse)
	s.Params.Currency = depositAddressParams.Currency
	fmt.Println(getPayload(s.Params))
	fmt.Println(getHMAC(getPayload(s.Params)))
	resp, err := s.sling.New().Set("HMAC", s.getHMAC()).Post(
		"api.php").BodyForm(s.Params).ReceiveSuccess(depositAddressResponse)
	return *depositAddressResponse, resp, err
}

func (s *DepositAddressService) getParams() {
	s.Params.Command = "get_deposit_address"
	s.Params.Key = s.ApiPublicKey
	s.Params.Version = "1"
}
