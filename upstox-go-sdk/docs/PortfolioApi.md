# {{classname}}

All URIs are relative to *https://api-v2.upstox.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertPositions**](PortfolioApi.md#ConvertPositions) | **Put** /v2/portfolio/convert-position | Convert Positions
[**GetHoldings**](PortfolioApi.md#GetHoldings) | **Get** /v2/portfolio/long-term-holdings | Get Holdings
[**GetMtfPositions**](PortfolioApi.md#GetMtfPositions) | **Get** /v3/portfolio/mtf-positions | Get MTF positions
[**GetPositions**](PortfolioApi.md#GetPositions) | **Get** /v2/portfolio/short-term-positions | Get Positions

# **ConvertPositions**
> ConvertPositionResponse ConvertPositions(ctx, body)
Convert Positions

Convert the margin product of an open position

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConvertPositionRequest**](ConvertPositionRequest.md)|  | 

### Return type

[**ConvertPositionResponse**](ConvertPositionResponse.md)

### Authorization

[OAUTH2](../README.md#OAUTH2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHoldings**
> GetHoldingsResponse GetHoldings(ctx, )
Get Holdings

Fetches the holdings which the user has bought/sold in previous trading sessions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetHoldingsResponse**](GetHoldingsResponse.md)

### Authorization

[OAUTH2](../README.md#OAUTH2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMtfPositions**
> GetPositionResponse GetMtfPositions(ctx, )
Get MTF positions

This API allows you to get MTF positions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetPositionResponse**](GetPositionResponse.md)

### Authorization

[OAUTH2](../README.md#OAUTH2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPositions**
> GetPositionResponse GetPositions(ctx, )
Get Positions

Fetches the current positions for the user for the current day.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetPositionResponse**](GetPositionResponse.md)

### Authorization

[OAUTH2](../README.md#OAUTH2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

