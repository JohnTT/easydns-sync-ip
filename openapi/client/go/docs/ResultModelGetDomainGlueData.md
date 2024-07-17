# ResultModelGetDomainGlueData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain the glue records are provided by | 
**Total** | **int32** | Total number of glue records returned for domain | 
**GlueRecords** | [**[]ResultModelGetDomainGlueItem**](ResultModelGetDomainGlueItem.md) | The list of glue records for a domain | 

## Methods

### NewResultModelGetDomainGlueData

`func NewResultModelGetDomainGlueData(domain string, total int32, glueRecords []ResultModelGetDomainGlueItem, ) *ResultModelGetDomainGlueData`

NewResultModelGetDomainGlueData instantiates a new ResultModelGetDomainGlueData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelGetDomainGlueDataWithDefaults

`func NewResultModelGetDomainGlueDataWithDefaults() *ResultModelGetDomainGlueData`

NewResultModelGetDomainGlueDataWithDefaults instantiates a new ResultModelGetDomainGlueData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ResultModelGetDomainGlueData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelGetDomainGlueData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelGetDomainGlueData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetTotal

`func (o *ResultModelGetDomainGlueData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResultModelGetDomainGlueData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResultModelGetDomainGlueData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetGlueRecords

`func (o *ResultModelGetDomainGlueData) GetGlueRecords() []ResultModelGetDomainGlueItem`

GetGlueRecords returns the GlueRecords field if non-nil, zero value otherwise.

### GetGlueRecordsOk

`func (o *ResultModelGetDomainGlueData) GetGlueRecordsOk() (*[]ResultModelGetDomainGlueItem, bool)`

GetGlueRecordsOk returns a tuple with the GlueRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlueRecords

`func (o *ResultModelGetDomainGlueData) SetGlueRecords(v []ResultModelGetDomainGlueItem)`

SetGlueRecords sets GlueRecords field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


