package network

import (
	"net/http"
	"net/url"
	"testing"
)

func setup(addSignatureToQueryParams QueryStringMapper) *Request {
	request := &Request{
		Method:            "GET",
		Header:            http.Header{},
		QueryValues:       url.Values{},
		BodyValues:        url.Values{},
		Path:              "v3/account",
		Url:               "http://localhost:9000/account",
		BaseURL:           "http://localhost:9000/",
		QueryStringMapper: addSignatureToQueryParams,
	}

	return request
}

func TestQueryString(test *testing.T) {
	request := setup(func(value string) string { return value })

	request.SetParam("id", "1223456789")
	request.SetParam("name", "Douglas")

	result := request.QueryString()

	if result != "id=1223456789&name=Douglas" {
		test.Errorf("QueryString is not returning the right value")
	}
}
