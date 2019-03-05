# \MessagesApi

All URIs are relative to *https://foursee-polyrhythm.herokuapp.com:443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessage**](MessagesApi.md#CreateMessage) | **Post** /messages | Create a Message
[**DeleteMessage**](MessagesApi.md#DeleteMessage) | **Delete** /messages/{messageId} | Delete a Message
[**GetMessage**](MessagesApi.md#GetMessage) | **Get** /messages/{messageId} | Get a Message
[**GetMessages**](MessagesApi.md#GetMessages) | **Get** /messages | List All Messages


# **CreateMessage**
> CreateMessage(ctx, message)
Create a Message

Creates a new instance of a `Message`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **message** | [**Message**](Message.md)| A new &#x60;Message&#x60; to be created. | 

### Return type

 (empty response body)

### Authorization

[agent_auth_key](../README.md#agent_auth_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMessage**
> DeleteMessage(ctx, messageId)
Delete a Message

Deletes an existing `Message`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **messageId** | **string**| A unique identifier for a &#x60;Message&#x60;. | 

### Return type

 (empty response body)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessage**
> Message GetMessage(ctx, messageId)
Get a Message

Gets the details of a single instance of a `Message`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **messageId** | **string**| A unique identifier for a &#x60;Message&#x60;. | 

### Return type

[**Message**](Message.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessages**
> MessagesList GetMessages(ctx, )
List All Messages

Gets a list of all `Message` entities.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MessagesList**](MessagesList.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

