package balancer

import (
	"context"
	"crypto-balancer/src/feature/asset"
	"crypto-balancer/src/feature/binance"
	log "github.com/sirupsen/logrus"
	"math"
)

const (
	BuyTriggerValue  = 0.15
	SellTriggerValue = 0.1
)

type AssetBalancer struct {
	Symbol     string
	PairSymbol string
	// Between 0 and 1
	Weight float64
}

func (asset *AssetBalancer) BasePairSymbol() string {
	return asset.Symbol + asset.PairSymbol
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

func newBaseUsdAssetWrapper(
	account *binance.Account,
	assetBalancer AssetBalancer,
) (*UsdAssetWrapper, error) {
	price, err := asset.GetPrice(assetBalancer.BasePairSymbol())

	if err != nil {
		return nil, err
	}

	return newUsdAssetWrapper(account, assetBalancer, price)
}

func newUsdAssetWrapper(
	account *binance.Account,
	assetBalancer AssetBalancer,
	price float64,
) (*UsdAssetWrapper, error) {
	baseAssetAmount, err := asset.GetAmount(account.Balances, assetBalancer.Symbol)

	if err != nil {
		return nil, err
	}

	return &UsdAssetWrapper{
		asset:           assetBalancer,
		assetAmount:     baseAssetAmount,
		usdPricePerUnit: price,
		totalUsdPrice:   baseAssetAmount * price,
	}, nil
}

func BalanceBetweenTwoAssets(account *binance.Account, baseAsset AssetBalancer, subAsset AssetBalancer) error {
	if err := asset.CanTrade(account, []string{baseAsset.Symbol, subAsset.Symbol}); err != nil {
		return err
	}

	pairSymbol := GetPairSymbol(baseAsset, subAsset)

	baseAssetWrapper, err := newBaseUsdAssetWrapper(account, baseAsset)

	if err != nil {
		return err
	}

	subAssetWrapper, err := newUsdAssetWrapper(account, subAsset, 1)

	if err != nil {
		return err
	}

	if ShouldBuy(baseAssetWrapper, subAssetWrapper) {
		amountToBuy := AmountToBuy(baseAssetWrapper, subAssetWrapper)
		Log("BUY", baseAssetWrapper, subAssetWrapper, amountToBuy)
		err = Buy(amountToBuy, pairSymbol)
	} else {
		log.Warning("The Balancer will not BUY anything")
	}

	if err == nil && ShouldSell(baseAssetWrapper, subAssetWrapper) {
		amountToSell := AmountToSell(baseAssetWrapper, subAssetWrapper)
		Log("SELL", baseAssetWrapper, subAssetWrapper, amountToSell)
		err = Sell(
			amountToSell,
			pairSymbol,
		)
	} else {
		log.Warning("The Balancer will not SELL anything")
	}

	return err
}

func Log(
	transactionType string,
	baseWrapper *UsdAssetWrapper,
	subWrapper *UsdAssetWrapper,
	transactionAmount float64,
) {
	log.Infof("----------------------------------------")
	log.Infof("-----------| TRANSACTION |--------------")
	log.Infof("----------------------------------------")
	log.WithFields(log.Fields{
		"Asset":      baseWrapper.asset.Symbol,
		"Price":      baseWrapper.usdPricePerUnit,
		"Total":      baseWrapper.totalUsdPrice,
		"Weight":     baseWrapper.asset.Weight,
		"PairSymbol": baseWrapper.asset.PairSymbol,
	}).Info("BASE")

	log.WithFields(log.Fields{
		"Asset":      subWrapper.asset.Symbol,
		"Weight":     subWrapper.asset.Weight,
		"PairSymbol": subWrapper.asset.PairSymbol,
		"Total":      subWrapper.totalUsdPrice,
		"Price":      subWrapper.usdPricePerUnit,
	}).Info("SUB")

	log.Infof("Total: %f", subWrapper.totalUsdPrice+baseWrapper.totalUsdPrice)
	log.Infof("%s: %f", baseWrapper.asset.Symbol, baseWrapper.totalUsdPrice)
	log.Infof("%s: %f", subWrapper.asset.Symbol, subWrapper.totalUsdPrice)
	log.Infof("Amount in USD of %s to %s: %f", baseWrapper.asset.Symbol, transactionType, transactionAmount)
	log.Infof("-----------| - |--------------")
}

func GetPairSymbol(baseAsset AssetBalancer, subAsset AssetBalancer) string {
	return baseAsset.Symbol + subAsset.Symbol
}

func Buy(amount float64, pairSymbol string) error {
	_, err := binance.NewBinanceClient().
		NewCreateOrderGateway().
		Symbol(pairSymbol).
		Type(binance.OrderTypeMarket).
		Side(binance.SideTypeBuy).
		QuoteOrderQty(amount).
		Validate().
		Do(context.Background())

	return err
}

func Sell(amount float64, pairSymbol string) error {
	_, err := binance.NewBinanceClient().
		NewCreateOrderGateway().
		Symbol(pairSymbol).
		Type(binance.OrderTypeMarket).
		Side(binance.SideTypeSell).
		QuoteOrderQty(amount).
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

	return math.Abs(sellAmount)
}

func ShouldBuy(baseWrapper *UsdAssetWrapper, subWrapper *UsdAssetWrapper) bool {
	totalUsdPrice := baseWrapper.totalUsdPrice + subWrapper.totalUsdPrice
	baseRation := baseWrapper.Ratio(totalUsdPrice)
	diff := math.Abs(baseRation - baseWrapper.asset.Weight)
	return baseRation < baseWrapper.asset.Weight && diff >= BuyTriggerValue
}

func ShouldSell(baseWrapper *UsdAssetWrapper, subWrapper *UsdAssetWrapper) bool {
	totalUsdPrice := baseWrapper.totalUsdPrice + subWrapper.totalUsdPrice
	diff := subWrapper.Ratio(totalUsdPrice) - subWrapper.asset.Weight
	return subWrapper.Ratio(totalUsdPrice) < subWrapper.asset.Weight && diff >= SellTriggerValue
}
