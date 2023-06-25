package test

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/antihax/optional"
	"github.com/google/uuid"

	"github.com/qct/bitmex-go/restful"
	"github.com/qct/bitmex-go/swagger"

	"io/ioutil"
	"log"
	"strings"
	"testing"
)

var (
	apiKey    = ""
	apiSecret = ""
	apiClient = swagger.NewAPIClient(&swagger.Configuration{
		BasePath:      "https://testnet.bitmex.com/api/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
		ExpireTime:    5, //seconds
	})
	auth = context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
		Key:    apiKey,
		Secret: apiSecret,
	})
)

func Test_all(t *testing.T) {
	//Test_testMargin(t)
	//Test_testOrderBook(t)
	//Test_testPosition(t)
	//Test_testWallet(t)
	//Test_testGetOrder(t)
	//Test_testBuy(t)
	//Test_testSell(t)

	//testCancel("24e81603-e31f-431e-b06c-3b93226c8913,77a45818-ba95-451e-84c7-6667cd3b0912")
	//testChat()
}

func Test_testBuy(t *testing.T) {
	log.Println("-----------test buy------------")
	orderApi := restful.NewOrderApi(apiClient.OrderApi, auth)

	resp, orderId, err := orderApi.LimitBuy("XBTUSD", 100.0, 36000, generateClOrdID())
	if err != nil {
		log.Println("error: ", string(err.(swagger.GenericSwaggerError).Body()))
	}

	log.Printf("response: %s, orderId: %s\n", resp.Status, orderId)
	testCancel(orderId)
}

func Test_testGetOrder(t *testing.T) {
	log.Println("-----------test get order------------")
	orderApi := apiClient.OrderApi
	orders, response, err := orderApi.OrderGetOrders(auth, &swagger.OrderApiOrderGetOrdersOpts{
		Symbol:  optional.NewString("XBTUSD"),
		Filter:  optional.NewString("{\"open\":true}"),
		Count:   optional.NewFloat32(50),
		Reverse: optional.NewBool(true),
	})
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%s\n", string(bodyBytes))
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println("orders: ", len(orders))
	for _, order := range orders {
		log.Println(order.OrderID)
	}
}

func Test_testMargin(t *testing.T) {
	log.Println("-----------test margin------------")
	userApi := apiClient.UserApi
	margin, response, err := userApi.UserGetMargin(auth, nil)
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println(response.Status)
	log.Printf("margin: %+v\n", margin)
}

func Test_testOrderBook(t *testing.T) {
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

func Test_testPosition(t *testing.T) {
	log.Println("-----------test position------------")
	positionApi := apiClient.PositionApi
	positions, response, err := positionApi.PositionGet(auth, &swagger.PositionApiPositionGetOpts{
		Filter:  optional.NewString("{\"symbol\": \"XBTUSD\"}"),
		Columns: optional.NewString(""),
		Count:   optional.NewFloat32(10),
	})
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println(response.Status)
	log.Printf("%+v\n", positions)
}

func Test_testSell(t *testing.T) {
	log.Println("-----------test buy------------")
	orderApi := restful.NewOrderApi(apiClient.OrderApi, auth)

	resp, orderId, err := orderApi.LimitSell("XBTUSD", 100.0, 60000, generateClOrdID())
	if err != nil {
		log.Println("error: ", err)
	}

	log.Printf("result: %s, orderId: %s\n", resp.Status, orderId)
	testCancel(orderId)
}

func Test_testWallet(t *testing.T) {
	log.Println("-----------test wallet------------")
	userApi := apiClient.UserApi
	wallet, response, err := userApi.UserGetWallet(auth, &swagger.UserApiUserGetWalletOpts{
		Currency: optional.NewString("XBt"),
	})
	if err != nil {
		log.Println("error: ", err)
		log.Println("error: ", string(err.(swagger.GenericSwaggerError).Body()))
	}
	log.Println(response.Status)
	log.Printf("wallet: %+v\n", wallet)
}

func Test_testChat(t *testing.T) {
	log.Println("-----------test chat------------")
	chatApi := apiClient.ChatApi
	chat, response, err := chatApi.ChatNew(auth, "hello", &swagger.ChatApiChatNewOpts{
		ChannelID: optional.NewFloat64(2.0),
	})
	if err != nil {
		log.Println("error: ", string(err.(swagger.GenericSwaggerError).Body()))
	}
	log.Println(response.Status)
	if b, err := ioutil.ReadAll(response.Body); err == nil {
		log.Println("response body: ", string(b))
	} else {
		log.Println(err)
	}
	log.Printf("%+v\n", chat)
}

func testCancel(orderID string) {
	log.Println("-----------testCancel------------")
	cancel, resp, err := apiClient.OrderApi.OrderCancel(auth, &swagger.OrderApiOrderCancelOpts{
		OrderID: optional.NewString(orderID),
		Text:    optional.NewString("Cancelled by API"),
	})
	if err != nil {
		if resp != nil {
			log.Println(resp.Status)
		}
		log.Println("CancelledOrder err: ", string(err.(swagger.GenericSwaggerError).Body()))
	}
	log.Println("cancelled: ", len(cancel))
}

func generateClOrdID() string {
	u := uuid.NewString()
	bytes := []byte(u)[:16]
	s := strings.Replace(base64.StdEncoding.EncodeToString(bytes), "=", "", -1)
	return "qct_f_f_" + s
}
