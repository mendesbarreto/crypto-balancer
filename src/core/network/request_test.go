package network

import (
	"context"
	"net/http"
	"net/url"
	"testing"
)

func setup(addSignatureToQueryParams QueryStringMapper) *Request {
	return setupWithUrl(
		"http://localhost:9000/account",
		"http://localhost:9000/",
		addSignatureToQueryParams,
	)
}

func setupWithUrl(urlString string, baseUrl string, addSignatureToQueryParams QueryStringMapper) *Request {
	request := &Request{
		Method:            "GET",
		Header:            http.Header{},
		QueryValues:       url.Values{},
		BodyValues:        url.Values{},
		Path:              "v3/account",
		Url:               urlString,
		BaseURL:           baseUrl,
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

func TestAddParam(test *testing.T) {
	request := setup(func(value string) string { return value })

	request.AddParam("id", "1223456789")
	request.AddParam("name", "Douglas")

	result := request.QueryString()

	if result != "id=1223456789&name=Douglas" {
		test.Errorf("QueryString is not returning the right value")
	}
}

func TestSetParams(test *testing.T) {
	request := setup(func(value string) string { return value })

	request.SetParams(map[string]interface{}{
		"name": "Douglas",
		"id":   "1223456789",
	})

	result := request.QueryString()

	if result != "id=1223456789&name=Douglas" {
		test.Errorf("QueryString is not returning the right value")
	}

}

func TestToHttpRequest(test *testing.T) {
	request := setup(func(value string) string { return value })
	request.SetParam("id", "1223456789")
	request.SetParam("name", "Douglas")

	httpRequest, err := request.ToHttpRequest(context.Background())

	if err != nil {
		test.Error(err)
	}

	expectedURL := "localhost:9000"

	if httpRequest.URL.Host != expectedURL {
		test.Errorf("The http request should has the host %s", expectedURL)
	}

	expectedPath := "/account"

	if httpRequest.URL.Path != expectedPath {
		test.Errorf("The http request should has the host %s", expectedURL)
	}

	expectedScheme := "http"

	if httpRequest.URL.Scheme != expectedScheme {
		test.Errorf("The http request should has the schme %s", expectedScheme)
	}

	expectedQuery := "id=1223456789&name=Douglas"

	if httpRequest.URL.RawQuery != expectedQuery {
		test.Errorf("The http request should has the query %s", expectedQuery)
	}
}
