package client

import (
	"bytes"
	"context"
	"crypto-balancer/src/core/datetime"
	"crypto-balancer/src/core/environment"
	crytpoLog "crypto-balancer/src/core/log"
	"crypto-balancer/src/core/network"
	"crypto-balancer/src/feature/binance/signature"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type ExecuteFunc func(req *http.Request) (*http.Response, error)

type WorkingType string

const (
	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"

	WorkingTypeMarkPrice     WorkingType = "MARK_PRICE"
	WorkingTypeContractPrice WorkingType = "CONTRACT_PRICE"
)

type BinanceClient struct {
	ApiKey      string
	SecretKey   string
	BaseURL     string
	UserAgent   string
	HTTPClient  *http.Client
	TimesOffset int64
}

func NewBinanceClient() *BinanceClient {
	return &BinanceClient{
		ApiKey:     environment.BinanceApiKey(),
		SecretKey:  environment.BinanceAPiSecretKey(),
		BaseURL:    environment.BinanceApiBaseUrl(),
		UserAgent:  environment.UserAgent(),
		HTTPClient: http.DefaultClient,
	}
}

type SectionApiKeyType int

const (
	SectionNone SectionApiKeyType = iota
	SectionAPIKey
	SectionSigned
)

func (client *BinanceClient) NewHeader(sectionType SectionApiKeyType) http.Header {
	header := http.Header{}

	header.Set("Content-Type", "application/json")

	if sectionType == SectionAPIKey || sectionType == SectionSigned {
		header.Set("X-MBX-APIKEY", client.ApiKey)
	}

	return header
}

func (client *BinanceClient) createURL(endpoint string) string {
	return fmt.Sprintf("%s%s", client.BaseURL, endpoint)
}

func (client *BinanceClient) NewRequest(method string, endpoint string, sectionType SectionApiKeyType) *network.Request {
	request := &network.Request{
		Method:            method,
		Header:            client.NewHeader(sectionType),
		QueryValues:       url.Values{},
		BodyValues:        url.Values{},
		Body:              &bytes.Buffer{},
		Path:              endpoint,
		Url:               client.createURL(endpoint),
		BaseURL:           client.BaseURL,
		QueryStringMapper: AddRequiredParams(client.SecretKey, sectionType, time.Now),
	}

	return request
}

func (client *BinanceClient) Call(ctx context.Context, request *network.Request, sectionType SectionApiKeyType) (data []byte, err error) {
	httpRequest, err := request.ToHttpRequest(ctx)

	if err != nil {
		return nil, err
	}

	response, err := client.HTTPClient.Do(httpRequest)

	if err != nil {
		return nil, err
	}

	crytpoLog.LogDebug("response: %#v", response)

	defer func() {
		closeError := response.Body.Close()
		if err == nil && closeError != nil {
			err = closeError
		}
	}()

	return HttpResponseHandler(response)
}

func HttpResponseHandler(response *http.Response) (data []byte, err error) {
	crytpoLog.LogDebug("status code: %d", response.StatusCode)

	data, err = ioutil.ReadAll(response.Body)

	if response.StatusCode >= 400 {
		binanceApiError := new(network.APIError)
		jsonError := json.Unmarshal(data, binanceApiError)

		if jsonError != nil {
			crytpoLog.LogError("failed to unmarshal json: %s", jsonError)
		}

		return nil, binanceApiError
	}

	return data, err
}

func AddRequiredParams(secretKey string, sectionType SectionApiKeyType, now func() time.Time) func(value string) string {
	if sectionType != SectionAPIKey && sectionType != SectionSigned {
		return func(value string) string {
			return value
		}
	}

	return func(value string) string {
		queryString := ""

		if value != "" {
			queryString = fmt.Sprintf("%s&", value)
		}

		timestamp := fmt.Sprintf("timestamp=%d", datetime.Timestamp(now))

		signatureString := ""

		queryString = fmt.Sprintf("%s%s", queryString, timestamp)

		if sectionType == SectionSigned {
			mac := signature.Generate(secretKey)

			if _, err := mac.Write([]byte(queryString)); err != nil {
				log.Fatal(err)
			}

			signatureHex := hex.EncodeToString(mac.Sum(nil))
			signatureString = fmt.Sprintf("&%s=%s", signatureKey, signatureHex)
		}

		return fmt.Sprintf("%s%s", queryString, signatureString)
	}
}

func (client *BinanceClient) NewAccountRequestBuilder() *GetAccountBuilder {
	return &GetAccountBuilder{
		client: client,
	}
}

func (client *BinanceClient) NewOrderRequestBuilder() *OrderBuilder {
	return &OrderBuilder{
		client: client,
	}
}
