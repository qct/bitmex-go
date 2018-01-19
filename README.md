# Golang SDK for [bitmex](https://www.bitmex.com)

inspired by [https://github.com/BitMEX/api-connectors](https://github.com/BitMEX/api-connectors), [https://www.bitmex.com/api/explorer](https://www.bitmex.com/api/explorer) and [https://github.com/jxc6698/bitcoin-exchange-api](https://github.com/jxc6698/bitcoin-exchange-api)

all structs and APIs are from the bitmex official api connectors, based on that, add authentication and fix bugs.

## Why
the generated connectors by bitmex have too many mistakes to use, this SDK fix these bugs to ensure bitmex API can get right results.

## Installation
`go get github.com/qct/bitmex-go`

## Examples
1. Get order book
```
apiClient = swagger.NewAPIClient(swagger.NewConfiguration())
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
```

2. Get your position
```
apiClient = swagger.NewAPIClient(swagger.NewConfiguration())
auth      = context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
        Key:    apiKey,
        Secret: secretKey,
    })

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
```

3. Get your wallet
```
apiClient = swagger.NewAPIClient(swagger.NewConfiguration())
auth      = context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
        Key:    apiKey,
        Secret: secretKey,
    })

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
```

4. Get margin info
```
apiClient = swagger.NewAPIClient(swagger.NewConfiguration())
auth      = context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
        Key:    apiKey,
        Secret: secretKey,
    })

userApi := apiClient.UserApi
margin, response, err := userApi.UserGetMargin(auth, nil)
if err != nil {
    log.Println("error: ", err)
}
log.Println(response.Status)
log.Printf("margin: %+v\n", margin)
```

5. Buy & Sell
```
apiClient = swagger.NewAPIClient(swagger.NewConfiguration())
auth      = context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
        Key:    apiKey,
        Secret: secretKey,
    })

// buy
orderApi := restful.NewOrderApi(apiClient.OrderApi, auth)
resp, orderId, err := orderApi.LimitBuy("XBTUSD", 1.0, 13000, "qct_f_f_")
if err != nil {
    log.Println("error: ", err)
}
log.Printf("response: %s, orderId: %s\n", resp.Status, orderId)

// sell
resp, orderId, err := orderApi.LimitSell("XBTUSD", 1.0, 20000, "qct_f_f_")
if err != nil {
    log.Println("error: ", err)
}
log.Printf("result: %s, orderId: %s\n", resp.Status, orderId)
```

6. Get your order
```
apiClient = swagger.NewAPIClient(swagger.NewConfiguration())
auth      = context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
        Key:    apiKey,
        Secret: secretKey,
    })

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
orders, response, err := orderApi.OrderGetOrders(auth, params)
if err != nil {
    log.Println("error: ", err)
}
log.Println(response.Status)
log.Printf("orders: %+v\n", orders)
```

7. Even chat through RESTFul api
```
apiClient = swagger.NewAPIClient(swagger.NewConfiguration())
auth      = context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
        Key:    apiKey,
        Secret: secretKey,
    })

chatApi := apiClient.ChatApi
params := map[string]interface{}{
    "channelID": 2.0,
}
chat, response, err := chatApi.ChatNew(auth, "hello", params)
if err != nil {
    log.Println("error: ", err)
}
log.Println(response.Status)
log.Printf("%+v\n", chat)
```

## Roadmap
* I only tested some API I will use, in the near future, I will keep adding and fixing to make more API available.
* make a PR to bitmex api collectors to generate a ready to use golang client.
