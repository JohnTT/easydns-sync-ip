# \GlueAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckRegistryGlue**](GlueAPI.md#CheckRegistryGlue) | **Get** /domains/glue/{domain}/{host}/status | Gets the status of a glue record at the registry to indicate if the glue record has been put in place by the registry
[**CreateGlue**](GlueAPI.md#CreateGlue) | **Put** /domains/glue/{domain} | Creates glue record information at the registry
[**DeleteGlue**](GlueAPI.md#DeleteGlue) | **Delete** /domains/glue/{domain}/{host} | Deletes a glue record that is configured at the registry and is no longer needed. All requests to delete glue records that are still in use by a domain with the same TLD will fail until all associations are removed!
[**GetDomainGlue**](GlueAPI.md#GetDomainGlue) | **Get** /domains/glue/{domain} | Returns any glue records associated with a domain
[**UpdateGlue**](GlueAPI.md#UpdateGlue) | **Post** /domains/glue/{domain} | Updates an existing glue record configured at the registry with new address information



## CheckRegistryGlue

> ResultModelCheckRegistryGlue CheckRegistryGlue(ctx, domain, host).Execute()

Gets the status of a glue record at the registry to indicate if the glue record has been put in place by the registry

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
	domain := "example.com" // string | Just the domain part of the nameserver FQDN
	host := "dns1" // string | Just the hostname part of the nameserver FQDN. If glue record is for root of domain just specify '@'.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlueAPI.CheckRegistryGlue(context.Background(), domain, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlueAPI.CheckRegistryGlue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckRegistryGlue`: ResultModelCheckRegistryGlue
	fmt.Fprintf(os.Stdout, "Response from `GlueAPI.CheckRegistryGlue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Just the domain part of the nameserver FQDN | 
**host** | **string** | Just the hostname part of the nameserver FQDN. If glue record is for root of domain just specify &#39;@&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckRegistryGlueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResultModelCheckRegistryGlue**](ResultModelCheckRegistryGlue.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlue

> ResultModelCreateDomainGlue CreateGlue(ctx, domain).RequestModelCreateDomainGlue(requestModelCreateDomainGlue).Execute()

Creates glue record information at the registry

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
	domain := "example.com" // string | 
	requestModelCreateDomainGlue := *openapiclient.NewRequestModelCreateDomainGlue("dns1.example.com") // RequestModelCreateDomainGlue |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlueAPI.CreateGlue(context.Background(), domain).RequestModelCreateDomainGlue(requestModelCreateDomainGlue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlueAPI.CreateGlue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGlue`: ResultModelCreateDomainGlue
	fmt.Fprintf(os.Stdout, "Response from `GlueAPI.CreateGlue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestModelCreateDomainGlue** | [**RequestModelCreateDomainGlue**](RequestModelCreateDomainGlue.md) |  | 

### Return type

[**ResultModelCreateDomainGlue**](ResultModelCreateDomainGlue.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlue

> ResultModelDeleteGlue DeleteGlue(ctx, domain, host).Execute()

Deletes a glue record that is configured at the registry and is no longer needed. All requests to delete glue records that are still in use by a domain with the same TLD will fail until all associations are removed!

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
	domain := "example.com" // string | Just the domain part of the nameserver FQDN.
	host := "dns1" // string | Just the hostname part of the nameserver FQDN. If glue record is for root of domain just specify '@'.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlueAPI.DeleteGlue(context.Background(), domain, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlueAPI.DeleteGlue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGlue`: ResultModelDeleteGlue
	fmt.Fprintf(os.Stdout, "Response from `GlueAPI.DeleteGlue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Just the domain part of the nameserver FQDN. | 
**host** | **string** | Just the hostname part of the nameserver FQDN. If glue record is for root of domain just specify &#39;@&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResultModelDeleteGlue**](ResultModelDeleteGlue.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainGlue

> ResultModelGetDomainGlue GetDomainGlue(ctx, domain).Execute()

Returns any glue records associated with a domain

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
	domain := "example.com" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlueAPI.GetDomainGlue(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlueAPI.GetDomainGlue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainGlue`: ResultModelGetDomainGlue
	fmt.Fprintf(os.Stdout, "Response from `GlueAPI.GetDomainGlue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainGlueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelGetDomainGlue**](ResultModelGetDomainGlue.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlue

> ResultModelCreateDomainGlue UpdateGlue(ctx, domain).RequestModelCreateDomainGlue(requestModelCreateDomainGlue).Execute()

Updates an existing glue record configured at the registry with new address information

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
	domain := "example.com" // string | Just the domain part of the nameserver FQDN.
	requestModelCreateDomainGlue := *openapiclient.NewRequestModelCreateDomainGlue("dns1.example.com") // RequestModelCreateDomainGlue |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlueAPI.UpdateGlue(context.Background(), domain).RequestModelCreateDomainGlue(requestModelCreateDomainGlue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlueAPI.UpdateGlue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGlue`: ResultModelCreateDomainGlue
	fmt.Fprintf(os.Stdout, "Response from `GlueAPI.UpdateGlue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Just the domain part of the nameserver FQDN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestModelCreateDomainGlue** | [**RequestModelCreateDomainGlue**](RequestModelCreateDomainGlue.md) |  | 

### Return type

[**ResultModelCreateDomainGlue**](ResultModelCreateDomainGlue.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

