# ResultModelCheckRegistryGlueData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain the glue records are provided by | 
**Fqdn** | **string** | The hostname for the glue record that was looked up | 
**Exists** | **int32** | Boolean integer that indicates if the glue is configured at the registry | 

## Methods

### NewResultModelCheckRegistryGlueData

`func NewResultModelCheckRegistryGlueData(domain string, fqdn string, exists int32, ) *ResultModelCheckRegistryGlueData`

NewResultModelCheckRegistryGlueData instantiates a new ResultModelCheckRegistryGlueData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelCheckRegistryGlueDataWithDefaults

`func NewResultModelCheckRegistryGlueDataWithDefaults() *ResultModelCheckRegistryGlueData`

NewResultModelCheckRegistryGlueDataWithDefaults instantiates a new ResultModelCheckRegistryGlueData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ResultModelCheckRegistryGlueData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelCheckRegistryGlueData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelCheckRegistryGlueData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetFqdn

`func (o *ResultModelCheckRegistryGlueData) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ResultModelCheckRegistryGlueData) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ResultModelCheckRegistryGlueData) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetExists

`func (o *ResultModelCheckRegistryGlueData) GetExists() int32`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *ResultModelCheckRegistryGlueData) GetExistsOk() (*int32, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *ResultModelCheckRegistryGlueData) SetExists(v int32)`

SetExists sets Exists field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


