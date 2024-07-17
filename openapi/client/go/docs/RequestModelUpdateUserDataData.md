# RequestModelUpdateUserDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The first name of the user | [optional] 
**LastName** | Pointer to **string** | The last name of the user | [optional] 
**OrgName** | Pointer to **string** | The organization of the user | [optional] 
**Address1** | Pointer to **string** | The address of the user | [optional] 
**Address2** | Pointer to **string** |  | [optional] 
**Address3** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** | The city of the address provided | [optional] 
**State** | Pointer to **string** | The state/province code of the address | [optional] 
**Country** | Pointer to **string** | The 2-letter ISO country code of the address | [optional] 
**PostalCode** | Pointer to **string** | The zip/postal code of the address | [optional] 
**Currency** | Pointer to **string** | The default currency to use for the account | [optional] 
**Phone** | Pointer to **string** | Phone number that can be used to contact customer | [optional] 
**Cellphone** | Pointer to **string** | Cellular number that can be used to contact customer | [optional] 
**Fax** | Pointer to **string** | Fax number that can be used to contact customer | [optional] 
**Email** | Pointer to **string** | Main contact email for user | [optional] 
**Email2** | Pointer to **string** | Secondary contact email for user | [optional] 
**NoticesEmail** | Pointer to **string** | The email address(es) where generic notices should be sent | [optional] 
**PublicEmail** | Pointer to **string** | The email address(es) that can be made public for contacting the user | [optional] 
**AlertsEmail** | Pointer to **string** | The email address(es) where alert notice should be sent | [optional] 
**Url** | Pointer to **string** | The URL associated with a user (Not currently used anywhere) | [optional] 
**OptOut** | Pointer to **int32** | Has the user opted out of any non-essential communications | [optional] 
**Beta** | Pointer to **int32** | Does the user access to BETA level features | [optional] 
**Fraud** | Pointer to **string** | Is the account currently marked as fraud | [optional] [default to "N"]
**Vip** | Pointer to **int32** | Is the user currently marked as a VIP user | [optional] [default to 0]
**Password** | Pointer to **string** | Provides a new password for the user that will replace the users current password | [optional] 
**Esm** | Pointer to [**RequestModelUpdateUserDataDataEsm**](RequestModelUpdateUserDataDataEsm.md) |  | [optional] 

## Methods

### NewRequestModelUpdateUserDataData

`func NewRequestModelUpdateUserDataData() *RequestModelUpdateUserDataData`

NewRequestModelUpdateUserDataData instantiates a new RequestModelUpdateUserDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestModelUpdateUserDataDataWithDefaults

`func NewRequestModelUpdateUserDataDataWithDefaults() *RequestModelUpdateUserDataData`

NewRequestModelUpdateUserDataDataWithDefaults instantiates a new RequestModelUpdateUserDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *RequestModelUpdateUserDataData) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RequestModelUpdateUserDataData) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RequestModelUpdateUserDataData) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *RequestModelUpdateUserDataData) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *RequestModelUpdateUserDataData) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RequestModelUpdateUserDataData) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RequestModelUpdateUserDataData) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *RequestModelUpdateUserDataData) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetOrgName

`func (o *RequestModelUpdateUserDataData) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *RequestModelUpdateUserDataData) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *RequestModelUpdateUserDataData) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *RequestModelUpdateUserDataData) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetAddress1

`func (o *RequestModelUpdateUserDataData) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *RequestModelUpdateUserDataData) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *RequestModelUpdateUserDataData) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *RequestModelUpdateUserDataData) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *RequestModelUpdateUserDataData) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *RequestModelUpdateUserDataData) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *RequestModelUpdateUserDataData) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *RequestModelUpdateUserDataData) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetAddress3

