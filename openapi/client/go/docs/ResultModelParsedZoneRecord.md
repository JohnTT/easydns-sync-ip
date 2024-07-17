# ResultModelParsedZoneRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Zone record unique ID. Records that parse to multiple zone records will share the same ID! | 
**Domain** | **string** | The domain name that contains the zone records | 
**Host** | **string** | The hostname of the record (without the domain name). &#39;@&#39; indicates the record is for the root of the zone. | 
**Ttl** | Pointer to **int32** | (OPTIONAL) The Time-to-live (TTL) of the zone record. If not returned, assume the default. | [optional] [default to 600]
**Prio** | Pointer to **int32** | (OPTIONAL) The priority of the record. This is only required for records that include the prio field (such as MX). | [optional] [default to 0]
**GeozoneId** | Pointer to **int32** | (OPTIONAL) The ID of the geolocation this record applies to. | [optional] [default to 0]
**Type** | **string** | The zone record type of the record. PLEASE NOTE: Some record types are specific to easyDNS and may require translation on other providers. | 
**Rdata** | **string** | The contents of the rdata for a zone record. | 
**LastMod** | Pointer to **string** | The time this zone record was last updated. | [optional] 
**Url** | Pointer to **string** | The original URL that a record uses as a target. This applies to custom records like URL/STEALTH. | [optional] 
**OrigRdata** | Pointer to **string** | The data originally stored in the rdata field for custom types and records containing keywords | [optional] 

## Methods

### NewResultModelParsedZoneRecord

`func NewResultModelParsedZoneRecord(id int32, domain string, host string, type_ string, rdata string, ) *ResultModelParsedZoneRecord`

NewResultModelParsedZoneRecord instantiates a new ResultModelParsedZoneRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelParsedZoneRecordWithDefaults

`func NewResultModelParsedZoneRecordWithDefaults() *ResultModelParsedZoneRecord`

NewResultModelParsedZoneRecordWithDefaults instantiates a new ResultModelParsedZoneRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResultModelParsedZoneRecord) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultModelParsedZoneRecord) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultModelParsedZoneRecord) SetId(v int32)`

SetId sets Id field to given value.


### GetDomain

`func (o *ResultModelParsedZoneRecord) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelParsedZoneRecord) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelParsedZoneRecord) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetHost

`func (o *ResultModelParsedZoneRecord) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ResultModelParsedZoneRecord) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ResultModelParsedZoneRecord) SetHost(v string)`

SetHost sets Host field to given value.


### GetTtl

`func (o *ResultModelParsedZoneRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ResultModelParsedZoneRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ResultModelParsedZoneRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ResultModelParsedZoneRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetPrio

`func (o *ResultModelParsedZoneRecord) GetPrio() int32`

GetPrio returns the Prio field if non-nil, zero value otherwise.

### GetPrioOk

`func (o *ResultModelParsedZoneRecord) GetPrioOk() (*int32, bool)`

GetPrioOk returns a tuple with the Prio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrio

`func (o *ResultModelParsedZoneRecord) SetPrio(v int32)`

SetPrio sets Prio field to given value.

### HasPrio

`func (o *ResultModelParsedZoneRecord) HasPrio() bool`

HasPrio returns a boolean if a field has been set.

### GetGeozoneId

`func (o *ResultModelParsedZoneRecord) GetGeozoneId() int32`

GetGeozoneId returns the GeozoneId field if non-nil, zero value otherwise.

### GetGeozoneIdOk

`func (o *ResultModelParsedZoneRecord) GetGeozoneIdOk() (*int32, bool)`

GetGeozoneIdOk returns a tuple with the GeozoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeozoneId

`func (o *ResultModelParsedZoneRecord) SetGeozoneId(v int32)`

SetGeozoneId sets GeozoneId field to given value.

### HasGeozoneId

`func (o *ResultModelParsedZoneRecord) HasGeozoneId() bool`

HasGeozoneId returns a boolean if a field has been set.

### GetType

`func (o *ResultModelParsedZoneRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResultModelParsedZoneRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResultModelParsedZoneRecord) SetType(v string)`

SetType sets Type field to given value.


### GetRdata

`func (o *ResultModelParsedZoneRecord) GetRdata() string`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *ResultModelParsedZoneRecord) GetRdataOk() (*string, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *ResultModelParsedZoneRecord) SetRdata(v string)`

SetRdata sets Rdata field to given value.


### GetLastMod

`func (o *ResultModelParsedZoneRecord) GetLastMod() string`

GetLastMod returns the LastMod field if non-nil, zero value otherwise.

### GetLastModOk

`func (o *ResultModelParsedZoneRecord) GetLastModOk() (*string, bool)`

GetLastModOk returns a tuple with the LastMod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMod

`func (o *ResultModelParsedZoneRecord) SetLastMod(v string)`

SetLastMod sets LastMod field to given value.

### HasLastMod

`func (o *ResultModelParsedZoneRecord) HasLastMod() bool`

HasLastMod returns a boolean if a field has been set.

### GetUrl

`func (o *ResultModelParsedZoneRecord) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResultModelParsedZoneRecord) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResultModelParsedZoneRecord) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ResultModelParsedZoneRecord) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOrigRdata

`func (o *ResultModelParsedZoneRecord) GetOrigRdata() string`

GetOrigRdata returns the OrigRdata field if non-nil, zero value otherwise.

### GetOrigRdataOk

`func (o *ResultModelParsedZoneRecord) GetOrigRdataOk() (*string, bool)`

GetOrigRdataOk returns a tuple with the OrigRdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigRdata

`func (o *ResultModelParsedZoneRecord) SetOrigRdata(v string)`

SetOrigRdata sets OrigRdata field to given value.

### HasOrigRdata

`func (o *ResultModelParsedZoneRecord) HasOrigRdata() bool`

HasOrigRdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


