# ResultModelCreateUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **string** |  | 
**Currency** | **string** | The default currency used for the new user account | 
**State** | **string** | New users state/province (2 letter code) | 
**Country** | **string** | New users country (2 letter ISO country code) | 
**Email** | **string** | Primary contact email for new user account | 
**PostalCode** | Pointer to **string** | The postal/zip code of the new user account | [optional] 
**Phone** | Pointer to **string** | Primary contact number for new user account | [optional] 
**FirstName** | **string** | First name of new user | 
**LastName** | **string** | Last name of new user | 
**Token** | **string** | API token value for the new user account. Required in combination with key to authenticate with API as new user. | 
**Key** | **string** | API key value for the new user account. Required in combination with token to authenticate with API as new user. This value should be kept secret. | 

## Methods

### NewResultModelCreateUserData

`func NewResultModelCreateUserData(user string, currency string, state string, country string, email string, firstName string, lastName string, token string, key string, ) *ResultModelCreateUserData`

NewResultModelCreateUserData instantiates a new ResultModelCreateUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultModelCreateUserDataWithDefaults

`func NewResultModelCreateUserDataWithDefaults() *ResultModelCreateUserData`

NewResultModelCreateUserDataWithDefaults instantiates a new ResultModelCreateUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ResultModelCreateUserData) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ResultModelCreateUserData) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ResultModelCreateUserData) SetUser(v string)`

SetUser sets User field to given value.


### GetCurrency

`func (o *ResultModelCreateUserData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ResultModelCreateUserData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ResultModelCreateUserData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetState

`func (o *ResultModelCreateUserData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResultModelCreateUserData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResultModelCreateUserData) SetState(v string)`

SetState sets State field to given value.


### GetCountry

`func (o *ResultModelCreateUserData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ResultModelCreateUserData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ResultModelCreateUserData) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *ResultModelCreateUserData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResultModelCreateUserData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResultModelCreateUserData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPostalCode

`func (o *ResultModelCreateUserData) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ResultModelCreateUserData) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ResultModelCreateUserData) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ResultModelCreateUserData) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetPhone

`func (o *ResultModelCreateUserData) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ResultModelCreateUserData) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ResultModelCreateUserData) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ResultModelCreateUserData) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetFirstName

`func (o *ResultModelCreateUserData) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResultModelCreateUserData) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResultModelCreateUserData) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ResultModelCreateUserData) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResultModelCreateUserData) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResultModelCreateUserData) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetToken

`func (o *ResultModelCreateUserData) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResultModelCreateUserData) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ResultModelCreateUserData) SetToken(v string)`

SetToken sets Token field to given value.


### GetKey

`func (o *ResultModelCreateUserData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ResultModelCreateUserData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ResultModelCreateUserData) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


