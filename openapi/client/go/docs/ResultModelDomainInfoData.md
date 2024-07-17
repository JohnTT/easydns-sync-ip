# ResultModelDomainInfoData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The exact identifier for domain. Almost always just the domain name | [optional] 
**Domain** | **string** | The domain name the returned information is associated with | 
**Exists** | **string** | Does the domain currently exist at the registry | 
**Onsystem** | **string** | Does the domain currently exist on the system | 
**Expiry** | **string** | Expiry date of the domain if it differs from the next_due date. Domains without registration will return a value of \&quot;false\&quot; | 
**NextDue** | **string** | The date the service for the domain is next due | 
**ClonedTo** | Pointer to **string** | When cloning is enabled this will indicate the domain that is being cloned | [optional] 
**Service** | **int32** | The service ID that the domain is currently using | 
**SubBlock** | Pointer to **int32** | Indicates the ID of the subscription block the domain is associated with when applicable | [optional] 

## Methods

### NewResultModelDomainInfoData

`func NewResultModelDomainInfoData(domain string, exists string, onsystem string, expiry string, nextDue string, service int32, ) *ResultModelDomainInfoData`

NewResultModelDomainInfoData instantiates a new ResultModelDomainInfoData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelDomainInfoDataWithDefaults

`func NewResultModelDomainInfoDataWithDefaults() *ResultModelDomainInfoData`

NewResultModelDomainInfoDataWithDefaults instantiates a new ResultModelDomainInfoData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResultModelDomainInfoData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultModelDomainInfoData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultModelDomainInfoData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResultModelDomainInfoData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *ResultModelDomainInfoData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelDomainInfoData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelDomainInfoData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetExists

`func (o *ResultModelDomainInfoData) GetExists() string`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *ResultModelDomainInfoData) GetExistsOk() (*string, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *ResultModelDomainInfoData) SetExists(v string)`

SetExists sets Exists field to given value.


### GetOnsystem

`func (o *ResultModelDomainInfoData) GetOnsystem() string`

GetOnsystem returns the Onsystem field if non-nil, zero value otherwise.

### GetOnsystemOk

`func (o *ResultModelDomainInfoData) GetOnsystemOk() (*string, bool)`

GetOnsystemOk returns a tuple with the Onsystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnsystem

`func (o *ResultModelDomainInfoData) SetOnsystem(v string)`

SetOnsystem sets Onsystem field to given value.


### GetExpiry

`func (o *ResultModelDomainInfoData) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ResultModelDomainInfoData) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ResultModelDomainInfoData) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.


### GetNextDue

`func (o *ResultModelDomainInfoData) GetNextDue() string`

GetNextDue returns the NextDue field if non-nil, zero value otherwise.

### GetNextDueOk

`func (o *ResultModelDomainInfoData) GetNextDueOk() (*string, bool)`

GetNextDueOk returns a tuple with the NextDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDue

`func (o *ResultModelDomainInfoData) SetNextDue(v string)`

SetNextDue sets NextDue field to given value.


### GetClonedTo

`func (o *ResultModelDomainInfoData) GetClonedTo() string`

GetClonedTo returns the ClonedTo field if non-nil, zero value otherwise.

### GetClonedToOk

`func (o *ResultModelDomainInfoData) GetClonedToOk() (*string, bool)`

GetClonedToOk returns a tuple with the ClonedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedTo

`func (o *ResultModelDomainInfoData) SetClonedTo(v string)`

SetClonedTo sets ClonedTo field to given value.

### HasClonedTo

`func (o *ResultModelDomainInfoData) HasClonedTo() bool`

HasClonedTo returns a boolean if a field has been set.

### GetService

`func (o *ResultModelDomainInfoData) GetService() int32`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ResultModelDomainInfoData) GetServiceOk() (*int32, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ResultModelDomainInfoData) SetService(v int32)`

SetService sets Service field to given value.


### GetSubBlock

`func (o *ResultModelDomainInfoData) GetSubBlock() int32`

GetSubBlock returns the SubBlock field if non-nil, zero value otherwise.

### GetSubBlockOk

`func (o *ResultModelDomainInfoData) GetSubBlockOk() (*int32, bool)`

GetSubBlockOk returns a tuple with the SubBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubBlock

`func (o *ResultModelDomainInfoData) SetSubBlock(v int32)`

SetSubBlock sets SubBlock field to given value.

### HasSubBlock

`func (o *ResultModelDomainInfoData) HasSubBlock() bool`

HasSubBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


