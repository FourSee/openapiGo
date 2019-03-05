# \PairingRequestsApi

All URIs are relative to *https://foursee-polyrhythm.herokuapp.com:443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePairingRequest**](PairingRequestsApi.md#CreatePairingRequest) | **Post** /pairing_requests | Create a PairingRequest
[**DeletePairingRequest**](PairingRequestsApi.md#DeletePairingRequest) | **Delete** /pairing_requests/{pairingRequestId} | Delete a PairingRequest
[**GetPairingRequest**](PairingRequestsApi.md#GetPairingRequest) | **Get** /pairing_requests/{pairingRequestId} | Get a PairingRequest
[**GetPairingRequests**](PairingRequestsApi.md#GetPairingRequests) | **Get** /pairing_requests | List All PairingRequests
[**UpdatePairingRequest**](PairingRequestsApi.md#UpdatePairingRequest) | **Put** /pairing_requests/{pairingRequestId} | Update a PairingRequest


# **CreatePairingRequest**
> PairingRequest CreatePairingRequest(ctx, pairingRequestCreate)
Create a PairingRequest

Creates a new instance of a `PairingRequest`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pairingRequestCreate** | [**PairingRequestCreate**](PairingRequestCreate.md)| A new &#x60;PairingRequest&#x60; to be created. | 

### Return type

[**PairingRequest**](PairingRequest.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePairingRequest**
> DeletePairingRequest(ctx, pairingRequestId)
Delete a PairingRequest

Deletes an existing `PairingRequest`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pairingRequestId** | **string**| A unique identifier for a &#x60;PairingRequest&#x60;. | 

### Return type

 (empty response body)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPairingRequest**
> PairingRequest GetPairingRequest(ctx, pairingRequestId)
Get a PairingRequest

Gets the details of a single instance of a `PairingRequest`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pairingRequestId** | **string**| A unique identifier for a &#x60;PairingRequest&#x60;. | 

### Return type

[**PairingRequest**](PairingRequest.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPairingRequests**
> PairingRequestsList GetPairingRequests(ctx, )
List All PairingRequests

Gets a list of all `PairingRequest` entities.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PairingRequestsList**](PairingRequestsList.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePairingRequest**
> PairingRequest UpdatePairingRequest(ctx, pairingRequestId, pairingRequest)
Update a PairingRequest

Updates an existing `PairingRequest`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pairingRequestId** | **string**| A unique identifier for a &#x60;PairingRequest&#x60;. | 
  **pairingRequest** | [**PairingRequest**](PairingRequest.md)| Updated &#x60;PairingRequest&#x60; information. | 

### Return type

[**PairingRequest**](PairingRequest.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

