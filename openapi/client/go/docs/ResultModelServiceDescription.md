# ResultModelServiceDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Tm** | **int32** | The unix timestamp of the request | 
**Msg** | **string** | Processing result message related to response | 
**Data** | [**ResultModelServiceDescriptionData**](ResultModelServiceDescriptionData.md) |  | 

## Methods

### NewResultModelServiceDescription

`func NewResultModelServiceDescription(status int32, tm int32, msg string, data ResultModelServiceDescriptionData, ) *ResultModelServiceDescription`

NewResultModelServiceDescription instantiates a new ResultModelServiceDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelServiceDescriptionWithDefaults

`func NewResultModelServiceDescriptionWithDefaults() *ResultModelServiceDescription`

NewResultModelServiceDescriptionWithDefaults instantiates a new ResultModelServiceDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResultModelServiceDescription) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultModelServiceDescription) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultModelServiceDescription) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTm

`func (o *ResultModelServiceDescription) GetTm() int32`

GetTm returns the Tm field if non-nil, zero value otherwise.

### GetTmOk

`func (o *ResultModelServiceDescription) GetTmOk() (*int32, bool)`

GetTmOk returns a tuple with the Tm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTm

`func (o *ResultModelServiceDescription) SetTm(v int32)`

SetTm sets Tm field to given value.


### GetMsg

`func (o *ResultModelServiceDescription) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ResultModelServiceDescription) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ResultModelServiceDescription) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *ResultModelServiceDescription) GetData() ResultModelServiceDescriptionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultModelServiceDescription) GetDataOk() (*ResultModelServiceDescriptionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultModelServiceDescription) SetData(v ResultModelServiceDescriptionData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


