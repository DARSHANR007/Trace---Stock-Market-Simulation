# PositionData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | [***Object**](.md) | Exchange to which the order is associated | [optional] [default to null]
**Multiplier** | [***Object**](.md) | The quantity/lot size multiplier used for calculating P&amp;Ls | [optional] [default to null]
**Value** | [***Object**](.md) | Net value of the position | [optional] [default to null]
**Pnl** | [***Object**](.md) | Profit and loss - net returns on the position | [optional] [default to null]
**Product** | [***Object**](.md) | Shows if the order was either Intraday, Delivery, CO or OCO | [optional] [default to null]
**InstrumentToken** | [***Object**](.md) | Key issued by Upstox for the instrument | [optional] [default to null]
**AveragePrice** | [***Object**](.md) | Average price at which the net position quantity was acquired | [optional] [default to null]
**BuyValue** | [***Object**](.md) | Net value of the bought quantities | [optional] [default to null]
**OvernightQuantity** | [***Object**](.md) | Quantity held previously and carried forward over night | [optional] [default to null]
**DayBuyValue** | [***Object**](.md) | Amount at which the quantity is bought during the day | [optional] [default to null]
**DayBuyPrice** | [***Object**](.md) | Average price at which the day qty was bought. Default is empty string | [optional] [default to null]
**OvernightBuyAmount** | [***Object**](.md) | Amount at which the quantity was bought in the previous session | [optional] [default to null]
**OvernightBuyQuantity** | [***Object**](.md) | Quantity bought in the previous session | [optional] [default to null]
**DayBuyQuantity** | [***Object**](.md) | Quantity bought during the day | [optional] [default to null]
**DaySellValue** | [***Object**](.md) | Amount at which the quantity is sold during the day | [optional] [default to null]
**DaySellPrice** | [***Object**](.md) | Average price at which the day quantity was sold | [optional] [default to null]
**OvernightSellAmount** | [***Object**](.md) | Amount at which the quantity was sold in the previous session | [optional] [default to null]
**OvernightSellQuantity** | [***Object**](.md) | Quantity sold short in the previous session | [optional] [default to null]
**DaySellQuantity** | [***Object**](.md) | Quantity sold during the day | [optional] [default to null]
**Quantity** | [***Object**](.md) | Quantity left after nullifying Day and CF buy quantity towards Day and CF sell quantity | [optional] [default to null]
**LastPrice** | [***Object**](.md) | Last traded market price of the instrument | [optional] [default to null]
**Unrealised** | [***Object**](.md) | Day PnL generated against open positions | [optional] [default to null]
**Realised** | [***Object**](.md) | Day PnL generated against closed positions | [optional] [default to null]
**SellValue** | [***Object**](.md) | Net value of the sold quantities | [optional] [default to null]
**Tradingsymbol** | [***Object**](.md) | Shows the trading symbol of the instrument | [optional] [default to null]
**TradingSymbol** | [***Object**](.md) | Shows the trading symbol of the instrument | [optional] [default to null]
**ClosePrice** | [***Object**](.md) | Closing price of the instrument from the last trading day | [optional] [default to null]
**BuyPrice** | [***Object**](.md) | Average price at which quantities were bought | [optional] [default to null]
**SellPrice** | [***Object**](.md) | Average price at which quantities were sold | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

