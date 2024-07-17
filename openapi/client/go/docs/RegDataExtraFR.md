# RegDataExtraFR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrantType** | **string** | Registrant type | 
**DateOfBirth** | Pointer to **string** | Date of Birth | [optional] 
**CountryOfBirth** | Pointer to **string** | Country of Birth (2 letter ISO country code) | [optional] 
**PlaceOfBirth** | Pointer to **string** | Place of Birth | [optional] 
**PostalCodeOfBirth** | Pointer to **string** | Postal Code of Birth | [optional] 
**RegistrantVatId** | Pointer to **string** | Value Added Tax (VAT) Number | [optional] 
**SirenSiret** | Pointer to **string** | SIREN or SIRET Code | [optional] 
**TrademarkNumber** | Pointer to **string** | Trademark Registration Number | [optional] 

## Methods

### NewRegDataExtraFR

`func NewRegDataExtraFR(registrantType string, ) *RegDataExtraFR`

NewRegDataExtraFR instantiates a new RegDataExtraFR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegDataExtraFRWithDefaults

`func NewRegDataExtraFRWithDefaults() *RegDataExtraFR`

NewRegDataExtraFRWithDefaults instantiates a new RegDataExtraFR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrantType

`func (o *RegDataExtraFR) GetRegistrantType() string`

GetRegistrantType returns the RegistrantType field if non-nil, zero value otherwise.

### GetRegistrantTypeOk

`func (o *RegDataExtraFR) GetRegistrantTypeOk() (*string, bool)`

GetRegistrantTypeOk returns a tuple with the RegistrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantType

`func (o *RegDataExtraFR) SetRegistrantType(v string)`

SetRegistrantType sets RegistrantType field to given value.


### GetDateOfBirth

`func (o *RegDataExtraFR) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *RegDataExtraFR) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *RegDataExtraFR) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *RegDataExtraFR) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetCountryOfBirth

`func (o *RegDataExtraFR) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *RegDataExtraFR) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *RegDataExtraFR) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.

### HasCountryOfBirth

`func (o *RegDataExtraFR) HasCountryOfBirth() bool`

HasCountryOfBirth returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *RegDataExtraFR) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *RegDataExtraFR) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *RegDataExtraFR) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *RegDataExtraFR) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### GetPostalCodeOfBirth

`func (o *RegDataExtraFR) GetPostalCodeOfBirth() string`

GetPostalCodeOfBirth returns the PostalCodeOfBirth field if non-nil, zero value otherwise.

### GetPostalCodeOfBirthOk

`func (o *RegDataExtraFR) GetPostalCodeOfBirthOk() (*string, bool)`

GetPostalCodeOfBirthOk returns a tuple with the PostalCodeOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCodeOfBirth

`func (o *RegDataExtraFR) SetPostalCodeOfBirth(v string)`

SetPostalCodeOfBirth sets PostalCodeOfBirth field to given value.

### HasPostalCodeOfBirth

`func (o *RegDataExtraFR) HasPostalCodeOfBirth() bool`

HasPostalCodeOfBirth returns a boolean if a field has been set.

### GetRegistrantVatId

`func (o *RegDataExtraFR) GetRegistrantVatId() string`

GetRegistrantVatId returns the RegistrantVatId field if non-nil, zero value otherwise.

### GetRegistrantVatIdOk

`func (o *RegDataExtraFR) GetRegistrantVatIdOk() (*string, bool)`

GetRegistrantVatIdOk returns a tuple with the RegistrantVatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantVatId

`func (o *RegDataExtraFR) SetRegistrantVatId(v string)`

SetRegistrantVatId sets RegistrantVatId field to given value.

### HasRegistrantVatId

`func (o *RegDataExtraFR) HasRegistrantVatId() bool`

HasRegistrantVatId returns a boolean if a field has been set.

### GetSirenSiret

`func (o *RegDataExtraFR) GetSirenSiret() string`

GetSirenSiret returns the SirenSiret field if non-nil, zero value otherwise.

### GetSirenSiretOk

`func (o *RegDataExtraFR) GetSirenSiretOk() (*string, bool)`

GetSirenSiretOk returns a tuple with the SirenSiret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirenSiret

`func (o *RegDataExtraFR) SetSirenSiret(v string)`

SetSirenSiret sets SirenSiret field to given value.

### HasSirenSiret

`func (o *RegDataExtraFR) HasSirenSiret() bool`

HasSirenSiret returns a boolean if a field has been set.

### GetTrademarkNumber

`func (o *RegDataExtraFR) GetTrademarkNumber() string`

GetTrademarkNumber returns the TrademarkNumber field if non-nil, zero value otherwise.

### GetTrademarkNumberOk

`func (o *RegDataExtraFR) GetTrademarkNumberOk() (*string, bool)`

GetTrademarkNumberOk returns a tuple with the TrademarkNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrademarkNumber

`func (o *RegDataExtraFR) SetTrademarkNumber(v string)`

SetTrademarkNumber sets TrademarkNumber field to given value.

### HasTrademarkNumber

`func (o *RegDataExtraFR) HasTrademarkNumber() bool`

HasTrademarkNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


