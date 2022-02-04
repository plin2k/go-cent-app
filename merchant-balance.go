package go_cent_app

import (
	"encoding/json"
	"net/url"
)

const (
	merchantBalanceURL = apiURL + "/api/v1/merchant/balance"
)

type merchantBalanceResponse struct {
	Balances []balance `json:"balances"`       // Balance information
	Success  bool      `json:"success,string"` // Request result
}

// You can request information about your current balance state using this API.
// https://cent.app/en/merchant/api#merchant-balance
func (api *api) MerchantBalance() (merchantBalanceResponse, error) {
	var response merchantBalanceResponse

	jsonString, err := api.request("GET", merchantBalanceURL, url.Values{})
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
