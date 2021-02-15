package main

import (
	"context"
	"crypto-balancer/src/core/environment" //nolint:gci
	"crypto-balancer/src/core/log"
	"crypto-balancer/src/feature/balancer"
	"crypto-balancer/src/feature/binance/api/client"
	"os"
	//nolint:gci
)

func main() {
	environment.LoadVariables()
	binanceClient := client.NewBinanceClient()
	account, err := binanceClient.NewGetAccountGateway().Do(context.Background())

	if err != nil {
		println(err)
		os.Exit(-1)
	}

	log.LogInfo("%#v", account)

	asset, err := binanceClient.NewMarketAverageGateway().Symbol("LTCBTC").Do(context.Background())

	if err != nil {
		log.LogDebug("%s", err.Error())
		os.Exit(-1)
	}

	log.LogInfo("%#v", asset)

	err = balancer.BalanceBetweenTwoAssets(account, New1InchBalancer(), NewUsdtBalancer())

	if err != nil {
		log.LogDebug("%s", err.Error())
		os.Exit(-1)
	}

	log.LogInfo("%#v", asset)
}

func New1InchBalancer() balancer.AssetBalancer {
	return balancer.AssetBalancer{
		Symbol: "BTC",
		Weight: 0.5,
	}
}

func NewUsdtBalancer() balancer.AssetBalancer {
	return balancer.AssetBalancer{
		Symbol: "USDT",
		Weight: 0.5,
	}
}
