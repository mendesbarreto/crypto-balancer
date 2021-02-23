package coingecko

import (
	"context"
	"crypto-balancer/src/core/network"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	MinItemsPerPage     = 1
	MaxItemsPerPage     = 250
	CoinsMarketEndpoint = "/coins/markets"
)

type MarketGateway struct {
	client                *Client
	vsCurrency            string
	ids                   []string
	order                 OrderType
	itemsPerPage          int
	page                  int
	sparkline             bool
	priceChangePercentage []string
}

func (gateway *MarketGateway) AddVsCurrency(vsCurrency string) *MarketGateway {
	gateway.vsCurrency = vsCurrency
	return gateway
}

func (gateway *MarketGateway) AddCoinId(id string) *MarketGateway {
	gateway.ids = append(gateway.ids, id)
	return gateway
}

func (gateway *MarketGateway) ItemsPerPage(value int) *MarketGateway {
	gateway.itemsPerPage = value
	return gateway
}

func (gateway *MarketGateway) Order(order OrderType) *MarketGateway {
	gateway.order = order
	return gateway
}

func (gateway *MarketGateway) Page(page int) *MarketGateway {
	gateway.page = page
	return gateway
}

func (gateway *MarketGateway) Sparkline(sparkline bool) *MarketGateway {
	gateway.sparkline = sparkline
	return gateway
}

func (gateway *MarketGateway) AddPriceChangePercentage(percentage string) *MarketGateway {
	gateway.priceChangePercentage = append(gateway.priceChangePercentage, percentage)
	return gateway
}

func (gateway *MarketGateway) Build(endpoint string) (*network.Request, error) {
	params := network.Params{}

	if len(gateway.vsCurrency) == 0 {
		return nil, network.NewApiError(ApiErrorCode, "vs_currency is required")
	}
	params["vs_currency"] = gateway.vsCurrency

	if len(gateway.order) == 0 {
		gateway.order = OrderTypeMarketCapDesc
	}

	params["order"] = string(gateway.order)

	if len(gateway.ids) != 0 {
		idsParam := strings.Join(gateway.ids[:], ",")
		params["ids"] = idsParam
	}

	if gateway.itemsPerPage < MinItemsPerPage || gateway.itemsPerPage > MaxItemsPerPage {
		return nil, network.NewApiError(
			ApiErrorCode,
			fmt.Sprintf("Invalid items per page the min: %v max: %v got: %v",
				MinItemsPerPage, MaxItemsPerPage, gateway.itemsPerPage,
			),
		)
	}

	params["per_page"] = strconv.Itoa(gateway.itemsPerPage)
	params["page"] = strconv.Itoa(gateway.page)
	params["sparkline"] = strconv.FormatBool(gateway.sparkline)

	if len(gateway.priceChangePercentage) != 0 {
		priceChangePercentageParam := strings.Join(gateway.priceChangePercentage[:], ",")
		params["price_change_percentage"] = priceChangePercentageParam
	}

	request := gateway.client.NewRequest(http.MethodGet, endpoint)
	request.SetParams(params)

	return request, nil
}

func (gateway *MarketGateway) Do(ctx context.Context) (response *CoinsMarket, err error) {
	request, err := gateway.Build(CoinsMarketEndpoint)

	if err != nil {
		return nil, err
	}

	data, err := gateway.client.Call(ctx, request)

	if err != nil {
		return nil, err
	}

	response = new(CoinsMarket)
	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
