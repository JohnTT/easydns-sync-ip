# RequestModelSetRegStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain name to update | 
**Reglock** | Pointer to **bool** | Turn on/off reglock at the registry | [optional] 
**Renewal** | Pointer to **string** | Set renewal action of domain | [optional] 

## Methods

### NewRequestModelSetRegStatus

`func NewRequestModelSetRegStatus(domain string, ) *RequestModelSetRegStatus`

NewRequestModelSetRegStatus instantiates a new RequestModelSetRegStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestModelSetRegStatusWithDefaults

`func NewRequestModelSetRegStatusWithDefaults() *RequestModelSetRegStatus`

NewRequestModelSetRegStatusWithDefaults instantiates a new RequestModelSetRegStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RequestModelSetRegStatus) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RequestModelSetRegStatus) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RequestModelSetRegStatus) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetReglock

`func (o *RequestModelSetRegStatus) GetReglock() bool`

GetReglock returns the Reglock field if non-nil, zero value otherwise.

### GetReglockOk

`func (o *RequestModelSetRegStatus) GetReglockOk() (*bool, bool)`

GetReglockOk returns a tuple with the Reglock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReglock

`func (o *RequestModelSetRegStatus) SetReglock(v bool)`

SetReglock sets Reglock field to given value.

### HasReglock

`func (o *RequestModelSetRegStatus) HasReglock() bool`

HasReglock returns a boolean if a field has been set.

### GetRenewal

`func (o *RequestModelSetRegStatus) GetRenewal() string`

GetRenewal returns the Renewal field if non-nil, zero value otherwise.

### GetRenewalOk

`func (o *RequestModelSetRegStatus) GetRenewalOk() (*string, bool)`

GetRenewalOk returns a tuple with the Renewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewal

`func (o *RequestModelSetRegStatus) SetRenewal(v string)`

SetRenewal sets Renewal field to given value.

### HasRenewal

`func (o *RequestModelSetRegStatus) HasRenewal() bool`

HasRenewal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


