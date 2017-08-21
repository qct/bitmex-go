package restful

import (
    "../swagger"
    "log"
    "sort"
)

const (
    SIDE_BUY  = "Buy"
    SIDE_SELL = "Sell"
)

type OrderBooks struct {
    AskList,
    BidList []swagger.OrderBookL2
}

type OrderBookApi struct {
    swaggerOrderBookApi *swagger.OrderBookApi
}

func NewOrderBookApi() *OrderBookApi {
    configuration := swagger.NewConfiguration()
    swaggerOrderBookApi := swagger.OrderBookApi{Configuration: *configuration}
    return &OrderBookApi{swaggerOrderBookApi: &swaggerOrderBookApi}
}

func NewOrderBookApiWithBasePath(basePath string) *OrderBookApi {
    configuration := swagger.NewConfiguration()
    configuration.BasePath = basePath

    swaggerOrderBookApi := swagger.OrderBookApi{Configuration: *configuration}
    return &OrderBookApi{swaggerOrderBookApi: &swaggerOrderBookApi}
}

func (a OrderBookApi) OrderBookGetL2(symbol string, depth float32) (*OrderBooks, error) {
    rawOrderBooks, response, err := a.swaggerOrderBookApi.OrderBookGetL2(symbol, depth)
    if err != nil {
        log.Println("error wihle getting orderbook: ", err)
        return nil, err
    }
    if response.StatusCode != 200 {
        log.Println("response code incorrect, expect 200, but ", response.StatusCode)
        return nil, err
    }
    orderBooks := new(OrderBooks)
    for _, v := range rawOrderBooks {
        if v.Side == SIDE_BUY {
            orderBooks.BidList = append(orderBooks.BidList, v)
        } else if v.Side == SIDE_SELL {
            orderBooks.AskList = append(orderBooks.AskList, v)
        }
    }
    sort.Slice(orderBooks.AskList, func(i, j int) bool {
        return orderBooks.AskList[i].Price < orderBooks.AskList[j].Price
    })
    return orderBooks, nil
}
