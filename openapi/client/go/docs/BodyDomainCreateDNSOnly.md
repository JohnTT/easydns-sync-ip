# BodyDomainCreateDNSOnly

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

## Methods

### NewBodyDomainCreateDNSOnly

`func NewBodyDomainCreateDNSOnly(service string, term int32, currency string, dnsOnly int32, ) *BodyDomainCreateDNSOnly`

NewBodyDomainCreateDNSOnly instantiates a new BodyDomainCreateDNSOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodyDomainCreateDNSOnlyWithDefaults

`func NewBodyDomainCreateDNSOnlyWithDefaults() *BodyDomainCreateDNSOnly`

NewBodyDomainCreateDNSOnlyWithDefaults instantiates a new BodyDomainCreateDNSOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *BodyDomainCreateDNSOnly) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BodyDomainCreateDNSOnly) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BodyDomainCreateDNSOnly) SetService(v string)`

SetService sets Service field to given value.


### GetTerm

`func (o *BodyDomainCreateDNSOnly) GetTerm() int32`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *BodyDomainCreateDNSOnly) GetTermOk() (*int32, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *BodyDomainCreateDNSOnly) SetTerm(v int32)`

SetTerm sets Term field to given value.


### GetCurrency

`func (o *BodyDomainCreateDNSOnly) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BodyDomainCreateDNSOnly) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BodyDomainCreateDNSOnly) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDnsOnly

`func (o *BodyDomainCreateDNSOnly) GetDnsOnly() int32`

GetDnsOnly returns the DnsOnly field if non-nil, zero value otherwise.

### GetDnsOnlyOk

`func (o *BodyDomainCreateDNSOnly) GetDnsOnlyOk() (*int32, bool)`

GetDnsOnlyOk returns a tuple with the DnsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsOnly

`func (o *BodyDomainCreateDNSOnly) SetDnsOnly(v int32)`

SetDnsOnly sets DnsOnly field to given value.


### GetNameservers

`func (o *BodyDomainCreateDNSOnly) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *BodyDomainCreateDNSOnly) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *BodyDomainCreateDNSOnly) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *BodyDomainCreateDNSOnly) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetPortfolio

`func (o *BodyDomainCreateDNSOnly) GetPortfolio() string`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *BodyDomainCreateDNSOnly) GetPortfolioOk() (*string, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *BodyDomainCreateDNSOnly) SetPortfolio(v string)`

SetPortfolio sets Portfolio field to given value.

### HasPortfolio

`func (o *BodyDomainCreateDNSOnly) HasPortfolio() bool`

HasPortfolio returns a boolean if a field has been set.

### GetDomainGroup

`func (o *BodyDomainCreateDNSOnly) GetDomainGroup() string`

GetDomainGroup returns the DomainGroup field if non-nil, zero value otherwise.

### GetDomainGroupOk

`func (o *BodyDomainCreateDNSOnly) GetDomainGroupOk() (*string, bool)`

GetDomainGroupOk returns a tuple with the DomainGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroup

`func (o *BodyDomainCreateDNSOnly) SetDomainGroup(v string)`

SetDomainGroup sets DomainGroup field to given value.

### HasDomainGroup

`func (o *BodyDomainCreateDNSOnly) HasDomainGroup() bool`

HasDomainGroup returns a boolean if a field has been set.

### GetPrimaryNs

`func (o *BodyDomainCreateDNSOnly) GetPrimaryNs() string`

GetPrimaryNs returns the PrimaryNs field if non-nil, zero value otherwise.

### GetPrimaryNsOk

`func (o *BodyDomainCreateDNSOnly) GetPrimaryNsOk() (*string, bool)`

GetPrimaryNsOk returns a tuple with the PrimaryNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryNs

`func (o *BodyDomainCreateDNSOnly) SetPrimaryNs(v string)`

SetPrimaryNs sets PrimaryNs field to given value.

### HasPrimaryNs

`func (o *BodyDomainCreateDNSOnly) HasPrimaryNs() bool`

HasPrimaryNs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


