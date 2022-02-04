package go_cent_app

import (
	"encoding/json"
	"net/url"
	"strconv"
)

const (
	payoutRegularCreateURL = apiURL + "/api/v1/payout/regular/create"
)

type PayoutRegularCreateRequest struct {
	Amount            float64 // Payout amount
	Currency          string  // Currency
	AccountType       string  // Account type for payout
	AccountIdentifier string  // Account ID
	CardHolder        string  // Cardholder name. Only for account_type=credit_card. Should be the same as on the cad.
}

type payoutRegularCreateResponse struct {
	Data    []payout `json:"data"`    // Payout information
	Success bool     `json:"success"` // Result of request
}

func (app *app) PayoutRegularCreate(req *PayoutRegularCreateRequest) (payoutRegularCreateResponse, error) {
	var response payoutRegularCreateResponse

	jsonString, err := app.request("POST", payoutRegularCreateURL, req.constructURL())
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (req *PayoutRegularCreateRequest) constructURL() url.Values {
	params := url.Values{}

	params.Add("amount", strconv.FormatFloat(req.Amount, 'E', -1, 64))
	params.Add("currency", req.Currency)
	params.Add("account_type", req.AccountType)
	params.Add("account_identifier", req.AccountIdentifier)
	params.Add("card_holder", req.CardHolder)

	return params
}
