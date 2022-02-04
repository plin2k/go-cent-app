package go_cent_app

import (
	"encoding/json"
	"net/url"
	"time"
)

const (
	payoutSearchURL = apiURL + "/api/v1/payout/search"
)

type PayoutSearchRequest struct {
	StartDate  time.Time // Start date for search
	FinishDate time.Time // End date for search
}

type payoutSearchResponse struct {
	Data    []payout `json:"data"`    // Payout information
	Success bool     `json:"success"` // Result of request
}

func (app *app) PayoutSearch(req *PayoutSearchRequest) (payoutSearchResponse, error) {
	var response payoutSearchResponse

	jsonString, err := app.request("GET", payoutSearchURL, req.constructURL())
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (req *PayoutSearchRequest) constructURL() url.Values {
	params := url.Values{}

	if req.StartDate.Year() != 1 {
		params.Add("start_date", req.StartDate.Format("2006-01-02"))
	}
	if req.FinishDate.Year() != 1 {
		params.Add("finish_date", req.FinishDate.Format("2006-01-02"))
	}

	return params
}
