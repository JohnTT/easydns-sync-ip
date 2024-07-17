# \GeoAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGeoRegions**](GeoAPI.md#ListGeoRegions) | **Get** /zones/geo/region/list | Get a paginated list of all geo region IDs available for geo records



## ListGeoRegions

> ResultModelGeoList ListGeoRegions(ctx).Start(start).Max(max).Execute()

Get a paginated list of all geo region IDs available for geo records

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
	start := int32(100) // int32 | The number of the record to start the page on when using paginated results (optional)
	max := int32(10) // int32 | The maximum number of the records to return when using paginated results (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeoAPI.ListGeoRegions(context.Background()).Start(start).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.ListGeoRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGeoRegions`: ResultModelGeoList
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.ListGeoRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGeoRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | The number of the record to start the page on when using paginated results | 
 **max** | **int32** | The maximum number of the records to return when using paginated results | 

### Return type

[**ResultModelGeoList**](ResultModelGeoList.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

