# ResultModelDomainCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** |  | 
**Data** | Pointer to [**[]ResultModelDomainCreateData**](ResultModelDomainCreateData.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewResultModelDomainCreate

`func NewResultModelDomainCreate(msg string, status string, ) *ResultModelDomainCreate`

NewResultModelDomainCreate instantiates a new ResultModelDomainCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelDomainCreateWithDefaults

`func NewResultModelDomainCreateWithDefaults() *ResultModelDomainCreate`

NewResultModelDomainCreateWithDefaults instantiates a new ResultModelDomainCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ResultModelDomainCreate) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ResultModelDomainCreate) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ResultModelDomainCreate) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *ResultModelDomainCreate) GetData() []ResultModelDomainCreateData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultModelDomainCreate) GetDataOk() (*[]ResultModelDomainCreateData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultModelDomainCreate) SetData(v []ResultModelDomainCreateData)`

SetData sets Data field to given value.

### HasData

`func (o *ResultModelDomainCreate) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *ResultModelDomainCreate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultModelDomainCreate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultModelDomainCreate) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


