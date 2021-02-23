package environment

import (
	"fmt"
	"os"
)

const (
	AppNameKey                   = "APP_NAME"
	ProgrammingLanguageKey       = "PROGRAMMING_LANGUAGE"
	BinanceApiKeyStringKey       = "BINANCE_API_KEY"
	BinanceApiSecretKeyStringKey = "BINANCE_API_SECRET_KEY" //nolint:gosec
	BinanceApiBaseUrlStringKey   = "BINANCE_API_BASE_URL"
	CoinGeckoApiBaseUrlKey       = "COINGECKO_API_BASE_URL"
)

func BinanceApiKey() string {
	return os.Getenv(BinanceApiKeyStringKey)
}

func BinanceAPiSecretKey() string {
	return os.Getenv(BinanceApiSecretKeyStringKey)
}

func BinanceApiBaseUrl() string {
	return os.Getenv(BinanceApiBaseUrlStringKey)
}

func CoinGeckoApiBaseUrl() string {
	return os.Getenv(CoinGeckoApiBaseUrlKey)
}

func AppName() string {
	return os.Getenv(AppNameKey)
}

func ProgrammingLanguage() string {
	return os.Getenv(ProgrammingLanguageKey)
}

func UserAgent() string {
	return fmt.Sprintf("%s/%s", AppName(), ProgrammingLanguage())
}
