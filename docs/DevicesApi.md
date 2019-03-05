# \DevicesApi

All URIs are relative to *https://foursee-polyrhythm.herokuapp.com:443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDevice**](DevicesApi.md#CreateDevice) | **Post** /devices | Create a Device
[**DeleteDevice**](DevicesApi.md#DeleteDevice) | **Delete** /devices/{deviceId} | Delete a Device
[**GetDevice**](DevicesApi.md#GetDevice) | **Get** /devices/{deviceId} | Get a Device
[**GetDevices**](DevicesApi.md#GetDevices) | **Get** /devices | List All Devices
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Put** /devices/{deviceId} | Update a Device


# **CreateDevice**
> CreateDevice(ctx, device)
Create a Device

Creates a new instance of a `Device`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **device** | [**Device**](Device.md)| A new &#x60;Device&#x60; to be created. | 

### Return type

 (empty response body)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDevice**
> DeleteDevice(ctx, deviceId)
Delete a Device

Deletes an existing `Device`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **string**| A unique identifier for a &#x60;Device&#x60;. | 

### Return type

 (empty response body)

### Authorization

[agent_auth_key](../README.md#agent_auth_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevice**
> Device GetDevice(ctx, deviceId)
Get a Device

Gets the details of a single instance of a `Device`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **string**| A unique identifier for a &#x60;Device&#x60;. | 

### Return type

[**Device**](Device.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevices**
> DevicesList GetDevices(ctx, )
List All Devices

Gets a list of all `Device` entities.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DevicesList**](DevicesList.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDevice**
> Device UpdateDevice(ctx, deviceId, device)
Update a Device

Updates an existing `Device`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceId** | **string**| A unique identifier for a &#x60;Device&#x60;. | 
  **device** | [**Device**](Device.md)| Updated &#x60;Device&#x60; information. | 

### Return type

[**Device**](Device.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

