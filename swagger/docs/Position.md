# Position

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **float32** |  | [default to null]
**Symbol** | **string** |  | [default to null]
**Currency** | **string** |  | [optional] [default to null]
**Underlying** | **string** |  | [optional] [default to null]
**QuoteCurrency** | **string** |  | [optional] [default to null]
**Commission** | **float64** |  | [optional] [default to 0.0]
**InitMarginReq** | **float64** |  | [optional] [default to 0.0]
**MaintMarginReq** | **float64** |  | [optional] [default to 0.0]
**RiskLimit** | **float32** |  | [optional] [default to null]
**Leverage** | **float64** |  | [optional] [default to 0.0]
**CrossMargin** | **bool** |  | [optional] [default to null]
**DeleveragePercentile** | **float64** |  | [optional] [default to 0.0]
**RebalancedPnl** | **float32** |  | [optional] [default to null]
**PrevRealisedPnl** | **float32** |  | [optional] [default to null]
**PrevUnrealisedPnl** | **float32** |  | [optional] [default to null]
**PrevClosePrice** | **float64** |  | [optional] [default to 0.0]
**OpeningTimestamp** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**OpeningQty** | **float32** |  | [optional] [default to null]
**OpeningCost** | **float32** |  | [optional] [default to null]
**OpeningComm** | **float32** |  | [optional] [default to null]
**OpenOrderBuyQty** | **float32** |  | [optional] [default to null]
**OpenOrderBuyCost** | **float32** |  | [optional] [default to null]
**OpenOrderBuyPremium** | **float32** |  | [optional] [default to null]
**OpenOrderSellQty** | **float32** |  | [optional] [default to null]
**OpenOrderSellCost** | **float32** |  | [optional] [default to null]
**OpenOrderSellPremium** | **float32** |  | [optional] [default to null]
**ExecBuyQty** | **float32** |  | [optional] [default to null]
**ExecBuyCost** | **float32** |  | [optional] [default to null]
**ExecSellQty** | **float32** |  | [optional] [default to null]
**ExecSellCost** | **float32** |  | [optional] [default to null]
**ExecQty** | **float32** |  | [optional] [default to null]
**ExecCost** | **float32** |  | [optional] [default to null]
**ExecComm** | **float32** |  | [optional] [default to null]
**CurrentTimestamp** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**CurrentQty** | **float32** |  | [optional] [default to null]
**CurrentCost** | **float32** |  | [optional] [default to null]
**CurrentComm** | **float32** |  | [optional] [default to null]
**RealisedCost** | **float32** |  | [optional] [default to null]
**UnrealisedCost** | **float32** |  | [optional] [default to null]
**GrossOpenCost** | **float32** |  | [optional] [default to null]
**GrossOpenPremium** | **float32** |  | [optional] [default to null]
**GrossExecCost** | **float32** |  | [optional] [default to null]
**IsOpen** | **bool** |  | [optional] [default to null]
**MarkPrice** | **float64** |  | [optional] [default to 0.0]
**MarkValue** | **float32** |  | [optional] [default to null]
**RiskValue** | **float32** |  | [optional] [default to null]
**HomeNotional** | **float64** |  | [optional] [default to 0.0]
**ForeignNotional** | **float64** |  | [optional] [default to 0.0]
**PosState** | **string** |  | [optional] [default to null]
**PosCost** | **float32** |  | [optional] [default to null]
**PosCost2** | **float32** |  | [optional] [default to null]
**PosCross** | **float32** |  | [optional] [default to null]
**PosInit** | **float32** |  | [optional] [default to null]
**PosComm** | **float32** |  | [optional] [default to null]
**PosLoss** | **float32** |  | [optional] [default to null]
**PosMargin** | **float32** |  | [optional] [default to null]
**PosMaint** | **float32** |  | [optional] [default to null]
**PosAllowance** | **float32** |  | [optional] [default to null]
**TaxableMargin** | **float32** |  | [optional] [default to null]
**InitMargin** | **float32** |  | [optional] [default to null]
**MaintMargin** | **float32** |  | [optional] [default to null]
**SessionMargin** | **float32** |  | [optional] [default to null]
**TargetExcessMargin** | **float32** |  | [optional] [default to null]
**VarMargin** | **float32** |  | [optional] [default to null]
**RealisedGrossPnl** | **float32** |  | [optional] [default to null]
**RealisedTax** | **float32** |  | [optional] [default to null]
**RealisedPnl** | **float32** |  | [optional] [default to null]
**UnrealisedGrossPnl** | **float32** |  | [optional] [default to null]
**LongBankrupt** | **float32** |  | [optional] [default to null]
**ShortBankrupt** | **float32** |  | [optional] [default to null]
**TaxBase** | **float32** |  | [optional] [default to null]
**IndicativeTaxRate** | **float64** |  | [optional] [default to 0.0]
**IndicativeTax** | **float32** |  | [optional] [default to null]
**UnrealisedTax** | **float32** |  | [optional] [default to null]
**UnrealisedPnl** | **float32** |  | [optional] [default to null]
**UnrealisedPnlPcnt** | **float64** |  | [optional] [default to 0.0]
**UnrealisedRoePcnt** | **float64** |  | [optional] [default to 0.0]
**SimpleQty** | **float64** |  | [optional] [default to 0.0]
**SimpleCost** | **float64** |  | [optional] [default to 0.0]
**SimpleValue** | **float64** |  | [optional] [default to 0.0]
**SimplePnl** | **float64** |  | [optional] [default to 0.0]
**SimplePnlPcnt** | **float64** |  | [optional] [default to 0.0]
**AvgCostPrice** | **float64** |  | [optional] [default to 0.0]
**AvgEntryPrice** | **float64** |  | [optional] [default to 0.0]
**BreakEvenPrice** | **float64** |  | [optional] [default to 0.0]
**MarginCallPrice** | **float64** |  | [optional] [default to 0.0]
**LiquidationPrice** | **float64** |  | [optional] [default to 0.0]
**BankruptPrice** | **float64** |  | [optional] [default to 0.0]
**Timestamp** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**LastPrice** | **float64** |  | [optional] [default to 0.0]
**LastValue** | **float32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


