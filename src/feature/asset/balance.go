package asset

import (
	"crypto-balancer/src/core/network"
	"crypto-balancer/src/feature/binance/api/client"
	"fmt"
)

func Filter(balances []client.AccountBalance, lambda func(client.AccountBalance) bool) []client.AccountBalance {
	list := make([]client.AccountBalance, 0)
	for _, balance := range balances {
		if lambda(balance) {
			list = append(list, balance)
		}
	}

	return list
}

func GetBalance(balances []client.AccountBalance, asset string) (balance *client.AccountBalance, err *network.APIError) {

	result := Filter(balances, func(balance client.AccountBalance) bool {
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
