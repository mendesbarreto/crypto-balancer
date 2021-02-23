package binance

import (
	"crypto-balancer/src/core/environment"
	"crypto-balancer/src/core/network"
	"fmt"
	"net/url"
	"testing"
)

func setupCreateOrder() *CreateOrderGateway {
	environment.LoadVariables()
	return NewBinanceClient().NewCreateOrderGateway()
}

func TestSymbols(test *testing.T) {
	createOrderGateway := setupCreateOrder()
	value := "BTC"

	createOrderGateway.Symbol(value)

	if createOrderGateway.symbol != "BTC" {
		test.Errorf("The value to symbol should the the same as mtehod parameter %s", value)
	}
}

func TestSide(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	createOrderGateway.Side(SideTypeBuy)
	if createOrderGateway.side != SideTypeBuy {
		test.Errorf("The value should the the same as mtehod parameter %s", SideTypeBuy)
	}

	createOrderGateway.Side(SideTypeSell)
	if createOrderGateway.side != SideTypeSell {
		test.Errorf("The value should the the same as mtehod parameter %s", SideTypeSell)
	}
}

func TestPosition(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	createOrderGateway.PositionSide(PositionSideTypeBoth)
	if *createOrderGateway.positionSide != PositionSideTypeBoth {
		test.Errorf("The value should the the same as mtehod parameter %s", PositionSideTypeBoth)
	}

	createOrderGateway.PositionSide(PositionSideTypeLong)
	if *createOrderGateway.positionSide != PositionSideTypeLong {
		test.Errorf("The value should the the same as mtehod parameter %s", PositionSideTypeLong)
	}

	createOrderGateway.PositionSide(PositionSideTypeShort)
	if *createOrderGateway.positionSide != PositionSideTypeShort {
		test.Errorf("The value should the the same as mtehod parameter %s", PositionSideTypeShort)
	}
}

func TestType(test *testing.T) {
	createOrderGateway := setupCreateOrder()
	createOrderGateway.Type(OrderTypeLimit)

	if createOrderGateway.orderType != OrderTypeLimit {
		test.Errorf("The value should the the same as mtehod parameter %s", OrderTypeLimit)
	}

	createOrderGateway.Type(OrderTypeMarket)

	if createOrderGateway.orderType != OrderTypeMarket {
		test.Errorf("The value should the the same as mtehod parameter %s", OrderTypeMarket)
	}

	createOrderGateway.Type(OrderTypeStop)

	if createOrderGateway.orderType != OrderTypeStop {
		test.Errorf("The value should the the same as mtehod parameter %s", OrderTypeStop)
	}

	createOrderGateway.Type(OrderTypeStopMarket)

	if createOrderGateway.orderType != OrderTypeStopMarket {
		test.Errorf("The value should the the same as mtehod parameter %s", OrderTypeStopMarket)
	}

	createOrderGateway.Type(OrderTypeTakeProfit)

	if createOrderGateway.orderType != OrderTypeTakeProfit {
		test.Errorf("The value should the the same as mtehod parameter %s", OrderTypeTakeProfit)
	}

	createOrderGateway.Type(OrderTypeTakeProfitMarket)

	if createOrderGateway.orderType != OrderTypeTakeProfitMarket {
		test.Errorf("The value should the the same as mtehod parameter %s", OrderTypeTakeProfitMarket)
	}

	createOrderGateway.Type(OrderTypeTrailingStopMarket)

	if createOrderGateway.orderType != OrderTypeTrailingStopMarket {
		test.Errorf("The value should the the same as mtehod parameter %s", OrderTypeTrailingStopMarket)
	}
}

func TestTimeInForce(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	createOrderGateway.TimeInForce(TimeInForceTypeGTC)
	if *createOrderGateway.timeInForce != TimeInForceTypeGTC {
		test.Errorf("The value should the the same as mtehod parameter %s", TimeInForceTypeGTC)
	}

	createOrderGateway.TimeInForce(TimeInForceTypeIOC)

	if *createOrderGateway.timeInForce != TimeInForceTypeIOC {
		test.Errorf("The value should the the same as mtehod parameter %s", TimeInForceTypeIOC)
	}

	createOrderGateway.TimeInForce(TimeInForceTypeFOK)

	if *createOrderGateway.timeInForce != TimeInForceTypeFOK {
		test.Errorf("The value should the the same as mtehod parameter %s", TimeInForceTypeFOK)
	}

	createOrderGateway.TimeInForce(TimeInForceTypeGTX)

	if *createOrderGateway.timeInForce != TimeInForceTypeGTX {
		test.Errorf("The value should the the same as mtehod parameter %s", TimeInForceTypeGTX)
	}
}

