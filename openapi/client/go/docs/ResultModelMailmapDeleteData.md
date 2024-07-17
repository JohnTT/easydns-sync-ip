# ResultModelMailmapDeleteData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain the mailmap was created for | 
**MailmapId** | **int32** | The unique ID of the deleted mailmap | 

## Methods

### NewResultModelMailmapDeleteData

`func NewResultModelMailmapDeleteData(domain string, mailmapId int32, ) *ResultModelMailmapDeleteData`

NewResultModelMailmapDeleteData instantiates a new ResultModelMailmapDeleteData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelMailmapDeleteDataWithDefaults

`func NewResultModelMailmapDeleteDataWithDefaults() *ResultModelMailmapDeleteData`

NewResultModelMailmapDeleteDataWithDefaults instantiates a new ResultModelMailmapDeleteData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ResultModelMailmapDeleteData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelMailmapDeleteData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelMailmapDeleteData) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetMailmapId

`func (o *ResultModelMailmapDeleteData) GetMailmapId() int32`

GetMailmapId returns the MailmapId field if non-nil, zero value otherwise.

### GetMailmapIdOk

`func (o *ResultModelMailmapDeleteData) GetMailmapIdOk() (*int32, bool)`

GetMailmapIdOk returns a tuple with the MailmapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailmapId

`func (o *ResultModelMailmapDeleteData) SetMailmapId(v int32)`

SetMailmapId sets MailmapId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


