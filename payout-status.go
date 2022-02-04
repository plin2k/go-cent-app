package go_cent_app

import (
	"encoding/json"
	"net/url"
)

const (
	payoutStatusURL = apiURL + "/api/v1/payout/status"
)

type PayoutStatusRequest struct {
	ID string // Unique payout ID
}

type payoutStatusResponse struct {
	payout
}

func (app *app) PayoutStatus(req *PayoutStatusRequest) (payoutStatusResponse, error) {
	var response payoutStatusResponse

	jsonString, err := app.request("GET", payoutStatusURL, req.constructURL())
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (req *PayoutStatusRequest) constructURL() url.Values {
	params := url.Values{}

	params.Add("id", req.ID)

	return params
}
