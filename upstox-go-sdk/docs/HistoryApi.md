# {{classname}}

All URIs are relative to *https://api-v2.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricalCandleData2**](HistoryApi.md#GetHistoricalCandleData2) | **Get** /v2/historical-candle/{instrumentKey}/{interval}/{to_date} | Historical candle data
[**GetHistoricalCandleData3**](HistoryApi.md#GetHistoricalCandleData3) | **Get** /v2/historical-candle/{instrumentKey}/{interval}/{to_date}/{from_date} | Historical candle data
[**GetIntraDayCandleData1**](HistoryApi.md#GetIntraDayCandleData1) | **Get** /v2/historical-candle/intraday/{instrumentKey}/{interval} | Intra day candle data

# **GetHistoricalCandleData2**
> GetHistoricalCandleResponse GetHistoricalCandleData2(ctx, instrumentKey, interval, toDate)
Historical candle data

Get OHLC values for all instruments across various timeframes. Historical data can be fetched for the following durations. 1minute: last 1 month candles till endDate 30minute: last 1 year candles till endDate day: last 1 year candles till endDate week: last 10 year candles till endDate month: last 10 year candles till endDate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | [**Object**](.md)|  | 
  **interval** | [**Object**](.md)|  | 
  **toDate** | [**Object**](.md)|  | 

### Return type

[**GetHistoricalCandleResponse**](GetHistoricalCandleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoricalCandleData3**
> GetHistoricalCandleResponse GetHistoricalCandleData3(ctx, instrumentKey, interval, toDate, fromDate)
Historical candle data

Get OHLC values for all instruments across various timeframes. Historical data can be fetched for the following durations. 1minute: last 1 month candles till endDate 30minute: last 1 year candles till endDate day: last 1 year candles till endDate week: last 10 year candles till endDate month: last 10 year candles till endDate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | [**Object**](.md)|  | 
  **interval** | [**Object**](.md)|  | 
  **toDate** | [**Object**](.md)|  | 
  **fromDate** | [**Object**](.md)|  | 

### Return type

[**GetHistoricalCandleResponse**](GetHistoricalCandleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIntraDayCandleData1**
> GetIntraDayCandleResponse GetIntraDayCandleData1(ctx, instrumentKey, interval)
Intra day candle data

Get OHLC values for all instruments for the present trading day

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | [**Object**](.md)|  | 
  **interval** | [**Object**](.md)|  | 

### Return type

[**GetIntraDayCandleResponse**](GetIntraDayCandleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

