# ResultModelSubscriptionDescriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **int32** | The ID of the subscription block included in the request | 
**Name** | **string** | The name of the service | 
**ServiceId** | **int32** | The ID of the service provided by the requested subscription block | 
**Period** | **int32** | The term the subscription block is active before requiring renewal. Almost always presented in months | 
**Enterprise** | **int32** | Is the service an enterprise level service | 
**Description** | **string** | A description of the service offering | 
**Size** | **int32** | The number of domains the subscription block supports | 

## Methods

### NewResultModelSubscriptionDescriptionData

`func NewResultModelSubscriptionDescriptionData(subscriptionId int32, name string, serviceId int32, period int32, enterprise int32, description string, size int32, ) *ResultModelSubscriptionDescriptionData`

NewResultModelSubscriptionDescriptionData instantiates a new ResultModelSubscriptionDescriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelSubscriptionDescriptionDataWithDefaults

`func NewResultModelSubscriptionDescriptionDataWithDefaults() *ResultModelSubscriptionDescriptionData`

NewResultModelSubscriptionDescriptionDataWithDefaults instantiates a new ResultModelSubscriptionDescriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *ResultModelSubscriptionDescriptionData) GetSubscriptionId() int32`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ResultModelSubscriptionDescriptionData) GetSubscriptionIdOk() (*int32, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ResultModelSubscriptionDescriptionData) SetSubscriptionId(v int32)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetName

`func (o *ResultModelSubscriptionDescriptionData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResultModelSubscriptionDescriptionData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResultModelSubscriptionDescriptionData) SetName(v string)`

SetName sets Name field to given value.


### GetServiceId

`func (o *ResultModelSubscriptionDescriptionData) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ResultModelSubscriptionDescriptionData) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ResultModelSubscriptionDescriptionData) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.


### GetPeriod

`func (o *ResultModelSubscriptionDescriptionData) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ResultModelSubscriptionDescriptionData) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ResultModelSubscriptionDescriptionData) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetEnterprise

`func (o *ResultModelSubscriptionDescriptionData) GetEnterprise() int32`

GetEnterprise returns the Enterprise field if non-nil, zero value otherwise.

### GetEnterpriseOk

`func (o *ResultModelSubscriptionDescriptionData) GetEnterpriseOk() (*int32, bool)`

GetEnterpriseOk returns a tuple with the Enterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprise

`func (o *ResultModelSubscriptionDescriptionData) SetEnterprise(v int32)`

SetEnterprise sets Enterprise field to given value.


### GetDescription

`func (o *ResultModelSubscriptionDescriptionData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResultModelSubscriptionDescriptionData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResultModelSubscriptionDescriptionData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSize

`func (o *ResultModelSubscriptionDescriptionData) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResultModelSubscriptionDescriptionData) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResultModelSubscriptionDescriptionData) SetSize(v int32)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


