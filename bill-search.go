package go_cent_app

import (
	"encoding/json"
	"net/url"
	"time"
)

const (
	billSearchURL = apiURL + "/api/v1/bill/search"
)

// https://cent.app/en/merchant/api#bill-search
type BillSearchRequest struct {
	StartDate  time.Time // Start date
	FinishDate time.Time // End date
	ShopID     string    // Unique shop ID
}

type billSearchResponse struct {
	Data    []bill `json:"data"`    // Information about bill
	Success bool   `json:"success"` // Result
}

// Search by bills.
// https://cent.app/en/merchant/api#bill-search
func (api *api) BillSearch(req *BillSearchRequest) (billSearchResponse, error) {
	var response billSearchResponse

	jsonString, err := api.request("GET", billSearchURL, req.constructURL())
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (req *BillSearchRequest) constructURL() url.Values {
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
