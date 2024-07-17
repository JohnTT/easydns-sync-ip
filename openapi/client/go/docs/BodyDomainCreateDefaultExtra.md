# BodyDomainCreateDefaultExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrRegisterNumber** | **string** |  | 
**RegistrantType** | **string** | Registrant type | 
**IntendedUse** | **string** |  | 
**DateOfBirth** | Pointer to **string** | Date of Birth | [optional] 
**CountryOfBirth** | Pointer to **string** | Country of Birth (2 letter ISO country code) | [optional] 
**PlaceOfBirth** | Pointer to **string** | Place of Birth | [optional] 
**PostalCodeOfBirth** | Pointer to **string** | Postal Code of Birth | [optional] 
**RegistrantVatId** | Pointer to **string** | Value Added Tax (VAT) number | [optional] 
**SirenSiret** | Pointer to **string** | SIREN or SIRET Code | [optional] 
**TrademarkNumber** | Pointer to **string** | Trademark Registration Number | [optional] 
**EntityType** | **int32** | Entity type | 
**NationalityCode** | **string** | Nationality code (2 letter ISO country code) | 
**RegCode** | **string** | VAT or Codice Fiscale or n.a. | 
**QliAccreditationId** | **string** | Qualified Lawyer&#39;s Accreditation ID | 
**QliAccreditationBody** | **string** | Qualified Lawyer&#39;s Accreditation Body | 
**QliJurisdictionCountry** | **string** | Accreditation Jurisdiction Country (2 letter ISO country code) | 
**QliJurisdictionState** | Pointer to **string** | Accreditation Jurisdiction State/Province | [optional] 
**QliAccreditationYear** | **int32** | Qualified Lawyer&#39;s Accreditation Year (20XX) | 
**IdCardNumber** | **string** | ID Card Number | 
**RegistrationNumber** | **string** | Business number | 
**ParisNexus** | **string** | Link to Paris community | 
**AppPurpose** | **string** | Application purpose | 
**Category** | **string** | Application category | 
**Validator** | **string** | Country of Citizenship (2 letter ISO country code) | 

## Methods

### NewBodyDomainCreateDefaultExtra

`func NewBodyDomainCreateDefaultExtra(brRegisterNumber string, registrantType string, intendedUse string, entityType int32, nationalityCode string, regCode string, qliAccreditationId string, qliAccreditationBody string, qliJurisdictionCountry string, qliAccreditationYear int32, idCardNumber string, registrationNumber string, parisNexus string, appPurpose string, category string, validator string, ) *BodyDomainCreateDefaultExtra`

NewBodyDomainCreateDefaultExtra instantiates a new BodyDomainCreateDefaultExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodyDomainCreateDefaultExtraWithDefaults

`func NewBodyDomainCreateDefaultExtraWithDefaults() *BodyDomainCreateDefaultExtra`

NewBodyDomainCreateDefaultExtraWithDefaults instantiates a new BodyDomainCreateDefaultExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrRegisterNumber

`func (o *BodyDomainCreateDefaultExtra) GetBrRegisterNumber() string`

GetBrRegisterNumber returns the BrRegisterNumber field if non-nil, zero value otherwise.

### GetBrRegisterNumberOk

`func (o *BodyDomainCreateDefaultExtra) GetBrRegisterNumberOk() (*string, bool)`

GetBrRegisterNumberOk returns a tuple with the BrRegisterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrRegisterNumber

`func (o *BodyDomainCreateDefaultExtra) SetBrRegisterNumber(v string)`

SetBrRegisterNumber sets BrRegisterNumber field to given value.


### GetRegistrantType

`func (o *BodyDomainCreateDefaultExtra) GetRegistrantType() string`

GetRegistrantType returns the RegistrantType field if non-nil, zero value otherwise.

### GetRegistrantTypeOk

`func (o *BodyDomainCreateDefaultExtra) GetRegistrantTypeOk() (*string, bool)`

GetRegistrantTypeOk returns a tuple with the RegistrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantType

`func (o *BodyDomainCreateDefaultExtra) SetRegistrantType(v string)`

