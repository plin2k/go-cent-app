package go_cent_app_test

import (
	"fmt"
	cent_app "github.com/plin2k/go-cent-app"
	"time"
)

func ExampleApi_BillCreate() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	request := &cent_app.BillCreateRequest{
		Amount:              100.0,
		CurrencyIn:          cent_app.CurrencyUSD,
		ShopID:              "LXZv3R7Q8B",
		PayerPaysCommission: true,
		OrderID:             fmt.Sprintf("%d", time.Now().UnixMicro()),
	}

	res, err := api.BillCreate(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	//	Success: "true",
	//	LinkURL: "https://cent.app/link/GkLWvKx3",
	//	LinkPageURL: "https://cent.app/transfer/GkLWvKx3",
	//	BillID: "GkLWvKx3"
	//}
}

func ExampleApi_BillToggleActivity() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	request := &cent_app.BillToggleActivityRequest{
		ID:     "LXZv3R7Q8B",
		Active: true,
	}

	res, err := api.BillToggleActivity(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	//	ID: "LXZv3R7Q5B",
	//	Activity: "false",
	//	Status: "NEW",
	//	Type: "MULTI",
	//	Amount: 100.05,
	//	CurrencyIn: "USD",
	//	CreatedAt: "2020-11-11 14:46:20",
	//	Success: true,
	//}
}

func ExampleApi_BillPayments() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	request := &cent_app.BillPaymentsRequest{
		ID: "LXZv3R7Q8B",
	}

	res, err := api.BillPayments(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	//Data: [
	//			{
	//				ID: "LXZv3R7Q8B",
	// 				BillID: "do5G93m",
	//				Status: "NEW",
	//				Amount: 0,
	//				AmountReceived: 100,
	//				FromCard: "676454******7272",
	//				CurrencyIn: "RUB",
	//				CreatedAt: "2020-11-03 06:43:36",
	//			},
	//		],
	//	Success: true
	//}
}

func ExampleApi_BillSearch() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	request := &cent_app.BillSearchRequest{
		ShopID:    "LXZv3R7Q8B",
		StartDate: time.Now().Add(time.Hour - 5),
	}

	res, err := api.BillSearch(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	//Data: [
	//			{
	//				ID: "LXZv3R7Q5B",
	//				Status: "NEW",
	//				Type: "MULTI",
	//				Amount: 100.05,
	//				CurrencyIn: "RUB",
	//				CreatedAt: "2020-11-11 14:46:20",
	//			},
	//		],
	//	Success: true
	//}
}

func ExampleApi_BillStatus() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	request := &cent_app.BillStatusRequest{
		ID: "LXZv3R7Q8B",
	}

	res, err := api.BillStatus(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	//	ID: "LXZv3R7Q5B",
	//	Status: "NEW",
	//	Type: "MULTI",
	//	Amount: 100.05,
	//	CurrencyIn: "RUB",
	//	CreatedAt: "2020-11-11 14:46:20",
	//	Success: true
	//}
}

func ExampleApi_PaymentSearch() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	request := &cent_app.PaymentSearchRequest{
		ShopID:    "LXZv3R7Q8B",
		StartDate: time.Now().Add(time.Hour - 5),
	}

	res, err := api.PaymentSearch(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	// Data: [
	//				{
	//					ID: "LXZv3R7Q8B",
	// 					BillID: "do5G93m",
	//					Status: "NEW",
	//					Amount: 0,
	//					AmountReceived: 100,
	//					FromCard: "676454******7272",
	//					CurrencyIn: "RUB",
	//					CreatedAt: "2020-11-03 06:43:36",
	//				},
	// 			],
	//	Success: true
	//}
}

func ExampleApi_PaymentStatus() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	request := &cent_app.PaymentStatusRequest{
		ID: "LXZv3R7Q8B",
	}

	res, err := api.PaymentStatus(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	//	ID: "LXZv3R7Q8B",
	// 	BillID: "do5G93m",
	//	Status: "NEW",
	//	Amount: 0,
	//	AmountReceived: 100,
	//	FromCard: "676454******7272",
	//	CurrencyIn: "RUB",
	//	CreatedAt: "2020-11-03 06:43:36",
	//	Success: true
	//}
}

func ExampleApi_MerchantBalance() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	res, err := api.MerchantBalance()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	//	Success: "true",
	//	Balances: [
	//					{
	//						Currency: "RUB",
	//						BalanceAvailable: "93988.37000000",
	//						BalanceLocked: "0.00000000",
	//						BalanceHold: "0.00000000"
	//					}
	//				],
	//}
}

func ExampleApi_PayoutPersonalCreate() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	request := &cent_app.PayoutPersonalCreateRequest{
		Amount:          100.56,
		PayoutAccountID: "LXZv3R7Q8B",
	}

	res, err := api.PayoutPersonalCreate(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	//	Data: [
	//				{
	//					ID: "gY37jYr2b6",
	//					Status: "MODERATING",
	//					Amount: 1600,
	//					Commission: 100,
	//					AccountIdentifier: "676454******7272",
	//					Currency: "RUB",
	//					CreatedAt: "2020-10-19 17:00:00"
	//				},
	//			],
	//	Success: true
	//}
}

func ExampleApi_PayoutRegularCreate() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	request := &cent_app.PayoutRegularCreateRequest{
		Amount:            100.56,
		Currency:          cent_app.CurrencyUSD,
		AccountType:       "credit_card",
		AccountIdentifier: "6712544488737272",
		CardHolder:        "JOHN DOE",
	}

	res, err := api.PayoutRegularCreate(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	//	Data: [
	//				{
	//					ID: "gY37jYr2b6",
	//					Status: "MODERATING",
	//					Amount: 1600,
	//					Commission: 100,
	//					AccountIdentifier: "676454******7272",
	//					Currency: "RUB",
	//					CreatedAt: "2020-10-19 17:00:00"
	//				},
	//			],
	//	Success: true
	//}
}

func ExampleApi_PayoutSearch() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	request := &cent_app.PayoutSearchRequest{
		StartDate:  time.Now().Add(time.Hour - 5),
		FinishDate: time.Now(),
	}

	res, err := api.PayoutSearch(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	//	Data: [
	//				{
	//					ID: "gY37jYr2b6",
	//					Status: "MODERATING",
	//					Amount: 1600,
	//					Commission: 100,
	//					AccountIdentifier: "676454******7272",
	//					Currency: "RUB",
	//					CreatedAt: "2020-10-19 17:00:00"
	//				},
	//			],
	//	Success: true
	//}
}

func ExampleApi_PayoutStatus() {
	api := cent_app.New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	request := &cent_app.PayoutStatusRequest{
		ID: "LXZv3R7Q8B",
	}

	res, err := api.PayoutStatus(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Output:
	//{
	//	ID: "lGPmp4mYeE",
	//	Status: "MODERATING",
	//	Amount: 1600,
	//	Commission: 100,
	//	AccountIdentifier: "676454******7272",
	//	Currency: "RUB",
	//	CreatedAt: "2020-10-19 17:00:00",
	//	Success: true,
	//}
}
