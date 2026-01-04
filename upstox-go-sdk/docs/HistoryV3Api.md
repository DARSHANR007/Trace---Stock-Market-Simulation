# {{classname}}

All URIs are relative to *https://api-v2.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricalCandleData**](HistoryV3Api.md#GetHistoricalCandleData) | **Get** /v3/historical-candle/{instrumentKey}/{unit}/{interval}/{to_date} | Historical candle data
[**GetHistoricalCandleData1**](HistoryV3Api.md#GetHistoricalCandleData1) | **Get** /v3/historical-candle/{instrumentKey}/{unit}/{interval}/{to_date}/{from_date} | Historical candle data
[**GetIntraDayCandleData**](HistoryV3Api.md#GetIntraDayCandleData) | **Get** /v3/historical-candle/intraday/{instrumentKey}/{unit}/{interval} | Intra day candle data

# **GetHistoricalCandleData**
> GetHistoricalCandleResponse GetHistoricalCandleData(ctx, instrumentKey, unit, interval, toDate)
Historical candle data

Get OHLC values for all instruments for the present trading day with expanded interval options.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | [**Object**](.md)|  | 
  **unit** | [**Object**](.md)|  | 
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

# **GetHistoricalCandleData1**
> GetHistoricalCandleResponse GetHistoricalCandleData1(ctx, instrumentKey, unit, interval, toDate, fromDate)
Historical candle data

Get OHLC values for all instruments for the present trading day with expanded interval options

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | [**Object**](.md)|  | 
  **unit** | [**Object**](.md)|  | 
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

# **GetIntraDayCandleData**
> GetIntraDayCandleResponse GetIntraDayCandleData(ctx, instrumentKey, unit, interval)
Intra day candle data

Get OHLC values for all instruments for the present trading day

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instrumentKey** | [**Object**](.md)|  | 
  **unit** | [**Object**](.md)|  | 
  **interval** | [**Object**](.md)|  | 

### Return type

[**GetIntraDayCandleResponse**](GetIntraDayCandleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

