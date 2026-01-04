# {{classname}}

All URIs are relative to *https://api-v2.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLtp**](MarketQuoteV3Api.md#GetLtp) | **Get** /v3/market-quote/ltp | Market quotes and instruments - LTP quotes.
[**GetMarketQuoteOHLCV3**](MarketQuoteV3Api.md#GetMarketQuoteOHLCV3) | **Get** /v3/market-quote/ohlc | Market quotes and instruments - OHLC quotes
[**GetMarketQuoteOptionGreek**](MarketQuoteV3Api.md#GetMarketQuoteOptionGreek) | **Get** /v3/market-quote/option-greek | Market quotes and instruments - Option Greek

# **GetLtp**
> GetMarketQuoteLastTradedPriceResponseV3 GetLtp(ctx, optional)
Market quotes and instruments - LTP quotes.

This API provides the functionality to retrieve the LTP quotes for one or more instruments.This API returns the LTPs of up to 500 instruments in one go.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MarketQuoteV3ApiGetLtpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketQuoteV3ApiGetLtpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instrumentKey** | [**optional.Interface of Object**](.md)| Comma separated list of instrument keys | 

### Return type

[**GetMarketQuoteLastTradedPriceResponseV3**](GetMarketQuoteLastTradedPriceResponseV3.md)

### Authorization

[OAUTH2](../README.md#OAUTH2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketQuoteOHLCV3**
> GetMarketQuoteOhlcResponseV3 GetMarketQuoteOHLCV3(ctx, interval, optional)
Market quotes and instruments - OHLC quotes

This API provides the functionality to retrieve the OHLC quotes for one or more instruments.This API returns the OHLC snapshots of up to 500 instruments in one go.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interval** | [**Object**](.md)| Interval to get ohlc data | 
 **optional** | ***MarketQuoteV3ApiGetMarketQuoteOHLCV3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketQuoteV3ApiGetMarketQuoteOHLCV3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instrumentKey** | [**optional.Interface of Object**](.md)| Comma separated list of instrument keys | 

### Return type

[**GetMarketQuoteOhlcResponseV3**](GetMarketQuoteOHLCResponseV3.md)

### Authorization

[OAUTH2](../README.md#OAUTH2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarketQuoteOptionGreek**
> GetMarketQuoteOptionGreekResponseV3 GetMarketQuoteOptionGreek(ctx, optional)
Market quotes and instruments - Option Greek

This API provides the functionality to retrieve the Option Greek data for one or more instruments.This API returns the Option Greek data of up to 500 instruments in one go.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MarketQuoteV3ApiGetMarketQuoteOptionGreekOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarketQuoteV3ApiGetMarketQuoteOptionGreekOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instrumentKey** | [**optional.Interface of Object**](.md)| Comma separated list of instrument keys | 

### Return type

[**GetMarketQuoteOptionGreekResponseV3**](GetMarketQuoteOptionGreekResponseV3.md)

### Authorization

[OAUTH2](../README.md#OAUTH2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