SetRegistrantType sets RegistrantType field to given value.


### GetIntendedUse

`func (o *BodyDomainCreateDefaultExtra) GetIntendedUse() string`

GetIntendedUse returns the IntendedUse field if non-nil, zero value otherwise.

### GetIntendedUseOk

`func (o *BodyDomainCreateDefaultExtra) GetIntendedUseOk() (*string, bool)`

GetIntendedUseOk returns a tuple with the IntendedUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedUse

`func (o *BodyDomainCreateDefaultExtra) SetIntendedUse(v string)`

SetIntendedUse sets IntendedUse field to given value.


### GetDateOfBirth

`func (o *BodyDomainCreateDefaultExtra) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *BodyDomainCreateDefaultExtra) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *BodyDomainCreateDefaultExtra) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *BodyDomainCreateDefaultExtra) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetCountryOfBirth

`func (o *BodyDomainCreateDefaultExtra) GetCountryOfBirth() string`

GetCountryOfBirth returns the CountryOfBirth field if non-nil, zero value otherwise.

### GetCountryOfBirthOk

`func (o *BodyDomainCreateDefaultExtra) GetCountryOfBirthOk() (*string, bool)`

GetCountryOfBirthOk returns a tuple with the CountryOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfBirth

`func (o *BodyDomainCreateDefaultExtra) SetCountryOfBirth(v string)`

SetCountryOfBirth sets CountryOfBirth field to given value.

### HasCountryOfBirth

`func (o *BodyDomainCreateDefaultExtra) HasCountryOfBirth() bool`

HasCountryOfBirth returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *BodyDomainCreateDefaultExtra) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *BodyDomainCreateDefaultExtra) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *BodyDomainCreateDefaultExtra) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *BodyDomainCreateDefaultExtra) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### GetPostalCodeOfBirth

`func (o *BodyDomainCreateDefaultExtra) GetPostalCodeOfBirth() string`

GetPostalCodeOfBirth returns the PostalCodeOfBirth field if non-nil, zero value otherwise.

### GetPostalCodeOfBirthOk

`func (o *BodyDomainCreateDefaultExtra) GetPostalCodeOfBirthOk() (*string, bool)`

GetPostalCodeOfBirthOk returns a tuple with the PostalCodeOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCodeOfBirth

`func (o *BodyDomainCreateDefaultExtra) SetPostalCodeOfBirth(v string)`

SetPostalCodeOfBirth sets PostalCodeOfBirth field to given value.

### HasPostalCodeOfBirth

`func (o *BodyDomainCreateDefaultExtra) HasPostalCodeOfBirth() bool`

HasPostalCodeOfBirth returns a boolean if a field has been set.

### GetRegistrantVatId

`func (o *BodyDomainCreateDefaultExtra) GetRegistrantVatId() string`

GetRegistrantVatId returns the RegistrantVatId field if non-nil, zero value otherwise.

### GetRegistrantVatIdOk

`func (o *BodyDomainCreateDefaultExtra) GetRegistrantVatIdOk() (*string, bool)`

GetRegistrantVatIdOk returns a tuple with the RegistrantVatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantVatId

`func (o *BodyDomainCreateDefaultExtra) SetRegistrantVatId(v string)`

SetRegistrantVatId sets RegistrantVatId field to given value.

### HasRegistrantVatId

`func (o *BodyDomainCreateDefaultExtra) HasRegistrantVatId() bool`

HasRegistrantVatId returns a boolean if a field has been set.

### GetSirenSiret

`func (o *BodyDomainCreateDefaultExtra) GetSirenSiret() string`

GetSirenSiret returns the SirenSiret field if non-nil, zero value otherwise.

### GetSirenSiretOk

`func (o *BodyDomainCreateDefaultExtra) GetSirenSiretOk() (*string, bool)`

GetSirenSiretOk returns a tuple with the SirenSiret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSirenSiret

`func (o *BodyDomainCreateDefaultExtra) SetSirenSiret(v string)`

SetSirenSiret sets SirenSiret field to given value.

### HasSirenSiret

`func (o *BodyDomainCreateDefaultExtra) HasSirenSiret() bool`

HasSirenSiret returns a boolean if a field has been set.

### GetTrademarkNumber

`func (o *BodyDomainCreateDefaultExtra) GetTrademarkNumber() string`

GetTrademarkNumber returns the TrademarkNumber field if non-nil, zero value otherwise.

### GetTrademarkNumberOk

`func (o *BodyDomainCreateDefaultExtra) GetTrademarkNumberOk() (*string, bool)`

GetTrademarkNumberOk returns a tuple with the TrademarkNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrademarkNumber

`func (o *BodyDomainCreateDefaultExtra) SetTrademarkNumber(v string)`

SetTrademarkNumber sets TrademarkNumber field to given value.

### HasTrademarkNumber

`func (o *BodyDomainCreateDefaultExtra) HasTrademarkNumber() bool`

HasTrademarkNumber returns a boolean if a field has been set.

### GetEntityType

`func (o *BodyDomainCreateDefaultExtra) GetEntityType() int32`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BodyDomainCreateDefaultExtra) GetEntityTypeOk() (*int32, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BodyDomainCreateDefaultExtra) SetEntityType(v int32)`

