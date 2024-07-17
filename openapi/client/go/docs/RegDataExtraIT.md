# RegDataExtraIT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **int32** | Entity type | 
**NationalityCode** | **string** | Nationality code (2 letter ISO country code) | 
**RegCode** | **string** | VAT or Codice Fiscale or n.a. | 

## Methods

### NewRegDataExtraIT

`func NewRegDataExtraIT(entityType int32, nationalityCode string, regCode string, ) *RegDataExtraIT`

NewRegDataExtraIT instantiates a new RegDataExtraIT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegDataExtraITWithDefaults

`func NewRegDataExtraITWithDefaults() *RegDataExtraIT`

NewRegDataExtraITWithDefaults instantiates a new RegDataExtraIT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *RegDataExtraIT) GetEntityType() int32`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *RegDataExtraIT) GetEntityTypeOk() (*int32, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *RegDataExtraIT) SetEntityType(v int32)`

SetEntityType sets EntityType field to given value.


### GetNationalityCode

`func (o *RegDataExtraIT) GetNationalityCode() string`

GetNationalityCode returns the NationalityCode field if non-nil, zero value otherwise.

### GetNationalityCodeOk

`func (o *RegDataExtraIT) GetNationalityCodeOk() (*string, bool)`

GetNationalityCodeOk returns a tuple with the NationalityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityCode

`func (o *RegDataExtraIT) SetNationalityCode(v string)`

SetNationalityCode sets NationalityCode field to given value.


### GetRegCode

`func (o *RegDataExtraIT) GetRegCode() string`

GetRegCode returns the RegCode field if non-nil, zero value otherwise.

### GetRegCodeOk

`func (o *RegDataExtraIT) GetRegCodeOk() (*string, bool)`

GetRegCodeOk returns a tuple with the RegCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegCode

`func (o *RegDataExtraIT) SetRegCode(v string)`

SetRegCode sets RegCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


