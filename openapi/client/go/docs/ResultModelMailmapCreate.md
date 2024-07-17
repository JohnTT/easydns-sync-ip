# ResultModelMailmapCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Tm** | **int32** | The unix timestamp of the request | 
**Msg** | **string** | Processing result message related to response | 
**Data** | [**ResultModelMailmapCreateData**](ResultModelMailmapCreateData.md) |  | 

## Methods

### NewResultModelMailmapCreate

`func NewResultModelMailmapCreate(status int32, tm int32, msg string, data ResultModelMailmapCreateData, ) *ResultModelMailmapCreate`

NewResultModelMailmapCreate instantiates a new ResultModelMailmapCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelMailmapCreateWithDefaults

`func NewResultModelMailmapCreateWithDefaults() *ResultModelMailmapCreate`

NewResultModelMailmapCreateWithDefaults instantiates a new ResultModelMailmapCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResultModelMailmapCreate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultModelMailmapCreate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultModelMailmapCreate) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTm

`func (o *ResultModelMailmapCreate) GetTm() int32`

GetTm returns the Tm field if non-nil, zero value otherwise.

### GetTmOk

`func (o *ResultModelMailmapCreate) GetTmOk() (*int32, bool)`

GetTmOk returns a tuple with the Tm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTm

`func (o *ResultModelMailmapCreate) SetTm(v int32)`

SetTm sets Tm field to given value.


### GetMsg

`func (o *ResultModelMailmapCreate) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ResultModelMailmapCreate) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ResultModelMailmapCreate) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *ResultModelMailmapCreate) GetData() ResultModelMailmapCreateData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultModelMailmapCreate) GetDataOk() (*ResultModelMailmapCreateData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultModelMailmapCreate) SetData(v ResultModelMailmapCreateData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


