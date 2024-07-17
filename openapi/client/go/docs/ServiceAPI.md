# \ServiceAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceDescription**](ServiceAPI.md#GetServiceDescription) | **Get** /services/{service_id}/description | Returns a description of the provided service
[**GetSubscriptionServiceDescription**](ServiceAPI.md#GetSubscriptionServiceDescription) | **Get** /services/subscription/{subscription_id}/description | Returns a description of the service used by the provided subscription block



## GetServiceDescription

> ResultModelServiceDescription GetServiceDescription(ctx, serviceId).Execute()

Returns a description of the provided service

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serviceId := int32(99) // int32 | The service that the description is being requested for in request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.GetServiceDescription(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetServiceDescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceDescription`: ResultModelServiceDescription
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetServiceDescription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **int32** | The service that the description is being requested for in request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelServiceDescription**](ResultModelServiceDescription.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionServiceDescription

> ResultModelSubscriptionDescription GetSubscriptionServiceDescription(ctx, subscriptionId).Execute()

Returns a description of the service used by the provided subscription block

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	subscriptionId := int32(9001) // int32 | The subscription block ID to return the description for in request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.GetSubscriptionServiceDescription(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetSubscriptionServiceDescription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptionServiceDescription`: ResultModelSubscriptionDescription
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetSubscriptionServiceDescription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int32** | The subscription block ID to return the description for in request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionServiceDescriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelSubscriptionDescription**](ResultModelSubscriptionDescription.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

