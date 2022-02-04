package go_cent_app

import (
	"encoding/json"
	"net/url"
)

const (
	billToggleActivityURL = apiURL + "/api/v1/bill/toggle_activity"
)

// https://cent.app/en/merchant/api#bill-toggle
type BillToggleActivityRequest struct {
	ID     string // Unique bill id
	Active bool   // false - deactivate bill true - activate bill
}

type billToggleActivityResponse struct {
	bill
	Active  bool `json:"active"`  // Bill activity flag
	Success bool `json:"success"` // This flag indicates status of request
}

// You can deactivate and activate bills using this APO.
// https://cent.app/en/merchant/api#bill-toggle
func (api *Api) BillToggleActivity(req *BillToggleActivityRequest) (billToggleActivityResponse, error) {
	var response billToggleActivityResponse

	jsonString, err := api.request("POST", billToggleActivityURL, req.constructURL())
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (req *BillToggleActivityRequest) constructURL() url.Values {
	params := url.Values{}

	params.Add("id", req.ID)

	if req.Active {
		params.Add("active", "1")
	} else {
		params.Add("active", "0")
	}

	return params
}
