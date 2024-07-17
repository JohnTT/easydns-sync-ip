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

// checks if the ResultSetPrimaryNSData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultSetPrimaryNSData{}

// ResultSetPrimaryNSData struct for ResultSetPrimaryNSData
type ResultSetPrimaryNSData struct {
	// The domain name that was updated
	Domain string `json:"domain"`
	// The value for the primary NS that was used
	Master string `json:"master"`
}

type _ResultSetPrimaryNSData ResultSetPrimaryNSData

// NewResultSetPrimaryNSData instantiates a new ResultSetPrimaryNSData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultSetPrimaryNSData(domain string, master string) *ResultSetPrimaryNSData {
	this := ResultSetPrimaryNSData{}
	this.Domain = domain
	this.Master = master
	return &this
}

// NewResultSetPrimaryNSDataWithDefaults instantiates a new ResultSetPrimaryNSData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultSetPrimaryNSDataWithDefaults() *ResultSetPrimaryNSData {
	this := ResultSetPrimaryNSData{}
	return &this
}

// GetDomain returns the Domain field value
func (o *ResultSetPrimaryNSData) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ResultSetPrimaryNSData) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ResultSetPrimaryNSData) SetDomain(v string) {
	o.Domain = v
}

// GetMaster returns the Master field value
func (o *ResultSetPrimaryNSData) GetMaster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Master
}

// GetMasterOk returns a tuple with the Master field value
// and a boolean to check if the value has been set.
func (o *ResultSetPrimaryNSData) GetMasterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Master, true
}

// SetMaster sets field value
func (o *ResultSetPrimaryNSData) SetMaster(v string) {
	o.Master = v
}

func (o ResultSetPrimaryNSData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultSetPrimaryNSData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	toSerialize["master"] = o.Master
	return toSerialize, nil
}

func (o *ResultSetPrimaryNSData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"master",
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

	varResultSetPrimaryNSData := _ResultSetPrimaryNSData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultSetPrimaryNSData)

	if err != nil {
		return err
	}

	*o = ResultSetPrimaryNSData(varResultSetPrimaryNSData)

	return err
}

type NullableResultSetPrimaryNSData struct {
	value *ResultSetPrimaryNSData
	isSet bool
}

func (v NullableResultSetPrimaryNSData) Get() *ResultSetPrimaryNSData {
	return v.value
}

func (v *NullableResultSetPrimaryNSData) Set(val *ResultSetPrimaryNSData) {
	v.value = val
	v.isSet = true
}

func (v NullableResultSetPrimaryNSData) IsSet() bool {
	return v.isSet
}

func (v *NullableResultSetPrimaryNSData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultSetPrimaryNSData(val *ResultSetPrimaryNSData) *NullableResultSetPrimaryNSData {
	return &NullableResultSetPrimaryNSData{value: val, isSet: true}
}

func (v NullableResultSetPrimaryNSData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultSetPrimaryNSData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

