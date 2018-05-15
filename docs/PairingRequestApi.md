# \PairingRequestApi

All URIs are relative to *https://localhost:3000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePairingRequest**](PairingRequestApi.md#CreatePairingRequest) | **Post** /pairing_requests | Create pairing request


# **CreatePairingRequest**
> PairingRequestCreateResponse CreatePairingRequest(ctx, pairingrequestinput)
Create pairing request

Initiates a pairing request. Assumed to be done by the PolyRhythym agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pairingrequestinput** | [**PairingRequestInput**](PairingRequestInput.md)| The created pairing request | 

### Return type

[**PairingRequestCreateResponse**](PairingRequestCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

