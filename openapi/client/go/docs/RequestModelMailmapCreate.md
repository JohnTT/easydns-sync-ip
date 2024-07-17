# RequestModelMailmapCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | Alias part of mailmap address (for example: &#39;test&#39; from test@example.com) | 
**Host** | **string** | The hostname (without domain name) the mailmap should be created under | 
**Destination** | **string** | The email address to forward the mailmaps emails to | 
**Active** | **int32** | Setting this to 0 will disable the mailmap from forwarding mail | 

## Methods

### NewRequestModelMailmapCreate

`func NewRequestModelMailmapCreate(alias string, host string, destination string, active int32, ) *RequestModelMailmapCreate`

NewRequestModelMailmapCreate instantiates a new RequestModelMailmapCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestModelMailmapCreateWithDefaults

`func NewRequestModelMailmapCreateWithDefaults() *RequestModelMailmapCreate`

NewRequestModelMailmapCreateWithDefaults instantiates a new RequestModelMailmapCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *RequestModelMailmapCreate) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *RequestModelMailmapCreate) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *RequestModelMailmapCreate) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetHost

`func (o *RequestModelMailmapCreate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RequestModelMailmapCreate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RequestModelMailmapCreate) SetHost(v string)`

SetHost sets Host field to given value.


### GetDestination

`func (o *RequestModelMailmapCreate) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RequestModelMailmapCreate) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RequestModelMailmapCreate) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetActive

`func (o *RequestModelMailmapCreate) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RequestModelMailmapCreate) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RequestModelMailmapCreate) SetActive(v int32)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


