package client

import (
	"crypto-balancer/src/core/datetime"
	"crypto-balancer/src/core/environment"
	"crypto-balancer/src/core/network"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

type BinanceClient struct {
	ApiKey      string
	SecretKey   string
	BaseURL     string
	UserAgent   string
	HTTPClient  *http.Client
	Logger      *log.Logger
	TimesOffset int64
}

func NewBinanceClient() *BinanceClient {
	return &BinanceClient{
		ApiKey:     environment.BinanceApiKey(),
		SecretKey:  environment.BinanceAPiSecretKey(),
		BaseURL:    environment.BinanceApiBaseUrl(),
		UserAgent:  environment.UserAgent(),
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, environment.AppName()+" ", log.LstdFlags),
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

	if sectionType == SectionAPIKey || sectionType == SectionSigned {
		header.Set("X-MBX-APIKEY", client.ApiKey)
	}

	return header
}

func (client *BinanceClient) createURL(endpoint string) string {
	return fmt.Sprintf("%s%s", client.BaseURL, endpoint)
}

func (client *BinanceClient) createQueryParams(sectionType SectionApiKeyType) url.Values {
	query := url.Values{}

	if sectionType == SectionSigned {
		query.Add(timestampKey, fmt.Sprintf("%v", datetime.Timestamp(time.Now)))
	}

	return query
}

func (client *BinanceClient) NewRequest(method string, endpoint string, sectionType SectionApiKeyType) *network.Request {
	request := &network.Request{
		Method:      method,
		Header:      client.NewHeader(sectionType),
		QueryValues: client.createQueryParams(sectionType),
		BodyValues:  url.Values{},
		Path:        endpoint,
		Url:         client.createURL(endpoint),
		BaseURL:     client.BaseURL,
	}

	return request
}
