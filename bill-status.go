package go_cent_app

import (
	"encoding/json"
	"net/url"
)

const (
	billStatusURL = apiURL + "/api/v1/bill/status"
)

// https://cent.app/en/merchant/api#bill-status
type BillStatusRequest struct {
	ID string // Unique bill ID
}

type billStatusResponse struct {
	bill
	Success bool `json:"success"` // Payment status
}

// Get bill info and status.
// https://cent.app/en/merchant/api#bill-status
func (api *Api) BillStatus(req *BillStatusRequest) (billStatusResponse, error) {
	var response billStatusResponse

	jsonString, err := api.request("GET", billStatusURL, req.constructURL())
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (req *BillStatusRequest) constructURL() url.Values {
	params := url.Values{}

	params.Add("id", req.ID)

	return params
}