`func (o *RequestModelUpdateUserDataData) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *RequestModelUpdateUserDataData) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *RequestModelUpdateUserDataData) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *RequestModelUpdateUserDataData) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### GetCity

`func (o *RequestModelUpdateUserDataData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RequestModelUpdateUserDataData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RequestModelUpdateUserDataData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *RequestModelUpdateUserDataData) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *RequestModelUpdateUserDataData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RequestModelUpdateUserDataData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RequestModelUpdateUserDataData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RequestModelUpdateUserDataData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *RequestModelUpdateUserDataData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RequestModelUpdateUserDataData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RequestModelUpdateUserDataData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RequestModelUpdateUserDataData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPostalCode

`func (o *RequestModelUpdateUserDataData) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *RequestModelUpdateUserDataData) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *RequestModelUpdateUserDataData) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *RequestModelUpdateUserDataData) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCurrency

`func (o *RequestModelUpdateUserDataData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RequestModelUpdateUserDataData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RequestModelUpdateUserDataData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *RequestModelUpdateUserDataData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPhone

`func (o *RequestModelUpdateUserDataData) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *RequestModelUpdateUserDataData) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *RequestModelUpdateUserDataData) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *RequestModelUpdateUserDataData) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCellphone

`func (o *RequestModelUpdateUserDataData) GetCellphone() string`

GetCellphone returns the Cellphone field if non-nil, zero value otherwise.

### GetCellphoneOk

`func (o *RequestModelUpdateUserDataData) GetCellphoneOk() (*string, bool)`

GetCellphoneOk returns a tuple with the Cellphone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellphone

`func (o *RequestModelUpdateUserDataData) SetCellphone(v string)`

SetCellphone sets Cellphone field to given value.

### HasCellphone

`func (o *RequestModelUpdateUserDataData) HasCellphone() bool`

HasCellphone returns a boolean if a field has been set.

### GetFax

`func (o *RequestModelUpdateUserDataData) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *RequestModelUpdateUserDataData) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *RequestModelUpdateUserDataData) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *RequestModelUpdateUserDataData) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetEmail

`func (o *RequestModelUpdateUserDataData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RequestModelUpdateUserDataData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RequestModelUpdateUserDataData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RequestModelUpdateUserDataData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmail2

`func (o *RequestModelUpdateUserDataData) GetEmail2() string`

GetEmail2 returns the Email2 field if non-nil, zero value otherwise.

### GetEmail2Ok

`func (o *RequestModelUpdateUserDataData) GetEmail2Ok() (*string, bool)`

GetEmail2Ok returns a tuple with the Email2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail2

`func (o *RequestModelUpdateUserDataData) SetEmail2(v string)`

SetEmail2 sets Email2 field to given value.

### HasEmail2

`func (o *RequestModelUpdateUserDataData) HasEmail2() bool`

HasEmail2 returns a boolean if a field has been set.

### GetNoticesEmail

`func (o *RequestModelUpdateUserDataData) GetNoticesEmail() string`

GetNoticesEmail returns the NoticesEmail field if non-nil, zero value otherwise.

### GetNoticesEmailOk

`func (o *RequestModelUpdateUserDataData) GetNoticesEmailOk() (*string, bool)`

GetNoticesEmailOk returns a tuple with the NoticesEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticesEmail

`func (o *RequestModelUpdateUserDataData) SetNoticesEmail(v string)`

SetNoticesEmail sets NoticesEmail field to given value.

### HasNoticesEmail

`func (o *RequestModelUpdateUserDataData) HasNoticesEmail() bool`

HasNoticesEmail returns a boolean if a field has been set.

### GetPublicEmail

`func (o *RequestModelUpdateUserDataData) GetPublicEmail() string`

GetPublicEmail returns the PublicEmail field if non-nil, zero value otherwise.

### GetPublicEmailOk

`func (o *RequestModelUpdateUserDataData) GetPublicEmailOk() (*string, bool)`

GetPublicEmailOk returns a tuple with the PublicEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEmail

`func (o *RequestModelUpdateUserDataData) SetPublicEmail(v string)`

SetPublicEmail sets PublicEmail field to given value.

### HasPublicEmail

`func (o *RequestModelUpdateUserDataData) HasPublicEmail() bool`

HasPublicEmail returns a boolean if a field has been set.

### GetAlertsEmail

`func (o *RequestModelUpdateUserDataData) GetAlertsEmail() string`

GetAlertsEmail returns the AlertsEmail field if non-nil, zero value otherwise.

### GetAlertsEmailOk

`func (o *RequestModelUpdateUserDataData) GetAlertsEmailOk() (*string, bool)`

GetAlertsEmailOk returns a tuple with the AlertsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsEmail

`func (o *RequestModelUpdateUserDataData) SetAlertsEmail(v string)`

SetAlertsEmail sets AlertsEmail field to given value.

### HasAlertsEmail

`func (o *RequestModelUpdateUserDataData) HasAlertsEmail() bool`

HasAlertsEmail returns a boolean if a field has been set.

### GetUrl

`func (o *RequestModelUpdateUserDataData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequestModelUpdateUserDataData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequestModelUpdateUserDataData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RequestModelUpdateUserDataData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOptOut

`func (o *RequestModelUpdateUserDataData) GetOptOut() int32`

GetOptOut returns the OptOut field if non-nil, zero value otherwise.

### GetOptOutOk

`func (o *RequestModelUpdateUserDataData) GetOptOutOk() (*int32, bool)`

GetOptOutOk returns a tuple with the OptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptOut

`func (o *RequestModelUpdateUserDataData) SetOptOut(v int32)`

SetOptOut sets OptOut field to given value.

### HasOptOut

`func (o *RequestModelUpdateUserDataData) HasOptOut() bool`

HasOptOut returns a boolean if a field has been set.

### GetBeta

`func (o *RequestModelUpdateUserDataData) GetBeta() int32`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *RequestModelUpdateUserDataData) GetBetaOk() (*int32, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *RequestModelUpdateUserDataData) SetBeta(v int32)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *RequestModelUpdateUserDataData) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### GetFraud

