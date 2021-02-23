package binance

import (
	"context"
	"crypto-balancer/src/core/network"
	"encoding/json"
	"net/http"
)

type SideType string
type PositionSideType string
type OrderType string
type TimeInForceType string
type NewOrderRespType string
type OrderStatusType string

// Global enums
const (
	OrderEndpoint string   = "/api/v3/order"
	SideTypeBuy   SideType = "BUY"
	SideTypeSell  SideType = "SELL"

	PositionSideTypeBoth  PositionSideType = "BOTH"
	PositionSideTypeLong  PositionSideType = "LONG"
	PositionSideTypeShort PositionSideType = "SHORT"

	OrderTypeLimit              OrderType = "LIMIT"
	OrderTypeMarket             OrderType = "MARKET"
	OrderTypeStop               OrderType = "STOP"
	OrderTypeStopMarket         OrderType = "STOP_MARKET"
	OrderTypeTakeProfit         OrderType = "TAKE_PROFIT"
	OrderTypeTakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"
	OrderTypeTrailingStopMarket OrderType = "TRAILING_STOP_MARKET"

	TimeInForceTypeGTC TimeInForceType = "GTC" // Good Till Cancel
	TimeInForceTypeIOC TimeInForceType = "IOC" // Immediate or Cancel
	TimeInForceTypeFOK TimeInForceType = "FOK" // Fill or Kill
	TimeInForceTypeGTX TimeInForceType = "GTX" // Good Till Crossing (Post Only)

	NewOrderRespTypeACK    NewOrderRespType = "ACK"
	NewOrderRespTypeRESULT NewOrderRespType = "RESULT"
	NewOrderRespTypeFULL   NewOrderRespType = "FULL"

	OrderStatusTypeNew             OrderStatusType = "NEW"
	OrderStatusTypePartiallyFilled OrderStatusType = "PARTIALLY_FILLED"
	OrderStatusTypeFilled          OrderStatusType = "FILLED"
	OrderStatusTypeCanceled        OrderStatusType = "CANCELED"
	OrderStatusTypeRejected        OrderStatusType = "REJECTED"
	OrderStatusTypeExpired         OrderStatusType = "EXPIRED"
)

// CreateOrderResponse define create order response
type CreateOrderResponse struct {
	Symbol                  string `json:"symbol"`
	OrderID                 int64  `json:"orderId"`
	ClientOrderID           string `json:"clientOrderId"`
	TransactTime            int64  `json:"transactTime"`
	Price                   string `json:"price"`
	OrigQuantity            string `json:"origQty"`
	ExecutedQuantity        string `json:"executedQty"`
	CumulativeQuoteQuantity string `json:"cummulativeQuoteQty"`
	IsIsolated              bool   `json:"isIsolated"` // for isolated margin

	Status      OrderStatusType `json:"status"`
	TimeInForce TimeInForceType `json:"timeInForce"`
	Type        OrderType       `json:"type"`
	Side        SideType        `json:"side"`

	// for order response is set to FULL
	Fills                 []*Fill `json:"fills"`
	MarginBuyBorrowAmount string  `json:"marginBuyBorrowAmount"` // for margin
	MarginBuyBorrowAsset  string  `json:"marginBuyBorrowAsset"`
}

// Fill may be returned in an array of fills in a CreateOrderResponse.
type Fill struct {
	Price           string `json:"price"`
	Quantity        string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
}

// CreateOrderService create order
type CreateOrderGateway struct {
	client           *Client
	symbol           string
	side             SideType
	positionSide     *PositionSideType
	orderType        OrderType
	timeInForce      *TimeInForceType
	quantity         *float64
	quoteOrderQty    *float64
	reduceOnly       *bool
	price            *string
	newClientOrderID *string
	stopPrice        *string
	workingType      *WorkingType
	activationPrice  *string
	callbackRate     *string
	newOrderRespType *NewOrderRespType
	closePosition    *bool
}

func (gateway *CreateOrderGateway) Symbol(symbol string) *CreateOrderGateway {
	gateway.symbol = symbol
	return gateway
}

func (gateway *CreateOrderGateway) Side(side SideType) *CreateOrderGateway {
	gateway.side = side
	return gateway
}

func (gateway *CreateOrderGateway) PositionSide(positionSide PositionSideType) *CreateOrderGateway {
	gateway.positionSide = &positionSide
	return gateway
}

func (gateway *CreateOrderGateway) Type(orderType OrderType) *CreateOrderGateway {
	gateway.orderType = orderType
	return gateway
}

func (gateway *CreateOrderGateway) TimeInForce(timeInForce TimeInForceType) *CreateOrderGateway {
	gateway.timeInForce = &timeInForce
	return gateway
}

