# RegDataExtraSE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrantType** | **string** | Registrant type | 
**IdCardNumber** | **string** | ID Card Number | 
**RegistrantVatId** | Pointer to **string** | Value Added Tax (VAT) number | [optional] 
**RegistrationNumber** | **string** | Business number | 

## Methods

### NewRegDataExtraSE

`func NewRegDataExtraSE(registrantType string, idCardNumber string, registrationNumber string, ) *RegDataExtraSE`

NewRegDataExtraSE instantiates a new RegDataExtraSE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegDataExtraSEWithDefaults

`func NewRegDataExtraSEWithDefaults() *RegDataExtraSE`

NewRegDataExtraSEWithDefaults instantiates a new RegDataExtraSE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrantType

`func (o *RegDataExtraSE) GetRegistrantType() string`

GetRegistrantType returns the RegistrantType field if non-nil, zero value otherwise.

### GetRegistrantTypeOk

`func (o *RegDataExtraSE) GetRegistrantTypeOk() (*string, bool)`

GetRegistrantTypeOk returns a tuple with the RegistrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantType

`func (o *RegDataExtraSE) SetRegistrantType(v string)`

SetRegistrantType sets RegistrantType field to given value.


### GetIdCardNumber

`func (o *RegDataExtraSE) GetIdCardNumber() string`

GetIdCardNumber returns the IdCardNumber field if non-nil, zero value otherwise.

### GetIdCardNumberOk

`func (o *RegDataExtraSE) GetIdCardNumberOk() (*string, bool)`

GetIdCardNumberOk returns a tuple with the IdCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCardNumber

`func (o *RegDataExtraSE) SetIdCardNumber(v string)`

SetIdCardNumber sets IdCardNumber field to given value.


### GetRegistrantVatId

`func (o *RegDataExtraSE) GetRegistrantVatId() string`

GetRegistrantVatId returns the RegistrantVatId field if non-nil, zero value otherwise.

### GetRegistrantVatIdOk

`func (o *RegDataExtraSE) GetRegistrantVatIdOk() (*string, bool)`

GetRegistrantVatIdOk returns a tuple with the RegistrantVatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantVatId

`func (o *RegDataExtraSE) SetRegistrantVatId(v string)`

SetRegistrantVatId sets RegistrantVatId field to given value.

### HasRegistrantVatId

`func (o *RegDataExtraSE) HasRegistrantVatId() bool`

HasRegistrantVatId returns a boolean if a field has been set.

### GetRegistrationNumber

`func (o *RegDataExtraSE) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *RegDataExtraSE) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *RegDataExtraSE) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


