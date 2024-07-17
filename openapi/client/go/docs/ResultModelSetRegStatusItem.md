# ResultModelSetRegStatusItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain to change the regstatus on | 
**Reglock** | Pointer to **bool** | (OPTIONAL) The reglock setting provided by the user | [optional] 
**Renewal** | Pointer to **string** | (OPTIONAL) The renew setting provided by the caller | [optional] 
**ChangeLock** | Pointer to **bool** | (OPTIONAL) Set to true if a lock change was attempted | [optional] 
**ChangeLockFailed** | Pointer to **bool** | (OPTIONAL) Set to true if a lock change was attempted but failed | [optional] 
**ChangeRenew** | Pointer to **bool** | (OPTIONAL) Set to true if a renewal action change was attempted | [optional] 
**ChangeRenewFailed** | Pointer to **bool** | (OPTIONAL) Set to true if a renewal action change was attempted but failed | [optional] 
**ChangeLetExpire** | Pointer to **bool** | (OPTIONAL) Set to true if a let_expire change was attempted | [optional] 
**ChangeLetExpireFailed** | Pointer to **bool** | (OPTIONAL) Set to true if a let_expire change was attempted but failed | [optional] 
**ConnectFailed** | Pointer to **bool** | (OPTIONAL) Attempts to connect to the registry for this domain failed | [optional] 
**InvalidRenewalState** | Pointer to **bool** | (OPTIONAL) Based on the domain provided the renewal action change was invalid. For example trying to set a .CA domain to expire | [optional] 

## Methods

### NewResultModelSetRegStatusItem

`func NewResultModelSetRegStatusItem(domain string, ) *ResultModelSetRegStatusItem`

NewResultModelSetRegStatusItem instantiates a new ResultModelSetRegStatusItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelSetRegStatusItemWithDefaults

`func NewResultModelSetRegStatusItemWithDefaults() *ResultModelSetRegStatusItem`

NewResultModelSetRegStatusItemWithDefaults instantiates a new ResultModelSetRegStatusItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ResultModelSetRegStatusItem) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ResultModelSetRegStatusItem) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ResultModelSetRegStatusItem) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetReglock

`func (o *ResultModelSetRegStatusItem) GetReglock() bool`

GetReglock returns the Reglock field if non-nil, zero value otherwise.

### GetReglockOk

`func (o *ResultModelSetRegStatusItem) GetReglockOk() (*bool, bool)`

GetReglockOk returns a tuple with the Reglock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReglock

`func (o *ResultModelSetRegStatusItem) SetReglock(v bool)`

SetReglock sets Reglock field to given value.

### HasReglock

`func (o *ResultModelSetRegStatusItem) HasReglock() bool`

HasReglock returns a boolean if a field has been set.

### GetRenewal

`func (o *ResultModelSetRegStatusItem) GetRenewal() string`

GetRenewal returns the Renewal field if non-nil, zero value otherwise.

### GetRenewalOk

`func (o *ResultModelSetRegStatusItem) GetRenewalOk() (*string, bool)`

GetRenewalOk returns a tuple with the Renewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewal

`func (o *ResultModelSetRegStatusItem) SetRenewal(v string)`

SetRenewal sets Renewal field to given value.

### HasRenewal

`func (o *ResultModelSetRegStatusItem) HasRenewal() bool`

HasRenewal returns a boolean if a field has been set.

### GetChangeLock

`func (o *ResultModelSetRegStatusItem) GetChangeLock() bool`

GetChangeLock returns the ChangeLock field if non-nil, zero value otherwise.

### GetChangeLockOk

`func (o *ResultModelSetRegStatusItem) GetChangeLockOk() (*bool, bool)`

GetChangeLockOk returns a tuple with the ChangeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLock

`func (o *ResultModelSetRegStatusItem) SetChangeLock(v bool)`

SetChangeLock sets ChangeLock field to given value.

### HasChangeLock

`func (o *ResultModelSetRegStatusItem) HasChangeLock() bool`

HasChangeLock returns a boolean if a field has been set.

### GetChangeLockFailed

`func (o *ResultModelSetRegStatusItem) GetChangeLockFailed() bool`

GetChangeLockFailed returns the ChangeLockFailed field if non-nil, zero value otherwise.

### GetChangeLockFailedOk

`func (o *ResultModelSetRegStatusItem) GetChangeLockFailedOk() (*bool, bool)`

GetChangeLockFailedOk returns a tuple with the ChangeLockFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLockFailed

`func (o *ResultModelSetRegStatusItem) SetChangeLockFailed(v bool)`

SetChangeLockFailed sets ChangeLockFailed field to given value.

### HasChangeLockFailed

`func (o *ResultModelSetRegStatusItem) HasChangeLockFailed() bool`