func (gateway *CreateOrderGateway) Quantity(quantity float64) *CreateOrderGateway {
	gateway.quantity = &quantity
	return gateway
}

func (gateway *CreateOrderGateway) QuoteOrderQty(quoteOrderQty float64) *CreateOrderGateway {
	gateway.quoteOrderQty = &quoteOrderQty
	return gateway
}

func (gateway *CreateOrderGateway) ReduceOnly(reduceOnly bool) *CreateOrderGateway {
	gateway.reduceOnly = &reduceOnly
	return gateway
}

func (gateway *CreateOrderGateway) Price(price string) *CreateOrderGateway {
	gateway.price = &price
	return gateway
}

func (gateway *CreateOrderGateway) NewClientOrderID(newClientOrderID string) *CreateOrderGateway {
	gateway.newClientOrderID = &newClientOrderID
	return gateway
}

func (gateway *CreateOrderGateway) StopPrice(stopPrice string) *CreateOrderGateway {
	gateway.stopPrice = &stopPrice
	return gateway
}

func (gateway *CreateOrderGateway) WorkingType(workingType WorkingType) *CreateOrderGateway {
	gateway.workingType = &workingType
	return gateway
}

func (gateway *CreateOrderGateway) ActivationPrice(activationPrice string) *CreateOrderGateway {
	gateway.activationPrice = &activationPrice
	return gateway
}

func (gateway *CreateOrderGateway) CallbackRate(callbackRate string) *CreateOrderGateway {
	gateway.callbackRate = &callbackRate
	return gateway
}

func (gateway *CreateOrderGateway) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateOrderGateway {
	gateway.newOrderRespType = &newOrderResponseType
	return gateway
}

func (gateway *CreateOrderGateway) ClosePosition(closePosition bool) *CreateOrderGateway {
	gateway.closePosition = &closePosition
	return gateway
}

func (gateway *CreateOrderGateway) Validate() *CreateOrderGateway {
	return gateway
}

func (gateway *CreateOrderGateway) GetRequiredOrderParameters() (params network.Params, err error) {
	if gateway.symbol == "" {
		return nil, network.APIError{Code: -1, Message: "Missing Order param: symbol"}
	}

	if gateway.side == "" {
		return nil, network.APIError{Code: -1, Message: "Missing Order param: side"}
	}

	if gateway.orderType == "" {
		return nil, network.APIError{Code: -1, Message: "Missing Order param: orderType"}
	}

	requiredParams := network.Params{
		"symbol": gateway.symbol,
		"side":   gateway.side,
		"type":   gateway.orderType,
	}

	if gateway.quantity != nil {
		requiredParams["quantity"] = *gateway.quantity
	} else if gateway.quoteOrderQty != nil {
		requiredParams["quoteOrderQty"] = *gateway.quoteOrderQty
	} else {
		return nil, network.APIError{Code: -1, Message: "Missing Order param: quantity"}
	}

	return requiredParams, nil
}

func (gateway *CreateOrderGateway) Build(endpoint string) (*network.Request, error) {
	request := gateway.client.NewRequest(http.MethodPost, endpoint, SectionSigned)

	parameters, err := gateway.GetRequiredOrderParameters()

	if err != nil {
		return nil, err
	}

	if gateway.positionSide != nil {
		parameters["positionSide"] = *gateway.positionSide
	}
	if gateway.timeInForce != nil {
		parameters["timeInForce"] = *gateway.timeInForce
	}
	if gateway.reduceOnly != nil {
		parameters["reduceOnly"] = *gateway.reduceOnly
	}
	if gateway.price != nil {
		parameters["price"] = *gateway.price
	}
	if gateway.newClientOrderID != nil {
		parameters["newClientOrderId"] = *gateway.newClientOrderID
	}
	if gateway.stopPrice != nil {
		parameters["stopPrice"] = *gateway.stopPrice
	}
	if gateway.workingType != nil {
		parameters["workingType"] = *gateway.workingType
	}
	if gateway.activationPrice != nil {
		parameters["activationPrice"] = *gateway.activationPrice
	}
	if gateway.callbackRate != nil {
		parameters["callbackRate"] = *gateway.callbackRate
	}
	if gateway.closePosition != nil {
		parameters["closePosition"] = *gateway.closePosition
	}
	if gateway.newOrderRespType != nil {
		parameters["newOrderRespType"] = gateway.newOrderRespType
	}

	request.SetParams(parameters)

	return request, nil
}

func (gateway *CreateOrderGateway) Do(ctx context.Context) (response *CreateOrderResponse, err error) {
	request, err := gateway.Build(OrderEndpoint)

	if err != nil {
		return nil, err
	}

	data, err := gateway.client.Call(ctx, request)

	if err != nil {
		return nil, err
	}

	response = new(CreateOrderResponse)
	err = json.Unmarshal(data, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
