# ResultModelGetUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** |  | 
**Data** | Pointer to [**[]ResultModelGetUserInfoData**](ResultModelGetUserInfoData.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewResultModelGetUserInfo

`func NewResultModelGetUserInfo(msg string, status string, ) *ResultModelGetUserInfo`

NewResultModelGetUserInfo instantiates a new ResultModelGetUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelGetUserInfoWithDefaults

`func NewResultModelGetUserInfoWithDefaults() *ResultModelGetUserInfo`

NewResultModelGetUserInfoWithDefaults instantiates a new ResultModelGetUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ResultModelGetUserInfo) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ResultModelGetUserInfo) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ResultModelGetUserInfo) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *ResultModelGetUserInfo) GetData() []ResultModelGetUserInfoData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultModelGetUserInfo) GetDataOk() (*[]ResultModelGetUserInfoData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultModelGetUserInfo) SetData(v []ResultModelGetUserInfoData)`

SetData sets Data field to given value.

### HasData

`func (o *ResultModelGetUserInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *ResultModelGetUserInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultModelGetUserInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultModelGetUserInfo) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


