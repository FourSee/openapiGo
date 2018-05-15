# \MessageApi

All URIs are relative to *https://localhost:3000/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendMessage**](MessageApi.md#SendMessage) | **Post** /messages | Send a message


# **SendMessage**
> SendMessage(ctx, messageinput)
Send a message

Sends a message to a paired device. Recipient and sender are identified via encryption & signature metadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **messageinput** | [**MessageInput**](MessageInput.md)| The message to send | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

