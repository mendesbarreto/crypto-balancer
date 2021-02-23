package binance

import (
	"context"
	"crypto-balancer/src/core/network"
	"encoding/json"
	"net/http"
)

func UnmarshalAsset(data []byte) (*Asset, error) {
	asset := new(Asset)
	err := json.Unmarshal(data, &asset)
	return asset, err
}

func (asset *Asset) Marshal() ([]byte, error) {
	return json.Marshal(asset)
}

type Asset struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int64  `json:"firstId"`
	LastID             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}

const (
	Ticker24HrEndpoint string = "/api/v3/ticker/24hr"
)

type MarketTickerGateway struct {
	client *Client
	symbol string
}

func (gateway *MarketTickerGateway) Symbol(symbol string) *MarketTickerGateway {
	gateway.symbol = symbol
	return gateway
}

func (gateway *MarketTickerGateway) Build() *network.Request {
	request := gateway.client.NewRequest(http.MethodGet, Ticker24HrEndpoint, SectionNone)

	parameters := network.Params{
		"symbol": gateway.symbol,
	}

	return request.SetParams(parameters)
}

func (gateway *MarketTickerGateway) Do(ctx context.Context) (response *Asset, err error) {
	request := gateway.Build()

	//TODO: ADD tests: {"symbol":"LTCBTC","priceChange":"0.00000000","priceChangePercent":"0.000","weightedAvgPrice":"0.08900000","prevClosePrice":"0.08900000","lastPrice":"0.08900000","lastQty":"11.43705000","bidPrice":"0.00000000","bidQty":"0.00000000","askPrice":"0.08900000","askQty":"34.98036000","openPrice":"0.08900000","highPrice":"0.08900000","lowPrice":"0.08900000","volume":"11.43705000","quoteVolume":"1.01789745","openTime":1613248104387,"closeTime":1613334504387,"firstId":464,"lastId":514,"count":51}
	data, err := gateway.client.Call(ctx, request)

	if err != nil {
		return nil, err
	}

	response, err = UnmarshalAsset(data)

	if err != nil {
		return nil, err
	}

	return response, nil
}
