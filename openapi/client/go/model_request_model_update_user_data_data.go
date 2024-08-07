/*
EasyAPI REST Services API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
Contact: easyapi@easydns.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RequestModelUpdateUserDataData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestModelUpdateUserDataData{}

// RequestModelUpdateUserDataData Data must be encrypted when sent in to a single data object.
type RequestModelUpdateUserDataData struct {
	// The first name of the user
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the user
	LastName *string `json:"last_name,omitempty"`
	// The organization of the user
	OrgName *string `json:"org_name,omitempty"`
	// The address of the user
	Address1 *string `json:"address1,omitempty"`
	Address2 *string `json:"address2,omitempty"`
	Address3 *string `json:"address3,omitempty"`
	// The city of the address provided
	City *string `json:"city,omitempty"`
	// The state/province code of the address
	State *string `json:"state,omitempty"`
	// The 2-letter ISO country code of the address
	Country *string `json:"country,omitempty"`
	// The zip/postal code of the address
	PostalCode *string `json:"postal_code,omitempty"`
	// The default currency to use for the account
	Currency *string `json:"currency,omitempty"`
	// Phone number that can be used to contact customer
	Phone *string `json:"phone,omitempty"`
	// Cellular number that can be used to contact customer
	Cellphone *string `json:"cellphone,omitempty"`
	// Fax number that can be used to contact customer
	Fax *string `json:"fax,omitempty"`
	// Main contact email for user
	Email *string `json:"email,omitempty"`
	// Secondary contact email for user
	Email2 *string `json:"email2,omitempty"`
	// The email address(es) where generic notices should be sent
	NoticesEmail *string `json:"notices_email,omitempty"`
	// The email address(es) that can be made public for contacting the user
	PublicEmail *string `json:"public_email,omitempty"`
	// The email address(es) where alert notice should be sent
	AlertsEmail *string `json:"alerts_email,omitempty"`
	// The URL associated with a user (Not currently used anywhere)
	Url *string `json:"url,omitempty"`
	// Has the user opted out of any non-essential communications
	OptOut *int32 `json:"opt_out,omitempty"`
	// Does the user access to BETA level features
	Beta *int32 `json:"beta,omitempty"`
	// Is the account currently marked as fraud
	Fraud *string `json:"fraud,omitempty"`
	// Is the user currently marked as a VIP user
	Vip *int32 `json:"vip,omitempty"`
	// Provides a new password for the user that will replace the users current password
	Password *string `json:"password,omitempty"`
	Esm *RequestModelUpdateUserDataDataEsm `json:"esm,omitempty"`
}

// NewRequestModelUpdateUserDataData instantiates a new RequestModelUpdateUserDataData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestModelUpdateUserDataData() *RequestModelUpdateUserDataData {
	this := RequestModelUpdateUserDataData{}
	var fraud string = "N"
	this.Fraud = &fraud
	var vip int32 = 0
	this.Vip = &vip
	return &this
}

// NewRequestModelUpdateUserDataDataWithDefaults instantiates a new RequestModelUpdateUserDataData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestModelUpdateUserDataDataWithDefaults() *RequestModelUpdateUserDataData {
	this := RequestModelUpdateUserDataData{}
	var fraud string = "N"
	this.Fraud = &fraud
	var vip int32 = 0
	this.Vip = &vip
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *RequestModelUpdateUserDataData) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *RequestModelUpdateUserDataData) SetLastName(v string) {
	o.LastName = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetOrgName() string {
	if o == nil || IsNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetOrgNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrgName) {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasOrgName() bool {
	if o != nil && !IsNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *RequestModelUpdateUserDataData) SetOrgName(v string) {
	o.OrgName = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *RequestModelUpdateUserDataData) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *RequestModelUpdateUserDataData) SetAddress2(v string) {
	o.Address2 = &v
}

// GetAddress3 returns the Address3 field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetAddress3() string {
	if o == nil || IsNil(o.Address3) {
		var ret string
		return ret
	}
	return *o.Address3
}

// GetAddress3Ok returns a tuple with the Address3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetAddress3Ok() (*string, bool) {
	if o == nil || IsNil(o.Address3) {
		return nil, false
	}
	return o.Address3, true
}

// HasAddress3 returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasAddress3() bool {
	if o != nil && !IsNil(o.Address3) {
		return true
	}

	return false
}

// SetAddress3 gets a reference to the given string and assigns it to the Address3 field.
func (o *RequestModelUpdateUserDataData) SetAddress3(v string) {
	o.Address3 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *RequestModelUpdateUserDataData) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RequestModelUpdateUserDataData) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *RequestModelUpdateUserDataData) SetCountry(v string) {
	o.Country = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *RequestModelUpdateUserDataData) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *RequestModelUpdateUserDataData) SetCurrency(v string) {
	o.Currency = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *RequestModelUpdateUserDataData) SetPhone(v string) {
	o.Phone = &v
}

// GetCellphone returns the Cellphone field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetCellphone() string {
	if o == nil || IsNil(o.Cellphone) {
		var ret string
		return ret
	}
	return *o.Cellphone
}

// GetCellphoneOk returns a tuple with the Cellphone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetCellphoneOk() (*string, bool) {
	if o == nil || IsNil(o.Cellphone) {
		return nil, false
	}
	return o.Cellphone, true
}

// HasCellphone returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasCellphone() bool {
	if o != nil && !IsNil(o.Cellphone) {
		return true
	}

	return false
}

// SetCellphone gets a reference to the given string and assigns it to the Cellphone field.
func (o *RequestModelUpdateUserDataData) SetCellphone(v string) {
	o.Cellphone = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetFax() string {
	if o == nil || IsNil(o.Fax) {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetFaxOk() (*string, bool) {
	if o == nil || IsNil(o.Fax) {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasFax() bool {
	if o != nil && !IsNil(o.Fax) {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *RequestModelUpdateUserDataData) SetFax(v string) {
	o.Fax = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RequestModelUpdateUserDataData) SetEmail(v string) {
	o.Email = &v
}

// GetEmail2 returns the Email2 field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetEmail2() string {
	if o == nil || IsNil(o.Email2) {
		var ret string
		return ret
	}
	return *o.Email2
}

// GetEmail2Ok returns a tuple with the Email2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetEmail2Ok() (*string, bool) {
	if o == nil || IsNil(o.Email2) {
		return nil, false
	}
	return o.Email2, true
}

// HasEmail2 returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasEmail2() bool {
	if o != nil && !IsNil(o.Email2) {
		return true
	}

	return false
}

// SetEmail2 gets a reference to the given string and assigns it to the Email2 field.
func (o *RequestModelUpdateUserDataData) SetEmail2(v string) {
	o.Email2 = &v
}

// GetNoticesEmail returns the NoticesEmail field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetNoticesEmail() string {
	if o == nil || IsNil(o.NoticesEmail) {
		var ret string
		return ret
	}
	return *o.NoticesEmail
}

// GetNoticesEmailOk returns a tuple with the NoticesEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetNoticesEmailOk() (*string, bool) {
	if o == nil || IsNil(o.NoticesEmail) {
		return nil, false
	}
	return o.NoticesEmail, true
}

// HasNoticesEmail returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasNoticesEmail() bool {
	if o != nil && !IsNil(o.NoticesEmail) {
		return true
	}

	return false
}

// SetNoticesEmail gets a reference to the given string and assigns it to the NoticesEmail field.
func (o *RequestModelUpdateUserDataData) SetNoticesEmail(v string) {
	o.NoticesEmail = &v
}

// GetPublicEmail returns the PublicEmail field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetPublicEmail() string {
	if o == nil || IsNil(o.PublicEmail) {
		var ret string
		return ret
	}
	return *o.PublicEmail
}

// GetPublicEmailOk returns a tuple with the PublicEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetPublicEmailOk() (*string, bool) {
	if o == nil || IsNil(o.PublicEmail) {
		return nil, false
	}
	return o.PublicEmail, true
}

// HasPublicEmail returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasPublicEmail() bool {
	if o != nil && !IsNil(o.PublicEmail) {
		return true
	}

	return false
}

// SetPublicEmail gets a reference to the given string and assigns it to the PublicEmail field.
func (o *RequestModelUpdateUserDataData) SetPublicEmail(v string) {
	o.PublicEmail = &v
}

// GetAlertsEmail returns the AlertsEmail field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetAlertsEmail() string {
	if o == nil || IsNil(o.AlertsEmail) {
		var ret string
		return ret
	}
	return *o.AlertsEmail
}

// GetAlertsEmailOk returns a tuple with the AlertsEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetAlertsEmailOk() (*string, bool) {
	if o == nil || IsNil(o.AlertsEmail) {
		return nil, false
	}
	return o.AlertsEmail, true
}

// HasAlertsEmail returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasAlertsEmail() bool {
	if o != nil && !IsNil(o.AlertsEmail) {
		return true
	}

	return false
}

// SetAlertsEmail gets a reference to the given string and assigns it to the AlertsEmail field.
func (o *RequestModelUpdateUserDataData) SetAlertsEmail(v string) {
	o.AlertsEmail = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RequestModelUpdateUserDataData) SetUrl(v string) {
	o.Url = &v
}

// GetOptOut returns the OptOut field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetOptOut() int32 {
	if o == nil || IsNil(o.OptOut) {
		var ret int32
		return ret
	}
	return *o.OptOut
}

// GetOptOutOk returns a tuple with the OptOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetOptOutOk() (*int32, bool) {
	if o == nil || IsNil(o.OptOut) {
		return nil, false
	}
	return o.OptOut, true
}

// HasOptOut returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasOptOut() bool {
	if o != nil && !IsNil(o.OptOut) {
		return true
	}

	return false
}

// SetOptOut gets a reference to the given int32 and assigns it to the OptOut field.
func (o *RequestModelUpdateUserDataData) SetOptOut(v int32) {
	o.OptOut = &v
}

// GetBeta returns the Beta field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetBeta() int32 {
	if o == nil || IsNil(o.Beta) {
		var ret int32
		return ret
	}
	return *o.Beta
}

// GetBetaOk returns a tuple with the Beta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetBetaOk() (*int32, bool) {
	if o == nil || IsNil(o.Beta) {
		return nil, false
	}
	return o.Beta, true
}

// HasBeta returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasBeta() bool {
	if o != nil && !IsNil(o.Beta) {
		return true
	}

	return false
}

// SetBeta gets a reference to the given int32 and assigns it to the Beta field.
func (o *RequestModelUpdateUserDataData) SetBeta(v int32) {
	o.Beta = &v
}

// GetFraud returns the Fraud field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetFraud() string {
	if o == nil || IsNil(o.Fraud) {
		var ret string
		return ret
	}
	return *o.Fraud
}

// GetFraudOk returns a tuple with the Fraud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetFraudOk() (*string, bool) {
	if o == nil || IsNil(o.Fraud) {
		return nil, false
	}
	return o.Fraud, true
}

// HasFraud returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasFraud() bool {
	if o != nil && !IsNil(o.Fraud) {
		return true
	}

	return false
}

// SetFraud gets a reference to the given string and assigns it to the Fraud field.
func (o *RequestModelUpdateUserDataData) SetFraud(v string) {
	o.Fraud = &v
}

// GetVip returns the Vip field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetVip() int32 {
	if o == nil || IsNil(o.Vip) {
		var ret int32
		return ret
	}
	return *o.Vip
}

// GetVipOk returns a tuple with the Vip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetVipOk() (*int32, bool) {
	if o == nil || IsNil(o.Vip) {
		return nil, false
	}
	return o.Vip, true
}

// HasVip returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasVip() bool {
	if o != nil && !IsNil(o.Vip) {
		return true
	}

	return false
}

// SetVip gets a reference to the given int32 and assigns it to the Vip field.
func (o *RequestModelUpdateUserDataData) SetVip(v int32) {
	o.Vip = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *RequestModelUpdateUserDataData) SetPassword(v string) {
	o.Password = &v
}

// GetEsm returns the Esm field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataData) GetEsm() RequestModelUpdateUserDataDataEsm {
	if o == nil || IsNil(o.Esm) {
		var ret RequestModelUpdateUserDataDataEsm
		return ret
	}
	return *o.Esm
}

// GetEsmOk returns a tuple with the Esm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataData) GetEsmOk() (*RequestModelUpdateUserDataDataEsm, bool) {
	if o == nil || IsNil(o.Esm) {
		return nil, false
	}
	return o.Esm, true
}

// HasEsm returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataData) HasEsm() bool {
	if o != nil && !IsNil(o.Esm) {
		return true
	}

	return false
}

// SetEsm gets a reference to the given RequestModelUpdateUserDataDataEsm and assigns it to the Esm field.
func (o *RequestModelUpdateUserDataData) SetEsm(v RequestModelUpdateUserDataDataEsm) {
	o.Esm = &v
}

func (o RequestModelUpdateUserDataData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestModelUpdateUserDataData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.OrgName) {
		toSerialize["org_name"] = o.OrgName
	}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.Address3) {
		toSerialize["address3"] = o.Address3
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Cellphone) {
		toSerialize["cellphone"] = o.Cellphone
	}
	if !IsNil(o.Fax) {
		toSerialize["fax"] = o.Fax
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Email2) {
		toSerialize["email2"] = o.Email2
	}
	if !IsNil(o.NoticesEmail) {
		toSerialize["notices_email"] = o.NoticesEmail
	}
	if !IsNil(o.PublicEmail) {
		toSerialize["public_email"] = o.PublicEmail
	}
	if !IsNil(o.AlertsEmail) {
		toSerialize["alerts_email"] = o.AlertsEmail
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.OptOut) {
		toSerialize["opt_out"] = o.OptOut
	}
	if !IsNil(o.Beta) {
		toSerialize["beta"] = o.Beta
	}
	if !IsNil(o.Fraud) {
		toSerialize["fraud"] = o.Fraud
	}
	if !IsNil(o.Vip) {
		toSerialize["vip"] = o.Vip
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Esm) {
		toSerialize["esm"] = o.Esm
	}
	return toSerialize, nil
}

type NullableRequestModelUpdateUserDataData struct {
	value *RequestModelUpdateUserDataData
	isSet bool
}

func (v NullableRequestModelUpdateUserDataData) Get() *RequestModelUpdateUserDataData {
	return v.value
}

func (v *NullableRequestModelUpdateUserDataData) Set(val *RequestModelUpdateUserDataData) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestModelUpdateUserDataData) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestModelUpdateUserDataData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestModelUpdateUserDataData(val *RequestModelUpdateUserDataData) *NullableRequestModelUpdateUserDataData {
	return &NullableRequestModelUpdateUserDataData{value: val, isSet: true}
}

func (v NullableRequestModelUpdateUserDataData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestModelUpdateUserDataData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


