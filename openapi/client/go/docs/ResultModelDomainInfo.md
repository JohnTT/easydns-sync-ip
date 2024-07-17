# ResultModelDomainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** | Successful result message | 
**Tm** | **int64** | The unix epoch timestamp of the request | 
**Data** | [**ResultModelDomainInfoData**](ResultModelDomainInfoData.md) |  | 
**Status** | **int64** | The HTTP status code for the result | [default to 200]

## Methods

### NewResultModelDomainInfo

`func NewResultModelDomainInfo(msg string, tm int64, data ResultModelDomainInfoData, status int64, ) *ResultModelDomainInfo`

NewResultModelDomainInfo instantiates a new ResultModelDomainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelDomainInfoWithDefaults

`func NewResultModelDomainInfoWithDefaults() *ResultModelDomainInfo`

NewResultModelDomainInfoWithDefaults instantiates a new ResultModelDomainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ResultModelDomainInfo) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ResultModelDomainInfo) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ResultModelDomainInfo) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetTm

`func (o *ResultModelDomainInfo) GetTm() int64`

GetTm returns the Tm field if non-nil, zero value otherwise.

### GetTmOk

`func (o *ResultModelDomainInfo) GetTmOk() (*int64, bool)`

GetTmOk returns a tuple with the Tm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTm

`func (o *ResultModelDomainInfo) SetTm(v int64)`

SetTm sets Tm field to given value.


### GetData

`func (o *ResultModelDomainInfo) GetData() ResultModelDomainInfoData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultModelDomainInfo) GetDataOk() (*ResultModelDomainInfoData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultModelDomainInfo) SetData(v ResultModelDomainInfoData)`

SetData sets Data field to given value.


### GetStatus

`func (o *ResultModelDomainInfo) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultModelDomainInfo) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultModelDomainInfo) SetStatus(v int64)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


