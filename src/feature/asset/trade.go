package asset

import (
	"crypto-balancer/src/core/network"
	"crypto-balancer/src/feature/binance/api/client"
	"fmt"
)

const (
	USDTSymbol    = "USDT"
	BTCSymbol     = "BTC"
	ETHSymbol     = "ETH"
	DOTSymbol     = "DOT"
	OneInchSymbol = "1INCH"
)

func CanTrade(account *client.Account, assets []string) error {
	for _, asset := range assets {
		amount, err := GetAmount(account.Balances, asset)

		if err != nil {
			return err
		}

		if amount <= 0 {
			return network.NewApiError(-1,
				fmt.Sprintf("%s does not have amount available for trade", asset),
			)
		}
	}

	//TODO: Put here the endpoint to validate the API status

	return nil
}
