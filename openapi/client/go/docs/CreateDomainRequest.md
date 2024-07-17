# CreateDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** |  | 
**Term** | **int32** | The term (in years) to create the domain for. If domain is being registered this is also the term it will be registered for | 
**Currency** | **string** | The currency to use when processing billing for the new domain | 
**DnsOnly** | **int32** | Should this domain be added for DNS service only (excluding registration)? Requests that do not set this value will attempt registration of the provided domain name | 
**Nameservers** | Pointer to **[]string** | The nameservers to use when creating a domains NS records. If you are looking to set the primary NS for a secondary domain please use the primary_ns field | [optional] 
**Portfolio** | Pointer to **string** | The name of the portfolio to automatically assign this domain to when domain is created. The portfolio provided MUST already exist and originate from the same user as the domain is being added to. *This parameter has been deprecated by the domain_group parameter.* | [optional] 
**DomainGroup** | Pointer to **string** | The name of the domain group to automatically assign this domain to when domain is created. The domain group provided MUST already exist and originate from the same user as the user who will control the domain | [optional] 
**PrimaryNs** | Pointer to **string** | The primary nameserver(s) to use when adding a new secondary domain to the system. Providing this value will cause the domain to be set to secondary automatically and the zone transfer process to start. Multiple masters should be separated by a semi-colon and no spaces (i.e. &#39;1.2.3.4;5.6.7.8&#39;). | [optional] 
**Contacts** | Pointer to [**ContactDataSet**](ContactDataSet.md) |  | [optional] 
**Extra** | Pointer to [**NullableBodyDomainCreateDefaultExtra**](BodyDomainCreateDefaultExtra.md) |  | [optional] 

## Methods

### NewCreateDomainRequest

`func NewCreateDomainRequest(service string, term int32, currency string, dnsOnly int32, ) *CreateDomainRequest`

NewCreateDomainRequest instantiates a new CreateDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainRequestWithDefaults

`func NewCreateDomainRequestWithDefaults() *CreateDomainRequest`

NewCreateDomainRequestWithDefaults instantiates a new CreateDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CreateDomainRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateDomainRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateDomainRequest) SetService(v string)`

SetService sets Service field to given value.


### GetTerm

`func (o *CreateDomainRequest) GetTerm() int32`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *CreateDomainRequest) GetTermOk() (*int32, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *CreateDomainRequest) SetTerm(v int32)`

SetTerm sets Term field to given value.


### GetCurrency

`func (o *CreateDomainRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateDomainRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateDomainRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDnsOnly

`func (o *CreateDomainRequest) GetDnsOnly() int32`

GetDnsOnly returns the DnsOnly field if non-nil, zero value otherwise.

### GetDnsOnlyOk

`func (o *CreateDomainRequest) GetDnsOnlyOk() (*int32, bool)`

GetDnsOnlyOk returns a tuple with the DnsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsOnly

`func (o *CreateDomainRequest) SetDnsOnly(v int32)`

SetDnsOnly sets DnsOnly field to given value.


### GetNameservers

`func (o *CreateDomainRequest) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *CreateDomainRequest) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *CreateDomainRequest) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *CreateDomainRequest) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetPortfolio

`func (o *CreateDomainRequest) GetPortfolio() string`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *CreateDomainRequest) GetPortfolioOk() (*string, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *CreateDomainRequest) SetPortfolio(v string)`

SetPortfolio sets Portfolio field to given value.

### HasPortfolio

`func (o *CreateDomainRequest) HasPortfolio() bool`

HasPortfolio returns a boolean if a field has been set.

### GetDomainGroup

`func (o *CreateDomainRequest) GetDomainGroup() string`

GetDomainGroup returns the DomainGroup field if non-nil, zero value otherwise.

### GetDomainGroupOk

`func (o *CreateDomainRequest) GetDomainGroupOk() (*string, bool)`

GetDomainGroupOk returns a tuple with the DomainGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroup

`func (o *CreateDomainRequest) SetDomainGroup(v string)`

SetDomainGroup sets DomainGroup field to given value.

### HasDomainGroup

`func (o *CreateDomainRequest) HasDomainGroup() bool`

HasDomainGroup returns a boolean if a field has been set.

### GetPrimaryNs

`func (o *CreateDomainRequest) GetPrimaryNs() string`

GetPrimaryNs returns the PrimaryNs field if non-nil, zero value otherwise.

### GetPrimaryNsOk

`func (o *CreateDomainRequest) GetPrimaryNsOk() (*string, bool)`

GetPrimaryNsOk returns a tuple with the PrimaryNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryNs

`func (o *CreateDomainRequest) SetPrimaryNs(v string)`

SetPrimaryNs sets PrimaryNs field to given value.

### HasPrimaryNs

`func (o *CreateDomainRequest) HasPrimaryNs() bool`

HasPrimaryNs returns a boolean if a field has been set.

### GetContacts

`func (o *CreateDomainRequest) GetContacts() ContactDataSet`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *CreateDomainRequest) GetContactsOk() (*ContactDataSet, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *CreateDomainRequest) SetContacts(v ContactDataSet)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *CreateDomainRequest) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetExtra

`func (o *CreateDomainRequest) GetExtra() BodyDomainCreateDefaultExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *CreateDomainRequest) GetExtraOk() (*BodyDomainCreateDefaultExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *CreateDomainRequest) SetExtra(v BodyDomainCreateDefaultExtra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *CreateDomainRequest) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### SetExtraNil

`func (o *CreateDomainRequest) SetExtraNil(b bool)`

 SetExtraNil sets the value for Extra to be an explicit nil

### UnsetExtra
`func (o *CreateDomainRequest) UnsetExtra()`

UnsetExtra ensures that no value is present for Extra, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


