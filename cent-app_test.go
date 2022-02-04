package go_cent_app

import (
	"net/url"
	"testing"
)

func TestApi_request(t *testing.T) {

	api_test := New("72|oBCB7Z3SmUm1gvkpEdRcSR2q1ERHpG4vD3DNBmuT")

	type args struct {
		method string
		url    string
		params url.Values
	}
	tests := []struct {
		name    string
		fields  *Api
		args    args
		want    string
		wantErr bool
	}{
		{
			"Bill Create",
			api_test,
			args{
				method: "POST",
				url:    billCreateURL,
				params: url.Values{
					"amount":  {"100.54"},
					"shop_id": {"LXZv3R7Q8B"},
				},
			},
			"{\"message\":\"Unauthenticated.\"}",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &Api{
				token:       tt.fields.token,
				bearerToken: tt.fields.bearerToken,
			}
			got, err := api.request(tt.args.method, tt.args.url, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("request() got = %v, want %v", got, tt.want)
			}
		})
	}
}
