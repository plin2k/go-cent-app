package go_cent_app

import (
	"encoding/json"
	"net/url"
)

const (
	paymentStatusURL = apiURL + "/api/v1/payment/status"
)

type PaymentStatusRequest struct {
	ID string // Unique payment ID
}

type paymentStatusResponse struct {
	payment
	Success bool `json:"success"` // Status
}

func (app *app) PaymentStatus(req *PaymentStatusRequest) (paymentStatusResponse, error) {
	var response paymentStatusResponse

	jsonString, err := app.request("GET", paymentStatusURL, req.constructURL())
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (req *PaymentStatusRequest) constructURL() url.Values {
	params := url.Values{}

	params.Add("id", req.ID)

	return params
}
