package client

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
type OrderBuilder struct {
	client           *BinanceClient
	symbol           string
	side             SideType
	positionSide     *PositionSideType
	orderType        OrderType
	timeInForce      *TimeInForceType
	quantity         string
	reduceOnly       *bool
	price            *string
	newClientOrderID *string
	stopPrice        *string
	workingType      *WorkingType
	activationPrice  *string
	callbackRate     *string
	newOrderRespType NewOrderRespType
	closePosition    *bool
}

func (builder *OrderBuilder) Symbol(symbol string) *OrderBuilder {
	builder.symbol = symbol
	return builder
}

func (builder *OrderBuilder) Side(side SideType) *OrderBuilder {
	builder.side = side
	return builder
}

func (builder *OrderBuilder) PositionSide(positionSide PositionSideType) *OrderBuilder {
	builder.positionSide = &positionSide
	return builder
}

func (builder *OrderBuilder) Type(orderType OrderType) *OrderBuilder {
	builder.orderType = orderType
	return builder
}

func (builder *OrderBuilder) TimeInForce(timeInForce TimeInForceType) *OrderBuilder {
	builder.timeInForce = &timeInForce
	return builder
}

func (builder *OrderBuilder) Quantity(quantity string) *OrderBuilder {
	builder.quantity = quantity
	return builder
}

func (builder *OrderBuilder) ReduceOnly(reduceOnly bool) *OrderBuilder {
	builder.reduceOnly = &reduceOnly
	return builder
}

func (builder *OrderBuilder) Price(price string) *OrderBuilder {
	builder.price = &price
	return builder
}

func (builder *OrderBuilder) NewClientOrderID(newClientOrderID string) *OrderBuilder {
	builder.newClientOrderID = &newClientOrderID
	return builder
}

func (builder *OrderBuilder) StopPrice(stopPrice string) *OrderBuilder {
	builder.stopPrice = &stopPrice
	return builder
}

func (builder *OrderBuilder) WorkingType(workingType WorkingType) *OrderBuilder {
	builder.workingType = &workingType
	return builder
}

func (builder *OrderBuilder) ActivationPrice(activationPrice string) *OrderBuilder {
	builder.activationPrice = &activationPrice
	return builder
}

func (builder *OrderBuilder) CallbackRate(callbackRate string) *OrderBuilder {
	builder.callbackRate = &callbackRate
	return builder
}

func (builder *OrderBuilder) NewOrderResponseType(newOrderResponseType NewOrderRespType) *OrderBuilder {
	builder.newOrderRespType = newOrderResponseType
	return builder
}

func (builder *OrderBuilder) Validate() *OrderBuilder {
	return builder
}

func (builder *OrderBuilder) Build(ctx context.Context, endpoint string) (data []byte, err error) {
	request := builder.client.NewRequest(http.MethodGet, endpoint, SectionSigned)

	parameters := network.Params{
		"symbol":           builder.symbol,
		"side":             builder.side,
		"type":             builder.orderType,
		"quantity":         builder.quantity,
		"newOrderRespType": builder.newOrderRespType,
	}

	if builder.positionSide != nil {
		parameters["positionSide"] = *builder.positionSide
	}
	if builder.timeInForce != nil {
		parameters["timeInForce"] = *builder.timeInForce
	}
	if builder.reduceOnly != nil {
		parameters["reduceOnly"] = *builder.reduceOnly
	}
	if builder.price != nil {
		parameters["price"] = *builder.price
	}
	if builder.newClientOrderID != nil {
		parameters["newClientOrderId"] = *builder.newClientOrderID
	}
	if builder.stopPrice != nil {
		parameters["stopPrice"] = *builder.stopPrice
	}
	if builder.workingType != nil {
		parameters["workingType"] = *builder.workingType
	}
	if builder.activationPrice != nil {
		parameters["activationPrice"] = *builder.activationPrice
	}
	if builder.callbackRate != nil {
		parameters["callbackRate"] = *builder.callbackRate
	}
	if builder.closePosition != nil {
		parameters["closePosition"] = *builder.closePosition
	}

	request.SetParams(parameters)

	data, err = builder.client.Call(ctx, request, SectionSigned)

	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func (builder *OrderBuilder) Do(ctx context.Context) (response *CreateOrderResponse, err error) {
	data, err := builder.Build(ctx, OrderEndpoint)

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
