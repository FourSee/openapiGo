# PairingRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for this pairing request. | 
**DeviceName** | **string** | Ideally, a human-readable name for the device that created the pairing request.  Will default to using the &#x60;hostname&#x60;. | 
**ShortId** | **string** | A URL-safe base62-encoded representation of the ID | 
**Status** | **string** | Status of the pairing request | 
**AcceptedCryptoKey** | **string** | If this pairing request has been accepted, this will be the base64 encoded public cyrpto key of the device that accepted the pairing reques. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | 
**AcceptUrl** | **string** | The URL to accept this pairing request.  Will use the &#x60;polyrhythm&#x60; scheme, and follow this pattern: &#x60;polyrhythm://pr/{short_id}&#x60; | 
**IpAddress** | **string** | IP address where the pairing request came from. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


