# ResultModelGetDomainGlueItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The hostname that the glue record is defined for | 
**Domain** | **string** | The domain the glue records are provided by | 
**Ipaddress** | Pointer to **string** | The IPv4 address defined for this glue record if one is defined | [optional] 
**Ipv6** | Pointer to **string** | The IPv6 address defined for this glue record if one is defined | [optional] 

## Methods

### NewResultModelGetDomainGlueItem

`func NewResultModelGetDomainGlueItem(name string, domain string, ) *ResultModelGetDomainGlueItem`

NewResultModelGetDomainGlueItem instantiates a new ResultModelGetDomainGlueItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelGetDomainGlueItemWithDefaults

`func NewResultModelGetDomainGlueItemWithDefaults() *ResultModelGetDomainGlueItem`

NewResultModelGetDomainGlueItemWithDefaults instantiates a new ResultModelGetDomainGlueItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResultModelGetDomainGlueItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResultModelGetDomainGlueItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResultModelGetDomainGlueItem) SetName(v string)`

SetName sets Name field to given value.


### GetDomain

`func (o *ResultModelGetDomainGlueItem) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelGetDomainGlueItem) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelGetDomainGlueItem) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetIpaddress

`func (o *ResultModelGetDomainGlueItem) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *ResultModelGetDomainGlueItem) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *ResultModelGetDomainGlueItem) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *ResultModelGetDomainGlueItem) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetIpv6

`func (o *ResultModelGetDomainGlueItem) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *ResultModelGetDomainGlueItem) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *ResultModelGetDomainGlueItem) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *ResultModelGetDomainGlueItem) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


