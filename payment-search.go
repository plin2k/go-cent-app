package go_cent_app

import (
	"encoding/json"
	"net/url"
	"time"
)

const (
	paymentSearchURL = apiURL + "/api/v1/payment/search"
)

// https://cent.app/en/merchant/api#payment-search
type PaymentSearchRequest struct {
	StartDate  time.Time // Start date
	FinishDate time.Time // End date
	ShopID     string    // Unique shop ID
}

type paymentSearchResponse struct {
	Data    []payment `json:"data"`    // Information about payments
	Success bool      `json:"success"` // Result of request
}

// Search payment.
// https://cent.app/en/merchant/api#payment-search
func (api *api) PaymentSearch(req *PaymentSearchRequest) (paymentSearchResponse, error) {
	var response paymentSearchResponse

	jsonString, err := api.request("GET", paymentSearchURL, req.constructURL())
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (req *PaymentSearchRequest) constructURL() url.Values {
	params := url.Values{}

	if req.StartDate.Year() != 1 {
		params.Add("start_date", req.StartDate.Format("2006-01-02"))
	}
	if req.FinishDate.Year() != 1 {
		params.Add("finish_date", req.FinishDate.Format("2006-01-02"))
	}
	if req.ShopID != "" {
		params.Add("shop_id", req.ShopID)
	}

	return params
}
