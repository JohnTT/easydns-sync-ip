# \DeleteAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DelUXZoneRec**](DeleteAPI.md#DelUXZoneRec) | **Delete** /zones/async/ux/records/{domain}/{id} | Delete an existing zone record using async zone reloading
[**DelZoneRec**](DeleteAPI.md#DelZoneRec) | **Delete** /zones/records/{domain}/{id} | Delete an existing zone record
[**DeleteDomain**](DeleteAPI.md#DeleteDomain) | **Delete** /domain/{domain} | Delete an existing domain from the system
[**DeleteMailmap**](DeleteAPI.md#DeleteMailmap) | **Delete** /mail/maps/{domain}/{mailmap_id} | This API call allows you to delete an existing mailmap for a domain. To obtain the mailmap_id for a mailmap you should perform a request to list your mailmaps



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
	r, err := apiClient.DeleteAPI.DelUXZoneRec(context.Background(), domain, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteAPI.DelUXZoneRec``: %v\n", err)
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
	r, err := apiClient.DeleteAPI.DelZoneRec(context.Background(), domain, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteAPI.DelZoneRec``: %v\n", err)
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


## DeleteDomain

> ResultModelDomainDelete DeleteDomain(ctx, domain).Execute()

Delete an existing domain from the system

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeleteAPI.DeleteDomain(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteAPI.DeleteDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDomain`: ResultModelDomainDelete
	fmt.Fprintf(os.Stdout, "Response from `DeleteAPI.DeleteDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelDomainDelete**](ResultModelDomainDelete.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMailmap

> ResultModelMailmapDelete DeleteMailmap(ctx, domain, mailmapId).Execute()

This API call allows you to delete an existing mailmap for a domain. To obtain the mailmap_id for a mailmap you should perform a request to list your mailmaps

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
	domain := "example.com" // string | The domain name to create the mailmap under
	mailmapId := int32(1234) // int32 | The unique ID of the mailmap

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeleteAPI.DeleteMailmap(context.Background(), domain, mailmapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeleteAPI.DeleteMailmap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMailmap`: ResultModelMailmapDelete
	fmt.Fprintf(os.Stdout, "Response from `DeleteAPI.DeleteMailmap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name to create the mailmap under | 
**mailmapId** | **int32** | The unique ID of the mailmap | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMailmapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResultModelMailmapDelete**](ResultModelMailmapDelete.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

