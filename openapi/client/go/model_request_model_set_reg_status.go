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

// checks if the RequestModelSetRegStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestModelSetRegStatus{}

// RequestModelSetRegStatus struct for RequestModelSetRegStatus
type RequestModelSetRegStatus struct {
	// The domain name to update
	Domain string `json:"domain"`
	// Turn on/off reglock at the registry
	Reglock *bool `json:"reglock,omitempty"`
	// Set renewal action of domain
	Renewal *string `json:"renewal,omitempty"`
}

type _RequestModelSetRegStatus RequestModelSetRegStatus

// NewRequestModelSetRegStatus instantiates a new RequestModelSetRegStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestModelSetRegStatus(domain string) *RequestModelSetRegStatus {
	this := RequestModelSetRegStatus{}
	this.Domain = domain
	return &this
}

// NewRequestModelSetRegStatusWithDefaults instantiates a new RequestModelSetRegStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestModelSetRegStatusWithDefaults() *RequestModelSetRegStatus {
	this := RequestModelSetRegStatus{}
	return &this
}

// GetDomain returns the Domain field value
func (o *RequestModelSetRegStatus) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *RequestModelSetRegStatus) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *RequestModelSetRegStatus) SetDomain(v string) {
	o.Domain = v
}

// GetReglock returns the Reglock field value if set, zero value otherwise.
func (o *RequestModelSetRegStatus) GetReglock() bool {
	if o == nil || IsNil(o.Reglock) {
		var ret bool
		return ret
	}
	return *o.Reglock
}

// GetReglockOk returns a tuple with the Reglock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelSetRegStatus) GetReglockOk() (*bool, bool) {
	if o == nil || IsNil(o.Reglock) {
		return nil, false
	}
	return o.Reglock, true
}

// HasReglock returns a boolean if a field has been set.
func (o *RequestModelSetRegStatus) HasReglock() bool {
	if o != nil && !IsNil(o.Reglock) {
		return true
	}

	return false
}

// SetReglock gets a reference to the given bool and assigns it to the Reglock field.
func (o *RequestModelSetRegStatus) SetReglock(v bool) {
	o.Reglock = &v
}

// GetRenewal returns the Renewal field value if set, zero value otherwise.
func (o *RequestModelSetRegStatus) GetRenewal() string {
	if o == nil || IsNil(o.Renewal) {
		var ret string
		return ret
	}
	return *o.Renewal
}

// GetRenewalOk returns a tuple with the Renewal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestModelSetRegStatus) GetRenewalOk() (*string, bool) {
	if o == nil || IsNil(o.Renewal) {
		return nil, false
	}
	return o.Renewal, true
}

// HasRenewal returns a boolean if a field has been set.
func (o *RequestModelSetRegStatus) HasRenewal() bool {
	if o != nil && !IsNil(o.Renewal) {
		return true
	}

	return false
}

// SetRenewal gets a reference to the given string and assigns it to the Renewal field.
func (o *RequestModelSetRegStatus) SetRenewal(v string) {
	o.Renewal = &v
}

func (o RequestModelSetRegStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestModelSetRegStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	if !IsNil(o.Reglock) {
		toSerialize["reglock"] = o.Reglock
	}
	if !IsNil(o.Renewal) {
		toSerialize["renewal"] = o.Renewal
	}
	return toSerialize, nil
}

func (o *RequestModelSetRegStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
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

	varRequestModelSetRegStatus := _RequestModelSetRegStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRequestModelSetRegStatus)

	if err != nil {
		return err
	}

	*o = RequestModelSetRegStatus(varRequestModelSetRegStatus)

	return err
}

type NullableRequestModelSetRegStatus struct {
	value *RequestModelSetRegStatus
	isSet bool
}

func (v NullableRequestModelSetRegStatus) Get() *RequestModelSetRegStatus {
	return v.value
}

func (v *NullableRequestModelSetRegStatus) Set(val *RequestModelSetRegStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestModelSetRegStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestModelSetRegStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestModelSetRegStatus(val *RequestModelSetRegStatus) *NullableRequestModelSetRegStatus {
	return &NullableRequestModelSetRegStatus{value: val, isSet: true}
}

func (v NullableRequestModelSetRegStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestModelSetRegStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

