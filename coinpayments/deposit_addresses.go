package coinpayments

import (
	"net/http"

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

type CallbackAddressParams struct {
	Currency string `url:"currency"`
	IPNUrl   string `url:"ipn_url"`
}

type CallbackAddressBodyParams struct {
	APIParams
	CallbackAddressParams
}

func newDepositAddressService(sling *sling.Sling, apiPublicKey string) *DepositAddressService {
	depositAddressService := &DepositAddressService{
		sling:        sling.Path("api.php"),
		ApiPublicKey: apiPublicKey,
	}
	depositAddressService.getParams()
	return depositAddressService
}

func (s *DepositAddressService) getHMAC(params interface{}) string {
	return getHMAC(getPayload(params))
}

func (s *DepositAddressService) GetNewAddress(depositAddressParams *DepositAddressParams) (DepositAddressResponse,
	*http.Response,
	error) {
	depositAddressResponse := new(DepositAddressResponse)
	s.Params.Currency = depositAddressParams.Currency
	resp, err := s.sling.New().Set("HMAC", s.getHMAC(s.Params)).Post(
		"api.php").BodyForm(s.Params).ReceiveSuccess(depositAddressResponse)
	return *depositAddressResponse, resp, err
}

func (s *DepositAddressService) GetNewCallbackAddress(callbackAddressParams *CallbackAddressParams) (DepositAddressResponse, *http.Response, error) {
	depositAddressResponse := new(DepositAddressResponse)
	callbackAddressBodyParams := s.getCallbackAddressParams()
	callbackAddressBodyParams.CallbackAddressParams = *callbackAddressParams
	resp, err := s.sling.New().Set("HMAC", s.getHMAC(callbackAddressBodyParams)).Post(
		"api.php").BodyForm(callbackAddressBodyParams).ReceiveSuccess(depositAddressResponse)
	return *depositAddressResponse, resp, err
}

func (s *DepositAddressService) getParams() {
	s.Params.Command = "get_deposit_address"
	s.Params.Key = s.ApiPublicKey
	s.Params.Version = "1"
}

func (s *DepositAddressService) getCallbackAddressParams() *CallbackAddressBodyParams {
	callbackAddressBodyParams := new(CallbackAddressBodyParams)
	callbackAddressBodyParams.Command = "get_callback_address"
	callbackAddressBodyParams.Key = s.ApiPublicKey
	callbackAddressBodyParams.Version = "1"
	return callbackAddressBodyParams
}
