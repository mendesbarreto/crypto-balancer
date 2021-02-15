package balancer

import (
	"context"
	"crypto-balancer/src/core/log"
	"crypto-balancer/src/feature/asset"
	"crypto-balancer/src/feature/binance/api/client"
)

const (
	ada = "ADA"
)

type AssetBalancer struct {
	Symbol string
	// Between 0 and 1
	Weight float64
}

type UsdAssetWrapper struct {
	asset           AssetBalancer
	assetAmount     float64
	usdPricePerUnit float64
	totalUsdPrice   float64
}

func (wrapper *UsdAssetWrapper) Ratio(total float64) float64 {
	return wrapper.totalUsdPrice / total
}

func newUsdAssetWrapper(
	account *client.Account,
	assetBalancer AssetBalancer,
	pairSymbol string,
) (*UsdAssetWrapper, error) {
	baseAssetAmount, err := asset.GetAmount(account.Balances, assetBalancer.Symbol)

	if err != nil {
		return nil, err
	}

	baseAssetUsdPrice, err := asset.GetPrice(pairSymbol)

	if err != nil {
		return nil, err
	}

	return &UsdAssetWrapper{
		asset:           assetBalancer,
		assetAmount:     baseAssetAmount,
		usdPricePerUnit: baseAssetUsdPrice,
		totalUsdPrice:   baseAssetAmount * baseAssetUsdPrice,
	}, nil
}

func BalanceBetweenTwoAssets(account *client.Account, baseAsset AssetBalancer, subAsset AssetBalancer) error {
	if err := asset.CanTrade(account, []string{baseAsset.Symbol, subAsset.Symbol}); err != nil {
		return err
	}

	pairSymbol := GetPairSymbol(baseAsset, subAsset)

	baseAssetWrapper, err := newUsdAssetWrapper(account, baseAsset, pairSymbol)

	if err != nil {
		return err
	}

	subAssetWrapper, err := newUsdAssetWrapper(account, subAsset, pairSymbol)

	if err != nil {
		return err
	}

	if ShouldBuy(baseAssetWrapper, subAssetWrapper) {
		amountToBuy := AmountToBuy(baseAssetWrapper, subAssetWrapper)
		Log("BUY", baseAssetWrapper, subAssetWrapper, pairSymbol, amountToBuy)
		err = Buy(amountToBuy, pairSymbol)
	}

	if err != nil && ShouldSell(baseAssetWrapper, subAssetWrapper) {
		amountToSell := AmountToBuy(baseAssetWrapper, subAssetWrapper)
		Log("SELL", baseAssetWrapper, subAssetWrapper, pairSymbol, amountToSell)
		err = Sell(
			amountToSell,
			pairSymbol,
		)
	}

	return err
}

func Log(
	transaction string,
	baseWrapper *UsdAssetWrapper,
	subWrapper *UsdAssetWrapper,
	pair string,
	transactionAmount float64,
) {
	log.LogInfo("---------------|%s|---------------------", transaction)
	log.LogInfo("-----------| PAIR:%s |------------------", pair)
	log.LogInfo("----------------------------------------")
	log.LogInfo(
		"BaseAsset: %#v | price: %#v | total: %#v | ratio:%#v |",
		baseWrapper.asset.Symbol,
		baseWrapper.usdPricePerUnit,
		baseWrapper.totalUsdPrice,
		baseWrapper.asset.Weight,
	)
	log.LogInfo(
		"SubAsset: %#v | price: %f | total: %f | ratio:%f |",
		subWrapper.asset.Symbol,
		subWrapper.usdPricePerUnit,
		subWrapper.totalUsdPrice,
		baseWrapper.asset.Weight,
	)
	log.LogInfo("Amount: %f", transactionAmount)
	log.LogInfo("--------------------------------------------")
}

func GetPairSymbol(baseAsset AssetBalancer, subAsset AssetBalancer) string {
	return baseAsset.Symbol + subAsset.Symbol
}

func Buy(amount float64, pairSymbol string) error {
	_, err := client.NewBinanceClient().
		NewCreateOrderGateway().
		Symbol(pairSymbol).
		Type(client.OrderTypeMarket).
		Side(client.SideTypeBuy).
		QuoteOrderQty(amount).
		Validate().
		Do(context.Background())

	return err
}

func Sell(amount float64, pairSymbol string) error {
	_, err := client.NewBinanceClient().
		NewCreateOrderGateway().
		Symbol(pairSymbol).
		Type(client.OrderTypeMarket).
		Side(client.SideTypeBuy).
		QuoteOrderQty(amount).
		Validate().
		Do(context.Background())

	return err
}

func AmountToBuy(wrapper *UsdAssetWrapper, subWrapper *UsdAssetWrapper) float64 {
	total := wrapper.totalUsdPrice + subWrapper.totalUsdPrice
	buyAmount := (wrapper.asset.Weight - wrapper.Ratio(total)) * total

	if buyAmount <= 0 {
		return 0
	}

	return buyAmount
}

func AmountToSell(wrapper *UsdAssetWrapper, subWrapper *UsdAssetWrapper) float64 {
	total := wrapper.totalUsdPrice + subWrapper.totalUsdPrice
	sellAmount := (wrapper.asset.Weight - wrapper.Ratio(total)) * total

	if sellAmount >= 0 {
		return 0
	}

	return sellAmount
}

func ShouldBuy(baseWrapper *UsdAssetWrapper, subWrapper *UsdAssetWrapper) bool {
	totalUsdPrice := baseWrapper.totalUsdPrice + subWrapper.totalUsdPrice
	return baseWrapper.Ratio(totalUsdPrice) < baseWrapper.asset.Weight
}

func ShouldSell(baseWrapper *UsdAssetWrapper, subWrapper *UsdAssetWrapper) bool {
	totalUsdPrice := baseWrapper.totalUsdPrice + subWrapper.totalUsdPrice
	return subWrapper.Ratio(totalUsdPrice) < subWrapper.asset.Weight
}
