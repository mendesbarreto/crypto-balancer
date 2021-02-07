package main

import (
	"crypto-balancer/src/core/environment" //nolint:gci
	"crypto-balancer/src/feature/binance/signature"
	"fmt" //nolint:gci
)

func main() {
	environment.LoadVariables()
	binanceSecretKey := environment.BinanceAPiSecretKey()
	fmt.Println(signature.Generate(binanceSecretKey))
}
