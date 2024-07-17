# ResultModelUserDomainListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **string** | The user the domains are associated with | 
**Index** | Pointer to [**[]ResultModelUserDomainListDataIndexInner**](ResultModelUserDomainListDataIndexInner.md) | An array of objects containing the domains | [optional] 

## Methods

### NewResultModelUserDomainListData

`func NewResultModelUserDomainListData(user string, ) *ResultModelUserDomainListData`

NewResultModelUserDomainListData instantiates a new ResultModelUserDomainListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelUserDomainListDataWithDefaults

`func NewResultModelUserDomainListDataWithDefaults() *ResultModelUserDomainListData`

NewResultModelUserDomainListDataWithDefaults instantiates a new ResultModelUserDomainListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ResultModelUserDomainListData) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ResultModelUserDomainListData) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ResultModelUserDomainListData) SetUser(v string)`

SetUser sets User field to given value.


### GetIndex

`func (o *ResultModelUserDomainListData) GetIndex() []ResultModelUserDomainListDataIndexInner`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ResultModelUserDomainListData) GetIndexOk() (*[]ResultModelUserDomainListDataIndexInner, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ResultModelUserDomainListData) SetIndex(v []ResultModelUserDomainListDataIndexInner)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ResultModelUserDomainListData) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


