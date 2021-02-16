package asset

import (
	"context"
	"crypto-balancer/src/feature/binance/api/client"
	"strconv"
)

func GetPrice(assetSymbol string) (float64, error) {
	if assetSymbol == USDTSymbol {
		return 1, nil
	}

	gateway := client.NewBinanceClient().NewMarketAverageGateway()
	market, err := gateway.Symbol(assetSymbol).Do(context.Background())

	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(market.LastPrice, 64)
}
