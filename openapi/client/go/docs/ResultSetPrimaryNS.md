# ResultSetPrimaryNS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** | Result message of request | 
**Data** | [**ResultSetPrimaryNSData**](ResultSetPrimaryNSData.md) |  | 
**Status** | **int32** | HTTP status code of result | [default to 200]

## Methods

### NewResultSetPrimaryNS

`func NewResultSetPrimaryNS(msg string, data ResultSetPrimaryNSData, status int32, ) *ResultSetPrimaryNS`

NewResultSetPrimaryNS instantiates a new ResultSetPrimaryNS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultSetPrimaryNSWithDefaults

`func NewResultSetPrimaryNSWithDefaults() *ResultSetPrimaryNS`

NewResultSetPrimaryNSWithDefaults instantiates a new ResultSetPrimaryNS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ResultSetPrimaryNS) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ResultSetPrimaryNS) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ResultSetPrimaryNS) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetData

`func (o *ResultSetPrimaryNS) GetData() ResultSetPrimaryNSData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultSetPrimaryNS) GetDataOk() (*ResultSetPrimaryNSData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultSetPrimaryNS) SetData(v ResultSetPrimaryNSData)`

SetData sets Data field to given value.


### GetStatus

`func (o *ResultSetPrimaryNS) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultSetPrimaryNS) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultSetPrimaryNS) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


