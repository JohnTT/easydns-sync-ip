# ResultModelServiceDescriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **int32** | The ID of the service included in the request | 
**Name** | **string** | The name of the service | 
**Period** | **int32** | The term the service is active before requiring renewal. Almost always presented in months | 
**Enterprise** | **int32** | Is the service an enterprise level service | 
**Description** | **string** | A description of the service offering | 

## Methods

### NewResultModelServiceDescriptionData

`func NewResultModelServiceDescriptionData(serviceId int32, name string, period int32, enterprise int32, description string, ) *ResultModelServiceDescriptionData`

NewResultModelServiceDescriptionData instantiates a new ResultModelServiceDescriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelServiceDescriptionDataWithDefaults

`func NewResultModelServiceDescriptionDataWithDefaults() *ResultModelServiceDescriptionData`

NewResultModelServiceDescriptionDataWithDefaults instantiates a new ResultModelServiceDescriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ResultModelServiceDescriptionData) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ResultModelServiceDescriptionData) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ResultModelServiceDescriptionData) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.


### GetName

`func (o *ResultModelServiceDescriptionData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResultModelServiceDescriptionData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResultModelServiceDescriptionData) SetName(v string)`

SetName sets Name field to given value.


### GetPeriod

`func (o *ResultModelServiceDescriptionData) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ResultModelServiceDescriptionData) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ResultModelServiceDescriptionData) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetEnterprise

`func (o *ResultModelServiceDescriptionData) GetEnterprise() int32`

GetEnterprise returns the Enterprise field if non-nil, zero value otherwise.

### GetEnterpriseOk

`func (o *ResultModelServiceDescriptionData) GetEnterpriseOk() (*int32, bool)`

GetEnterpriseOk returns a tuple with the Enterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprise

`func (o *ResultModelServiceDescriptionData) SetEnterprise(v int32)`

SetEnterprise sets Enterprise field to given value.


### GetDescription

`func (o *ResultModelServiceDescriptionData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResultModelServiceDescriptionData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResultModelServiceDescriptionData) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


