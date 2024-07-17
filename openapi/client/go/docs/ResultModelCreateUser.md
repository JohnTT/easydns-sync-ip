# ResultModelCreateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** | Result message of request | 
**Data** | Pointer to [**ResultModelCreateUserData**](ResultModelCreateUserData.md) |  | [optional] 
**Status** | **int32** | HTTP status code of result | [default to 201]

## Methods

### NewResultModelCreateUser

`func NewResultModelCreateUser(msg string, status int32, ) *ResultModelCreateUser`

NewResultModelCreateUser instantiates a new ResultModelCreateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelCreateUserWithDefaults

`func NewResultModelCreateUserWithDefaults() *ResultModelCreateUser`

NewResultModelCreateUserWithDefaults instantiates a new ResultModelCreateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ResultModelCreateUser) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ResultModelCreateUser) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ResultModelCreateUser) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *ResultModelCreateUser) GetData() ResultModelCreateUserData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultModelCreateUser) GetDataOk() (*ResultModelCreateUserData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultModelCreateUser) SetData(v ResultModelCreateUserData)`

SetData sets Data field to given value.

### HasData

`func (o *ResultModelCreateUser) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *ResultModelCreateUser) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultModelCreateUser) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultModelCreateUser) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


