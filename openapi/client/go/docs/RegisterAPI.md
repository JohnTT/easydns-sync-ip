# \RegisterAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomain**](RegisterAPI.md#CreateDomain) | **Put** /domains/add/{domain} | Add a new domain to the system including registering the domain at the registry if applicable. When registering domains additional fields may be required depending on the TLD



## CreateDomain

> ResultModelDomainCreate CreateDomain(ctx, domain).CreateDomainRequest(createDomainRequest).Execute()

Add a new domain to the system including registering the domain at the registry if applicable. When registering domains additional fields may be required depending on the TLD

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
	domain := "example.com" // string | The domain name to add to the system
	createDomainRequest := openapiclient.createDomain_request{BodyDomainCreateDNSOnly: openapiclient.NewBodyDomainCreateDNSOnly("Service_example", int32(123), "Currency_example", int32(123))} // CreateDomainRequest | Request data containing relevant domain creation information needed to add a domain to the system

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegisterAPI.CreateDomain(context.Background(), domain).CreateDomainRequest(createDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegisterAPI.CreateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDomain`: ResultModelDomainCreate
	fmt.Fprintf(os.Stdout, "Response from `RegisterAPI.CreateDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name to add to the system | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDomainRequest** | [**CreateDomainRequest**](CreateDomainRequest.md) | Request data containing relevant domain creation information needed to add a domain to the system | 

### Return type

[**ResultModelDomainCreate**](ResultModelDomainCreate.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

