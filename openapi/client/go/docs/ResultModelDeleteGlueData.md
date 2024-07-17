# ResultModelDeleteGlueData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain the glue records are provided by | 
**Fqdn** | **string** | The hostname for the glue record that was looked up | 

## Methods

### NewResultModelDeleteGlueData

`func NewResultModelDeleteGlueData(domain string, fqdn string, ) *ResultModelDeleteGlueData`

NewResultModelDeleteGlueData instantiates a new ResultModelDeleteGlueData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelDeleteGlueDataWithDefaults

`func NewResultModelDeleteGlueDataWithDefaults() *ResultModelDeleteGlueData`

NewResultModelDeleteGlueDataWithDefaults instantiates a new ResultModelDeleteGlueData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ResultModelDeleteGlueData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelDeleteGlueData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelDeleteGlueData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetFqdn

`func (o *ResultModelDeleteGlueData) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ResultModelDeleteGlueData) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ResultModelDeleteGlueData) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


