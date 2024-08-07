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

// checks if the RequestModelUpdateUserData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestModelUpdateUserData{}

// RequestModelUpdateUserData struct for RequestModelUpdateUserData
type RequestModelUpdateUserData struct {
	// Indicates that the data of the packet has been encrypted before transmission and placed in the data field.
	Encrypted *int32 `json:"_encrypted,omitempty"`
	Data *RequestModelUpdateUserDataData `json:"data,omitempty"`
}

// NewRequestModelUpdateUserData instantiates a new RequestModelUpdateUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestModelUpdateUserData() *RequestModelUpdateUserData {
	this := RequestModelUpdateUserData{}
	return &this
}

// NewRequestModelUpdateUserDataWithDefaults instantiates a new RequestModelUpdateUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestModelUpdateUserDataWithDefaults() *RequestModelUpdateUserData {
	this := RequestModelUpdateUserData{}
	return &this
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *RequestModelUpdateUserData) GetEncrypted() int32 {
	if o == nil || IsNil(o.Encrypted) {
		var ret int32
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserData) GetEncryptedOk() (*int32, bool) {
	if o == nil || IsNil(o.Encrypted) {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *RequestModelUpdateUserData) HasEncrypted() bool {
	if o != nil && !IsNil(o.Encrypted) {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given int32 and assigns it to the Encrypted field.
func (o *RequestModelUpdateUserData) SetEncrypted(v int32) {
	o.Encrypted = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RequestModelUpdateUserData) GetData() RequestModelUpdateUserDataData {
	if o == nil || IsNil(o.Data) {
		var ret RequestModelUpdateUserDataData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelUpdateUserData) GetDataOk() (*RequestModelUpdateUserDataData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RequestModelUpdateUserData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RequestModelUpdateUserDataData and assigns it to the Data field.
func (o *RequestModelUpdateUserData) SetData(v RequestModelUpdateUserDataData) {
	o.Data = &v
}

func (o RequestModelUpdateUserData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestModelUpdateUserData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Encrypted) {
		toSerialize["_encrypted"] = o.Encrypted
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRequestModelUpdateUserData struct {
	value *RequestModelUpdateUserData
	isSet bool
}

func (v NullableRequestModelUpdateUserData) Get() *RequestModelUpdateUserData {
	return v.value
}

func (v *NullableRequestModelUpdateUserData) Set(val *RequestModelUpdateUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestModelUpdateUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestModelUpdateUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestModelUpdateUserData(val *RequestModelUpdateUserData) *NullableRequestModelUpdateUserData {
	return &NullableRequestModelUpdateUserData{value: val, isSet: true}
}

func (v NullableRequestModelUpdateUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestModelUpdateUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


