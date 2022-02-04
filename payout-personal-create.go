package go_cent_app

import (
	"encoding/json"
	"net/url"
	"strconv"
)

const (
	payoutPersonalCreateURL = apiURL + "/api/v1/payout/personal/create"
)

type PayoutPersonalCreateRequest struct {
	Amount          float64 // Payout amount
	PayoutAccountID string  // Unique ID of payout account. Money will be send to this account
}

type payoutPersonalCreateResponse struct {
	Data    []payout `json:"data"`    // Payout information
	Success bool     `json:"success"` // Result of request
}

func (app *app) PayoutPersonalCreate(req *PayoutPersonalCreateRequest) (payoutPersonalCreateResponse, error) {
	var response payoutPersonalCreateResponse

	jsonString, err := app.request("POST", payoutPersonalCreateURL, req.constructURL())
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

	params.Add("amount", strconv.FormatFloat(req.Amount, 'E', -1, 64))
	params.Add("payout_account_id", req.PayoutAccountID)

	return params
}
