package asset

import (
	"crypto-balancer/src/core/network"
	"crypto-balancer/src/feature/binance"
	"fmt"
)

func Filter(balances []binance.AccountBalance, lambda func(binance.AccountBalance) bool) []binance.AccountBalance {
	list := make([]binance.AccountBalance, 0)
	for _, balance := range balances {
		if lambda(balance) {
			list = append(list, balance)
		}
	}

	return list
}

func GetBalance(balances []binance.AccountBalance, asset string) (balance *binance.AccountBalance, err *network.APIError) {

	result := Filter(balances, func(balance binance.AccountBalance) bool {
		return balance.Asset == asset
	})

	resultLength := len(result)

	if resultLength == 0 {
		return nil, network.NewApiError(
			-1,
			fmt.Sprintf("%s seems to be invalid, %s was not found in this account", asset, asset),
		)
	}

	if resultLength > 1 {
		return nil, network.NewApiError(-1, fmt.Sprintf("Asset %s has more then one entry in the list", asset))
	}

	return &result[0], nil
}
