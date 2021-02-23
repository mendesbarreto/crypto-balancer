package asset

import (
	"crypto-balancer/src/feature/binance"
	"strconv"
)

func GetAmount(balances []binance.AccountBalance, asset string) (float64, error) {
	balance, err := GetBalance(balances, asset)

	if err != nil {
		return 0, err
	}

	value, floatErr := strconv.ParseFloat(balance.Free, 64)

	if floatErr != nil {
		return 0, floatErr
	}

	return value, nil
}
