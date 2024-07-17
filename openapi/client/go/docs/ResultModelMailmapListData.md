# ResultModelMailmapListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain name the mailmaps below to | 
**Mailmaps** | [**[]ResultModelMailmapListItem**](ResultModelMailmapListItem.md) | List of mailmaps defined for a domain | 

## Methods

### NewResultModelMailmapListData

`func NewResultModelMailmapListData(domain string, mailmaps []ResultModelMailmapListItem, ) *ResultModelMailmapListData`

NewResultModelMailmapListData instantiates a new ResultModelMailmapListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelMailmapListDataWithDefaults

`func NewResultModelMailmapListDataWithDefaults() *ResultModelMailmapListData`

NewResultModelMailmapListDataWithDefaults instantiates a new ResultModelMailmapListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ResultModelMailmapListData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelMailmapListData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelMailmapListData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetMailmaps

`func (o *ResultModelMailmapListData) GetMailmaps() []ResultModelMailmapListItem`

GetMailmaps returns the Mailmaps field if non-nil, zero value otherwise.

### GetMailmapsOk

`func (o *ResultModelMailmapListData) GetMailmapsOk() (*[]ResultModelMailmapListItem, bool)`

GetMailmapsOk returns a tuple with the Mailmaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailmaps

`func (o *ResultModelMailmapListData) SetMailmaps(v []ResultModelMailmapListItem)`

SetMailmaps sets Mailmaps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


