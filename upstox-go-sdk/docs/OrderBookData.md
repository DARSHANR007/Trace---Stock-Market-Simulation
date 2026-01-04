# OrderBookData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | [***Object**](.md) | Exchange to which the order is associated | [optional] [default to null]
**Product** | [***Object**](.md) | Shows if the order was either Intraday, Delivery, CoverOrder or OneCancelsOther | [optional] [default to null]
**Price** | [***Object**](.md) | Price at which the order was placed | [optional] [default to null]
**Quantity** | [***Object**](.md) | Quantity with which the order was placed | [optional] [default to null]
**Status** | [***Object**](.md) | Indicates the current status of the order. Valid order status’ are outlined in the table below | [optional] [default to null]
**Guid** | [***Object**](.md) |  | [optional] [default to null]
**Tag** | [***Object**](.md) | Tag to uniquely identify an order | [optional] [default to null]
**InstrumentToken** | [***Object**](.md) | Identifier issued by Upstox used for subscribing to live market quotes | [optional] [default to null]
**PlacedBy** | [***Object**](.md) | Uniquely identifies the user | [optional] [default to null]
**Tradingsymbol** | [***Object**](.md) | Shows the trading symbol of the instrument | [optional] [default to null]
**TradingSymbol** | [***Object**](.md) | Shows the trading symbol of the instrument | [optional] [default to null]
**OrderType** | [***Object**](.md) | Type of order. It can be one of the following MARKET refers to market order&lt;br&gt;LIMIT refers to Limit Order&lt;br&gt;SL refers to Stop Loss Limit&lt;br&gt;SL-M refers to Stop loss market | [optional] [default to null]
**Validity** | [***Object**](.md) | Order validity (DAY- Day and IOC- Immediate or Cancel (IOC) order) | [optional] [default to null]
**TriggerPrice** | [***Object**](.md) | If the order was a stop loss order then the trigger price set is mentioned here | [optional] [default to null]
**DisclosedQuantity** | [***Object**](.md) | The quantity that should be disclosed in the market depth | [optional] [default to null]
**TransactionType** | [***Object**](.md) | Indicates whether the order was a buy or sell order | [optional] [default to null]
**AveragePrice** | [***Object**](.md) | Average price at which the qty got traded | [optional] [default to null]
**FilledQuantity** | [***Object**](.md) | The total quantity traded from this particular order | [optional] [default to null]
**PendingQuantity** | [***Object**](.md) | Pending quantity to be filled | [optional] [default to null]
**StatusMessage** | [***Object**](.md) | Indicates the reason when any order is rejected, not modified or cancelled | [optional] [default to null]
**StatusMessageRaw** | [***Object**](.md) | Description of the order&#x27;s status as received from RMS | [optional] [default to null]
**ExchangeOrderId** | [***Object**](.md) | Unique order ID assigned by the exchange for the order placed | [optional] [default to null]
**ParentOrderId** | [***Object**](.md) | In case the order is part of the second or third leg of a CO or OCO, the parent order ID is indicated here | [optional] [default to null]
**OrderId** | [***Object**](.md) | Unique order ID assigned internally for the order placed | [optional] [default to null]
**Variety** | [***Object**](.md) | Order complexity | [optional] [default to null]
**OrderTimestamp** | [***Object**](.md) | User readable timestamp at which the order was placed | [optional] [default to null]
**ExchangeTimestamp** | [***Object**](.md) | User readable time at which the order was placed or updated | [optional] [default to null]
**IsAmo** | [***Object**](.md) | Signifies if the order is an After Market Order | [optional] [default to null]
**OrderRequestId** | [***Object**](.md) | Apart from 1st order it shows the count of how many requests were sent | [optional] [default to null]
**OrderRefId** | [***Object**](.md) | The order reference ID for which the order must be modified | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

