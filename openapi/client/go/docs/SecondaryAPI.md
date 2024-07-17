# \SecondaryAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetPrimaryNS**](SecondaryAPI.md#SetPrimaryNS) | **Post** /domains/primary_ns/{domain} | Set the primary NS value for a domain. Will also update domain to make it secondary



## SetPrimaryNS

> ResultSetPrimaryNS SetPrimaryNS(ctx, domain).RequestPrimaryNSBody(requestPrimaryNSBody).Execute()

Set the primary NS value for a domain. Will also update domain to make it secondary

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
	domain := "domain_example" // string | 
	requestPrimaryNSBody := *openapiclient.NewRequestPrimaryNSBody("1.2.3.4") // RequestPrimaryNSBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecondaryAPI.SetPrimaryNS(context.Background(), domain).RequestPrimaryNSBody(requestPrimaryNSBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecondaryAPI.SetPrimaryNS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPrimaryNS`: ResultSetPrimaryNS
	fmt.Fprintf(os.Stdout, "Response from `SecondaryAPI.SetPrimaryNS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPrimaryNSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestPrimaryNSBody** | [**RequestPrimaryNSBody**](RequestPrimaryNSBody.md) |  | 

### Return type

[**ResultSetPrimaryNS**](ResultSetPrimaryNS.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

