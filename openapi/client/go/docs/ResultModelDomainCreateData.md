# ResultModelDomainCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain name processed as part of the request | 
**Term** | **int32** | The period (in years) to create the domain for on the system/registry (if applicable) | 
**Service** | **int32** | The service ID the domain was created as | 
**Tld** | **string** |  | 
**InvId** | **int32** | The ID of the invoice that was created and processed to fullfill the request | 
**Currency** | Pointer to **string** | The currency used to process the invoice | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewResultModelDomainCreateData

`func NewResultModelDomainCreateData(domain string, term int32, service int32, tld string, invId int32, ) *ResultModelDomainCreateData`

NewResultModelDomainCreateData instantiates a new ResultModelDomainCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelDomainCreateDataWithDefaults

`func NewResultModelDomainCreateDataWithDefaults() *ResultModelDomainCreateData`

NewResultModelDomainCreateDataWithDefaults instantiates a new ResultModelDomainCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ResultModelDomainCreateData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelDomainCreateData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelDomainCreateData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetTerm

`func (o *ResultModelDomainCreateData) GetTerm() int32`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *ResultModelDomainCreateData) GetTermOk() (*int32, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *ResultModelDomainCreateData) SetTerm(v int32)`

SetTerm sets Term field to given value.


### GetService

`func (o *ResultModelDomainCreateData) GetService() int32`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ResultModelDomainCreateData) GetServiceOk() (*int32, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ResultModelDomainCreateData) SetService(v int32)`

SetService sets Service field to given value.


### GetTld

`func (o *ResultModelDomainCreateData) GetTld() string`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *ResultModelDomainCreateData) GetTldOk() (*string, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *ResultModelDomainCreateData) SetTld(v string)`

SetTld sets Tld field to given value.


### GetInvId

`func (o *ResultModelDomainCreateData) GetInvId() int32`

GetInvId returns the InvId field if non-nil, zero value otherwise.

### GetInvIdOk

`func (o *ResultModelDomainCreateData) GetInvIdOk() (*int32, bool)`

GetInvIdOk returns a tuple with the InvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvId

`func (o *ResultModelDomainCreateData) SetInvId(v int32)`

SetInvId sets InvId field to given value.


### GetCurrency

`func (o *ResultModelDomainCreateData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ResultModelDomainCreateData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ResultModelDomainCreateData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ResultModelDomainCreateData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUser

`func (o *ResultModelDomainCreateData) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ResultModelDomainCreateData) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ResultModelDomainCreateData) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ResultModelDomainCreateData) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


