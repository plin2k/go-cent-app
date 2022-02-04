package go_cent_app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Api struct {
	token       string
	bearerToken string
}

const (
	apiURL = "https://cent.app"

	CurrencyRUB = "RUB"
	CurrencyUSD = "USD"
	CurrencyEUR = "EUR"

	TypeNormal = "NORMAL"
	TypeMulti  = "MULTI"
)

var errors = make(map[int]map[string]string)

type message struct {
	Message string `json:"message"`
}

type payment struct {
	ID             string  `json:"id"`              // Unique payment ID
	Status         string  `json:"status"`          // Status of payment
	Amount         float64 `json:"amount"`          // Total payment amount
	AmountReceived float64 `json:"amount_received"` // Received amount
	FromCard       string  `json:"from_card"`       // Payer's card
	CurrencyIn     string  `json:"currency_in"`     // Payment currency
	CreatedAt      Time    `json:"created_at"`      // Creation date and time
}

type bill struct {
	ID         string  `json:"id"`          // Unique bill id
	Status     string  `json:"status"`      // Bill status
	Amount     float64 `json:"amount"`      // Bill amount
	Type       string  `json:"type"`        // Type of bill. NORMAL is for onetime payments and MULTI is for infinity number of payments
	CurrencyIn string  `json:"currency_in"` // Currency
	CreatedAt  Time    `json:"created_at"`  // Bill creation date and time
}

type balance struct {
	Currency         string  `json:"currency"`                 // Currency of balance
	BalanceAvailable float64 `json:"balance_available,string"` // Available balance
	BalanceLocked    float64 `json:"balance_locked,string"`    // Locked balance for payout
	BalanceHold      float64 `json:"balance_hold,string"`      // Fees
}

type payout struct {
	ID                string  `json:"id"`                 // Unique ID of payment account
	Status            string  `json:"status"`             // Payout status
	Amount            float64 `json:"amount"`             // Payout amount
	Commission        float64 `json:"commission"`         // Fees
	AccountIdentifier string  `json:"account_identifier"` // Unique ID of payout account. Money will be send to this account
	Currency          string  `json:"currency"`           // Payout currency
	CreatedAt         Time    `json:"created_at"`         // Date and time
}

type Time struct {
	time.Time
}

func New(token string) *Api {

	errors[400] = map[string]string{
		"api:error.too_many_payments":             "You are trying to get too many payments in one request",
		"api:error.too_many_payouts":              "You are trying to request to many payouts. Maximum payouts for one request is 1000",
		"api:error.shop_not_found":                "Shop ID can't be found. This error only appears if you have sent the shop ID in the request",
		"api:error.access_denied":                 "Merchant doesn't have access to the shop",
		"api:error.too_many_bills":                "You are trying to request too many bills at once",
		"api:error.payout_account_not_found":      "Payout account is not found",
		"api:error.payout_account_banned":         "Payout account is blocked",
		"api:error.daily_payout_limit_exceeded":   "Exceeded daily limit",
		"api:error.monthly_payout_limit_exceeded": "Exceeded monthly limit",
		"api:error.balance_not_enough":            "Not enough balance for payout",
		"api:error.direction_not_available":       "Account is unavailable for payout",
		"api:error.merchant_not_verified":         "Merchant doesn't have Verified status",
	}
	errors[401] = map[string]string{
		"Unauthenticated": "Invalid API Token",
	}
	errors[403] = map[string]string{
		"This action is unauthorized":               "Merchant doesn't have access for API. Please contact Support Team",
		"api:error.invalid_amount":                  "Invalid amount",
		"api:error.merchant_banned":                 "Merchant is blocked",
		"api:error.merchant_not_found":              "Merchant is not found in the System",
		"api:error.merchant_subscription_inactive":  "Subscription is not active on your account",
		"api:error.merchant_subscription_not_found": "You account doesn't have subscription",
		"api:error.shop_not_found":                  "Shop is not found in the System",
		"api:error.shop_not_enabled":                "Merchant is deactivated",
		"api:error.access_denied":                   "Merchant doesn't have access to the shop",
		"api:error.rate-not-found":                  "Invalid currency or country",
		"api:error.bill_is_finished":                "Bill is paid in case of NORMAL bill",
	}
	errors[422] = map[string]string{
		"":                            "Invalid data in request",
		"api:error.bill_not_found":    "Bill doesn't exist",
		"api:error.payout_not_found":  "Payout is not found",
		"api:error.payment_not_found": "Payment is not found",
	}
	errors[500] = map[string]string{
		"api:error.general_error": "Internal error",
	}

	return &Api{
		token:       token,
		bearerToken: "Bearer " + token,
	}
}

func (api *Api) request(method, url string, params url.Values) (string, error) {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 60,
	}

	if method == "GET" {
		url += "?" + params.Encode()
	}

	req, err := http.NewRequest(method, url, strings.NewReader(params.Encode()))
	if err != nil {
		return "", fmt.Errorf("Got error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Set("Authorization", api.bearerToken)

	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Got error %s", err.Error())
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	jsonString := string(body)

	return jsonString, errorHandle(jsonString, response.StatusCode)
}

func errorHandle(jsonString string, code int) error {

	if val, exists := errors[code]; exists {
		var msg message

		err := json.Unmarshal([]byte(jsonString), &msg)
		if err != nil {
			return err
		}

		m, ok := val[msg.Message]
		if ok {
			return fmt.Errorf("%s", m)
		}

		return fmt.Errorf("%s", msg.Message)
	}

	return nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	ret, err := time.Parse("2006-01-02 15:04:05", strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}

	*t = Time{ret}

	return nil
}
