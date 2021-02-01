package main

import (
	"crypto-balancer/src/core/environment" //nolint:gci
	"crypto-balancer/src/feature/binance/signature"
	"fmt" //nolint:gci
)

func main() {
	environment.LoadVariables()
	binanceSecretKey := environment.GetBinanceAPiSecretKey()
	fmt.Println(signature.Generate(binanceSecretKey))
}
