# \UsersApi

All URIs are relative to *https://foursee-polyrhythm.herokuapp.com:443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /users | Create a User
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /users/{userId} | Delete a User
[**GetUser**](UsersApi.md#GetUser) | **Get** /users/{userId} | Get a User
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /users | List All Users
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /users/{userId} | Update a User


# **CreateUser**
> User CreateUser(ctx, user)
Create a User

Creates a new instance of a `User`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **user** | [**User**](User.md)| A new &#x60;User&#x60; to be created. | 

### Return type

[**User**](User.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, userId)
Delete a User

Deletes an existing `User`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| A unique identifier for a &#x60;User&#x60;. | 

### Return type

 (empty response body)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> User GetUser(ctx, userId)
Get a User

Gets the details of a single instance of a `User`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| A unique identifier for a &#x60;User&#x60;. | 

### Return type

[**User**](User.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> []User GetUsers(ctx, )
List All Users

Gets a list of all `User` entities.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]User**](User.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> User UpdateUser(ctx, userId, user)
Update a User

Updates an existing `User`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| A unique identifier for a &#x60;User&#x60;. | 
  **user** | [**User**](User.md)| Updated &#x60;User&#x60; information. | 

### Return type

[**User**](User.md)

### Authorization

[Auth0](../README.md#Auth0)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

