package main

import (
	"crypto-balancer/src/core/environment" //nolint:gci
	"fmt"
)

func main() {
	environment.LoadVariables()
	fmt.Println(environment.GetBinanceApiKey())
	fmt.Println(environment.GetBinanceAPiSecretKey())
}
