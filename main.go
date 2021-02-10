package main

import (
	"context"
	"crypto-balancer/src/core/environment" //nolint:gci
	"crypto-balancer/src/feature/binance/api/client"
	"os"
	//nolint:gci
)

//TODO: Add tests to generate secrete  key with binance: https://github.com/binance-exchange/binance-signature-examples

func main() {
	environment.LoadVariables()
	binanceClient := client.NewBinanceClient()
	account, err := binanceClient.NewAccountRequestBuilder().Do(context.Background())

	if err != nil {
		println(err)
		os.Exit(-1)
	}

	println("%#v", account)
}
