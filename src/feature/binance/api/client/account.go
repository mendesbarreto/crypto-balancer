package client

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	GetAccountEndpoint string = "/api/v3/account"
)

type Account struct {
	MakerCommission  int64            `json:"makerCommission"`
	TakerCommission  int64            `json:"takerCommission"`
	BuyerCommission  int64            `json:"buyerCommission"`
	SellerCommission int64            `json:"sellerCommission"`
	CanTrade         bool             `json:"canTrade"`
	CanWithdraw      bool             `json:"canWithdraw"`
	CanDeposit       bool             `json:"canDeposit"`
	Balances         []AccountBalance `json:"balances"`
}

type AccountBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// GetAccountService get asset info
type GetAccountGateway struct {
	client *BinanceClient
}

func (gateway *GetAccountGateway) Do(ctx context.Context) (res *Account, err error) {
	request := gateway.client.NewRequest(http.MethodGet, GetAccountEndpoint, SectionSigned)

	data, err := gateway.client.Call(ctx, request)

	if err != nil {
		return nil, err
	}

	res = new(Account)

	err = json.Unmarshal(data, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
