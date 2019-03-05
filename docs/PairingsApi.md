# \PairingsApi

All URIs are relative to *https://foursee-polyrhythm.herokuapp.com:443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePairing**](PairingsApi.md#CreatePairing) | **Post** /pairings | Create a Pairing
[**DeletePairing**](PairingsApi.md#DeletePairing) | **Delete** /pairings/{pairingId} | Delete a Pairing
[**GetPairing**](PairingsApi.md#GetPairing) | **Get** /pairings/{pairingId} | Get a Pairing
[**GetPairings**](PairingsApi.md#GetPairings) | **Get** /pairings | List All Pairings
[**UpdatePairing**](PairingsApi.md#UpdatePairing) | **Put** /pairings/{pairingId} | Update a Pairing


# **CreatePairing**
> CreatePairing(ctx, pairing)
Create a Pairing

Creates a new instance of a `Pairing`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pairing** | [**Pairing**](Pairing.md)| A new &#x60;Pairing&#x60; to be created. | 

### Return type

 (empty response body)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePairing**
> DeletePairing(ctx, pairingId)
Delete a Pairing

Deletes an existing `Pairing`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pairingId** | **string**| A unique identifier for a &#x60;Pairing&#x60;. | 

### Return type

 (empty response body)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPairing**
> Pairing GetPairing(ctx, pairingId)
Get a Pairing

Gets the details of a single instance of a `Pairing`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pairingId** | **string**| A unique identifier for a &#x60;Pairing&#x60;. | 

### Return type

[**Pairing**](Pairing.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPairings**
> PairingsList GetPairings(ctx, )
List All Pairings

Gets a list of all `Pairing` entities.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PairingsList**](PairingsList.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePairing**
> UpdatePairing(ctx, pairingId, pairing)
Update a Pairing

Updates an existing `Pairing`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pairingId** | **string**| A unique identifier for a &#x60;Pairing&#x60;. | 
  **pairing** | [**Pairing**](Pairing.md)| Updated &#x60;Pairing&#x60; information. | 

### Return type

 (empty response body)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

