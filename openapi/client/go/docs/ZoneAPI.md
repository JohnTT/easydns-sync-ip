# \ZoneAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUXZoneRec**](ZoneAPI.md#AddUXZoneRec) | **Put** /zones/async/ux/records/add/{domain}/{type} | Create a new record in a zone using async zone reloading
[**AddZoneRec**](ZoneAPI.md#AddZoneRec) | **Put** /zones/records/add/{domain}/{type} | Create a new record in a zone
[**DelUXZoneRec**](ZoneAPI.md#DelUXZoneRec) | **Delete** /zones/async/ux/records/{domain}/{id} | Delete an existing zone record using async zone reloading
[**DelZoneRec**](ZoneAPI.md#DelZoneRec) | **Delete** /zones/records/{domain}/{id} | Delete an existing zone record
[**ForceZoneReload**](ZoneAPI.md#ForceZoneReload) | **Get** /zones/reload/{domain}/force | Force regeneration of a zone immediately instead of allowing zone reloader to handle it
[**GetZoneSOA**](ZoneAPI.md#GetZoneSOA) | **Get** /zones/records/soa/{domain} | Get the current SOA value for a zone
[**ListGeoRegions**](ZoneAPI.md#ListGeoRegions) | **Get** /zones/geo/region/list | Get a paginated list of all geo region IDs available for geo records
[**ListParsedZone**](ZoneAPI.md#ListParsedZone) | **Get** /zones/records/parsed/{domain} | Get the current zone records for a domain fully parsed like you would see them in the zonefile.
[**ListZone**](ZoneAPI.md#ListZone) | **Get** /zones/records/all/{domain} | Get a listing of all available zone records for a domain
[**ModUXZoneRec**](ZoneAPI.md#ModUXZoneRec) | **Post** /zones/async/ux/records/{id} | Modify an existing zone record using async zone reloading
[**ModZoneRec**](ZoneAPI.md#ModZoneRec) | **Post** /zones/records/{id} | Modify an existing zone record
[**SearchZone**](ZoneAPI.md#SearchZone) | **Get** /zones/records/all/{domain}/search/{keyword} | Get a listing of all available zone records for a domain that match the keyword string



## AddUXZoneRec

> AddUXZoneRec(ctx, domain, type_).Body(body).Execute()

Create a new record in a zone using async zone reloading

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
	type_ := "type__example" // string | The type of resource record (RR) to create (i.e. A, CNAME, MX, etc)
	body := map[string]interface{}{ ... } // map[string]interface{} | The contents of the resource record (RR) to create in JSON encoded format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZoneAPI.AddUXZoneRec(context.Background(), domain, type_).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.AddUXZoneRec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The associated domain name | 
**type_** | **string** | The type of resource record (RR) to create (i.e. A, CNAME, MX, etc) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUXZoneRecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | The contents of the resource record (RR) to create in JSON encoded format | 

### Return type

 (empty response body)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddZoneRec

> AddZoneRec(ctx, domain, type_).RequestModelZoneBodyData(requestModelZoneBodyData).Execute()

Create a new record in a zone

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
	type_ := "type__example" // string | The type of resource record (RR) to create (i.e. A, CNAME, MX, etc)
	requestModelZoneBodyData := *openapiclient.NewRequestModelZoneBodyData("example.com", "test1", "A", "1.2.3.4") // RequestModelZoneBodyData | The contents of the resource record (RR) to create in JSON encoded format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZoneAPI.AddZoneRec(context.Background(), domain, type_).RequestModelZoneBodyData(requestModelZoneBodyData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.AddZoneRec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The associated domain name | 
**type_** | **string** | The type of resource record (RR) to create (i.e. A, CNAME, MX, etc) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddZoneRecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestModelZoneBodyData** | [**RequestModelZoneBodyData**](RequestModelZoneBodyData.md) | The contents of the resource record (RR) to create in JSON encoded format | 

### Return type

 (empty response body)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelUXZoneRec

> DelUXZoneRec(ctx, domain, id).Execute()

Delete an existing zone record using async zone reloading

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
	domain := "domain_example" // string | The domain name that owns the zone record
	id := "id_example" // string | The zone record ID to delete from the zone

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZoneAPI.DelUXZoneRec(context.Background(), domain, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.DelUXZoneRec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name that owns the zone record | 
**id** | **string** | The zone record ID to delete from the zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelUXZoneRecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelZoneRec

> DelZoneRec(ctx, domain, id).Execute()

Delete an existing zone record

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
	domain := "domain_example" // string | The domain name that owns the zone record
	id := "id_example" // string | The zone record ID to delete from the zone

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZoneAPI.DelZoneRec(context.Background(), domain, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.DelZoneRec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name that owns the zone record | 
**id** | **string** | The zone record ID to delete from the zone | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelZoneRecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForceZoneReload

> ForceZoneReload(ctx, domain).Execute()

Force regeneration of a zone immediately instead of allowing zone reloader to handle it

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
	domain := "domain_example" // string | The domain name to force a zone reload for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZoneAPI.ForceZoneReload(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.ForceZoneReload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name to force a zone reload for | 

### Other Parameters

Other parameters are passed through a pointer to a apiForceZoneReloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneSOA

> ResultModelZoneSOA GetZoneSOA(ctx, domain).Execute()

Get the current SOA value for a zone

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
	domain := "domain_example" // string | The domain name to retrieve the SOA value for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneAPI.GetZoneSOA(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.GetZoneSOA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneSOA`: ResultModelZoneSOA
	fmt.Fprintf(os.Stdout, "Response from `ZoneAPI.GetZoneSOA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name to retrieve the SOA value for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneSOARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelZoneSOA**](ResultModelZoneSOA.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.ZoneAPI.ListGeoRegions(context.Background()).Start(start).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.ListGeoRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGeoRegions`: ResultModelGeoList
	fmt.Fprintf(os.Stdout, "Response from `ZoneAPI.ListGeoRegions`: %v\n", resp)
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
	resp, r, err := apiClient.ZoneAPI.ListParsedZone(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.ListParsedZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListParsedZone`: ResultModelParsedZoneList
	fmt.Fprintf(os.Stdout, "Response from `ZoneAPI.ListParsedZone`: %v\n", resp)
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
	resp, r, err := apiClient.ZoneAPI.ListZone(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.ListZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListZone`: ResultModelZoneList
	fmt.Fprintf(os.Stdout, "Response from `ZoneAPI.ListZone`: %v\n", resp)
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


## ModUXZoneRec

> ResultModelZoneRecordBody ModUXZoneRec(ctx, id).RequestModelZoneBodyData(requestModelZoneBodyData).Execute()

Modify an existing zone record using async zone reloading

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
	id := int64(789) // int64 | The ID of the zone record to modify
	requestModelZoneBodyData := *openapiclient.NewRequestModelZoneBodyData("example.com", "test1", "A", "1.2.3.4") // RequestModelZoneBodyData | The contents of the resource record (RR) to modify in JSON encoded format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneAPI.ModUXZoneRec(context.Background(), id).RequestModelZoneBodyData(requestModelZoneBodyData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.ModUXZoneRec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModUXZoneRec`: ResultModelZoneRecordBody
	fmt.Fprintf(os.Stdout, "Response from `ZoneAPI.ModUXZoneRec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the zone record to modify | 

### Other Parameters

Other parameters are passed through a pointer to a apiModUXZoneRecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestModelZoneBodyData** | [**RequestModelZoneBodyData**](RequestModelZoneBodyData.md) | The contents of the resource record (RR) to modify in JSON encoded format | 

### Return type

[**ResultModelZoneRecordBody**](ResultModelZoneRecordBody.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModZoneRec

> ResultModelZoneRecordBody ModZoneRec(ctx, id).RequestModelZoneBodyData(requestModelZoneBodyData).Execute()

Modify an existing zone record

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
	id := int64(789) // int64 | The ID of the zone record to modify
	requestModelZoneBodyData := *openapiclient.NewRequestModelZoneBodyData("example.com", "test1", "A", "1.2.3.4") // RequestModelZoneBodyData | The contents of the resource record (RR) to create in JSON encoded format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZoneAPI.ModZoneRec(context.Background(), id).RequestModelZoneBodyData(requestModelZoneBodyData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.ModZoneRec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModZoneRec`: ResultModelZoneRecordBody
	fmt.Fprintf(os.Stdout, "Response from `ZoneAPI.ModZoneRec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the zone record to modify | 

### Other Parameters

Other parameters are passed through a pointer to a apiModZoneRecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestModelZoneBodyData** | [**RequestModelZoneBodyData**](RequestModelZoneBodyData.md) | The contents of the resource record (RR) to create in JSON encoded format | 

### Return type

[**ResultModelZoneRecordBody**](ResultModelZoneRecordBody.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.ZoneAPI.SearchZone(context.Background(), domain, keyword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZoneAPI.SearchZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchZone`: ResultModelZoneList
	fmt.Fprintf(os.Stdout, "Response from `ZoneAPI.SearchZone`: %v\n", resp)
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

