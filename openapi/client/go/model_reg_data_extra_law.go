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
	"bytes"
	"fmt"
)

// checks if the RegDataExtraLAW type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegDataExtraLAW{}

// RegDataExtraLAW struct for RegDataExtraLAW
type RegDataExtraLAW struct {
	// Qualified Lawyer's Accreditation ID
	QliAccreditationId string `json:"qli_accreditation_id"`
	// Qualified Lawyer's Accreditation Body
	QliAccreditationBody string `json:"qli_accreditation_body"`
	// Accreditation Jurisdiction Country (2 letter ISO country code)
	QliJurisdictionCountry string `json:"qli_jurisdiction_country"`
	// Accreditation Jurisdiction State/Province
	QliJurisdictionState *string `json:"qli_jurisdiction_state,omitempty"`
	// Qualified Lawyer's Accreditation Year (20XX)
	QliAccreditationYear int32 `json:"qli_accreditation_year"`
}

type _RegDataExtraLAW RegDataExtraLAW

// NewRegDataExtraLAW instantiates a new RegDataExtraLAW object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegDataExtraLAW(qliAccreditationId string, qliAccreditationBody string, qliJurisdictionCountry string, qliAccreditationYear int32) *RegDataExtraLAW {
	this := RegDataExtraLAW{}
	this.QliAccreditationId = qliAccreditationId
	this.QliAccreditationBody = qliAccreditationBody
	this.QliJurisdictionCountry = qliJurisdictionCountry
	this.QliAccreditationYear = qliAccreditationYear
	return &this
}

// NewRegDataExtraLAWWithDefaults instantiates a new RegDataExtraLAW object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegDataExtraLAWWithDefaults() *RegDataExtraLAW {
	this := RegDataExtraLAW{}
	return &this
}

// GetQliAccreditationId returns the QliAccreditationId field value
func (o *RegDataExtraLAW) GetQliAccreditationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QliAccreditationId
}

// GetQliAccreditationIdOk returns a tuple with the QliAccreditationId field value
// and a boolean to check if the value has been set.
func (o *RegDataExtraLAW) GetQliAccreditationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QliAccreditationId, true
}

// SetQliAccreditationId sets field value
func (o *RegDataExtraLAW) SetQliAccreditationId(v string) {
	o.QliAccreditationId = v
}

// GetQliAccreditationBody returns the QliAccreditationBody field value
func (o *RegDataExtraLAW) GetQliAccreditationBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QliAccreditationBody
}

// GetQliAccreditationBodyOk returns a tuple with the QliAccreditationBody field value
// and a boolean to check if the value has been set.
func (o *RegDataExtraLAW) GetQliAccreditationBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QliAccreditationBody, true
}

// SetQliAccreditationBody sets field value
func (o *RegDataExtraLAW) SetQliAccreditationBody(v string) {
	o.QliAccreditationBody = v
}

// GetQliJurisdictionCountry returns the QliJurisdictionCountry field value
func (o *RegDataExtraLAW) GetQliJurisdictionCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QliJurisdictionCountry
}

// GetQliJurisdictionCountryOk returns a tuple with the QliJurisdictionCountry field value
// and a boolean to check if the value has been set.
func (o *RegDataExtraLAW) GetQliJurisdictionCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QliJurisdictionCountry, true
}

// SetQliJurisdictionCountry sets field value
func (o *RegDataExtraLAW) SetQliJurisdictionCountry(v string) {
	o.QliJurisdictionCountry = v
}

// GetQliJurisdictionState returns the QliJurisdictionState field value if set, zero value otherwise.
func (o *RegDataExtraLAW) GetQliJurisdictionState() string {
	if o == nil || IsNil(o.QliJurisdictionState) {
		var ret string
		return ret
	}
	return *o.QliJurisdictionState
}

// GetQliJurisdictionStateOk returns a tuple with the QliJurisdictionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegDataExtraLAW) GetQliJurisdictionStateOk() (*string, bool) {
	if o == nil || IsNil(o.QliJurisdictionState) {
		return nil, false
	}
	return o.QliJurisdictionState, true
}

// HasQliJurisdictionState returns a boolean if a field has been set.
func (o *RegDataExtraLAW) HasQliJurisdictionState() bool {
	if o != nil && !IsNil(o.QliJurisdictionState) {
		return true
	}

	return false
}

// SetQliJurisdictionState gets a reference to the given string and assigns it to the QliJurisdictionState field.
func (o *RegDataExtraLAW) SetQliJurisdictionState(v string) {
	o.QliJurisdictionState = &v
}

// GetQliAccreditationYear returns the QliAccreditationYear field value
func (o *RegDataExtraLAW) GetQliAccreditationYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QliAccreditationYear
}

// GetQliAccreditationYearOk returns a tuple with the QliAccreditationYear field value
// and a boolean to check if the value has been set.
func (o *RegDataExtraLAW) GetQliAccreditationYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QliAccreditationYear, true
}

// SetQliAccreditationYear sets field value
func (o *RegDataExtraLAW) SetQliAccreditationYear(v int32) {
	o.QliAccreditationYear = v
}

func (o RegDataExtraLAW) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegDataExtraLAW) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["qli_accreditation_id"] = o.QliAccreditationId
	toSerialize["qli_accreditation_body"] = o.QliAccreditationBody
	toSerialize["qli_jurisdiction_country"] = o.QliJurisdictionCountry
	if !IsNil(o.QliJurisdictionState) {
		toSerialize["qli_jurisdiction_state"] = o.QliJurisdictionState
	}
	toSerialize["qli_accreditation_year"] = o.QliAccreditationYear
	return toSerialize, nil
}

func (o *RegDataExtraLAW) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"qli_accreditation_id",
		"qli_accreditation_body",
		"qli_jurisdiction_country",
		"qli_accreditation_year",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRegDataExtraLAW := _RegDataExtraLAW{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegDataExtraLAW)

	if err != nil {
		return err
	}

	*o = RegDataExtraLAW(varRegDataExtraLAW)

	return err
}

type NullableRegDataExtraLAW struct {
	value *RegDataExtraLAW
	isSet bool
}

func (v NullableRegDataExtraLAW) Get() *RegDataExtraLAW {
	return v.value
}

func (v *NullableRegDataExtraLAW) Set(val *RegDataExtraLAW) {
	v.value = val
	v.isSet = true
}

func (v NullableRegDataExtraLAW) IsSet() bool {
	return v.isSet
}

func (v *NullableRegDataExtraLAW) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegDataExtraLAW(val *RegDataExtraLAW) *NullableRegDataExtraLAW {
	return &NullableRegDataExtraLAW{value: val, isSet: true}
}

func (v NullableRegDataExtraLAW) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegDataExtraLAW) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

