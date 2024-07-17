# ResultModelCreateDomainGlueData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain the glue records are provided by | 
**NameserverName** | **string** | The hostname that the glue record is defined for | 
**Ipaddress** | Pointer to **string** | The IPv4 address defined for this glue record if one is defined | [optional] 
**Ipv6** | Pointer to **string** | The IPv6 address defined for this glue record if one is defined | [optional] 

## Methods

### NewResultModelCreateDomainGlueData

`func NewResultModelCreateDomainGlueData(domain string, nameserverName string, ) *ResultModelCreateDomainGlueData`

NewResultModelCreateDomainGlueData instantiates a new ResultModelCreateDomainGlueData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelCreateDomainGlueDataWithDefaults

`func NewResultModelCreateDomainGlueDataWithDefaults() *ResultModelCreateDomainGlueData`

NewResultModelCreateDomainGlueDataWithDefaults instantiates a new ResultModelCreateDomainGlueData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ResultModelCreateDomainGlueData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelCreateDomainGlueData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelCreateDomainGlueData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetNameserverName

`func (o *ResultModelCreateDomainGlueData) GetNameserverName() string`

GetNameserverName returns the NameserverName field if non-nil, zero value otherwise.

### GetNameserverNameOk

`func (o *ResultModelCreateDomainGlueData) GetNameserverNameOk() (*string, bool)`

GetNameserverNameOk returns a tuple with the NameserverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserverName

`func (o *ResultModelCreateDomainGlueData) SetNameserverName(v string)`

SetNameserverName sets NameserverName field to given value.


### GetIpaddress

`func (o *ResultModelCreateDomainGlueData) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *ResultModelCreateDomainGlueData) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *ResultModelCreateDomainGlueData) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *ResultModelCreateDomainGlueData) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetIpv6

`func (o *ResultModelCreateDomainGlueData) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *ResultModelCreateDomainGlueData) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *ResultModelCreateDomainGlueData) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *ResultModelCreateDomainGlueData) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


