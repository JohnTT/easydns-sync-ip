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

// checks if the RequestModelUpdateUserDataDataEsm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestModelUpdateUserDataDataEsm{}

// RequestModelUpdateUserDataDataEsm Enhanced Security Mode (ESM) user settings to update
type RequestModelUpdateUserDataDataEsm struct {
	// The new list of ESM keywords to set for the user
	Keywords []string `json:"keywords,omitempty"`
	// The new phone access code to set for the user
	PhoneCode *string `json:"phone_code,omitempty"`
}

// NewRequestModelUpdateUserDataDataEsm instantiates a new RequestModelUpdateUserDataDataEsm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestModelUpdateUserDataDataEsm() *RequestModelUpdateUserDataDataEsm {
	this := RequestModelUpdateUserDataDataEsm{}
	return &this
}

// NewRequestModelUpdateUserDataDataEsmWithDefaults instantiates a new RequestModelUpdateUserDataDataEsm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestModelUpdateUserDataDataEsmWithDefaults() *RequestModelUpdateUserDataDataEsm {
	this := RequestModelUpdateUserDataDataEsm{}
	return &this
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataDataEsm) GetKeywords() []string {
	if o == nil || IsNil(o.Keywords) {
		var ret []string
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataDataEsm) GetKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataDataEsm) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []string and assigns it to the Keywords field.
func (o *RequestModelUpdateUserDataDataEsm) SetKeywords(v []string) {
	o.Keywords = v
}

// GetPhoneCode returns the PhoneCode field value if set, zero value otherwise.
func (o *RequestModelUpdateUserDataDataEsm) GetPhoneCode() string {
	if o == nil || IsNil(o.PhoneCode) {
		var ret string
		return ret
	}
	return *o.PhoneCode
}

// GetPhoneCodeOk returns a tuple with the PhoneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserDataDataEsm) GetPhoneCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneCode) {
		return nil, false
	}
	return o.PhoneCode, true
}

// HasPhoneCode returns a boolean if a field has been set.
func (o *RequestModelUpdateUserDataDataEsm) HasPhoneCode() bool {
	if o != nil && !IsNil(o.PhoneCode) {
		return true
	}

	return false
}

// SetPhoneCode gets a reference to the given string and assigns it to the PhoneCode field.
func (o *RequestModelUpdateUserDataDataEsm) SetPhoneCode(v string) {
	o.PhoneCode = &v
}

func (o RequestModelUpdateUserDataDataEsm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestModelUpdateUserDataDataEsm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.PhoneCode) {
		toSerialize["phone_code"] = o.PhoneCode
	}
	return toSerialize, nil
}

type NullableRequestModelUpdateUserDataDataEsm struct {
	value *RequestModelUpdateUserDataDataEsm
	isSet bool
}

func (v NullableRequestModelUpdateUserDataDataEsm) Get() *RequestModelUpdateUserDataDataEsm {
	return v.value
}

func (v *NullableRequestModelUpdateUserDataDataEsm) Set(val *RequestModelUpdateUserDataDataEsm) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestModelUpdateUserDataDataEsm) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestModelUpdateUserDataDataEsm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestModelUpdateUserDataDataEsm(val *RequestModelUpdateUserDataDataEsm) *NullableRequestModelUpdateUserDataDataEsm {
	return &NullableRequestModelUpdateUserDataDataEsm{value: val, isSet: true}
}

func (v NullableRequestModelUpdateUserDataDataEsm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestModelUpdateUserDataDataEsm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