func TestQuantity(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	value := 100.0
	createOrderGateway.Quantity(value)
	if *createOrderGateway.quantity != value {
		test.Errorf("The value should the the same as mtehod parameter %f", value)
	}
}

func TestReduceOnly(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	createOrderGateway.ReduceOnly(true)

	if !*createOrderGateway.reduceOnly {
		test.Errorf("The value should the the same as mtehod parameter %v", true)
	}
}

func TestPrice(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	value := "200"
	createOrderGateway.Price(value)

	if *createOrderGateway.price != value {
		test.Errorf("The value should the the same as mtehod parameter %s", value)
	}
}

func TestNewClientOrderID(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	value := "ABC123"
	createOrderGateway.NewClientOrderID(value)

	if *createOrderGateway.newClientOrderID != value {
		test.Errorf("The value should the the same as mtehod parameter %s", value)
	}
}

func TestStopPrice(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	value := "100"
	createOrderGateway.StopPrice(value)

	if *createOrderGateway.stopPrice != value {
		test.Errorf("The value should the the same as mtehod parameter %s", value)
	}
}

func TestWorkingType(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	createOrderGateway.WorkingType(WorkingTypeMarkPrice)
	if *createOrderGateway.workingType != WorkingTypeMarkPrice {
		test.Errorf("The value should the the same as mtehod parameter %s", WorkingTypeMarkPrice)
	}

	createOrderGateway.WorkingType(WorkingTypeContractPrice)

	if *createOrderGateway.workingType != WorkingTypeContractPrice {
		test.Errorf("The value should the the same as mtehod parameter %s", WorkingTypeContractPrice)
	}
}

func TestActivationPrice(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	value := "100"
	createOrderGateway.ActivationPrice(value)

	if *createOrderGateway.activationPrice != value {
		test.Errorf("The value should the the same as mtehod parameter %s", value)
	}
}

func TestCallbackRate(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	value := "100.0"
	createOrderGateway.CallbackRate(value)

	if *createOrderGateway.callbackRate != value {
		test.Errorf("The value should the the same as mtehod parameter %s", value)
	}
}

func TestClosePosition(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	createOrderGateway.ClosePosition(true)

	if !*createOrderGateway.closePosition {
		test.Errorf("The value should the the same as mtehod parameter %v", true)
	}
}

func TestNewOrderResponseType(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	createOrderGateway.NewOrderResponseType(NewOrderRespTypeACK)
	if *createOrderGateway.newOrderRespType != NewOrderRespTypeACK {
		test.Errorf("The value should the the same as mtehod parameter %s", NewOrderRespTypeACK)
	}

	createOrderGateway.NewOrderResponseType(NewOrderRespTypeRESULT)
	if *createOrderGateway.newOrderRespType != NewOrderRespTypeRESULT {
		test.Errorf("The value should the the same as mtehod parameter %s", NewOrderRespTypeRESULT)
	}

	createOrderGateway.NewOrderResponseType(NewOrderRespTypeFULL)
	if *createOrderGateway.newOrderRespType != NewOrderRespTypeFULL {
		test.Errorf("The value should the the same as mtehod parameter %s", NewOrderRespTypeFULL)
	}
}

