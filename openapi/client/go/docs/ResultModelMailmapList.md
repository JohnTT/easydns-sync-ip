# ResultModelMailmapList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** |  | 
**Tm** | **int32** | The unix timestamp the request was made | 
**Data** | Pointer to [**ResultModelMailmapListData**](ResultModelMailmapListData.md) |  | [optional] 

## Methods

### NewResultModelMailmapList

`func NewResultModelMailmapList(status int32, tm int32, ) *ResultModelMailmapList`

NewResultModelMailmapList instantiates a new ResultModelMailmapList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelMailmapListWithDefaults

`func NewResultModelMailmapListWithDefaults() *ResultModelMailmapList`

NewResultModelMailmapListWithDefaults instantiates a new ResultModelMailmapList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResultModelMailmapList) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultModelMailmapList) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultModelMailmapList) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTm

`func (o *ResultModelMailmapList) GetTm() int32`

GetTm returns the Tm field if non-nil, zero value otherwise.

### GetTmOk

`func (o *ResultModelMailmapList) GetTmOk() (*int32, bool)`

GetTmOk returns a tuple with the Tm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTm

`func (o *ResultModelMailmapList) SetTm(v int32)`

SetTm sets Tm field to given value.


### GetData

`func (o *ResultModelMailmapList) GetData() ResultModelMailmapListData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultModelMailmapList) GetDataOk() (*ResultModelMailmapListData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultModelMailmapList) SetData(v ResultModelMailmapListData)`

SetData sets Data field to given value.

### HasData

`func (o *ResultModelMailmapList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