SetEntityType sets EntityType field to given value.


### GetNationalityCode

`func (o *BodyDomainCreateDefaultExtra) GetNationalityCode() string`

GetNationalityCode returns the NationalityCode field if non-nil, zero value otherwise.

### GetNationalityCodeOk

`func (o *BodyDomainCreateDefaultExtra) GetNationalityCodeOk() (*string, bool)`

GetNationalityCodeOk returns a tuple with the NationalityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityCode

`func (o *BodyDomainCreateDefaultExtra) SetNationalityCode(v string)`

SetNationalityCode sets NationalityCode field to given value.


### GetRegCode

`func (o *BodyDomainCreateDefaultExtra) GetRegCode() string`

GetRegCode returns the RegCode field if non-nil, zero value otherwise.

### GetRegCodeOk

`func (o *BodyDomainCreateDefaultExtra) GetRegCodeOk() (*string, bool)`

GetRegCodeOk returns a tuple with the RegCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegCode

`func (o *BodyDomainCreateDefaultExtra) SetRegCode(v string)`

SetRegCode sets RegCode field to given value.


### GetQliAccreditationId

`func (o *BodyDomainCreateDefaultExtra) GetQliAccreditationId() string`

GetQliAccreditationId returns the QliAccreditationId field if non-nil, zero value otherwise.

### GetQliAccreditationIdOk

`func (o *BodyDomainCreateDefaultExtra) GetQliAccreditationIdOk() (*string, bool)`

GetQliAccreditationIdOk returns a tuple with the QliAccreditationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQliAccreditationId

`func (o *BodyDomainCreateDefaultExtra) SetQliAccreditationId(v string)`

SetQliAccreditationId sets QliAccreditationId field to given value.


### GetQliAccreditationBody

`func (o *BodyDomainCreateDefaultExtra) GetQliAccreditationBody() string`

GetQliAccreditationBody returns the QliAccreditationBody field if non-nil, zero value otherwise.

### GetQliAccreditationBodyOk

`func (o *BodyDomainCreateDefaultExtra) GetQliAccreditationBodyOk() (*string, bool)`

GetQliAccreditationBodyOk returns a tuple with the QliAccreditationBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQliAccreditationBody

`func (o *BodyDomainCreateDefaultExtra) SetQliAccreditationBody(v string)`

SetQliAccreditationBody sets QliAccreditationBody field to given value.


### GetQliJurisdictionCountry

`func (o *BodyDomainCreateDefaultExtra) GetQliJurisdictionCountry() string`

GetQliJurisdictionCountry returns the QliJurisdictionCountry field if non-nil, zero value otherwise.

### GetQliJurisdictionCountryOk

`func (o *BodyDomainCreateDefaultExtra) GetQliJurisdictionCountryOk() (*string, bool)`

