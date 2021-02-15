package asset

import (
	"context"
	"crypto-balancer/src/feature/binance/api/client"
	"strconv"
)

func GetPrice(asset string) (float64, error) {
	gateway := client.NewBinanceClient().NewMarketAverageGateway()
	market, err := gateway.Symbol(asset).Do(context.Background())

	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(market.LastPrice, 64)
}
