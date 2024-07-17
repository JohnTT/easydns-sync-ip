# \NameServerAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDomainNameservers**](NameServerAPI.md#GetDomainNameservers) | **Get** /domains/ns/{domain} | Returns the list of name servers currently assigned to the provided domain
[**UpdateDomainNameservers**](NameServerAPI.md#UpdateDomainNameservers) | **Post** /domains/ns/{domain} | Updates the name servers assigned to the provided domain name



## GetDomainNameservers

> ResultModelDomainNameservers GetDomainNameservers(ctx, domain).Execute()

Returns the list of name servers currently assigned to the provided domain

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
	domain := "example.com" // string | The domain to get the nameservers for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NameServerAPI.GetDomainNameservers(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NameServerAPI.GetDomainNameservers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainNameservers`: ResultModelDomainNameservers
	fmt.Fprintf(os.Stdout, "Response from `NameServerAPI.GetDomainNameservers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain to get the nameservers for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainNameserversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelDomainNameservers**](ResultModelDomainNameservers.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainNameservers

> ResultModelDomainNameservers UpdateDomainNameservers(ctx, domain).RequestModelDomainNameservers(requestModelDomainNameservers).Execute()

Updates the name servers assigned to the provided domain name

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
	domain := "example.com" // string | The domain to update the nameservers for
	requestModelDomainNameservers := *openapiclient.NewRequestModelDomainNameservers([]string{"dns1.example.com"}) // RequestModelDomainNameservers |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NameServerAPI.UpdateDomainNameservers(context.Background(), domain).RequestModelDomainNameservers(requestModelDomainNameservers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NameServerAPI.UpdateDomainNameservers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDomainNameservers`: ResultModelDomainNameservers
	fmt.Fprintf(os.Stdout, "Response from `NameServerAPI.UpdateDomainNameservers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain to update the nameservers for | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainNameserversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestModelDomainNameservers** | [**RequestModelDomainNameservers**](RequestModelDomainNameservers.md) |  | 

### Return type

[**ResultModelDomainNameservers**](ResultModelDomainNameservers.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