GetQliJurisdictionCountryOk returns a tuple with the QliJurisdictionCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQliJurisdictionCountry

`func (o *BodyDomainCreateDefaultExtra) SetQliJurisdictionCountry(v string)`

SetQliJurisdictionCountry sets QliJurisdictionCountry field to given value.


### GetQliJurisdictionState

`func (o *BodyDomainCreateDefaultExtra) GetQliJurisdictionState() string`

GetQliJurisdictionState returns the QliJurisdictionState field if non-nil, zero value otherwise.

### GetQliJurisdictionStateOk

`func (o *BodyDomainCreateDefaultExtra) GetQliJurisdictionStateOk() (*string, bool)`

GetQliJurisdictionStateOk returns a tuple with the QliJurisdictionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQliJurisdictionState

`func (o *BodyDomainCreateDefaultExtra) SetQliJurisdictionState(v string)`

SetQliJurisdictionState sets QliJurisdictionState field to given value.

### HasQliJurisdictionState

`func (o *BodyDomainCreateDefaultExtra) HasQliJurisdictionState() bool`

HasQliJurisdictionState returns a boolean if a field has been set.

### GetQliAccreditationYear

`func (o *BodyDomainCreateDefaultExtra) GetQliAccreditationYear() int32`

GetQliAccreditationYear returns the QliAccreditationYear field if non-nil, zero value otherwise.

### GetQliAccreditationYearOk

`func (o *BodyDomainCreateDefaultExtra) GetQliAccreditationYearOk() (*int32, bool)`

GetQliAccreditationYearOk returns a tuple with the QliAccreditationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQliAccreditationYear

`func (o *BodyDomainCreateDefaultExtra) SetQliAccreditationYear(v int32)`

SetQliAccreditationYear sets QliAccreditationYear field to given value.


### GetIdCardNumber

`func (o *BodyDomainCreateDefaultExtra) GetIdCardNumber() string`

GetIdCardNumber returns the IdCardNumber field if non-nil, zero value otherwise.

### GetIdCardNumberOk

`func (o *BodyDomainCreateDefaultExtra) GetIdCardNumberOk() (*string, bool)`

GetIdCardNumberOk returns a tuple with the IdCardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCardNumber

`func (o *BodyDomainCreateDefaultExtra) SetIdCardNumber(v string)`

SetIdCardNumber sets IdCardNumber field to given value.


### GetRegistrationNumber

`func (o *BodyDomainCreateDefaultExtra) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *BodyDomainCreateDefaultExtra) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *BodyDomainCreateDefaultExtra) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.


### GetParisNexus

`func (o *BodyDomainCreateDefaultExtra) GetParisNexus() string`

GetParisNexus returns the ParisNexus field if non-nil, zero value otherwise.

### GetParisNexusOk

`func (o *BodyDomainCreateDefaultExtra) GetParisNexusOk() (*string, bool)`

GetParisNexusOk returns a tuple with the ParisNexus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParisNexus

`func (o *BodyDomainCreateDefaultExtra) SetParisNexus(v string)`

SetParisNexus sets ParisNexus field to given value.


### GetAppPurpose

`func (o *BodyDomainCreateDefaultExtra) GetAppPurpose() string`

GetAppPurpose returns the AppPurpose field if non-nil, zero value otherwise.

### GetAppPurposeOk

`func (o *BodyDomainCreateDefaultExtra) GetAppPurposeOk() (*string, bool)`

GetAppPurposeOk returns a tuple with the AppPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPurpose

`func (o *BodyDomainCreateDefaultExtra) SetAppPurpose(v string)`

SetAppPurpose sets AppPurpose field to given value.


### GetCategory

`func (o *BodyDomainCreateDefaultExtra) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BodyDomainCreateDefaultExtra) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BodyDomainCreateDefaultExtra) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetValidator

`func (o *BodyDomainCreateDefaultExtra) GetValidator() string`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *BodyDomainCreateDefaultExtra) GetValidatorOk() (*string, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *BodyDomainCreateDefaultExtra) SetValidator(v string)`

SetValidator sets Validator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


