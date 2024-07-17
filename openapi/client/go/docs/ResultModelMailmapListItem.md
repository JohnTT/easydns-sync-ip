# ResultModelMailmapListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **int32** | Status flag indicating if a mailmap is currently active | 
**Alias** | **string** | The alias (or email address) of the mailmap | 
**Destination** | **string** | A list of email addresses the mailmap forwards to | 
**Domain** | **string** | The domain the mailmap belongs to | 
**Host** | **string** | The hostname of the mailmap relative to the domain name | 
**LastModified** | **string** | The date the mailmap was last modified | 
**MailmapId** | **int32** | The unique ID that identifies this mailmap | 

## Methods

### NewResultModelMailmapListItem

`func NewResultModelMailmapListItem(active int32, alias string, destination string, domain string, host string, lastModified string, mailmapId int32, ) *ResultModelMailmapListItem`

NewResultModelMailmapListItem instantiates a new ResultModelMailmapListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelMailmapListItemWithDefaults

`func NewResultModelMailmapListItemWithDefaults() *ResultModelMailmapListItem`

NewResultModelMailmapListItemWithDefaults instantiates a new ResultModelMailmapListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ResultModelMailmapListItem) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResultModelMailmapListItem) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResultModelMailmapListItem) SetActive(v int32)`

SetActive sets Active field to given value.


### GetAlias

`func (o *ResultModelMailmapListItem) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ResultModelMailmapListItem) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ResultModelMailmapListItem) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetDestination

`func (o *ResultModelMailmapListItem) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ResultModelMailmapListItem) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ResultModelMailmapListItem) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetDomain

`func (o *ResultModelMailmapListItem) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelMailmapListItem) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelMailmapListItem) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetHost

`func (o *ResultModelMailmapListItem) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ResultModelMailmapListItem) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ResultModelMailmapListItem) SetHost(v string)`

SetHost sets Host field to given value.


### GetLastModified

`func (o *ResultModelMailmapListItem) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResultModelMailmapListItem) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResultModelMailmapListItem) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### GetMailmapId

`func (o *ResultModelMailmapListItem) GetMailmapId() int32`

GetMailmapId returns the MailmapId field if non-nil, zero value otherwise.

### GetMailmapIdOk

`func (o *ResultModelMailmapListItem) GetMailmapIdOk() (*int32, bool)`

GetMailmapIdOk returns a tuple with the MailmapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailmapId

`func (o *ResultModelMailmapListItem) SetMailmapId(v int32)`

SetMailmapId sets MailmapId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


