# \WriteAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUXZoneRec**](WriteAPI.md#AddUXZoneRec) | **Put** /zones/async/ux/records/add/{domain}/{type} | Create a new record in a zone using async zone reloading
[**AddZoneRec**](WriteAPI.md#AddZoneRec) | **Put** /zones/records/add/{domain}/{type} | Create a new record in a zone
[**CreateDomain**](WriteAPI.md#CreateDomain) | **Put** /domains/add/{domain} | Add a new domain to the system including registering the domain at the registry if applicable. When registering domains additional fields may be required depending on the TLD
[**CreateMailmap**](WriteAPI.md#CreateMailmap) | **Put** /mail/maps/{domain} | This API call allows an API user to create a new mailmap for a domain on the easyDNS system
[**CreateUser**](WriteAPI.md#CreateUser) | **Put** /users/{user} | Create a new user on the system using the provided user information and return the API credentials for the new user if successful
[**DelUXZoneRec**](WriteAPI.md#DelUXZoneRec) | **Delete** /zones/async/ux/records/{domain}/{id} | Delete an existing zone record using async zone reloading
[**DelZoneRec**](WriteAPI.md#DelZoneRec) | **Delete** /zones/records/{domain}/{id} | Delete an existing zone record
[**DeleteDomain**](WriteAPI.md#DeleteDomain) | **Delete** /domain/{domain} | Delete an existing domain from the system
[**DeleteMailmap**](WriteAPI.md#DeleteMailmap) | **Delete** /mail/maps/{domain}/{mailmap_id} | This API call allows you to delete an existing mailmap for a domain. To obtain the mailmap_id for a mailmap you should perform a request to list your mailmaps
[**ForceZoneReload**](WriteAPI.md#ForceZoneReload) | **Get** /zones/reload/{domain}/force | Force regeneration of a zone immediately instead of allowing zone reloader to handle it
[**ModUXZoneRec**](WriteAPI.md#ModUXZoneRec) | **Post** /zones/async/ux/records/{id} | Modify an existing zone record using async zone reloading
[**ModZoneRec**](WriteAPI.md#ModZoneRec) | **Post** /zones/records/{id} | Modify an existing zone record
[**SetPrimaryNS**](WriteAPI.md#SetPrimaryNS) | **Post** /domains/primary_ns/{domain} | Set the primary NS value for a domain. Will also update domain to make it secondary
[**SetRegStatus**](WriteAPI.md#SetRegStatus) | **Post** /domains/regstatus | Set the reglock status of a set of domains owned by a user
[**UpdateMailmap**](WriteAPI.md#UpdateMailmap) | **Post** /mail/maps/{domain}/{email} | This API call can be used to update the configuration of an existing mailmap for a domain. To update a mailmap you must provide all required parameters in the request. On a successful update the existing mailmap configuration will be replaced with the provided mailmap configuration in the request
[**UpdateUserDS**](WriteAPI.md#UpdateUserDS) | **Post** /users/{user}/info | Update information about a user on the system



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
	r, err := apiClient.WriteAPI.AddUXZoneRec(context.Background(), domain, type_).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.AddUXZoneRec``: %v\n", err)
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
	r, err := apiClient.WriteAPI.AddZoneRec(context.Background(), domain, type_).RequestModelZoneBodyData(requestModelZoneBodyData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.AddZoneRec``: %v\n", err)
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
	resp, r, err := apiClient.WriteAPI.CreateDomain(context.Background(), domain).CreateDomainRequest(createDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.CreateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDomain`: ResultModelDomainCreate
	fmt.Fprintf(os.Stdout, "Response from `WriteAPI.CreateDomain`: %v\n", resp)
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
	resp, r, err := apiClient.WriteAPI.CreateMailmap(context.Background(), domain).RequestModelMailmapCreate(requestModelMailmapCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.CreateMailmap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMailmap`: ResultModelMailmapCreate
	fmt.Fprintf(os.Stdout, "Response from `WriteAPI.CreateMailmap`: %v\n", resp)
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
	resp, r, err := apiClient.WriteAPI.CreateUser(context.Background(), user).RequestModelCreateUserData(requestModelCreateUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: ResultModelCreateUser
	fmt.Fprintf(os.Stdout, "Response from `WriteAPI.CreateUser`: %v\n", resp)
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
	r, err := apiClient.WriteAPI.DelUXZoneRec(context.Background(), domain, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.DelUXZoneRec``: %v\n", err)
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
	r, err := apiClient.WriteAPI.DelZoneRec(context.Background(), domain, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.DelZoneRec``: %v\n", err)
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
	resp, r, err := apiClient.WriteAPI.DeleteDomain(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.DeleteDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDomain`: ResultModelDomainDelete
	fmt.Fprintf(os.Stdout, "Response from `WriteAPI.DeleteDomain`: %v\n", resp)
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
	resp, r, err := apiClient.WriteAPI.DeleteMailmap(context.Background(), domain, mailmapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.DeleteMailmap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMailmap`: ResultModelMailmapDelete
	fmt.Fprintf(os.Stdout, "Response from `WriteAPI.DeleteMailmap`: %v\n", resp)
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
	r, err := apiClient.WriteAPI.ForceZoneReload(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.ForceZoneReload``: %v\n", err)
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
	resp, r, err := apiClient.WriteAPI.ModUXZoneRec(context.Background(), id).RequestModelZoneBodyData(requestModelZoneBodyData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.ModUXZoneRec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModUXZoneRec`: ResultModelZoneRecordBody
	fmt.Fprintf(os.Stdout, "Response from `WriteAPI.ModUXZoneRec`: %v\n", resp)
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
	resp, r, err := apiClient.WriteAPI.ModZoneRec(context.Background(), id).RequestModelZoneBodyData(requestModelZoneBodyData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.ModZoneRec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModZoneRec`: ResultModelZoneRecordBody
	fmt.Fprintf(os.Stdout, "Response from `WriteAPI.ModZoneRec`: %v\n", resp)
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
	resp, r, err := apiClient.WriteAPI.SetPrimaryNS(context.Background(), domain).RequestPrimaryNSBody(requestPrimaryNSBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.SetPrimaryNS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPrimaryNS`: ResultSetPrimaryNS
	fmt.Fprintf(os.Stdout, "Response from `WriteAPI.SetPrimaryNS`: %v\n", resp)
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
	resp, r, err := apiClient.WriteAPI.SetRegStatus(context.Background()).RequestModelSetRegStatus(requestModelSetRegStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.SetRegStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRegStatus`: ResultModelSetRegStatus
	fmt.Fprintf(os.Stdout, "Response from `WriteAPI.SetRegStatus`: %v\n", resp)
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


## UpdateMailmap

> ResultModelMailmapCreate UpdateMailmap(ctx, domain, email).RequestModelMailmapCreate(requestModelMailmapCreate).Execute()

This API call can be used to update the configuration of an existing mailmap for a domain. To update a mailmap you must provide all required parameters in the request. On a successful update the existing mailmap configuration will be replaced with the provided mailmap configuration in the request

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
	email := "email_example" // string | The email address of the mailmap to update
	requestModelMailmapCreate := *openapiclient.NewRequestModelMailmapCreate("test", "@", "test@example.com", int32(1)) // RequestModelMailmapCreate | The contents that the mailmap should be updated to in JSON encoded format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WriteAPI.UpdateMailmap(context.Background(), domain, email).RequestModelMailmapCreate(requestModelMailmapCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.UpdateMailmap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMailmap`: ResultModelMailmapCreate
	fmt.Fprintf(os.Stdout, "Response from `WriteAPI.UpdateMailmap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | The domain name the mailmap belongs to | 
**email** | **string** | The email address of the mailmap to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMailmapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestModelMailmapCreate** | [**RequestModelMailmapCreate**](RequestModelMailmapCreate.md) | The contents that the mailmap should be updated to in JSON encoded format | 

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


## UpdateUserDS

> ResultModelGetUserInfo UpdateUserDS(ctx, user).RequestModelUpdateUserData(requestModelUpdateUserData).Execute()

Update information about a user on the system

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
	user := "user_example" // string | 
	requestModelUpdateUserData := *openapiclient.NewRequestModelUpdateUserData() // RequestModelUpdateUserData | Updated user account information in JSON encoded format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WriteAPI.UpdateUserDS(context.Background(), user).RequestModelUpdateUserData(requestModelUpdateUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WriteAPI.UpdateUserDS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserDS`: ResultModelGetUserInfo
	fmt.Fprintf(os.Stdout, "Response from `WriteAPI.UpdateUserDS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserDSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestModelUpdateUserData** | [**RequestModelUpdateUserData**](RequestModelUpdateUserData.md) | Updated user account information in JSON encoded format | 

### Return type

[**ResultModelGetUserInfo**](ResultModelGetUserInfo.md)

### Authorization

[easyapi_basic](../README.md#easyapi_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

