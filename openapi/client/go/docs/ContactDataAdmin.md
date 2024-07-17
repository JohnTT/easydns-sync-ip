# ContactDataAdmin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**OrgName** | Pointer to **string** | Contact organization | [optional] 
**Address1** | **string** |  | 
**Address2** | Pointer to **string** |  | [optional] 
**City** | **string** |  | 
**State** | **string** | Province or state of contact | 
**Country** | **string** | 2 letter ISO country code | 
**PostalCode** | **string** | Postal code or zip code of contact | 
**Phone** | **string** |  | 
**Email** | **string** |  | 

## Methods

### NewContactDataAdmin

`func NewContactDataAdmin(firstName string, lastName string, address1 string, city string, state string, country string, postalCode string, phone string, email string, ) *ContactDataAdmin`

NewContactDataAdmin instantiates a new ContactDataAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactDataAdminWithDefaults

`func NewContactDataAdminWithDefaults() *ContactDataAdmin`

NewContactDataAdminWithDefaults instantiates a new ContactDataAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *ContactDataAdmin) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ContactDataAdmin) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ContactDataAdmin) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ContactDataAdmin) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ContactDataAdmin) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ContactDataAdmin) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetOrgName

`func (o *ContactDataAdmin) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ContactDataAdmin) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ContactDataAdmin) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *ContactDataAdmin) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetAddress1

`func (o *ContactDataAdmin) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ContactDataAdmin) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ContactDataAdmin) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *ContactDataAdmin) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ContactDataAdmin) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ContactDataAdmin) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ContactDataAdmin) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *ContactDataAdmin) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ContactDataAdmin) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ContactDataAdmin) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *ContactDataAdmin) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ContactDataAdmin) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ContactDataAdmin) SetState(v string)`

SetState sets State field to given value.


### GetCountry

`func (o *ContactDataAdmin) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ContactDataAdmin) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ContactDataAdmin) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetPostalCode

`func (o *ContactDataAdmin) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ContactDataAdmin) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ContactDataAdmin) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetPhone

`func (o *ContactDataAdmin) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ContactDataAdmin) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ContactDataAdmin) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetEmail

`func (o *ContactDataAdmin) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactDataAdmin) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactDataAdmin) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


