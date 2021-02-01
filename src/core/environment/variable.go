package environment

import "os"

const BinanceApiKeyStringKey = "BINANCE_API_KEY"
const BinanceApiSecretKeyStringKey = "BINANCE_API_SECRET_KEY" //nolint:gosec
const BinanceApiBaseUrlStringKey = "BINANCE_API_BASE_URL"

func GetBinanceApiKey() string {
	return os.Getenv(BinanceApiKeyStringKey)
}

func GetBinanceAPiSecretKey() string {
	return os.Getenv(BinanceApiSecretKeyStringKey)
}

func GetBinanceApiBaseUrl() string {
	return os.Getenv(BinanceApiBaseUrlStringKey)
}