HasChangeLockFailed returns a boolean if a field has been set.

### GetChangeRenew

`func (o *ResultModelSetRegStatusItem) GetChangeRenew() bool`

GetChangeRenew returns the ChangeRenew field if non-nil, zero value otherwise.

### GetChangeRenewOk

`func (o *ResultModelSetRegStatusItem) GetChangeRenewOk() (*bool, bool)`

GetChangeRenewOk returns a tuple with the ChangeRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRenew

`func (o *ResultModelSetRegStatusItem) SetChangeRenew(v bool)`

SetChangeRenew sets ChangeRenew field to given value.

### HasChangeRenew

`func (o *ResultModelSetRegStatusItem) HasChangeRenew() bool`

HasChangeRenew returns a boolean if a field has been set.

### GetChangeRenewFailed

`func (o *ResultModelSetRegStatusItem) GetChangeRenewFailed() bool`

GetChangeRenewFailed returns the ChangeRenewFailed field if non-nil, zero value otherwise.

### GetChangeRenewFailedOk

`func (o *ResultModelSetRegStatusItem) GetChangeRenewFailedOk() (*bool, bool)`

GetChangeRenewFailedOk returns a tuple with the ChangeRenewFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRenewFailed

`func (o *ResultModelSetRegStatusItem) SetChangeRenewFailed(v bool)`

SetChangeRenewFailed sets ChangeRenewFailed field to given value.

### HasChangeRenewFailed

`func (o *ResultModelSetRegStatusItem) HasChangeRenewFailed() bool`

HasChangeRenewFailed returns a boolean if a field has been set.

### GetChangeLetExpire

`func (o *ResultModelSetRegStatusItem) GetChangeLetExpire() bool`

GetChangeLetExpire returns the ChangeLetExpire field if non-nil, zero value otherwise.

### GetChangeLetExpireOk

`func (o *ResultModelSetRegStatusItem) GetChangeLetExpireOk() (*bool, bool)`

GetChangeLetExpireOk returns a tuple with the ChangeLetExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLetExpire

`func (o *ResultModelSetRegStatusItem) SetChangeLetExpire(v bool)`

SetChangeLetExpire sets ChangeLetExpire field to given value.

### HasChangeLetExpire

`func (o *ResultModelSetRegStatusItem) HasChangeLetExpire() bool`

HasChangeLetExpire returns a boolean if a field has been set.

### GetChangeLetExpireFailed

`func (o *ResultModelSetRegStatusItem) GetChangeLetExpireFailed() bool`

GetChangeLetExpireFailed returns the ChangeLetExpireFailed field if non-nil, zero value otherwise.

### GetChangeLetExpireFailedOk

`func (o *ResultModelSetRegStatusItem) GetChangeLetExpireFailedOk() (*bool, bool)`

GetChangeLetExpireFailedOk returns a tuple with the ChangeLetExpireFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLetExpireFailed

`func (o *ResultModelSetRegStatusItem) SetChangeLetExpireFailed(v bool)`

SetChangeLetExpireFailed sets ChangeLetExpireFailed field to given value.

### HasChangeLetExpireFailed

`func (o *ResultModelSetRegStatusItem) HasChangeLetExpireFailed() bool`

HasChangeLetExpireFailed returns a boolean if a field has been set.

### GetConnectFailed

`func (o *ResultModelSetRegStatusItem) GetConnectFailed() bool`

GetConnectFailed returns the ConnectFailed field if non-nil, zero value otherwise.

### GetConnectFailedOk

`func (o *ResultModelSetRegStatusItem) GetConnectFailedOk() (*bool, bool)`

GetConnectFailedOk returns a tuple with the ConnectFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectFailed

`func (o *ResultModelSetRegStatusItem) SetConnectFailed(v bool)`

SetConnectFailed sets ConnectFailed field to given value.

### HasConnectFailed

`func (o *ResultModelSetRegStatusItem) HasConnectFailed() bool`

HasConnectFailed returns a boolean if a field has been set.

### GetInvalidRenewalState

`func (o *ResultModelSetRegStatusItem) GetInvalidRenewalState() bool`

GetInvalidRenewalState returns the InvalidRenewalState field if non-nil, zero value otherwise.

### GetInvalidRenewalStateOk

`func (o *ResultModelSetRegStatusItem) GetInvalidRenewalStateOk() (*bool, bool)`

GetInvalidRenewalStateOk returns a tuple with the InvalidRenewalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidRenewalState

`func (o *ResultModelSetRegStatusItem) SetInvalidRenewalState(v bool)`

SetInvalidRenewalState sets InvalidRenewalState field to given value.

### HasInvalidRenewalState

`func (o *ResultModelSetRegStatusItem) HasInvalidRenewalState() bool`

HasInvalidRenewalState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


