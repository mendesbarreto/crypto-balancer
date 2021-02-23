package coindiscover

import (
	"context"
	"crypto-balancer/src/feature/coingecko"
	log "github.com/sirupsen/logrus"
)

func DiscoverGems(ratio float64) {
	client := coingecko.NewCoingeckoClient()
	assets := GetAllTokens(client)

	for _, asset := range assets {
		if asset.tvlRatio >= ratio {
			log.WithFields(log.Fields{
				"asset":  asset.item.Name,
				"symbol": asset.item.Symbol,
				"value":  asset.item.PriceChange24h,
				"tvl":    asset.tvlRatio,
			}).Info("Asset:")
		}
	}
}

type Asset struct {
	item     coingecko.CoinsMarketItem
	tvlRatio float64
}

func NewAsset(item coingecko.CoinsMarketItem) Asset {
	var tvlRatio = 0.0
	totalSuplyLocked := item.TotalSupply - item.CirculatingSupply
	tvlRatio = totalSuplyLocked / item.MarketCap
	return Asset{item: item, tvlRatio: tvlRatio}
}

func GetAllTokens(client *coingecko.Client) []Asset {
	assets := []Asset{}

	var pageNumber int = 1
	var response *coingecko.CoinsMarket
	var err error

	for {
		response, err = client.
			NewMarketGateway().
			AddVsCurrency("usd").
			ItemsPerPage(250).
			Page(pageNumber).
			Do(context.Background())

		if response != nil && err != nil {
			break
		}

		currentContList := *response

		for _, item := range currentContList {
			if item.MarketCap != 0 && item.TotalSupply != 0 {
				asset := NewAsset(item)
				assets = append(assets, asset)
			}
		}

		if len(currentContList) == 0 {
			break
		}

		pageNumber++
	}

	return assets
}
