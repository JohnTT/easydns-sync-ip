# ResultModelDomainNameserversData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain the glue records are provided by | 
**Nameservers** | **[]string** | The list of nameservers assigned to the domain | 

## Methods

### NewResultModelDomainNameserversData

`func NewResultModelDomainNameserversData(domain string, nameservers []string, ) *ResultModelDomainNameserversData`

NewResultModelDomainNameserversData instantiates a new ResultModelDomainNameserversData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelDomainNameserversDataWithDefaults

`func NewResultModelDomainNameserversDataWithDefaults() *ResultModelDomainNameserversData`

NewResultModelDomainNameserversDataWithDefaults instantiates a new ResultModelDomainNameserversData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ResultModelDomainNameserversData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelDomainNameserversData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelDomainNameserversData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetNameservers

`func (o *ResultModelDomainNameserversData) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *ResultModelDomainNameserversData) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *ResultModelDomainNameserversData) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


