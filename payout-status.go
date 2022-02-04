package go_cent_app

import (
	"encoding/json"
	"net/url"
)

const (
	payoutStatusURL = apiURL + "/api/v1/payout/status"
)

// https://cent.app/en/merchant/api#payout-status
type PayoutStatusRequest struct {
	ID string // Unique payout ID
}

type payoutStatusResponse struct {
	payout
}

// You can request a status of any payout operation.
// https://cent.app/en/merchant/api#payout-status
func (api *Api) PayoutStatus(req *PayoutStatusRequest) (payoutStatusResponse, error) {
	var response payoutStatusResponse

	jsonString, err := api.request("GET", payoutStatusURL, req.constructURL())
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
