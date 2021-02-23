package main

import (
	"context"
	"crypto-balancer/src/core/environment" //nolint:gci
	"crypto-balancer/src/feature/balancer"
	"crypto-balancer/src/feature/binance"
	"crypto-balancer/src/feature/coindiscover"
	log "github.com/sirupsen/logrus"
	"os"
	//nolint:gci
)

func main() {
	environment.LoadVariables()
	SetupLogger()

	binanceClient := binance.NewBinanceClient()
	account, err := binanceClient.NewGetAccountGateway().Do(context.Background())

	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	log.Info(account)

	//err = balancer.BalanceBetweenTwoAssets(account, New1InchBalancer(), NewUsdtBalancer())
	coindiscover.DiscoverGems(1.2)
	//err = balancer.BalanceBetweenTwoAssets(account, New1InchBalancer(), NewUsdtBalancer())

	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}

func SetupLogger() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func New1InchBalancer() balancer.AssetBalancer {
	return balancer.AssetBalancer{
		Symbol:     "BTC",
		PairSymbol: "USDT",
		Weight:     0.5,
	}
}

func NewUsdtBalancer() balancer.AssetBalancer {
	return balancer.AssetBalancer{
		Symbol:     "USDT",
		PairSymbol: "BTC",
		Weight:     0.5,
	}
}
