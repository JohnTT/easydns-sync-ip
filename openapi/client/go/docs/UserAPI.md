# \UserAPI

All URIs are relative to *https://sandbox.rest.easydns.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UserAPI.md#CreateUser) | **Put** /users/{user} | Create a new user on the system using the provided user information and return the API credentials for the new user if successful
[**GetRegStatus**](UserAPI.md#GetRegStatus) | **Get** /domains/regstatus | Get the reglock and renewal status for a users list of domains
[**GetUserInfo**](UserAPI.md#GetUserInfo) | **Get** /user | Get details about a user
[**ListUserDomains**](UserAPI.md#ListUserDomains) | **Get** /domains/list/{user} | Get a listing of all domains associated with a user
[**SetRegStatus**](UserAPI.md#SetRegStatus) | **Post** /domains/regstatus | Set the reglock status of a set of domains owned by a user
[**UpdateUserDS**](UserAPI.md#UpdateUserDS) | **Post** /users/{user}/info | Update information about a user on the system



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
	resp, r, err := apiClient.UserAPI.CreateUser(context.Background(), user).RequestModelCreateUserData(requestModelCreateUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: ResultModelCreateUser
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateUser`: %v\n", resp)
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
	resp, r, err := apiClient.UserAPI.GetRegStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetRegStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegStatus`: ResultModelRegStatus
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetRegStatus`: %v\n", resp)
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


## GetUserInfo

> ResultModelGetUserInfo GetUserInfo(ctx).Execute()

Get details about a user

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
	resp, r, err := apiClient.UserAPI.GetUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInfo`: ResultModelGetUserInfo
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoRequest struct via the builder pattern


### Return type

[**ResultModelGetUserInfo**](ResultModelGetUserInfo.md)

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
	resp, r, err := apiClient.UserAPI.ListUserDomains(context.Background(), user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ListUserDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserDomains`: ResultModelUserDomainList
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ListUserDomains`: %v\n", resp)
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
	resp, r, err := apiClient.UserAPI.SetRegStatus(context.Background()).RequestModelSetRegStatus(requestModelSetRegStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.SetRegStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRegStatus`: ResultModelSetRegStatus
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.SetRegStatus`: %v\n", resp)
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
	resp, r, err := apiClient.UserAPI.UpdateUserDS(context.Background(), user).RequestModelUpdateUserData(requestModelUpdateUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserDS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserDS`: ResultModelGetUserInfo
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUserDS`: %v\n", resp)
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

