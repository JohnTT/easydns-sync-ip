# RequestModelCreateUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First name of new user | [optional] 
**LastName** | Pointer to **string** | Last name of new user | [optional] 
**OrgName** | Pointer to **string** | Organization of new user account | [optional] 
**Address1** | Pointer to **string** | Address of new user | [optional] 
**Address2** | Pointer to **string** | Address of new user | [optional] 
**City** | Pointer to **string** | City of the new user | [optional] 
**State** | **string** | New users state/province (2 letter code) | 
**Country** | **string** | New users country (2 letter ISO country code) | 
**PostalCode** | Pointer to **string** | The postal/zip code of the new user account | [optional] 
**Phone** | Pointer to **string** | Primary contact number for new user account | [optional] 
**Email** | **string** | Primary contact email for new user account | 
**Currency** | **string** | The default currency to use for the new user account. This is required to allow creation of domains for a new user without first logging in to control panel (account will still need to be funded first) | 
**Password** | Pointer to **string** | Use the provided password for the new user account. If not provided a random password will be created for user and returned as part of result | [optional] 
**Vcode** | Pointer to **string** | Visitor code used during user creation | [optional] 
**RemoteCc** | Pointer to **string** | Country code of originator of request (2 letter ISO country code) | [optional] 
**RemoteIp** | Pointer to **string** | IP address of originator of request | [optional] 
**RemoteHost** | Pointer to **string** | Hostname of originator of request | [optional] 

## Methods

### NewRequestModelCreateUserData

`func NewRequestModelCreateUserData(state string, country string, email string, currency string, ) *RequestModelCreateUserData`

NewRequestModelCreateUserData instantiates a new RequestModelCreateUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestModelCreateUserDataWithDefaults

`func NewRequestModelCreateUserDataWithDefaults() *RequestModelCreateUserData`

NewRequestModelCreateUserDataWithDefaults instantiates a new RequestModelCreateUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *RequestModelCreateUserData) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RequestModelCreateUserData) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RequestModelCreateUserData) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *RequestModelCreateUserData) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *RequestModelCreateUserData) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RequestModelCreateUserData) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RequestModelCreateUserData) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *RequestModelCreateUserData) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetOrgName

`func (o *RequestModelCreateUserData) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *RequestModelCreateUserData) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *RequestModelCreateUserData) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *RequestModelCreateUserData) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetAddress1

`func (o *RequestModelCreateUserData) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *RequestModelCreateUserData) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *RequestModelCreateUserData) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *RequestModelCreateUserData) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *RequestModelCreateUserData) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *RequestModelCreateUserData) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *RequestModelCreateUserData) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *RequestModelCreateUserData) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *RequestModelCreateUserData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RequestModelCreateUserData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RequestModelCreateUserData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *RequestModelCreateUserData) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *RequestModelCreateUserData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RequestModelCreateUserData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RequestModelCreateUserData) SetState(v string)`

SetState sets State field to given value.


### GetCountry

`func (o *RequestModelCreateUserData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RequestModelCreateUserData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RequestModelCreateUserData) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetPostalCode

`func (o *RequestModelCreateUserData) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *RequestModelCreateUserData) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *RequestModelCreateUserData) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *RequestModelCreateUserData) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetPhone

`func (o *RequestModelCreateUserData) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *RequestModelCreateUserData) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *RequestModelCreateUserData) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *RequestModelCreateUserData) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *RequestModelCreateUserData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RequestModelCreateUserData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RequestModelCreateUserData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCurrency

`func (o *RequestModelCreateUserData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RequestModelCreateUserData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RequestModelCreateUserData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPassword

`func (o *RequestModelCreateUserData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RequestModelCreateUserData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RequestModelCreateUserData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RequestModelCreateUserData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetVcode

`func (o *RequestModelCreateUserData) GetVcode() string`

GetVcode returns the Vcode field if non-nil, zero value otherwise.

### GetVcodeOk

`func (o *RequestModelCreateUserData) GetVcodeOk() (*string, bool)`

GetVcodeOk returns a tuple with the Vcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcode

`func (o *RequestModelCreateUserData) SetVcode(v string)`

SetVcode sets Vcode field to given value.

### HasVcode

`func (o *RequestModelCreateUserData) HasVcode() bool`

HasVcode returns a boolean if a field has been set.

### GetRemoteCc

`func (o *RequestModelCreateUserData) GetRemoteCc() string`

GetRemoteCc returns the RemoteCc field if non-nil, zero value otherwise.

### GetRemoteCcOk

`func (o *RequestModelCreateUserData) GetRemoteCcOk() (*string, bool)`

GetRemoteCcOk returns a tuple with the RemoteCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCc

`func (o *RequestModelCreateUserData) SetRemoteCc(v string)`

SetRemoteCc sets RemoteCc field to given value.

### HasRemoteCc

`func (o *RequestModelCreateUserData) HasRemoteCc() bool`

HasRemoteCc returns a boolean if a field has been set.

### GetRemoteIp

`func (o *RequestModelCreateUserData) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *RequestModelCreateUserData) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *RequestModelCreateUserData) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *RequestModelCreateUserData) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteHost

`func (o *RequestModelCreateUserData) GetRemoteHost() string`

GetRemoteHost returns the RemoteHost field if non-nil, zero value otherwise.

### GetRemoteHostOk

`func (o *RequestModelCreateUserData) GetRemoteHostOk() (*string, bool)`

GetRemoteHostOk returns a tuple with the RemoteHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteHost

`func (o *RequestModelCreateUserData) SetRemoteHost(v string)`

SetRemoteHost sets RemoteHost field to given value.

### HasRemoteHost

`func (o *RequestModelCreateUserData) HasRemoteHost() bool`

HasRemoteHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


