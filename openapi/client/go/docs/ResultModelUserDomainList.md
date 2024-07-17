# ResultModelUserDomainList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Tm** | **int32** | The unix timestamp of the request | 
**Msg** | **string** | Processing result message related to response | 
**Data** | [**ResultModelUserDomainListData**](ResultModelUserDomainListData.md) |  | 

## Methods

### NewResultModelUserDomainList

`func NewResultModelUserDomainList(status int32, tm int32, msg string, data ResultModelUserDomainListData, ) *ResultModelUserDomainList`

NewResultModelUserDomainList instantiates a new ResultModelUserDomainList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelUserDomainListWithDefaults

`func NewResultModelUserDomainListWithDefaults() *ResultModelUserDomainList`

NewResultModelUserDomainListWithDefaults instantiates a new ResultModelUserDomainList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResultModelUserDomainList) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultModelUserDomainList) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultModelUserDomainList) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTm

`func (o *ResultModelUserDomainList) GetTm() int32`

GetTm returns the Tm field if non-nil, zero value otherwise.

### GetTmOk

`func (o *ResultModelUserDomainList) GetTmOk() (*int32, bool)`

GetTmOk returns a tuple with the Tm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTm

`func (o *ResultModelUserDomainList) SetTm(v int32)`

SetTm sets Tm field to given value.


### GetMsg

`func (o *ResultModelUserDomainList) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ResultModelUserDomainList) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ResultModelUserDomainList) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *ResultModelUserDomainList) GetData() ResultModelUserDomainListData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultModelUserDomainList) GetDataOk() (*ResultModelUserDomainListData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultModelUserDomainList) SetData(v ResultModelUserDomainListData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


