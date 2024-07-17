# \AsyncAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUXZoneRec**](AsyncAPI.md#AddUXZoneRec) | **Put** /zones/async/ux/records/add/{domain}/{type} | Create a new record in a zone using async zone reloading
[**DelUXZoneRec**](AsyncAPI.md#DelUXZoneRec) | **Delete** /zones/async/ux/records/{domain}/{id} | Delete an existing zone record using async zone reloading
[**ForceZoneReload**](AsyncAPI.md#ForceZoneReload) | **Get** /zones/reload/{domain}/force | Force regeneration of a zone immediately instead of allowing zone reloader to handle it
[**ModUXZoneRec**](AsyncAPI.md#ModUXZoneRec) | **Post** /zones/async/ux/records/{id} | Modify an existing zone record using async zone reloading



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
	r, err := apiClient.AsyncAPI.AddUXZoneRec(context.Background(), domain, type_).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsyncAPI.AddUXZoneRec``: %v\n", err)
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
	r, err := apiClient.AsyncAPI.DelUXZoneRec(context.Background(), domain, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsyncAPI.DelUXZoneRec``: %v\n", err)
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
	r, err := apiClient.AsyncAPI.ForceZoneReload(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsyncAPI.ForceZoneReload``: %v\n", err)
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
	resp, r, err := apiClient.AsyncAPI.ModUXZoneRec(context.Background(), id).RequestModelZoneBodyData(requestModelZoneBodyData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsyncAPI.ModUXZoneRec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModUXZoneRec`: ResultModelZoneRecordBody
	fmt.Fprintf(os.Stdout, "Response from `AsyncAPI.ModUXZoneRec`: %v\n", resp)
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

