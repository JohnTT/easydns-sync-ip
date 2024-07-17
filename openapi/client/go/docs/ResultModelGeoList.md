# ResultModelGeoList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** | Successful result message | 
**Tm** | **int64** | The unix epoch timestamp of the request | 
**Data** | [**[]ResultModelGeoListData**](ResultModelGeoListData.md) |  | 
**Count** | Pointer to **int32** | Number of records returned on page | [optional] 
**Total** | Pointer to **int32** | Total number of records on all pages | [optional] 
**Start** | Pointer to **int32** | Index of record to start paginated results | [optional] 
**Max** | Pointer to **int32** | Maximum number of records to return per page | [optional] 
**Status** | **int64** | The HTTP status code for the result | [default to 201]

## Methods

### NewResultModelGeoList

`func NewResultModelGeoList(msg string, tm int64, data []ResultModelGeoListData, status int64, ) *ResultModelGeoList`

NewResultModelGeoList instantiates a new ResultModelGeoList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelGeoListWithDefaults

`func NewResultModelGeoListWithDefaults() *ResultModelGeoList`

NewResultModelGeoListWithDefaults instantiates a new ResultModelGeoList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ResultModelGeoList) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ResultModelGeoList) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ResultModelGeoList) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetTm

`func (o *ResultModelGeoList) GetTm() int64`

GetTm returns the Tm field if non-nil, zero value otherwise.

### GetTmOk

`func (o *ResultModelGeoList) GetTmOk() (*int64, bool)`

GetTmOk returns a tuple with the Tm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTm

`func (o *ResultModelGeoList) SetTm(v int64)`

SetTm sets Tm field to given value.


### GetData

`func (o *ResultModelGeoList) GetData() []ResultModelGeoListData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultModelGeoList) GetDataOk() (*[]ResultModelGeoListData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultModelGeoList) SetData(v []ResultModelGeoListData)`

SetData sets Data field to given value.


### GetCount

`func (o *ResultModelGeoList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResultModelGeoList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResultModelGeoList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ResultModelGeoList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *ResultModelGeoList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResultModelGeoList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResultModelGeoList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ResultModelGeoList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetStart

`func (o *ResultModelGeoList) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResultModelGeoList) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResultModelGeoList) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ResultModelGeoList) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetMax

`func (o *ResultModelGeoList) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ResultModelGeoList) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ResultModelGeoList) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *ResultModelGeoList) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetStatus

`func (o *ResultModelGeoList) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultModelGeoList) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultModelGeoList) SetStatus(v int64)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


