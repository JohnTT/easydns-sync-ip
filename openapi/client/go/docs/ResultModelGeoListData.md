# ResultModelGeoListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The numeric ID of the described geo region | 
**GeoCode** | **string** | The short name for the described geo region | 
**Location** | **string** | Phyisical location of the described geo region | 

## Methods

### NewResultModelGeoListData

`func NewResultModelGeoListData(id int32, geoCode string, location string, ) *ResultModelGeoListData`

NewResultModelGeoListData instantiates a new ResultModelGeoListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelGeoListDataWithDefaults

`func NewResultModelGeoListDataWithDefaults() *ResultModelGeoListData`

NewResultModelGeoListDataWithDefaults instantiates a new ResultModelGeoListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResultModelGeoListData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultModelGeoListData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultModelGeoListData) SetId(v int32)`

SetId sets Id field to given value.


### GetGeoCode

`func (o *ResultModelGeoListData) GetGeoCode() string`

GetGeoCode returns the GeoCode field if non-nil, zero value otherwise.

### GetGeoCodeOk

`func (o *ResultModelGeoListData) GetGeoCodeOk() (*string, bool)`

GetGeoCodeOk returns a tuple with the GeoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoCode

`func (o *ResultModelGeoListData) SetGeoCode(v string)`

SetGeoCode sets GeoCode field to given value.


### GetLocation

`func (o *ResultModelGeoListData) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ResultModelGeoListData) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ResultModelGeoListData) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


