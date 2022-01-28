/*
 * BitMEX API
 *
 * ## REST API for the BitMEX Trading Platform  _If you are building automated tools, please subscribe to the_ _[BitMEX API RSS Feed](https://blog.bitmex.com/api_announcement/feed/) for changes. The feed will be updated_ _regularly and is the most reliable way to get downtime and update announcements._  [View Changelog](/app/apiChangelog)  ---  #### Getting Started  Base URI: [https://www.bitmex.com/api/v1](/api/v1)  ##### Fetching Data  All REST endpoints are documented below. You can try out any query right from this interface.  Most table queries accept `count`, `start`, and `reverse` params. Set `reverse=true` to get rows newest-first.  Additional documentation regarding filters, timestamps, and authentication is available in [the main API documentation](/app/restAPI).  _All_ table data is available via the [Websocket](/app/wsAPI). We highly recommend using the socket if you want to have the quickest possible data without being subject to ratelimits.  ##### Return Types  By default, all data is returned as JSON. Send `?_format=csv` to get CSV data or `?_format=xml` to get XML data.  ##### Trade Data Queries  _This is only a small subset of what is available, to get you started._  Fill in the parameters and click the `Try it out!` button to try any of these queries.  - [Pricing Data](#!/Quote/Quote_get)  - [Trade Data](#!/Trade/Trade_get)  - [OrderBook Data](#!/OrderBook/OrderBook_getL2)  - [Settlement Data](#!/Settlement/Settlement_get)  - [Exchange Statistics](#!/Stats/Stats_history)  Every function of the BitMEX.com platform is exposed here and documented. Many more functions are available.  ##### Swagger Specification  [⇩ Download Swagger JSON](swagger.json)  ---  ## All API Endpoints  Click to expand a section.
 *
 * API version: 1.0.0
 * Contact: support@bitmex.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// Raw Order and Balance Data
type Execution struct {
	ExecID                string    `json:"execID,omitempty"`
	OrderID               string    `json:"orderID,omitempty"`
	ClOrdID               string    `json:"clOrdID,omitempty"`
	ClOrdLinkID           string    `json:"clOrdLinkID,omitempty"`
	Account               float32   `json:"account,omitempty"`
	Symbol                string    `json:"symbol"`
	Side                  string    `json:"side,omitempty"`
	LastQty               float32   `json:"lastQty,omitempty"`
	LastPx                float64   `json:"lastPx,omitempty"`
	UnderlyingLastPx      float64   `json:"underlyingLastPx,omitempty"`
	LastMkt               string    `json:"lastMkt,omitempty"`
	LastLiquidityInd      string    `json:"lastLiquidityInd,omitempty"`
	SimpleOrderQty        float64   `json:"simpleOrderQty,omitempty"`
	OrderQty              float32   `json:"orderQty,omitempty"`
	Price                 float64   `json:"price,omitempty"`
	DisplayQty            float32   `json:"displayQty,omitempty"`
	StopPx                float64   `json:"stopPx,omitempty"`
	PegOffsetValue        float64   `json:"pegOffsetValue,omitempty"`
	PegPriceType          string    `json:"pegPriceType,omitempty"`
	Currency              string    `json:"currency,omitempty"`
	SettlCurrency         string    `json:"settlCurrency,omitempty"`
	ExecType              string    `json:"execType,omitempty"`
	OrdType               string    `json:"ordType,omitempty"`
	TimeInForce           string    `json:"timeInForce,omitempty"`
	ExecInst              string    `json:"execInst,omitempty"`
	ContingencyType       string    `json:"contingencyType,omitempty"`
	ExDestination         string    `json:"exDestination,omitempty"`
	OrdStatus             string    `json:"ordStatus,omitempty"`
	Triggered             string    `json:"triggered,omitempty"`
	WorkingIndicator      bool      `json:"workingIndicator,omitempty"`
	OrdRejReason          string    `json:"ordRejReason,omitempty"`
	SimpleLeavesQty       float64   `json:"simpleLeavesQty,omitempty"`
	LeavesQty             float32   `json:"leavesQty,omitempty"`
	SimpleCumQty          float64   `json:"simpleCumQty,omitempty"`
	CumQty                float32   `json:"cumQty,omitempty"`
	AvgPx                 float64   `json:"avgPx,omitempty"`
	Commission            float64   `json:"commission,omitempty"`
	TradePublishIndicator string    `json:"tradePublishIndicator,omitempty"`
	MultiLegReportingType string    `json:"multiLegReportingType,omitempty"`
	Text                  string    `json:"text,omitempty"`
	TrdMatchID            string    `json:"trdMatchID,omitempty"`
	ExecCost              float32   `json:"execCost,omitempty"`
	ExecComm              float32   `json:"execComm,omitempty"`
	HomeNotional          float64   `json:"homeNotional,omitempty"`
	ForeignNotional       float64   `json:"foreignNotional,omitempty"`
	TransactTime          time.Time `json:"transactTime,omitempty"`
	Timestamp             time.Time `json:"timestamp"`
}