# \ParsedAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListParsedZone**](ParsedAPI.md#ListParsedZone) | **Get** /zones/records/parsed/{domain} | Get the current zone records for a domain fully parsed like you would see them in the zonefile.



## ListParsedZone

> ResultModelParsedZoneList ListParsedZone(ctx, domain).Execute()

Get the current zone records for a domain fully parsed like you would see them in the zonefile.

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
	domain := "example.com" // string | The domain name the zone records belong to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParsedAPI.ListParsedZone(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParsedAPI.ListParsedZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListParsedZone`: ResultModelParsedZoneList
	fmt.Fprintf(os.Stdout, "Response from `ParsedAPI.ListParsedZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name the zone records belong to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListParsedZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelParsedZoneList**](ResultModelParsedZoneList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

