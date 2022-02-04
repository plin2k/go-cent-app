package go_cent_app

import (
	"encoding/json"
	"net/url"
	"strconv"
)

const (
	billCreateURL = apiURL + "/api/v1/bill/create"
)

type BillCreateRequest struct {
	Amount              float64 // Payment amount
	OrderID             string  // Unique order ID. Will be sent within Postback.
	Description         string  // Description of payment
	Type                string  // Type of payment link shows how many payments it could receive. 'normal' type means that only one successful payment could be received for this link. 'multi' type means that many payments could be received with one link.
	ShopID              string  // Unique shop ID.
	CurrencyIn          string  // Currency that customer sees during payment process. If you skip this parameter in your request, the default currency of your Shop will be used during the payment process. In case where shop_id doesn't exist, customer will pay in RUB.
	Custom              string  // You can send any string value in this field and it will be returned within postback.
	PayerPaysCommission bool    // Decides who will pay fees for incoming payment.
	Name                string  // Please specify the purpose of the payment. It will be shown on the payment form.
}

type billCreateResponse struct {
	Success     bool   `json:"success"`       // Payment status
	LinkURL     string `json:"link_url"`      // Link to the page with QR-code
	LinkPageURL string `json:"link_page_url"` // Link to the payment page
	BillID      string `json:"bill_id"`       // Unique bill ID
}

func (app *app) BillCreate(req *BillCreateRequest) (billCreateResponse, error) {
	var response billCreateResponse

	jsonString, err := app.request("POST", billCreateURL, req.constructURL())
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (req *BillCreateRequest) constructURL() url.Values {
	params := url.Values{}

	params.Add("amount", strconv.FormatFloat(req.Amount, 'E', -1, 64))

	if req.OrderID != "" {
		params.Add("order_id", req.OrderID)
	}

	if req.Description != "" {
		params.Add("description", req.Description)
	}

	if req.Type != "" {
		params.Add("type", req.Type)
	}

	params.Add("shop_id", req.ShopID)

	if req.CurrencyIn != "" {
		params.Add("currency_in", req.CurrencyIn)
	}

	if req.Custom != "" {
		params.Add("custom", req.Custom)
	}

	if req.PayerPaysCommission {
		params.Add("payer_pays_commission", "1")
	} else {
		params.Add("payer_pays_commission", "0")
	}

	if req.Name != "" {
		params.Add("name", req.Name)
	}

	return params
}
