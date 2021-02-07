package client

import (
	"crypto-balancer/src/core/environment"
	"crypto-balancer/src/core/network"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
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

	if binanceClient.Logger == nil {
		test.Error("The the logger should not be null")
	}

	if binanceClient.HTTPClient != http.DefaultClient {
		test.Error("The biance client http should be the same as the http default")
	}

}

func TestAddSignatureToQueryParams(test *testing.T) {
	binanceClient := setupBinance()
	addSignatureToQueryParams := AddSignatureToQueryParams(binanceClient.ApiKey, SectionAPIKey)
	queryValues := url.Values{}

	queryValues.Set("id", "1234")
	queryValues.Set("name", "Douglas")

	result := addSignatureToQueryParams(queryValues.Encode())
	expectedResult := queryValues.Encode()

	if result != expectedResult {
		test.Error("When the section is SectionAPIKey The result should be the same as query values added as parameter")
	}

	addSignatureToQueryParams = AddSignatureToQueryParams(binanceClient.ApiKey, SectionNone)

	result = addSignatureToQueryParams(queryValues.Encode())
	expectedResult = queryValues.Encode()

	if result != expectedResult {
		test.Error("When the section is SectionNone The result should be the same as query values added as parameter")
	}

	addSignatureToQueryParams = AddSignatureToQueryParams(binanceClient.ApiKey, SectionSigned)
	result = addSignatureToQueryParams(queryValues.Encode())
	expectedResult = "id=1234&name=Douglas&signature=6919eb2b0c6a227be61bd50e43ac810a8257749976dab3a01bae8de987460dd5"

	if len(result) == 0 {
		test.Error("The result should not empty")
	}

	if result != expectedResult {
		test.Error("The result should add signature on params and be the same as the expectedResult string")
	}
}

func TestCreateUrl(test *testing.T) {
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

	if len(request.Header.Get("X-Mbx-Apikey")) == 0 {
		test.Error("When the request section is Signed it should add an api key header")
	}
}
