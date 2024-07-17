# ResultModelDomainDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** |  | 
**Data** | Pointer to [**[]ResultModelDomainDeleteData**](ResultModelDomainDeleteData.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewResultModelDomainDelete

`func NewResultModelDomainDelete(msg string, status string, ) *ResultModelDomainDelete`

NewResultModelDomainDelete instantiates a new ResultModelDomainDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelDomainDeleteWithDefaults

`func NewResultModelDomainDeleteWithDefaults() *ResultModelDomainDelete`

NewResultModelDomainDeleteWithDefaults instantiates a new ResultModelDomainDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ResultModelDomainDelete) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ResultModelDomainDelete) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ResultModelDomainDelete) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *ResultModelDomainDelete) GetData() []ResultModelDomainDeleteData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultModelDomainDelete) GetDataOk() (*[]ResultModelDomainDeleteData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultModelDomainDelete) SetData(v []ResultModelDomainDeleteData)`

SetData sets Data field to given value.

### HasData

`func (o *ResultModelDomainDelete) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *ResultModelDomainDelete) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultModelDomainDelete) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultModelDomainDelete) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


