package go_cent_app

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const (
	payoutPersonalCreateURL = apiURL + "/api/v1/payout/personal/create"
)

// https://cent.app/en/merchant/api#personal-payout-create
type PayoutPersonalCreateRequest struct {
	Amount          float64 // Payout amount
	PayoutAccountID string  // Unique ID of payout account. Money will be send to this account
}

type payoutPersonalCreateResponse struct {
	Data    []payout `json:"data"`    // Payout information
	Success bool     `json:"success"` // Result of request
}

// In order to withdraw money you need to create a payout. The amount of payout can be split depending on payout account type. In this case you will get a list with payouts.
// https://cent.app/en/merchant/api#personal-payout-create
func (api *api) PayoutPersonalCreate(req *PayoutPersonalCreateRequest) (payoutPersonalCreateResponse, error) {
	var response payoutPersonalCreateResponse

	jsonString, err := api.request("POST", payoutPersonalCreateURL, req.constructURL())
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (req *PayoutPersonalCreateRequest) constructURL() url.Values {
	params := url.Values{}

	params.Add("amount", fmt.Sprintf("%g", req.Amount))
	params.Add("payout_account_id", req.PayoutAccountID)

	return params
}
