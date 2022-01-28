package test

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/antihax/optional"
	"github.com/qct/bitmex-go/restful"
	"github.com/qct/bitmex-go/swagger"
	uuid "github.com/satori/go.uuid"
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
	})
	auth = context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
		Key:    apiSecret,
		Prefix: apiKey,
	})
)

func Test_all(t *testing.T) {
	Test_testMargin(t)
	Test_testOrderBook(t)
	Test_testPosition(t)
	Test_testWallet(t)
	Test_testGetOrder(t)
	//testBuy()
	//testSell()
	//testBulkCreate()
	//testCancel()
	//testBulkAmend()
	//testChat()
}

func Test_testBuy(t *testing.T) {
	log.Println("-----------test buy------------")
	orderApi := restful.NewOrderApi(apiClient.OrderApi, auth)

	resp, orderId, err := orderApi.LimitBuy("XBTUSD", 1.0, 13000, generateClOrdID())
	if err != nil {
		log.Println("error: ", err)
	}

	log.Printf("response: %s, orderId: %s\n", resp.Status, orderId)
}

func Test_testGetOrder(t *testing.T) {
	log.Println("-----------test get order------------")
	orderApi := apiClient.OrderApi
	//time.Now().AddDate(0, 0, -1), time.Now()
	orders, response, err := orderApi.OrderGetOrders(auth, &swagger.OrderApiOrderGetOrdersOpts{
		Symbol:  optional.NewString("XBTUSD"),
		Filter:  optional.NewString("{\"opsen\":true}"),
		Count:   optional.NewFloat32(5),
		Reverse: optional.NewBool(true),
	})
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%s\n", string(bodyBytes))
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println(response.Status)
	log.Println("orders: ", len(orders))
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

	resp, orderId, err := orderApi.LimitSell("XBTUSD", 1.0, 20000, generateClOrdID())
	if err != nil {
		log.Println("error: ", err)
	}

	log.Printf("result: %s, orderId: %s\n", resp.Status, orderId)
}

func Test_testWallet(t *testing.T) {
	log.Println("-----------test wallet------------")
	userApi := apiClient.UserApi
	wallet, response, err := userApi.UserGetWallet(auth, &swagger.UserApiUserGetWalletOpts{
		Currency: optional.NewString("XBt"),
	})
	if err != nil {
		log.Println("error: ", err)
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

//func testBulkAmend() {
//	log.Println("-----------testBulkAmend------------")
//	bulkCreated := testBulkCreate()
//	var toAmend []swagger.Order
//	for _, o := range bulkCreated {
//		a := swagger.Order{
//			OrderQty: o.OrderQty + 1,
//			Price:    o.Price + 1,
//			OrderID:  o.OrderID,
//		}
//		toAmend = append(toAmend, a)
//	}
//	orderJson, _ := json.Marshal(toAmend)
//	amend, resp, err := apiClient.OrderApi.OrderAmendBulk(auth, map[string]interface{}{
//		"orders": string(orderJson),
//	})
//	if err != nil {
//		if resp != nil {
//			log.Println(resp.Status)
//		}
//		log.Println("AmendBulkOrders err: ", err)
//	}
//	log.Println("amended: ", len(amend))
//}

//func testCancel() {
//	log.Println("-----------testCancel------------")
//	bulkCreated := testBulkCreate()
//	var orderIds []string
//	for _, o := range bulkCreated {
//		orderIds = append(orderIds, o.OrderID)
//	}
//	cancel, resp, err := apiClient.OrderApi.OrderCancel(auth, map[string]interface{}{
//		"orderID": orderIds,
//		"text":    "Cancelled by API",
//	})
//	if err != nil {
//		if resp != nil {
//			log.Println(resp.Status)
//		}
//		log.Println("CancelBulkOrders err: ", err)
//	}
//	log.Println("cancelled: ", len(cancel))
//}

//func testBulkCreate() []swagger.Order {
//	log.Println("-----------testBulkCreate------------")
//	orders := []swagger.Order{
//		{
//			Symbol:   "XBTUSD",
//			ClOrdID:  generateClOrdID(),
//			Side:     "Sell",
//			Price:    20000,
//			OrderQty: 1,
//		},
//		{
//			Symbol:   "XBTUSD",
//			ClOrdID:  generateClOrdID(),
//			Side:     "Sell",
//			Price:    20001,
//			OrderQty: 2,
//		},
//	}
//	orderJson, _ := json.Marshal(orders)
//	bulk, resp, err := apiClient.OrderApi.OrderNewBulk(auth, map[string]interface{}{
//		"orders": string(orderJson),
//	})
//	if err != nil {
//		if resp != nil {
//			log.Println(resp.Status)
//		}
//		log.Println("CreateBulkOrders err: ", err)
//	}
//	return bulk
//}

func generateClOrdID() string {
	s := strings.Replace(base64.StdEncoding.EncodeToString(uuid.NewV4().Bytes()), "=", "", -1)
	return "qct_f_f_" + s
}
