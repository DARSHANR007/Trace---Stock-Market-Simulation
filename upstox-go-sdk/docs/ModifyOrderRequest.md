# ModifyOrderRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | [***Object**](.md) | Quantity with which the order was placed | [optional] [default to null]
**Validity** | [***Object**](.md) | Order validity (DAY- Day and IOC- Immediate or Cancel (IOC) order) | [default to null]
**Price** | [***Object**](.md) | Price at which the order was placed | [default to null]
**OrderId** | [***Object**](.md) | The order ID for which the order must be modified | [default to null]
**OrderType** | [***Object**](.md) | Type of order. It can be one of the following MARKET refers to market order LIMILT refers to Limit Order SL refers to Stop Loss Limit SL-M refers to Stop Loss Market | [default to null]
**DisclosedQuantity** | [***Object**](.md) | The quantity that should be disclosed in the market depth | [optional] [default to null]
**TriggerPrice** | [***Object**](.md) | If the order is a stop loss order then the trigger price to be set is mentioned here | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