`func (o *RequestModelUpdateUserDataData) GetFraud() string`

GetFraud returns the Fraud field if non-nil, zero value otherwise.

### GetFraudOk

`func (o *RequestModelUpdateUserDataData) GetFraudOk() (*string, bool)`

GetFraudOk returns a tuple with the Fraud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraud

`func (o *RequestModelUpdateUserDataData) SetFraud(v string)`

SetFraud sets Fraud field to given value.

### HasFraud

`func (o *RequestModelUpdateUserDataData) HasFraud() bool`

HasFraud returns a boolean if a field has been set.

### GetVip

`func (o *RequestModelUpdateUserDataData) GetVip() int32`

GetVip returns the Vip field if non-nil, zero value otherwise.

### GetVipOk

`func (o *RequestModelUpdateUserDataData) GetVipOk() (*int32, bool)`

GetVipOk returns a tuple with the Vip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVip

`func (o *RequestModelUpdateUserDataData) SetVip(v int32)`

SetVip sets Vip field to given value.

### HasVip

`func (o *RequestModelUpdateUserDataData) HasVip() bool`

HasVip returns a boolean if a field has been set.

### GetPassword

`func (o *RequestModelUpdateUserDataData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RequestModelUpdateUserDataData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RequestModelUpdateUserDataData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RequestModelUpdateUserDataData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEsm

`func (o *RequestModelUpdateUserDataData) GetEsm() RequestModelUpdateUserDataDataEsm`

GetEsm returns the Esm field if non-nil, zero value otherwise.

### GetEsmOk

`func (o *RequestModelUpdateUserDataData) GetEsmOk() (*RequestModelUpdateUserDataDataEsm, bool)`

GetEsmOk returns a tuple with the Esm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsm

`func (o *RequestModelUpdateUserDataData) SetEsm(v RequestModelUpdateUserDataDataEsm)`

SetEsm sets Esm field to given value.

### HasEsm

`func (o *RequestModelUpdateUserDataData) HasEsm() bool`

HasEsm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


