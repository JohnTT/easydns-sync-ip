# \ListAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGeoRegions**](ListAPI.md#ListGeoRegions) | **Get** /zones/geo/region/list | Get a paginated list of all geo region IDs available for geo records
[**ListMailmaps**](ListAPI.md#ListMailmaps) | **Get** /mail/maps/{domain} | To get a listing of the existing mailmaps for a domain this API call can be used. This call will return all pertinent details about each mailmap created for the domain.
[**ListParsedZone**](ListAPI.md#ListParsedZone) | **Get** /zones/records/parsed/{domain} | Get the current zone records for a domain fully parsed like you would see them in the zonefile.
[**ListUserDomains**](ListAPI.md#ListUserDomains) | **Get** /domains/list/{user} | Get a listing of all domains associated with a user
[**ListZone**](ListAPI.md#ListZone) | **Get** /zones/records/all/{domain} | Get a listing of all available zone records for a domain
[**SearchZone**](ListAPI.md#SearchZone) | **Get** /zones/records/all/{domain}/search/{keyword} | Get a listing of all available zone records for a domain that match the keyword string



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
	resp, r, err := apiClient.ListAPI.ListGeoRegions(context.Background()).Start(start).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.ListGeoRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGeoRegions`: ResultModelGeoList
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.ListGeoRegions`: %v\n", resp)
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


## ListMailmaps

> ResultModelMailmapList ListMailmaps(ctx, domain).Execute()

To get a listing of the existing mailmaps for a domain this API call can be used. This call will return all pertinent details about each mailmap created for the domain.

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
	domain := "domain_example" // string | The domain name the mailmap belongs to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAPI.ListMailmaps(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.ListMailmaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMailmaps`: ResultModelMailmapList
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.ListMailmaps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name the mailmap belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMailmapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelMailmapList**](ResultModelMailmapList.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.ListAPI.ListParsedZone(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.ListParsedZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListParsedZone`: ResultModelParsedZoneList
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.ListParsedZone`: %v\n", resp)
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


## ListUserDomains

> ResultModelUserDomainList ListUserDomains(ctx, user).Execute()

Get a listing of all domains associated with a user

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
	user := "myusername" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAPI.ListUserDomains(context.Background(), user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.ListUserDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserDomains`: ResultModelUserDomainList
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.ListUserDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelUserDomainList**](ResultModelUserDomainList.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListZone

> ResultModelZoneList ListZone(ctx, domain).Execute()

Get a listing of all available zone records for a domain

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListAPI.ListZone(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.ListZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListZone`: ResultModelZoneList
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.ListZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The associated domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelZoneList**](ResultModelZoneList.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.ListAPI.SearchZone(context.Background(), domain, keyword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.SearchZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchZone`: ResultModelZoneList
	fmt.Fprintf(os.Stdout, "Response from `ListAPI.SearchZone`: %v\n", resp)
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

