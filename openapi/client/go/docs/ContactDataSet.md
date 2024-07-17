# ContactDataSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | [**ContactDataOwner**](ContactDataOwner.md) |  | 
**Admin** | Pointer to [**ContactDataAdmin**](ContactDataAdmin.md) |  | [optional] 
**Tech** | Pointer to [**ContactDataTech**](ContactDataTech.md) |  | [optional] 
**Billing** | Pointer to [**ContactDataBilling**](ContactDataBilling.md) |  | [optional] 

## Methods

### NewContactDataSet

`func NewContactDataSet(owner ContactDataOwner, ) *ContactDataSet`

NewContactDataSet instantiates a new ContactDataSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactDataSetWithDefaults

`func NewContactDataSetWithDefaults() *ContactDataSet`

NewContactDataSetWithDefaults instantiates a new ContactDataSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *ContactDataSet) GetOwner() ContactDataOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ContactDataSet) GetOwnerOk() (*ContactDataOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ContactDataSet) SetOwner(v ContactDataOwner)`

SetOwner sets Owner field to given value.


### GetAdmin

`func (o *ContactDataSet) GetAdmin() ContactDataAdmin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ContactDataSet) GetAdminOk() (*ContactDataAdmin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ContactDataSet) SetAdmin(v ContactDataAdmin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ContactDataSet) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetTech

`func (o *ContactDataSet) GetTech() ContactDataTech`

GetTech returns the Tech field if non-nil, zero value otherwise.

### GetTechOk

`func (o *ContactDataSet) GetTechOk() (*ContactDataTech, bool)`

GetTechOk returns a tuple with the Tech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTech

`func (o *ContactDataSet) SetTech(v ContactDataTech)`

SetTech sets Tech field to given value.

### HasTech

`func (o *ContactDataSet) HasTech() bool`

HasTech returns a boolean if a field has been set.

### GetBilling

`func (o *ContactDataSet) GetBilling() ContactDataBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *ContactDataSet) GetBillingOk() (*ContactDataBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *ContactDataSet) SetBilling(v ContactDataBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *ContactDataSet) HasBilling() bool`

HasBilling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


