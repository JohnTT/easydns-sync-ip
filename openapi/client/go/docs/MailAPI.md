# \MailAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMailmap**](MailAPI.md#CreateMailmap) | **Put** /mail/maps/{domain} | This API call allows an API user to create a new mailmap for a domain on the easyDNS system
[**DeleteMailmap**](MailAPI.md#DeleteMailmap) | **Delete** /mail/maps/{domain}/{mailmap_id} | This API call allows you to delete an existing mailmap for a domain. To obtain the mailmap_id for a mailmap you should perform a request to list your mailmaps
[**ListMailmaps**](MailAPI.md#ListMailmaps) | **Get** /mail/maps/{domain} | To get a listing of the existing mailmaps for a domain this API call can be used. This call will return all pertinent details about each mailmap created for the domain.
[**UpdateMailmap**](MailAPI.md#UpdateMailmap) | **Post** /mail/maps/{domain}/{email} | This API call can be used to update the configuration of an existing mailmap for a domain. To update a mailmap you must provide all required parameters in the request. On a successful update the existing mailmap configuration will be replaced with the provided mailmap configuration in the request



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
	resp, r, err := apiClient.MailAPI.CreateMailmap(context.Background(), domain).RequestModelMailmapCreate(requestModelMailmapCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.CreateMailmap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMailmap`: ResultModelMailmapCreate
	fmt.Fprintf(os.Stdout, "Response from `MailAPI.CreateMailmap`: %v\n", resp)
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
	resp, r, err := apiClient.MailAPI.DeleteMailmap(context.Background(), domain, mailmapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.DeleteMailmap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMailmap`: ResultModelMailmapDelete
	fmt.Fprintf(os.Stdout, "Response from `MailAPI.DeleteMailmap`: %v\n", resp)
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
	resp, r, err := apiClient.MailAPI.ListMailmaps(context.Background(), domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.ListMailmaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMailmaps`: ResultModelMailmapList
	fmt.Fprintf(os.Stdout, "Response from `MailAPI.ListMailmaps`: %v\n", resp)
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
	resp, r, err := apiClient.MailAPI.UpdateMailmap(context.Background(), domain, email).RequestModelMailmapCreate(requestModelMailmapCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MailAPI.UpdateMailmap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMailmap`: ResultModelMailmapCreate
	fmt.Fprintf(os.Stdout, "Response from `MailAPI.UpdateMailmap`: %v\n", resp)
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