func TestBuild(test *testing.T) {
	createOrderGateway := setupCreateOrder()

	request, _ := createOrderGateway.
		Symbol("BTC").
		Side(SideTypeSell).
		Type(OrderTypeMarket).
		PositionSide(PositionSideTypeLong).
		Quantity(100).
		NewOrderResponseType(NewOrderRespTypeFULL).
		TimeInForce(TimeInForceTypeFOK).
		ReduceOnly(true).
		Price("200").
		NewClientOrderID("123").
		StopPrice("50").
		WorkingType(WorkingTypeMarkPrice).
		ActivationPrice("100").
		CallbackRate("3.4").
		ClosePosition(false).
		Build(OrderEndpoint)

	if request == nil {
		test.Errorf("The requst object should not be nil")
	}

	if len(request.QueryValues) == 14 {
		test.Errorf("The requst object should not be nil")
	}

	hasKeysAndValue := func(key string, value string, queryValues url.Values) {
		if result := queryValues[key]; result != nil {
			if result[0] != value {
				test.Errorf("Expected value %s got %s", queryValues[key][0], value)
			}
		} else {
			test.Errorf("Key %s not exists", key)
		}
	}

	hasKeysAndValue("positionSide", string(PositionSideTypeLong), request.QueryValues)
	hasKeysAndValue("timeInForce", string(TimeInForceTypeFOK), request.QueryValues)
	hasKeysAndValue("reduceOnly", fmt.Sprintf("%v", true), request.QueryValues)
	hasKeysAndValue("price", "200", request.QueryValues)
	hasKeysAndValue("newClientOrderId", "123", request.QueryValues)
	hasKeysAndValue("stopPrice", "50", request.QueryValues)
	hasKeysAndValue("workingType", string(WorkingTypeMarkPrice), request.QueryValues)
	hasKeysAndValue("activationPrice", "100", request.QueryValues)
	hasKeysAndValue("callbackRate", "3.4", request.QueryValues)
	hasKeysAndValue("closePosition", fmt.Sprintf("%v", false), request.QueryValues)
}

func TestGetRequiredOrderParameters(test *testing.T) {
	gateway := setupCreateOrder()

	gateway.
		Symbol("BTC").
		Side(SideTypeSell).
		Type(OrderTypeMarket).
		Quantity(100).
		NewOrderResponseType(NewOrderRespTypeFULL)

	params, err := gateway.GetRequiredOrderParameters()

	if err != nil {
		test.Errorf("The error should be nil")
	}

	if params == nil {
		test.Errorf("The params object should not be nil")
	}

	if len(params) != 4 {
		test.Errorf("The param object should has 5 keys")
	}

	if params["symbol"] != gateway.symbol && params["side"] != gateway.side &&
		params["type"] != gateway.orderType && params["quantity"] != gateway.quantity &&
		params["newOrderRespType"] != gateway.newOrderRespType {
		test.Errorf("All keys should exists and be the same as defined on gateway")
	}

	validateErrorMessage := func(err error, message string) {
		if err == nil && !network.IsAPIError(err) {
			test.Errorf("Error should not be nil or not be the type of APIError")
		} else {
			if err.(network.APIError).Message != message {
				test.Errorf("Error message should contains: %s", message)
			}
		}
	}

	gateway = setupCreateOrder()
	gateway.
		Side(SideTypeSell).
		Type(OrderTypeMarket).
		PositionSide(PositionSideTypeLong).
		Quantity(100).
		NewOrderResponseType(NewOrderRespTypeFULL)

	params, err = gateway.GetRequiredOrderParameters()

	validateErrorMessage(err, "Missing Order param: symbol")

	gateway = setupCreateOrder()
	gateway.
		Symbol("BTC").
		Type(OrderTypeMarket).
		Quantity(100).
		NewOrderResponseType(NewOrderRespTypeFULL)

	params, err = gateway.GetRequiredOrderParameters()

	validateErrorMessage(err, "Missing Order param: side")

	gateway = setupCreateOrder()
	gateway.
		Symbol("BTC").
		Side(SideTypeSell).
		Quantity(100).
		NewOrderResponseType(NewOrderRespTypeFULL)

	params, err = gateway.GetRequiredOrderParameters()

	validateErrorMessage(err, "Missing Order param: orderType")

	gateway = setupCreateOrder()
	gateway.
		Symbol("BTC").
		Side(SideTypeSell).
		Type(OrderTypeMarket).
		NewOrderResponseType(NewOrderRespTypeFULL)

	params, err = gateway.GetRequiredOrderParameters()
	validateErrorMessage(err, "Missing Order param: quantity")
}

func TestCreateOrderDo(test *testing.T) {

}
