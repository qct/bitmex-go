package restful

import (
    "encoding/base64"
    "errors"
    "github.com/qct/bitmex-go/swagger"
    "github.com/satori/go.uuid"
    "golang.org/x/net/context"
    "strings"
)

type OrderApi struct {
    swaggerOrderApi *swagger.OrderApiService
    ctx             context.Context
}

func NewOrderApi(swaggerOrderApi *swagger.OrderApiService, ctx context.Context) *OrderApi {
    return &OrderApi{swaggerOrderApi: swaggerOrderApi, ctx: ctx}
}

func (o *OrderApi) LimitBuy(symbol string, orderQty float64, price float64, orderIDPrefix string) (result bool, orderId string, err error) {
    if symbol == "" {
        return false, "", errors.New("symbol can NOT be empty")
    }
    if price <= 0 {
        return false, "", errors.New("price must be positive")
    }
    if orderIDPrefix != "" {
        s := strings.Replace(base64.StdEncoding.EncodeToString(uuid.NewV4().Bytes()), "=", "", -1)
        orderId = orderIDPrefix + s
    }

    params := map[string]interface{}{
        "symbol":   symbol,
        "ordType":  "Limit",
        "orderQty": float32(orderQty),
        "price":    price,
        "clOrdID":  orderId,
    }
    order, response, err := o.swaggerOrderApi.OrderNew(o.ctx, symbol, params)
    if err != nil || response == nil {
        return false, orderId, err
    }
    return true, order.OrderID, nil
}

func (o *OrderApi) LimitSell(symbol string, orderQty float64, price float64, orderIDPrefix string) (result bool, orderId string, err error) {
    if symbol == "" {
        return false, "", errors.New("symbol can NOT be empty")
    }
    if price <= 0 {
        return false, "", errors.New("price must be positive")
    }
    if orderIDPrefix != "" {
        s := strings.Replace(base64.StdEncoding.EncodeToString(uuid.NewV4().Bytes()), "=", "", -1)
        orderId = orderIDPrefix + s
    }

    params := map[string]interface{}{
        "symbol":   symbol,
        "orderQty": float32(-orderQty),
        "price":    price,
        "clOrdID":  orderId,
    }
    order, response, err := o.swaggerOrderApi.OrderNew(o.ctx, symbol, params)
    if err != nil {
        return false, orderId, err
    }
    if response.StatusCode != 200 {
        return false, orderId, errors.New("status: " + response.Status)
    }
    return true, order.OrderID, nil
}
