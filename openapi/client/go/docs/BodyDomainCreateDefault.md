# BodyDomainCreateDefault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** |  | 
**Term** | **int32** | The term (in years) to create the domain for. If domain is being registered this is also the term it will be registered for | 
**Currency** | **string** | The currency to use when processing billing for the new domain | 
**DnsOnly** | Pointer to **int32** | Should this domain be added for DNS service only (excluding registration)? Requests that do not set this value will attempt registration of the provided domain name | [optional] [default to 0]
**Nameservers** | Pointer to **[]string** |  | [optional] 
**Portfolio** | Pointer to **string** | The name of the portfolio to automatically assign this domain to when domain is created. The portfolio provided MUST already exist and originate from the same user as the user who will control the domain. *This parameter has been deprecated by the domain_group parameter.* | [optional] 
**DomainGroup** | Pointer to **string** | The name of the domain group to automatically assign this domain to when domain is created. The domain group provided MUST already exist and originate from the same user as the user who will control the domain | [optional] 
**PrimaryNs** | Pointer to **string** | The primary nameserver(s) to use when adding a new secondary domain to the system. Providing this value will cause the domain to be set to secondary automatically and the zone transfer process to start. Multiple masters should be separated by a semi-colon and no spaces (i.e. &#39;1.2.3.4;5.6.7.8&#39;). | [optional] 
**Contacts** | Pointer to [**ContactDataSet**](ContactDataSet.md) |  | [optional] 
**Extra** | Pointer to [**NullableBodyDomainCreateDefaultExtra**](BodyDomainCreateDefaultExtra.md) |  | [optional] 

## Methods

### NewBodyDomainCreateDefault

`func NewBodyDomainCreateDefault(service string, term int32, currency string, ) *BodyDomainCreateDefault`

NewBodyDomainCreateDefault instantiates a new BodyDomainCreateDefault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodyDomainCreateDefaultWithDefaults

`func NewBodyDomainCreateDefaultWithDefaults() *BodyDomainCreateDefault`

NewBodyDomainCreateDefaultWithDefaults instantiates a new BodyDomainCreateDefault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *BodyDomainCreateDefault) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BodyDomainCreateDefault) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BodyDomainCreateDefault) SetService(v string)`

SetService sets Service field to given value.


### GetTerm

`func (o *BodyDomainCreateDefault) GetTerm() int32`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *BodyDomainCreateDefault) GetTermOk() (*int32, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *BodyDomainCreateDefault) SetTerm(v int32)`

SetTerm sets Term field to given value.


### GetCurrency

`func (o *BodyDomainCreateDefault) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BodyDomainCreateDefault) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BodyDomainCreateDefault) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDnsOnly

`func (o *BodyDomainCreateDefault) GetDnsOnly() int32`

GetDnsOnly returns the DnsOnly field if non-nil, zero value otherwise.

### GetDnsOnlyOk

`func (o *BodyDomainCreateDefault) GetDnsOnlyOk() (*int32, bool)`

GetDnsOnlyOk returns a tuple with the DnsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsOnly

`func (o *BodyDomainCreateDefault) SetDnsOnly(v int32)`

SetDnsOnly sets DnsOnly field to given value.

### HasDnsOnly

`func (o *BodyDomainCreateDefault) HasDnsOnly() bool`

HasDnsOnly returns a boolean if a field has been set.

### GetNameservers

`func (o *BodyDomainCreateDefault) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *BodyDomainCreateDefault) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *BodyDomainCreateDefault) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *BodyDomainCreateDefault) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetPortfolio

`func (o *BodyDomainCreateDefault) GetPortfolio() string`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *BodyDomainCreateDefault) GetPortfolioOk() (*string, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *BodyDomainCreateDefault) SetPortfolio(v string)`

SetPortfolio sets Portfolio field to given value.

### HasPortfolio

`func (o *BodyDomainCreateDefault) HasPortfolio() bool`

HasPortfolio returns a boolean if a field has been set.

### GetDomainGroup

`func (o *BodyDomainCreateDefault) GetDomainGroup() string`

GetDomainGroup returns the DomainGroup field if non-nil, zero value otherwise.

### GetDomainGroupOk

`func (o *BodyDomainCreateDefault) GetDomainGroupOk() (*string, bool)`

GetDomainGroupOk returns a tuple with the DomainGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroup

`func (o *BodyDomainCreateDefault) SetDomainGroup(v string)`

SetDomainGroup sets DomainGroup field to given value.

### HasDomainGroup

`func (o *BodyDomainCreateDefault) HasDomainGroup() bool`

HasDomainGroup returns a boolean if a field has been set.

### GetPrimaryNs

`func (o *BodyDomainCreateDefault) GetPrimaryNs() string`

GetPrimaryNs returns the PrimaryNs field if non-nil, zero value otherwise.

### GetPrimaryNsOk

`func (o *BodyDomainCreateDefault) GetPrimaryNsOk() (*string, bool)`

GetPrimaryNsOk returns a tuple with the PrimaryNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryNs

`func (o *BodyDomainCreateDefault) SetPrimaryNs(v string)`

SetPrimaryNs sets PrimaryNs field to given value.

### HasPrimaryNs

`func (o *BodyDomainCreateDefault) HasPrimaryNs() bool`

HasPrimaryNs returns a boolean if a field has been set.

### GetContacts

`func (o *BodyDomainCreateDefault) GetContacts() ContactDataSet`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *BodyDomainCreateDefault) GetContactsOk() (*ContactDataSet, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *BodyDomainCreateDefault) SetContacts(v ContactDataSet)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *BodyDomainCreateDefault) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetExtra

`func (o *BodyDomainCreateDefault) GetExtra() BodyDomainCreateDefaultExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *BodyDomainCreateDefault) GetExtraOk() (*BodyDomainCreateDefaultExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *BodyDomainCreateDefault) SetExtra(v BodyDomainCreateDefaultExtra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *BodyDomainCreateDefault) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### SetExtraNil

`func (o *BodyDomainCreateDefault) SetExtraNil(b bool)`

 SetExtraNil sets the value for Extra to be an explicit nil

### UnsetExtra
`func (o *BodyDomainCreateDefault) UnsetExtra()`

UnsetExtra ensures that no value is present for Extra, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


