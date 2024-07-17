# ResultModelZoneRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Domain** | **string** |  | 
**Host** | **string** |  | 
**Ttl** | Pointer to **int32** |  | [optional] 
**Prio** | Pointer to **int32** |  | [optional] 
**GeozoneId** | Pointer to **int32** |  | [optional] 
**Type** | **string** |  | 
**Rdata** | **string** |  | 
**LastMod** | Pointer to **string** |  | [optional] 

## Methods

### NewResultModelZoneRecord

`func NewResultModelZoneRecord(id int32, domain string, host string, type_ string, rdata string, ) *ResultModelZoneRecord`

NewResultModelZoneRecord instantiates a new ResultModelZoneRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelZoneRecordWithDefaults

`func NewResultModelZoneRecordWithDefaults() *ResultModelZoneRecord`

NewResultModelZoneRecordWithDefaults instantiates a new ResultModelZoneRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResultModelZoneRecord) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultModelZoneRecord) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultModelZoneRecord) SetId(v int32)`

SetId sets Id field to given value.


### GetDomain

`func (o *ResultModelZoneRecord) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelZoneRecord) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelZoneRecord) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetHost

`func (o *ResultModelZoneRecord) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ResultModelZoneRecord) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ResultModelZoneRecord) SetHost(v string)`

SetHost sets Host field to given value.


### GetTtl

`func (o *ResultModelZoneRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ResultModelZoneRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ResultModelZoneRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ResultModelZoneRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetPrio

`func (o *ResultModelZoneRecord) GetPrio() int32`

GetPrio returns the Prio field if non-nil, zero value otherwise.

### GetPrioOk

`func (o *ResultModelZoneRecord) GetPrioOk() (*int32, bool)`

GetPrioOk returns a tuple with the Prio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrio

`func (o *ResultModelZoneRecord) SetPrio(v int32)`

SetPrio sets Prio field to given value.

### HasPrio

`func (o *ResultModelZoneRecord) HasPrio() bool`

HasPrio returns a boolean if a field has been set.

### GetGeozoneId

`func (o *ResultModelZoneRecord) GetGeozoneId() int32`

GetGeozoneId returns the GeozoneId field if non-nil, zero value otherwise.

### GetGeozoneIdOk

`func (o *ResultModelZoneRecord) GetGeozoneIdOk() (*int32, bool)`

GetGeozoneIdOk returns a tuple with the GeozoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeozoneId

`func (o *ResultModelZoneRecord) SetGeozoneId(v int32)`

SetGeozoneId sets GeozoneId field to given value.

### HasGeozoneId

`func (o *ResultModelZoneRecord) HasGeozoneId() bool`

HasGeozoneId returns a boolean if a field has been set.

### GetType

`func (o *ResultModelZoneRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResultModelZoneRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResultModelZoneRecord) SetType(v string)`

SetType sets Type field to given value.


### GetRdata

`func (o *ResultModelZoneRecord) GetRdata() string`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *ResultModelZoneRecord) GetRdataOk() (*string, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *ResultModelZoneRecord) SetRdata(v string)`

SetRdata sets Rdata field to given value.


### GetLastMod

`func (o *ResultModelZoneRecord) GetLastMod() string`

GetLastMod returns the LastMod field if non-nil, zero value otherwise.

### GetLastModOk

`func (o *ResultModelZoneRecord) GetLastModOk() (*string, bool)`

GetLastModOk returns a tuple with the LastMod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMod

`func (o *ResultModelZoneRecord) SetLastMod(v string)`

SetLastMod sets LastMod field to given value.

### HasLastMod

`func (o *ResultModelZoneRecord) HasLastMod() bool`

HasLastMod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


