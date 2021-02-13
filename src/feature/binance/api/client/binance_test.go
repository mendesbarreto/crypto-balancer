package client

import (
	"bytes"
	"crypto-balancer/src/core/environment"
	"crypto-balancer/src/core/network"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func setupBinance() (binanceClient *BinanceClient) {
	environment.LoadVariables()
	return NewBinanceClient()
}

func setupRequest(
	requestMethod string,
	requestEndpoint string,
	sectionType SectionApiKeyType) (*BinanceClient, *network.Request) {

	binanceClient := setupBinance()
	request := binanceClient.NewRequest(requestMethod, requestEndpoint, sectionType)

	return binanceClient, request
}

func TestNewBinanceClient(test *testing.T) {
	binanceClient := setupBinance()

	if binanceClient.ApiKey != environment.BinanceApiKey() {
		test.Error("The APi key created by method is not the same as the environment")
	}

	if binanceClient.BaseURL != environment.BinanceApiBaseUrl() {
		test.Error("The BaseURL created by method is not the same as the environment")
	}

	if binanceClient.UserAgent != environment.UserAgent() {
		test.Error("The UserAgent created by method is not the same as the environment")
	}

	if binanceClient.SecretKey != environment.BinanceAPiSecretKey() {
		test.Error("The UserAgent created by method is not the same as the environment")
	}

	if binanceClient.HTTPClient != http.DefaultClient {
		test.Error("The biance client http should be the same as the http default")
	}

}

func TestAddRequiredParams(test *testing.T) {
	binanceClient := setupBinance()

	nowFunction := func() time.Time {
		return time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	addSignatureToQueryParams := AddRequiredParams(binanceClient.ApiKey, SectionAPIKey, nowFunction)
	queryValues := url.Values{}
	queryValues.Set("id", "1234")
	queryValues.Set("name", "Douglas")

	result := addSignatureToQueryParams(queryValues.Encode())
	expectedResult := "id=1234&name=Douglas&timestamp=1577836800000"

	if result != expectedResult {
		test.Error("When the section is SectionAPIKey The result should be the same as query values added as parameter")
	}

	addSignatureToQueryParams = AddRequiredParams(binanceClient.ApiKey, SectionNone, nowFunction)

	result = addSignatureToQueryParams(queryValues.Encode())
	expectedResult = "id=1234&name=Douglas"

	if result != expectedResult {
		test.Error("When the section is SectionNone The result should be the same as query values added as parameter")
	}

	addSignatureToQueryParams = AddRequiredParams(binanceClient.ApiKey, SectionSigned, nowFunction)
	result = addSignatureToQueryParams(queryValues.Encode())
	expectedResult = "id=1234&name=Douglas&timestamp=1577836800000&signature=d829a0dd4821de0ebe70d12e48bc1e0dd687d5f0af67c795fd0e0abfc9f89f86"

	if len(result) == 0 {
		test.Error("The result should not empty")
	}

	if result != expectedResult {
		test.Error("The result should add signature on params and be the same as the expectedResult string")
	}
}

func TestNewRequest(test *testing.T) {
	method := http.MethodGet
	endpoint := "v2/hello"
	sectionType := SectionAPIKey

	binanceClient, request := setupRequest(method, endpoint, sectionType)

	if request.BaseURL != binanceClient.BaseURL {
		test.Error("The base url should be the same as binance client")
	}

	if request.BaseURL != binanceClient.BaseURL {
		test.Error("The base url should be the same as binance client")
	}

	if request.Method != method {
		test.Error("The method should be the same as NewRequest method parameter")
	}

	if request.Header == nil {
		test.Error("header should not be null")
	}

	if request.QueryValues == nil {
		test.Error("query values should not be null")
	}

	if request.BodyValues == nil {
		test.Error("body values should not be null")
	}

	if request.Path != endpoint {
		test.Error("path should be the same as endpoint")
	}

	if request.Url != fmt.Sprintf("%s%s", binanceClient.BaseURL, endpoint) {
		test.Error("URL should be the base URL + endpoint")
	}

	if len(request.Header.Get("X-Mbx-Apikey")) == 0 {
		test.Error("The Header value should not be empty")
	}

	if request.Header.Get("X-Mbx-Apikey") != os.Getenv("BINANCE_API_KEY") {
		test.Error("When a request is created and with section type APIKey should add a header: X-MBX-APIKEY")
	}

	if len(request.QueryValues.Get("timestamp")) > 0 {
		test.Error("When a request is created and with section type APIKey should add timestamp as query parameter")
	}

	_, request = setupRequest(method, endpoint, SectionNone)

	if len(request.Header.Get("X-Mbx-Apikey")) > 0 {
		test.Error("When request section is none, should not apply APIKey on header")
	}

	if len(request.QueryValues.Get("timestamp")) > 0 {
		test.Error("When request section is none, should not apply timestamp query parameter")
	}

	_, request = setupRequest(method, endpoint, SectionSigned)

	if len(request.Header.Get("X-Mbx-Apikey")) == 0 {
		test.Error("When the request section is Signed it should add an api key header")
	}

	queryString := request.QueryString()
	if strings.Contains(queryString, "timestamp=") &&
		strings.Contains(queryString, "signature=496f06959bcfff47746396f7bcaf6353662a78dba950762d0c8ee671b0b6bc40") {
		test.Error("When the request section is Signed it should add timestamp and signature parameters")
	}
}

func TestHttpResponseHandler(test *testing.T) {
	response := &http.Response{}

	response.StatusCode = 200
	response.Body = ioutil.NopCloser(bytes.NewReader([]byte("OK")))

	data, err := HttpResponseHandler(response)

	if err != nil {
		test.Error("When status code is between 200 and 399 " +
			"and the response body is a strig should not raise an error")
	}

	if string(data) != "OK" {
		test.Error("the content of the body should be OK")
	}

	response.StatusCode = 400
	response.Body = ioutil.NopCloser(bytes.NewReader([]byte("{\"code\": 1000, \"msg\": \"bad request\"}")))

	data, err = HttpResponseHandler(response)

	if data != nil {
		test.Error("When status code is greater than 399 data should be null")
	}

	if err == nil {
		test.Error("When status code is greater than 399 data should not be null")
	}

	if !network.IsAPIError(err) {
		test.Error("The error should should of type APIError")
	}

	apiError := err.(*network.APIError)

	if apiError.Message != "bad request" {
		test.Error("The error should should has the message bad request")
	}

	if apiError.Code != 1000 {
		test.Error("The error should should has the code 1000")
	}

	if apiError.Error() != "<APIError> code=1000, msg=bad request" {
		test.Error("The error should should has the code 1000")
	}
}
