package restful

import (
	"errors"
	"github.com/qct/bitmex-go/swagger"
	"golang.org/x/net/context"
	"net/http"
)

type OrderApi struct {
	swaggerOrderApi *swagger.OrderApiService
	ctx             context.Context
}

func NewOrderApi(swaggerOrderApi *swagger.OrderApiService, ctx context.Context) *OrderApi {
	return &OrderApi{swaggerOrderApi: swaggerOrderApi, ctx: ctx}
}

func (o *OrderApi) LimitBuy(symbol string, orderQty float64, price float64, clOrdID string) (resp *http.Response, orderId string, err error) {
	if symbol == "" {
		return nil, "", errors.New("symbol can NOT be empty")
	}
	if price <= 0 {
		return nil, "", errors.New("price must be positive")
	}
	params := map[string]interface{}{
		"symbol":   symbol,
		"ordType":  "Limit",
		"orderQty": float32(orderQty),
		"price":    price,
		"clOrdID":  clOrdID,
	}
	order, response, err := o.swaggerOrderApi.OrderNew(o.ctx, symbol, params)
	if err != nil || response.StatusCode != 200 {
		return response, order.OrderID, err
	}
	return response, order.OrderID, nil
}

func (o *OrderApi) LimitSell(symbol string, orderQty float64, price float64, clOrdID string) (resp *http.Response, orderId string, err error) {
	if symbol == "" {
		return nil, "", errors.New("symbol can NOT be empty")
	}
	if price <= 0 {
		return nil, "", errors.New("price must be positive")
	}
	params := map[string]interface{}{
		"symbol":   symbol,
		"orderQty": float32(-orderQty),
		"price":    price,
		"clOrdID":  clOrdID,
	}
	order, response, err := o.swaggerOrderApi.OrderNew(o.ctx, symbol, params)
	if err != nil || response.StatusCode != 200 {
		return response, order.OrderID, err
	}
	return response, order.OrderID, nil
}
