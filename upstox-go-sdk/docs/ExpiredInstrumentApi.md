# {{classname}}

All URIs are relative to *https://api-v2.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExpiredFutureContracts**](ExpiredInstrumentApi.md#GetExpiredFutureContracts) | **Get** /v2/expired-instruments/future/contract | Expired instruments - Get future contracts
[**GetExpiredHistoricalCandleData**](ExpiredInstrumentApi.md#GetExpiredHistoricalCandleData) | **Get** /v2/expired-instruments/historical-candle/{expired_instrument_key}/{interval}/{to_date}/{from_date} | Expired Historical candle data
[**GetExpiredOptionContracts**](ExpiredInstrumentApi.md#GetExpiredOptionContracts) | **Get** /v2/expired-instruments/option/contract | Get expired option contracts
[**GetExpiriesResponse**](ExpiredInstrumentApi.md#GetExpiriesResponse) | **Get** /v2/expired-instruments/expiries | Expired instruments - Get expiries

# **GetExpiredFutureContracts**
> GetExpiredFuturesContractResponse GetExpiredFutureContracts(ctx, instrumentKey, expiryDate)
Expired instruments - Get future contracts

This API provides the functionality to retrieve expired future contracts for a given instrument key and expiry date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | [**Object**](.md)| Instrument Key of asset | 
  **expiryDate** | [**Object**](.md)| Expiry date of the instrument | 

### Return type

[**GetExpiredFuturesContractResponse**](GetExpiredFuturesContractResponse.md)

### Authorization

[OAUTH2](../README.md#OAUTH2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExpiredHistoricalCandleData**
> GetHistoricalCandleResponse GetExpiredHistoricalCandleData(ctx, expiredInstrumentKey, interval, toDate, fromDate)
Expired Historical candle data

Get Expired OHLC values for all instruments across various timeframes. Expired Historical data can be fetched for the following durations. 1minute: last 1 month candles till endDate 30minute: last 1 year candles till endDate day: last 1 year candles till endDate week: last 10 year candles till endDate month: last 10 year candles till endDate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **expiredInstrumentKey** | [**Object**](.md)| Expired Instrument Key of asset | 
  **interval** | [**Object**](.md)| Interval to get expired ohlc data | 
  **toDate** | [**Object**](.md)| to date | 
  **fromDate** | [**Object**](.md)| from date | 

### Return type

[**GetHistoricalCandleResponse**](GetHistoricalCandleResponse.md)

### Authorization

[OAUTH2](../README.md#OAUTH2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExpiredOptionContracts**
> GetOptionContractResponse GetExpiredOptionContracts(ctx, instrumentKey, expiryDate)
Get expired option contracts

This API provides the functionality to retrieve the expired option contracts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | [**Object**](.md)| Instrument key for an underlying symbol | 
  **expiryDate** | [**Object**](.md)| Expiry date in format: YYYY-mm-dd | 

### Return type

[**GetOptionContractResponse**](GetOptionContractResponse.md)

### Authorization

[OAUTH2](../README.md#OAUTH2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExpiriesResponse**
> GetExpiriesResponse GetExpiriesResponse(ctx, instrumentKey)
Expired instruments - Get expiries

This API provides the functionality to retrieve expiry dates for a given instrument key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | [**Object**](.md)| Instrument Key of asset | 

### Return type

[**GetExpiriesResponse**](GetExpiriesResponse.md)

### Authorization

[OAUTH2](../README.md#OAUTH2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

