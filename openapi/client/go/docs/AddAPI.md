# \AddAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUXZoneRec**](AddAPI.md#AddUXZoneRec) | **Put** /zones/async/ux/records/add/{domain}/{type} | Create a new record in a zone using async zone reloading
[**CreateDomain**](AddAPI.md#CreateDomain) | **Put** /domains/add/{domain} | Add a new domain to the system including registering the domain at the registry if applicable. When registering domains additional fields may be required depending on the TLD
[**CreateMailmap**](AddAPI.md#CreateMailmap) | **Put** /mail/maps/{domain} | This API call allows an API user to create a new mailmap for a domain on the easyDNS system
[**CreateUser**](AddAPI.md#CreateUser) | **Put** /users/{user} | Create a new user on the system using the provided user information and return the API credentials for the new user if successful



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
	r, err := apiClient.AddAPI.AddUXZoneRec(context.Background(), domain, type_).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddAPI.AddUXZoneRec``: %v\n", err)
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
	resp, r, err := apiClient.AddAPI.CreateDomain(context.Background(), domain).CreateDomainRequest(createDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddAPI.CreateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDomain`: ResultModelDomainCreate
	fmt.Fprintf(os.Stdout, "Response from `AddAPI.CreateDomain`: %v\n", resp)
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


## CreateMailmap

> ResultModelMailmapCreate CreateMailmap(ctx, domain).RequestModelMailmapCreate(requestModelMailmapCreate).Execute()

This API call allows an API user to create a new mailmap for a domain on the easyDNS system

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
	requestModelMailmapCreate := *openapiclient.NewRequestModelMailmapCreate("test", "@", "test@example.com", int32(1)) // RequestModelMailmapCreate | The contents of the mailmap to create in JSON encoded format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddAPI.CreateMailmap(context.Background(), domain).RequestModelMailmapCreate(requestModelMailmapCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddAPI.CreateMailmap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMailmap`: ResultModelMailmapCreate
	fmt.Fprintf(os.Stdout, "Response from `AddAPI.CreateMailmap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name the mailmap belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMailmapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestModelMailmapCreate** | [**RequestModelMailmapCreate**](RequestModelMailmapCreate.md) | The contents of the mailmap to create in JSON encoded format | 

### Return type

[**ResultModelMailmapCreate**](ResultModelMailmapCreate.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> ResultModelCreateUser CreateUser(ctx, user).RequestModelCreateUserData(requestModelCreateUserData).Execute()

Create a new user on the system using the provided user information and return the API credentials for the new user if successful

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
	user := "mynewuser123" // string | The username to create
	requestModelCreateUserData := *openapiclient.NewRequestModelCreateUserData("ON", "CA", "support@easydns.com", "CAD") // RequestModelCreateUserData | New user account information in JSON encoded format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddAPI.CreateUser(context.Background(), user).RequestModelCreateUserData(requestModelCreateUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: ResultModelCreateUser
	fmt.Fprintf(os.Stdout, "Response from `AddAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | The username to create | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestModelCreateUserData** | [**RequestModelCreateUserData**](RequestModelCreateUserData.md) | New user account information in JSON encoded format | 

### Return type

[**ResultModelCreateUser**](ResultModelCreateUser.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

