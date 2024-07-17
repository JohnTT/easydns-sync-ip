# ResultModelRegStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reglock** | **bool** | Is reglock for domain set to on or off | 
**Renewal** | **string** | Indicates the current renewal action for the domain | 
**AutoRenew** | **bool** | Does the domain have automated renewal payments set up for this domain | 
**AutoRenewCardId** | Pointer to **string** | If a domain is set to auto_renew this field will contain the stripe card id that will be used for this domain | [optional] 
**LetExpire** | **bool** | Is the domain set to automatically expire at HRS/SRS | 
**LetExpireFailed** | Pointer to **bool** | This field is only present when attempts to get the let_expire value at HRS/SRS failed. Most often caused by invalid authentication info on the system for the domain | [optional] 
**Expiry** | **string** | The expiry date for the domain on the system. Either the expiry date for the domain or the date service is next due | 
**LocalRegistrar** | **bool** | Indicates if the domain is registered by us or the registrar is someone else | 
**SupportsReglock** | **bool** | Indicates if the TLD for this domain supports reglocking | 

## Methods

### NewResultModelRegStatus

`func NewResultModelRegStatus(reglock bool, renewal string, autoRenew bool, letExpire bool, expiry string, localRegistrar bool, supportsReglock bool, ) *ResultModelRegStatus`

NewResultModelRegStatus instantiates a new ResultModelRegStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelRegStatusWithDefaults

`func NewResultModelRegStatusWithDefaults() *ResultModelRegStatus`

NewResultModelRegStatusWithDefaults instantiates a new ResultModelRegStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReglock

`func (o *ResultModelRegStatus) GetReglock() bool`

GetReglock returns the Reglock field if non-nil, zero value otherwise.

### GetReglockOk

`func (o *ResultModelRegStatus) GetReglockOk() (*bool, bool)`

GetReglockOk returns a tuple with the Reglock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReglock

`func (o *ResultModelRegStatus) SetReglock(v bool)`

SetReglock sets Reglock field to given value.


### GetRenewal

`func (o *ResultModelRegStatus) GetRenewal() string`

GetRenewal returns the Renewal field if non-nil, zero value otherwise.

### GetRenewalOk

`func (o *ResultModelRegStatus) GetRenewalOk() (*string, bool)`

GetRenewalOk returns a tuple with the Renewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewal

`func (o *ResultModelRegStatus) SetRenewal(v string)`

SetRenewal sets Renewal field to given value.


### GetAutoRenew

`func (o *ResultModelRegStatus) GetAutoRenew() bool`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *ResultModelRegStatus) GetAutoRenewOk() (*bool, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *ResultModelRegStatus) SetAutoRenew(v bool)`

SetAutoRenew sets AutoRenew field to given value.


### GetAutoRenewCardId

`func (o *ResultModelRegStatus) GetAutoRenewCardId() string`

GetAutoRenewCardId returns the AutoRenewCardId field if non-nil, zero value otherwise.

### GetAutoRenewCardIdOk

`func (o *ResultModelRegStatus) GetAutoRenewCardIdOk() (*string, bool)`

GetAutoRenewCardIdOk returns a tuple with the AutoRenewCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewCardId

`func (o *ResultModelRegStatus) SetAutoRenewCardId(v string)`

SetAutoRenewCardId sets AutoRenewCardId field to given value.

### HasAutoRenewCardId

`func (o *ResultModelRegStatus) HasAutoRenewCardId() bool`

HasAutoRenewCardId returns a boolean if a field has been set.

### GetLetExpire

`func (o *ResultModelRegStatus) GetLetExpire() bool`

GetLetExpire returns the LetExpire field if non-nil, zero value otherwise.

### GetLetExpireOk

`func (o *ResultModelRegStatus) GetLetExpireOk() (*bool, bool)`

GetLetExpireOk returns a tuple with the LetExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetExpire

`func (o *ResultModelRegStatus) SetLetExpire(v bool)`

SetLetExpire sets LetExpire field to given value.


### GetLetExpireFailed

`func (o *ResultModelRegStatus) GetLetExpireFailed() bool`

GetLetExpireFailed returns the LetExpireFailed field if non-nil, zero value otherwise.

### GetLetExpireFailedOk

`func (o *ResultModelRegStatus) GetLetExpireFailedOk() (*bool, bool)`

GetLetExpireFailedOk returns a tuple with the LetExpireFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetExpireFailed

`func (o *ResultModelRegStatus) SetLetExpireFailed(v bool)`

SetLetExpireFailed sets LetExpireFailed field to given value.

### HasLetExpireFailed

`func (o *ResultModelRegStatus) HasLetExpireFailed() bool`

HasLetExpireFailed returns a boolean if a field has been set.

### GetExpiry

`func (o *ResultModelRegStatus) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ResultModelRegStatus) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ResultModelRegStatus) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.


### GetLocalRegistrar

`func (o *ResultModelRegStatus) GetLocalRegistrar() bool`

GetLocalRegistrar returns the LocalRegistrar field if non-nil, zero value otherwise.

### GetLocalRegistrarOk

`func (o *ResultModelRegStatus) GetLocalRegistrarOk() (*bool, bool)`

GetLocalRegistrarOk returns a tuple with the LocalRegistrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRegistrar

`func (o *ResultModelRegStatus) SetLocalRegistrar(v bool)`

SetLocalRegistrar sets LocalRegistrar field to given value.


### GetSupportsReglock

`func (o *ResultModelRegStatus) GetSupportsReglock() bool`

GetSupportsReglock returns the SupportsReglock field if non-nil, zero value otherwise.

### GetSupportsReglockOk

`func (o *ResultModelRegStatus) GetSupportsReglockOk() (*bool, bool)`

GetSupportsReglockOk returns a tuple with the SupportsReglock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsReglock

`func (o *ResultModelRegStatus) SetSupportsReglock(v bool)`

SetSupportsReglock sets SupportsReglock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


