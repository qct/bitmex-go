package main

import (
	"log"
	"github.com/qct/bitmex-go/swagger"
	//. "github.com/qct/bitmex-go/restful"
	"golang.org/x/net/context"
)

var (
	apiKey    = ""
	secretKey = ""
	config    = swagger.NewConfiguration()
	apiClient = swagger.NewAPIClient(config)
)

func main() {

	testChat()
	//testOrderBook()
	//testPosition()
	//testWallet()
	//testMargin()
}

func testChat() {
	log.Println("-----------test chat------------")
	chatApi := apiClient.ChatApi
	channel := make(map[string]interface{})
	channel["channelID"] = 2.0
	auth := context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
		Key: "APIKEY",
		Secret:"secret",
		Prefix: "Bearer", // Omit if not necessary.
	})
	chat, response, err := chatApi.ChatNew(auth,"hello", channel)
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println(response.Status)
	log.Println(chat)
}

/*func testGetOrder() {
	log.Println("-----------test get order------------")
	orderApi := swagger.NewOrderApiWithConfig(config)
	orders, response, err := orderApi.OrderGetOrders("XBTUSD", "", "",
		10, 0, true, time.Now().AddDate(0, 0, -1), time.Now())
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println(response.Status)
	log.Println("margin:", orders[0])
}
func testMargin() {
	log.Println("-----------test margin------------")
	userApi := swagger.NewUserApiWithConfig(config)
	margin, response, err := userApi.UserGetMargin("")
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println(response.Status)
	log.Println("margin:", *margin)
}
func testWallet() {
	log.Println("-----------test wallet------------")
	userApi := swagger.NewUserApiWithConfig(config)
	wallet, response, err := userApi.UserGetWallet("")
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println(response.Status)
	log.Println("wallet: ", *wallet)
}

func testPosition() {
	log.Println("-----------test position------------")
	positionApi := swagger.NewPositionApiWithConfig(config)
	positions, response, err := positionApi.PositionGet("{\"symbol\": \"XBTUSD\"}", "", 10)
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println(response.Status)
	log.Println(positions[0])
}

func testOrderBook() {
	log.Println("-----------test order book------------")
	orderBookApi := NewOrderBookApiWithBasePath("https://www.bitmex.com/api/v1")
	orderBooks, err := orderBookApi.OrderBookGetL2("XBTUSD", 5)
	if err != nil {
		log.Println("error wihle get orderbook: ", err)
	}
	for _, v := range orderBooks.AskList {
		log.Println(v)
	}
	for _, v := range orderBooks.BidList {
		log.Println(v)
	}
}*/
