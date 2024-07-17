# \DomainAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckRegistryGlue**](DomainAPI.md#CheckRegistryGlue) | **Get** /domains/glue/{domain}/{host}/status | Gets the status of a glue record at the registry to indicate if the glue record has been put in place by the registry
[**CreateDomain**](DomainAPI.md#CreateDomain) | **Put** /domains/add/{domain} | Add a new domain to the system including registering the domain at the registry if applicable. When registering domains additional fields may be required depending on the TLD
[**CreateGlue**](DomainAPI.md#CreateGlue) | **Put** /domains/glue/{domain} | Creates glue record information at the registry
[**DeleteDomain**](DomainAPI.md#DeleteDomain) | **Delete** /domain/{domain} | Delete an existing domain from the system
[**DeleteGlue**](DomainAPI.md#DeleteGlue) | **Delete** /domains/glue/{domain}/{host} | Deletes a glue record that is configured at the registry and is no longer needed. All requests to delete glue records that are still in use by a domain with the same TLD will fail until all associations are removed!
[**GetDomainGlue**](DomainAPI.md#GetDomainGlue) | **Get** /domains/glue/{domain} | Returns any glue records associated with a domain
[**GetDomainInfo**](DomainAPI.md#GetDomainInfo) | **Get** /domain/{domain} | Get details about a domain on the system
[**GetDomainNameservers**](DomainAPI.md#GetDomainNameservers) | **Get** /domains/ns/{domain} | Returns the list of name servers currently assigned to the provided domain
[**GetRegStatus**](DomainAPI.md#GetRegStatus) | **Get** /domains/regstatus | Get the reglock and renewal status for a users list of domains
[**ListUserDomains**](DomainAPI.md#ListUserDomains) | **Get** /domains/list/{user} | Get a listing of all domains associated with a user
[**SetPrimaryNS**](DomainAPI.md#SetPrimaryNS) | **Post** /domains/primary_ns/{domain} | Set the primary NS value for a domain. Will also update domain to make it secondary
[**SetRegStatus**](DomainAPI.md#SetRegStatus) | **Post** /domains/regstatus | Set the reglock status of a set of domains owned by a user
[**UpdateDomainNameservers**](DomainAPI.md#UpdateDomainNameservers) | **Post** /domains/ns/{domain} | Updates the name servers assigned to the provided domain name
[**UpdateGlue**](DomainAPI.md#UpdateGlue) | **Post** /domains/glue/{domain} | Updates an existing glue record configured at the registry with new address information



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
	resp, r, err := apiClient.DomainAPI.CheckRegistryGlue(context.Background(), domain, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.CheckRegistryGlue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckRegistryGlue`: ResultModelCheckRegistryGlue
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.CheckRegistryGlue`: %v\n", resp)
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
	resp, r, err := apiClient.DomainAPI.CreateDomain(context.Background(), domain).CreateDomainRequest(createDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.CreateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDomain`: ResultModelDomainCreate
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.CreateDomain`: %v\n", resp)
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
	resp, r, err := apiClient.DomainAPI.CreateGlue(context.Background(), domain).RequestModelCreateDomainGlue(requestModelCreateDomainGlue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.CreateGlue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGlue`: ResultModelCreateDomainGlue
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.CreateGlue`: %v\n", resp)
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
	resp, r, err := apiClient.DomainAPI.DeleteDomain(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DeleteDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDomain`: ResultModelDomainDelete
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DeleteDomain`: %v\n", resp)
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
	resp, r, err := apiClient.DomainAPI.DeleteGlue(context.Background(), domain, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DeleteGlue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGlue`: ResultModelDeleteGlue
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DeleteGlue`: %v\n", resp)
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
	resp, r, err := apiClient.DomainAPI.GetDomainGlue(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.GetDomainGlue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainGlue`: ResultModelGetDomainGlue
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.GetDomainGlue`: %v\n", resp)
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


## GetDomainInfo

> ResultModelDomainInfo GetDomainInfo(ctx, domain).Execute()

Get details about a domain on the system

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
	resp, r, err := apiClient.DomainAPI.GetDomainInfo(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.GetDomainInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainInfo`: ResultModelDomainInfo
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.GetDomainInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelDomainInfo**](ResultModelDomainInfo.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainNameservers

> ResultModelDomainNameservers GetDomainNameservers(ctx, domain).Execute()

Returns the list of name servers currently assigned to the provided domain

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
	domain := "example.com" // string | The domain to get the nameservers for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.GetDomainNameservers(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.GetDomainNameservers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainNameservers`: ResultModelDomainNameservers
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.GetDomainNameservers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain to get the nameservers for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainNameserversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResultModelDomainNameservers**](ResultModelDomainNameservers.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegStatus

> ResultModelRegStatus GetRegStatus(ctx).Execute()

Get the reglock and renewal status for a users list of domains

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.GetRegStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.GetRegStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegStatus`: ResultModelRegStatus
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.GetRegStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegStatusRequest struct via the builder pattern


### Return type

[**ResultModelRegStatus**](ResultModelRegStatus.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

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
	resp, r, err := apiClient.DomainAPI.ListUserDomains(context.Background(), user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.ListUserDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserDomains`: ResultModelUserDomainList
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.ListUserDomains`: %v\n", resp)
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
	resp, r, err := apiClient.DomainAPI.SetPrimaryNS(context.Background(), domain).RequestPrimaryNSBody(requestPrimaryNSBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.SetPrimaryNS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPrimaryNS`: ResultSetPrimaryNS
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.SetPrimaryNS`: %v\n", resp)
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


## SetRegStatus

> ResultModelSetRegStatus SetRegStatus(ctx).RequestModelSetRegStatus(requestModelSetRegStatus).Execute()

Set the reglock status of a set of domains owned by a user

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
	requestModelSetRegStatus := []openapiclient.RequestModelSetRegStatus{*openapiclient.NewRequestModelSetRegStatus("Domain_example")} // []RequestModelSetRegStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.SetRegStatus(context.Background()).RequestModelSetRegStatus(requestModelSetRegStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.SetRegStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRegStatus`: ResultModelSetRegStatus
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.SetRegStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRegStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestModelSetRegStatus** | [**[]RequestModelSetRegStatus**](RequestModelSetRegStatus.md) |  | 

### Return type

[**ResultModelSetRegStatus**](ResultModelSetRegStatus.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomainNameservers

> ResultModelDomainNameservers UpdateDomainNameservers(ctx, domain).RequestModelDomainNameservers(requestModelDomainNameservers).Execute()

Updates the name servers assigned to the provided domain name

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
	domain := "example.com" // string | The domain to update the nameservers for
	requestModelDomainNameservers := *openapiclient.NewRequestModelDomainNameservers([]string{"dns1.example.com"}) // RequestModelDomainNameservers |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.UpdateDomainNameservers(context.Background(), domain).RequestModelDomainNameservers(requestModelDomainNameservers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.UpdateDomainNameservers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDomainNameservers`: ResultModelDomainNameservers
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.UpdateDomainNameservers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain to update the nameservers for | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainNameserversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestModelDomainNameservers** | [**RequestModelDomainNameservers**](RequestModelDomainNameservers.md) |  | 

### Return type

[**ResultModelDomainNameservers**](ResultModelDomainNameservers.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
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
	resp, r, err := apiClient.DomainAPI.UpdateGlue(context.Background(), domain).RequestModelCreateDomainGlue(requestModelCreateDomainGlue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.UpdateGlue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGlue`: ResultModelCreateDomainGlue
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.UpdateGlue`: %v\n", resp)
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

