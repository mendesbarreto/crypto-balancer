package client

import (
	"crypto-balancer/src/core/environment"
	"log"
	"net/http"
	"os"
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

func NewClient() *BinanceClient {
	return &BinanceClient{
		ApiKey:     environment.GetBinanceApiKey(),
		SecretKey:  environment.GetBinanceAPiSecretKey(),
		BaseURL:    environment.GetBinanceApiBaseUrl(),
		UserAgent:  "Crypto-Balancer/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Crypto-Balancer: ", log.LstdFlags),
	}
}
