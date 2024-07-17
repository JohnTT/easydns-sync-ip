# \SearchAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchZone**](SearchAPI.md#SearchZone) | **Get** /zones/records/all/{domain}/search/{keyword} | Get a listing of all available zone records for a domain that match the keyword string



## SearchZone

> ResultModelZoneList SearchZone(ctx, domain, keyword).Execute()

Get a listing of all available zone records for a domain that match the keyword string

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
	domain := "domain_example" // string | The associated domain name
	keyword := "keyword_example" // string | The keyword or string to match against existing resource records

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchZone(context.Background(), domain, keyword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchZone`: ResultModelZoneList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The associated domain name | 
**keyword** | **string** | The keyword or string to match against existing resource records | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResultModelZoneList**](ResultModelZoneList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

