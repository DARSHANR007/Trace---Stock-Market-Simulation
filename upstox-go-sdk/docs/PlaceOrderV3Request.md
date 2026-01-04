# PlaceOrderV3Request

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | [***Object**](.md) | Quantity with which the order is to be placed | [default to null]
**Product** | [***Object**](.md) | Signifies if the order was either Intraday, Delivery, CO or OCO | [default to null]
**Validity** | [***Object**](.md) | It can be one of the following - DAY(default), IOC | [default to null]
**Price** | [***Object**](.md) | Price at which the order will be placed | [default to null]
**Tag** | [***Object**](.md) |  | [optional] [default to null]
**Slice** | [***Object**](.md) |  | [optional] [default to null]
**InstrumentToken** | [***Object**](.md) | Key of the instrument | [default to null]
**OrderType** | [***Object**](.md) | Type of order. It can be one of the following MARKET refers to market order LIMIT refers to Limit Order SL refers to Stop Loss Limit SL-M refers to Stop Loss Market | [default to null]
**TransactionType** | [***Object**](.md) | Indicates whether its a buy or sell order | [default to null]
**DisclosedQuantity** | [***Object**](.md) | The quantity that should be disclosed in the market depth | [default to null]
**TriggerPrice** | [***Object**](.md) | If the order is a stop loss order then the trigger price to be set is mentioned here | [default to null]
**IsAmo** | [***Object**](.md) | Signifies if the order is an After Market Order | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

