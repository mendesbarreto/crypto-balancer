package coingecko

import (
	"bytes"
	"context"
	"crypto-balancer/src/core/environment"
	"crypto-balancer/src/core/network"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ApiErrorCode = 4000
)

type ApiError struct {
	Error string `json:"error"`
}

type Client struct {
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
}

func NewCoingeckoClient() *Client {
	return &Client{
		BaseURL:    environment.CoinGeckoApiBaseUrl(),
		UserAgent:  environment.UserAgent(),
		HTTPClient: http.DefaultClient,
	}
}

func NewCoingeckoClientClientWithHttp(client *http.Client) *Client {
	return &Client{
		BaseURL:    environment.CoinGeckoApiBaseUrl(),
		UserAgent:  environment.UserAgent(),
		HTTPClient: client,
	}
}

func HttpResponseHandler(response *http.Response) (data []byte, err error) {
	log.Infof("status code: %d", response.StatusCode)

	data, err = ioutil.ReadAll(response.Body)

	if response.StatusCode >= 400 {
		coingeckoApiError := new(ApiError)
		jsonError := json.Unmarshal(data, coingeckoApiError)

		if jsonError != nil {
			log.Error(jsonError)
			return nil, jsonError
		}

		return nil, network.NewApiError(ApiErrorCode, coingeckoApiError.Error)
	}

	return data, err
}

func (client *Client) NewHeader() http.Header {
	header := http.Header{}
	header.Set("Content-Type", "application/json")
	return header
}

func (client *Client) createURL(endpoint string) string {
	return fmt.Sprintf("%s%s", client.BaseURL, endpoint)
}

func (client *Client) NewRequest(method string, endpoint string) *network.Request {
	request := &network.Request{
		Method:      method,
		Header:      client.NewHeader(),
		QueryValues: url.Values{},
		BodyValues:  url.Values{},
		Body:        &bytes.Buffer{},
		Path:        endpoint,
		Url:         client.createURL(endpoint),
		BaseURL:     client.BaseURL,
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

func (client *Client) NewMarketGateway() *MarketGateway {
	return &MarketGateway{
		client:                NewCoingeckoClient(),
		vsCurrency:            "",
		ids:                   []string{},
		order:                 "",
		itemsPerPage:          0,
		page:                  0,
		sparkline:             false,
		priceChangePercentage: []string{},
	}
}
