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

// checks if the ResultModelCheckRegistryGlueData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultModelCheckRegistryGlueData{}

// ResultModelCheckRegistryGlueData Response data related to request
type ResultModelCheckRegistryGlueData struct {
	// The domain the glue records are provided by
	Domain string `json:"domain"`
	// The hostname for the glue record that was looked up
	Fqdn string `json:"fqdn"`
	// Boolean integer that indicates if the glue is configured at the registry
	Exists int32 `json:"exists"`
}

type _ResultModelCheckRegistryGlueData ResultModelCheckRegistryGlueData

// NewResultModelCheckRegistryGlueData instantiates a new ResultModelCheckRegistryGlueData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultModelCheckRegistryGlueData(domain string, fqdn string, exists int32) *ResultModelCheckRegistryGlueData {
	this := ResultModelCheckRegistryGlueData{}
	this.Domain = domain
	this.Fqdn = fqdn
	this.Exists = exists
	return &this
}

// NewResultModelCheckRegistryGlueDataWithDefaults instantiates a new ResultModelCheckRegistryGlueData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultModelCheckRegistryGlueDataWithDefaults() *ResultModelCheckRegistryGlueData {
	this := ResultModelCheckRegistryGlueData{}
	return &this
}

// GetDomain returns the Domain field value
func (o *ResultModelCheckRegistryGlueData) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ResultModelCheckRegistryGlueData) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ResultModelCheckRegistryGlueData) SetDomain(v string) {
	o.Domain = v
}

// GetFqdn returns the Fqdn field value
func (o *ResultModelCheckRegistryGlueData) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *ResultModelCheckRegistryGlueData) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *ResultModelCheckRegistryGlueData) SetFqdn(v string) {
	o.Fqdn = v
}

// GetExists returns the Exists field value
func (o *ResultModelCheckRegistryGlueData) GetExists() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Exists
}

// GetExistsOk returns a tuple with the Exists field value
// and a boolean to check if the value has been set.
func (o *ResultModelCheckRegistryGlueData) GetExistsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exists, true
}

// SetExists sets field value
func (o *ResultModelCheckRegistryGlueData) SetExists(v int32) {
	o.Exists = v
}

func (o ResultModelCheckRegistryGlueData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultModelCheckRegistryGlueData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	toSerialize["fqdn"] = o.Fqdn
	toSerialize["exists"] = o.Exists
	return toSerialize, nil
}

func (o *ResultModelCheckRegistryGlueData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"fqdn",
		"exists",
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

	varResultModelCheckRegistryGlueData := _ResultModelCheckRegistryGlueData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultModelCheckRegistryGlueData)

	if err != nil {
		return err
	}

	*o = ResultModelCheckRegistryGlueData(varResultModelCheckRegistryGlueData)

	return err
}

type NullableResultModelCheckRegistryGlueData struct {
	value *ResultModelCheckRegistryGlueData
	isSet bool
}

func (v NullableResultModelCheckRegistryGlueData) Get() *ResultModelCheckRegistryGlueData {
	return v.value
}

func (v *NullableResultModelCheckRegistryGlueData) Set(val *ResultModelCheckRegistryGlueData) {
	v.value = val
	v.isSet = true
}

func (v NullableResultModelCheckRegistryGlueData) IsSet() bool {
	return v.isSet
}

func (v *NullableResultModelCheckRegistryGlueData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultModelCheckRegistryGlueData(val *ResultModelCheckRegistryGlueData) *NullableResultModelCheckRegistryGlueData {
	return &NullableResultModelCheckRegistryGlueData{value: val, isSet: true}
}

func (v NullableResultModelCheckRegistryGlueData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultModelCheckRegistryGlueData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


