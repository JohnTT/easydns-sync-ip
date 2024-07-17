# ResultModelDeleteGlue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Tm** | **int32** | The unix timestamp of the request | 
**Msg** | **string** | Processing result message related to response | 
**Data** | [**ResultModelDeleteGlueData**](ResultModelDeleteGlueData.md) |  | 

## Methods

### NewResultModelDeleteGlue

`func NewResultModelDeleteGlue(status int32, tm int32, msg string, data ResultModelDeleteGlueData, ) *ResultModelDeleteGlue`

NewResultModelDeleteGlue instantiates a new ResultModelDeleteGlue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelDeleteGlueWithDefaults

`func NewResultModelDeleteGlueWithDefaults() *ResultModelDeleteGlue`

NewResultModelDeleteGlueWithDefaults instantiates a new ResultModelDeleteGlue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResultModelDeleteGlue) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultModelDeleteGlue) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultModelDeleteGlue) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTm

`func (o *ResultModelDeleteGlue) GetTm() int32`

GetTm returns the Tm field if non-nil, zero value otherwise.

### GetTmOk

`func (o *ResultModelDeleteGlue) GetTmOk() (*int32, bool)`

GetTmOk returns a tuple with the Tm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTm

`func (o *ResultModelDeleteGlue) SetTm(v int32)`

SetTm sets Tm field to given value.


### GetMsg

`func (o *ResultModelDeleteGlue) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ResultModelDeleteGlue) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ResultModelDeleteGlue) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *ResultModelDeleteGlue) GetData() ResultModelDeleteGlueData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultModelDeleteGlue) GetDataOk() (*ResultModelDeleteGlueData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultModelDeleteGlue) SetData(v ResultModelDeleteGlueData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


