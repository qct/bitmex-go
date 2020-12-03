package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/qct/bitmex-go/restful"
	"github.com/qct/bitmex-go/swagger"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"strings"
)

var (
	apiKey    = ""
	apiSecret = ""
	apiClient = swagger.NewAPIClient(swagger.NewConfiguration())
	auth      = context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
		Key:    apiKey,
		Secret: apiSecret,
	})
)

func main() {
	//testMargin()
	//testOrderBook()
	//testPosition()
	//testWallet()
	//testGetOrder()
	//testBuy()
	//testSell()
	//testBulkCreate()
	//testCancel()
	testBulkAmend()

	//testChat()
}

func testBulkAmend() {
	log.Println("-----------testBulkAmend------------")
	bulkCreated := testBulkCreate()
	var toAmend []swagger.Order
	for _, o := range bulkCreated {
		a := swagger.Order{
			OrderQty: o.OrderQty + 1,
			Price:    o.Price + 1,
			OrderID:  o.OrderID,
		}
		toAmend = append(toAmend, a)
	}
	orderJson, _ := json.Marshal(toAmend)
	amend, resp, err := apiClient.OrderApi.OrderAmendBulk(auth, map[string]interface{}{
		"orders": string(orderJson),
	})
	if err != nil {
		if resp != nil {
			log.Println(resp.Status)
		}
		log.Println("AmendBulkOrders err: ", err)
	}
	log.Println("amended: ", len(amend))
}

func testCancel() {
	log.Println("-----------testCancel------------")
	bulkCreated := testBulkCreate()
	var orderIds []string
	for _, o := range bulkCreated {
		orderIds = append(orderIds, o.OrderID)
	}
	cancel, resp, err := apiClient.OrderApi.OrderCancel(auth, map[string]interface{}{
		"orderID": orderIds,
		"text":    "Cancelled by API",
	})
	if err != nil {
		if resp != nil {
			log.Println(resp.Status)
		}
		log.Println("CancelBulkOrders err: ", err)
	}
	log.Println("cancelled: ", len(cancel))
}

func testBulkCreate() []swagger.Order {
	log.Println("-----------testBulkCreate------------")
	orders := []swagger.Order{
		{
			Symbol:   "XBTUSD",
			ClOrdID:  generateClOrdID(),
			Side:     "Sell",
			Price:    20000,
			OrderQty: 1,
		},
		{
			Symbol:   "XBTUSD",
			ClOrdID:  generateClOrdID(),
			Side:     "Sell",
			Price:    20001,
			OrderQty: 2,
		},
	}
	orderJson, _ := json.Marshal(orders)
	bulk, resp, err := apiClient.OrderApi.OrderNewBulk(auth, map[string]interface{}{
		"orders": string(orderJson),
	})
	if err != nil {
		if resp != nil {
			log.Println(resp.Status)
		}
		log.Println("CreateBulkOrders err: ", err)
	}
	return bulk
}

func testSell() {
	log.Println("-----------test buy------------")
	orderApi := restful.NewOrderApi(apiClient.OrderApi, auth)

	resp, orderId, err := orderApi.LimitSell("XBTUSD", 1.0, 20000, generateClOrdID())
	if err != nil {
		log.Println("error: ", err)
	}

	log.Printf("result: %s, orderId: %s\n", resp.Status, orderId)
}

func testBuy() {
	log.Println("-----------test buy------------")
	orderApi := restful.NewOrderApi(apiClient.OrderApi, auth)

	resp, orderId, err := orderApi.LimitBuy("XBTUSD", 1.0, 13000, generateClOrdID())
	if err != nil {
		log.Println("error: ", err)
	}

	log.Printf("response: %s, orderId: %s\n", resp.Status, orderId)
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
	log.Printf("orders: %+v\n", orders)
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

func generateClOrdID() string {
	s := strings.Replace(base64.StdEncoding.EncodeToString(uuid.NewV4().Bytes()), "=", "", -1)
	return "qct_f_f_" + s
}
