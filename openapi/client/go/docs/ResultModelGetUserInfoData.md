# ResultModelGetUserInfoData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** | The username of the returned user data | [optional] 
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

## Methods

### NewResultModelGetUserInfoData

`func NewResultModelGetUserInfoData() *ResultModelGetUserInfoData`

NewResultModelGetUserInfoData instantiates a new ResultModelGetUserInfoData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelGetUserInfoDataWithDefaults

`func NewResultModelGetUserInfoDataWithDefaults() *ResultModelGetUserInfoData`

NewResultModelGetUserInfoDataWithDefaults instantiates a new ResultModelGetUserInfoData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ResultModelGetUserInfoData) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ResultModelGetUserInfoData) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ResultModelGetUserInfoData) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ResultModelGetUserInfoData) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetFirstName

`func (o *ResultModelGetUserInfoData) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResultModelGetUserInfoData) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResultModelGetUserInfoData) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResultModelGetUserInfoData) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ResultModelGetUserInfoData) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResultModelGetUserInfoData) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResultModelGetUserInfoData) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResultModelGetUserInfoData) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetOrgName

`func (o *ResultModelGetUserInfoData) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ResultModelGetUserInfoData) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ResultModelGetUserInfoData) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *ResultModelGetUserInfoData) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetAddress1

`func (o *ResultModelGetUserInfoData) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ResultModelGetUserInfoData) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ResultModelGetUserInfoData) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *ResultModelGetUserInfoData) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *ResultModelGetUserInfoData) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ResultModelGetUserInfoData) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ResultModelGetUserInfoData) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ResultModelGetUserInfoData) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetAddress3

`func (o *ResultModelGetUserInfoData) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *ResultModelGetUserInfoData) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *ResultModelGetUserInfoData) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *ResultModelGetUserInfoData) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### GetCity

`func (o *ResultModelGetUserInfoData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ResultModelGetUserInfoData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ResultModelGetUserInfoData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ResultModelGetUserInfoData) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *ResultModelGetUserInfoData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResultModelGetUserInfoData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResultModelGetUserInfoData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResultModelGetUserInfoData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *ResultModelGetUserInfoData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ResultModelGetUserInfoData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ResultModelGetUserInfoData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ResultModelGetUserInfoData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPostalCode

`func (o *ResultModelGetUserInfoData) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ResultModelGetUserInfoData) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ResultModelGetUserInfoData) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ResultModelGetUserInfoData) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCurrency

`func (o *ResultModelGetUserInfoData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ResultModelGetUserInfoData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ResultModelGetUserInfoData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ResultModelGetUserInfoData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPhone

`func (o *ResultModelGetUserInfoData) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ResultModelGetUserInfoData) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ResultModelGetUserInfoData) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ResultModelGetUserInfoData) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCellphone

`func (o *ResultModelGetUserInfoData) GetCellphone() string`

GetCellphone returns the Cellphone field if non-nil, zero value otherwise.

### GetCellphoneOk

`func (o *ResultModelGetUserInfoData) GetCellphoneOk() (*string, bool)`

GetCellphoneOk returns a tuple with the Cellphone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellphone

`func (o *ResultModelGetUserInfoData) SetCellphone(v string)`

SetCellphone sets Cellphone field to given value.

### HasCellphone

`func (o *ResultModelGetUserInfoData) HasCellphone() bool`

HasCellphone returns a boolean if a field has been set.

### GetFax

`func (o *ResultModelGetUserInfoData) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *ResultModelGetUserInfoData) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *ResultModelGetUserInfoData) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *ResultModelGetUserInfoData) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetEmail

`func (o *ResultModelGetUserInfoData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResultModelGetUserInfoData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResultModelGetUserInfoData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ResultModelGetUserInfoData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmail2

`func (o *ResultModelGetUserInfoData) GetEmail2() string`

GetEmail2 returns the Email2 field if non-nil, zero value otherwise.

### GetEmail2Ok

`func (o *ResultModelGetUserInfoData) GetEmail2Ok() (*string, bool)`

GetEmail2Ok returns a tuple with the Email2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail2

`func (o *ResultModelGetUserInfoData) SetEmail2(v string)`

SetEmail2 sets Email2 field to given value.

### HasEmail2

`func (o *ResultModelGetUserInfoData) HasEmail2() bool`

HasEmail2 returns a boolean if a field has been set.

### GetNoticesEmail

`func (o *ResultModelGetUserInfoData) GetNoticesEmail() string`

GetNoticesEmail returns the NoticesEmail field if non-nil, zero value otherwise.

### GetNoticesEmailOk

`func (o *ResultModelGetUserInfoData) GetNoticesEmailOk() (*string, bool)`

GetNoticesEmailOk returns a tuple with the NoticesEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticesEmail

`func (o *ResultModelGetUserInfoData) SetNoticesEmail(v string)`

SetNoticesEmail sets NoticesEmail field to given value.

### HasNoticesEmail

`func (o *ResultModelGetUserInfoData) HasNoticesEmail() bool`

HasNoticesEmail returns a boolean if a field has been set.

### GetPublicEmail

`func (o *ResultModelGetUserInfoData) GetPublicEmail() string`

GetPublicEmail returns the PublicEmail field if non-nil, zero value otherwise.

### GetPublicEmailOk

`func (o *ResultModelGetUserInfoData) GetPublicEmailOk() (*string, bool)`

GetPublicEmailOk returns a tuple with the PublicEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEmail

`func (o *ResultModelGetUserInfoData) SetPublicEmail(v string)`

SetPublicEmail sets PublicEmail field to given value.

### HasPublicEmail

`func (o *ResultModelGetUserInfoData) HasPublicEmail() bool`

HasPublicEmail returns a boolean if a field has been set.

### GetAlertsEmail

`func (o *ResultModelGetUserInfoData) GetAlertsEmail() string`

GetAlertsEmail returns the AlertsEmail field if non-nil, zero value otherwise.

### GetAlertsEmailOk

`func (o *ResultModelGetUserInfoData) GetAlertsEmailOk() (*string, bool)`

GetAlertsEmailOk returns a tuple with the AlertsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsEmail

`func (o *ResultModelGetUserInfoData) SetAlertsEmail(v string)`

SetAlertsEmail sets AlertsEmail field to given value.

### HasAlertsEmail

`func (o *ResultModelGetUserInfoData) HasAlertsEmail() bool`

HasAlertsEmail returns a boolean if a field has been set.

### GetUrl

`func (o *ResultModelGetUserInfoData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResultModelGetUserInfoData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResultModelGetUserInfoData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ResultModelGetUserInfoData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOptOut

`func (o *ResultModelGetUserInfoData) GetOptOut() int32`

GetOptOut returns the OptOut field if non-nil, zero value otherwise.

### GetOptOutOk

`func (o *ResultModelGetUserInfoData) GetOptOutOk() (*int32, bool)`

GetOptOutOk returns a tuple with the OptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptOut

`func (o *ResultModelGetUserInfoData) SetOptOut(v int32)`

SetOptOut sets OptOut field to given value.

### HasOptOut

`func (o *ResultModelGetUserInfoData) HasOptOut() bool`

HasOptOut returns a boolean if a field has been set.

### GetBeta

`func (o *ResultModelGetUserInfoData) GetBeta() int32`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *ResultModelGetUserInfoData) GetBetaOk() (*int32, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *ResultModelGetUserInfoData) SetBeta(v int32)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *ResultModelGetUserInfoData) HasBeta() bool`

HasBeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


