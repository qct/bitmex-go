package main

import (
    "log"
    "github.com/qct/bitmex-go/swagger"
    "golang.org/x/net/context"
    "github.com/qct/bitmex-go/restful"
    "io/ioutil"
)

var (
    apiKey    = ""
    secretKey = ""
    apiClient = swagger.NewAPIClient(swagger.NewConfiguration())
    auth      = context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
        Key:    apiKey,
        Secret: secretKey,
    })
)

func main() {

    //testMargin()
    //testOrderBook()
    //testPosition()
    //testWallet()
    //testGetOrder()
    testBuy()
    testSell()

    //testChat()
}

func testSell() {
    log.Println("-----------test buy------------")
    orderApi := restful.NewOrderApi(apiClient.OrderApi, auth)
    result, orderId, err := orderApi.LimitSell("XBTUSD", 1.0, 20000, "qct_f_f_")
    if err != nil {
        log.Println("error: ", err)
    }

    log.Printf("result: %t, orderId: %s\n", result, orderId)
}

func testBuy() {
    log.Println("-----------test buy------------")
    orderApi := restful.NewOrderApi(apiClient.OrderApi, auth)
    result, orderId, err := orderApi.LimitBuy("XBTUSD", 1.0, 13000, "qct_f_f_")
    if err != nil {
        log.Println("error: ", err)
    }

    log.Printf("result: %t, orderId: %s\n", result, orderId)
}

func testMargin() {
    log.Println("-----------test margin------------")
    userApi := apiClient.UserApi
    margin, response, err := userApi.UserGetMargin(auth, nil)
    if err != nil {
        log.Println("error: ", err)
    }
    log.Println(response.Status)
    log.Printf("margin: %+v\n", margin)
}

func testOrderBook() {
    log.Println("-----------test order book------------")
    orderBookApi := restful.NewOrderBookApi(apiClient.OrderBookApi)
    orderBooks, err := orderBookApi.OrderBookGetL2("XBTUSD", 5)
    if err != nil {
        log.Println("error wihle get orderbook: ", err)
    }
    for _, v := range orderBooks.AskList {
        log.Printf("%+v\n", v)
    }
    for _, v := range orderBooks.BidList {
        log.Printf("%+v\n", v)
    }
}

func testPosition() {
    log.Println("-----------test position------------")
    positionApi := apiClient.PositionApi
    params := map[string]interface{}{
        "filter":  "{\"symbol\": \"XBTUSD\"}",
        "columns": "",
        "count":   float32(10),
    }
    positions, response, err := positionApi.PositionGet(auth, params)
    if err != nil {
        log.Println("error: ", err)
    }
    log.Println(response.Status)
    log.Printf("%+v\n", positions)
}

func testWallet() {
    log.Println("-----------test wallet------------")
    userApi := apiClient.UserApi
    params := map[string]interface{}{
        "currency": "",
    }
    wallet, response, err := userApi.UserGetWallet(auth, params)
    if err != nil {
        log.Println("error: ", err)
    }
    log.Println(response.Status)
    log.Printf("wallet: %+v\n", wallet)
}

func testGetOrder() {
    log.Println("-----------test get order------------")
    orderApi := apiClient.OrderApi
    params := map[string]interface{}{
        "symbol":    "XBTUSD",
        "filter":    "",
        "columns":   "",
        "count":     float32(5),
        "start":     nil,
        "reverse":   true,
        "startTime": nil,
        "endTime":   nil,
    }
    //time.Now().AddDate(0, 0, -1), time.Now()
    orders, response, err := orderApi.OrderGetOrders(auth, params)
    if err != nil {
        log.Println("error: ", err)
    }
    log.Println(response.Status)
    log.Printf("margin: %+v\n", orders)
}

func testChat() {
    log.Println("-----------test chat------------")
    chatApi := apiClient.ChatApi
    params := map[string]interface{}{
        "channelID": 2.0,
    }
    chat, response, err := chatApi.ChatNew(auth, "hello", params)
    if err != nil {
        log.Println("error: ", err)
    }
    log.Println(response.Status)
    if b, err := ioutil.ReadAll(response.Body); err == nil {
        log.Println("response body: ", string(b))
    } else {
        log.Println(err)
    }
    log.Printf("%+v\n", chat)
}
