# RequestModelCreateDomainGlue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NameserverName** | **string** | The hostname of the glue record being defined in FQDN format | 
**Ipaddress** | Pointer to **string** | The IPv4 address to define for this glue record if one is provided | [optional] 
**Ipv6** | Pointer to **string** | The IPv6 address to define for this glue record if one is provided | [optional] 

## Methods

### NewRequestModelCreateDomainGlue

`func NewRequestModelCreateDomainGlue(nameserverName string, ) *RequestModelCreateDomainGlue`

NewRequestModelCreateDomainGlue instantiates a new RequestModelCreateDomainGlue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestModelCreateDomainGlueWithDefaults

`func NewRequestModelCreateDomainGlueWithDefaults() *RequestModelCreateDomainGlue`

NewRequestModelCreateDomainGlueWithDefaults instantiates a new RequestModelCreateDomainGlue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameserverName

`func (o *RequestModelCreateDomainGlue) GetNameserverName() string`

GetNameserverName returns the NameserverName field if non-nil, zero value otherwise.

### GetNameserverNameOk

`func (o *RequestModelCreateDomainGlue) GetNameserverNameOk() (*string, bool)`

GetNameserverNameOk returns a tuple with the NameserverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserverName

`func (o *RequestModelCreateDomainGlue) SetNameserverName(v string)`

SetNameserverName sets NameserverName field to given value.


### GetIpaddress

`func (o *RequestModelCreateDomainGlue) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *RequestModelCreateDomainGlue) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *RequestModelCreateDomainGlue) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *RequestModelCreateDomainGlue) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetIpv6

`func (o *RequestModelCreateDomainGlue) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *RequestModelCreateDomainGlue) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *RequestModelCreateDomainGlue) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *RequestModelCreateDomainGlue) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


