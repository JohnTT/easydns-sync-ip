# RequestModelZoneBodyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain (FQDN) the zone record belongs to | 
**Host** | **string** | The hostname (minus the domain) the zone record is for | 
**Ttl** | Pointer to **int32** | The TTL (time-to-live) value for the zone record | [optional] [default to 600]
**Prio** | Pointer to **int32** | The priority of the zone record (Only applicable to record types that require it such as MX records) | [optional] [default to 0]
**Type** | **string** | The record type (A, AAAA, CNAME, etc) to use for this zone record | 
**Rdata** | **string** | The rdata value to use for the zone record | 
**GeozoneId** | Pointer to **int32** | The ID of the geozone based on georegion that this record should apply to | [optional] [default to 0]

## Methods

### NewRequestModelZoneBodyData

`func NewRequestModelZoneBodyData(domain string, host string, type_ string, rdata string, ) *RequestModelZoneBodyData`

NewRequestModelZoneBodyData instantiates a new RequestModelZoneBodyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestModelZoneBodyDataWithDefaults

`func NewRequestModelZoneBodyDataWithDefaults() *RequestModelZoneBodyData`

NewRequestModelZoneBodyDataWithDefaults instantiates a new RequestModelZoneBodyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RequestModelZoneBodyData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RequestModelZoneBodyData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RequestModelZoneBodyData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetHost

`func (o *RequestModelZoneBodyData) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RequestModelZoneBodyData) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RequestModelZoneBodyData) SetHost(v string)`

SetHost sets Host field to given value.


### GetTtl

`func (o *RequestModelZoneBodyData) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RequestModelZoneBodyData) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RequestModelZoneBodyData) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RequestModelZoneBodyData) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetPrio

`func (o *RequestModelZoneBodyData) GetPrio() int32`

GetPrio returns the Prio field if non-nil, zero value otherwise.

### GetPrioOk

`func (o *RequestModelZoneBodyData) GetPrioOk() (*int32, bool)`

GetPrioOk returns a tuple with the Prio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrio

`func (o *RequestModelZoneBodyData) SetPrio(v int32)`

SetPrio sets Prio field to given value.

### HasPrio

`func (o *RequestModelZoneBodyData) HasPrio() bool`

HasPrio returns a boolean if a field has been set.

### GetType

`func (o *RequestModelZoneBodyData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestModelZoneBodyData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestModelZoneBodyData) SetType(v string)`

SetType sets Type field to given value.


### GetRdata

`func (o *RequestModelZoneBodyData) GetRdata() string`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *RequestModelZoneBodyData) GetRdataOk() (*string, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *RequestModelZoneBodyData) SetRdata(v string)`

SetRdata sets Rdata field to given value.


### GetGeozoneId

`func (o *RequestModelZoneBodyData) GetGeozoneId() int32`

GetGeozoneId returns the GeozoneId field if non-nil, zero value otherwise.

### GetGeozoneIdOk

`func (o *RequestModelZoneBodyData) GetGeozoneIdOk() (*int32, bool)`

GetGeozoneIdOk returns a tuple with the GeozoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeozoneId

`func (o *RequestModelZoneBodyData) SetGeozoneId(v int32)`

SetGeozoneId sets GeozoneId field to given value.

### HasGeozoneId

`func (o *RequestModelZoneBodyData) HasGeozoneId() bool`

HasGeozoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


