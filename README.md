#CENT.APP - GO Package 

###Official documentation - https://cent.app/en/merchant/api
The package is fully tested and works on API version v1

##Package usage example
For each request, a separate structure and a separate api method are created.
```GO
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
```
