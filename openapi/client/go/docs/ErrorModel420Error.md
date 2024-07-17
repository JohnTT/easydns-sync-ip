# ErrorModel420Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | HTTP status code indicating error type returned | 
**Message** | **string** | Resulting error message for a failed request indicating invalid/missing credentials sent with request or too many requests sent in too short a period of time | 

## Methods

### NewErrorModel420Error

`func NewErrorModel420Error(code int32, message string, ) *ErrorModel420Error`

NewErrorModel420Error instantiates a new ErrorModel420Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorModel420ErrorWithDefaults

`func NewErrorModel420ErrorWithDefaults() *ErrorModel420Error`

NewErrorModel420ErrorWithDefaults instantiates a new ErrorModel420Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorModel420Error) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorModel420Error) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorModel420Error) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ErrorModel420Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorModel420Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorModel420Error) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


