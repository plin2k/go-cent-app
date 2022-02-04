package go_cent_app

import (
	"encoding/json"
	"net/url"
)

const (
	merchantBalanceURL = apiURL + "/api/v1/merchant/balance"
)

type merchantBalanceResponse struct {
	Balances []balance `json:"balances"` // Balance information
	Success  bool      `json:"success"`  // Request result
}

func (app *app) MerchantBalance() (merchantBalanceResponse, error) {
	var response merchantBalanceResponse

	jsonString, err := app.request("GET", merchantBalanceURL, url.Values{})
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
