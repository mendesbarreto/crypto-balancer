package binance

import (
	"bytes"
	"context"
	"crypto-balancer/src/core/datetime"
	"crypto-balancer/src/core/environment"
	"crypto-balancer/src/core/network"
	"crypto-balancer/src/feature/binance/signature"
	"encoding/hex"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
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

type Client struct {
	ApiKey      string
	SecretKey   string
	BaseURL     string
	UserAgent   string
	HTTPClient  *http.Client
	TimesOffset int64
}

func NewBinanceClientWithHttp(client *http.Client) *Client {
	return &Client{
		ApiKey:     environment.BinanceApiKey(),
		SecretKey:  environment.BinanceAPiSecretKey(),
		BaseURL:    environment.BinanceApiBaseUrl(),
		UserAgent:  environment.UserAgent(),
		HTTPClient: client,
	}
}

func NewBinanceClient() *Client {
	return &Client{
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

func (client *Client) NewHeader(sectionType SectionApiKeyType) http.Header {
	header := http.Header{}

	header.Set("Content-Type", "application/json")

	if sectionType == SectionAPIKey || sectionType == SectionSigned {
		header.Set("X-MBX-APIKEY", client.ApiKey)
	}

	return header
}

func (client *Client) createURL(endpoint string) string {
	return fmt.Sprintf("%s%s", client.BaseURL, endpoint)
}

func (client *Client) NewRequest(method string, endpoint string, sectionType SectionApiKeyType) *network.Request {
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

func (client *Client) Call(ctx context.Context, request *network.Request) (data []byte, err error) {
	httpRequest, err := request.ToHttpRequest(ctx)

	if err != nil {
		return nil, err
	}

	log.Infof("Start Request to: %s", httpRequest.URL)
	response, err := client.HTTPClient.Do(httpRequest)

	if err != nil {
		return nil, err
	}

	log.Info(response)

	defer func() {
		closeError := response.Body.Close()
		if err == nil && closeError != nil {
			err = closeError
		}
	}()

	return HttpResponseHandler(response)
}

func HttpResponseHandler(response *http.Response) (data []byte, err error) {
	log.Infof("status code: %d", response.StatusCode)

	data, err = ioutil.ReadAll(response.Body)

	if response.StatusCode >= 400 {
		binanceApiError := new(network.APIError)
		jsonError := json.Unmarshal(data, binanceApiError)

		if jsonError != nil {
			log.Error(jsonError)
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

func (client *Client) NewGetAccountGateway() *GetAccountGateway {
	return &GetAccountGateway{
		client: client,
	}
}

func (client *Client) NewCreateOrderGateway() *CreateOrderGateway {
	return &CreateOrderGateway{
		client: client,
	}
}

func (client *Client) NewMarketAverageGateway() *MarketTickerGateway {
	return &MarketTickerGateway{
		client: client,
	}
}
