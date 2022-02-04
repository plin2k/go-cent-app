package go_cent_app

import (
	"encoding/json"
	"net/url"
)

const (
	billPaymentsURL = apiURL + "/api/v1/bill/payments"
)

// https://cent.app/en/merchant/api#bill-payments
type BillPaymentsRequest struct {
	ID string // Unique bill ID
}

type billPaymentsResponse struct {
	Data    []payment `json:"data"`    // Information about payments
	Success bool      `json:"success"` // Result
}

// Get information about payments for one bill.
// https://cent.app/en/merchant/api#bill-payments
func (api *Api) BillPayments(req *BillPaymentsRequest) (billPaymentsResponse, error) {
	var response billPaymentsResponse

	jsonString, err := api.request("GET", billPaymentsURL, req.constructURL())
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (req *BillPaymentsRequest) constructURL() url.Values {
	params := url.Values{}

	params.Add("id", req.ID)

	return params
}
