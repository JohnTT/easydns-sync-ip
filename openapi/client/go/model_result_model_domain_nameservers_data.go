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

// checks if the ResultModelDomainNameserversData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultModelDomainNameserversData{}

// ResultModelDomainNameserversData Response data related to request
type ResultModelDomainNameserversData struct {
	// The domain the glue records are provided by
	Domain string `json:"domain"`
	// The list of nameservers assigned to the domain
	Nameservers []string `json:"nameservers"`
}

type _ResultModelDomainNameserversData ResultModelDomainNameserversData

// NewResultModelDomainNameserversData instantiates a new ResultModelDomainNameserversData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultModelDomainNameserversData(domain string, nameservers []string) *ResultModelDomainNameserversData {
	this := ResultModelDomainNameserversData{}
	this.Domain = domain
	this.Nameservers = nameservers
	return &this
}

// NewResultModelDomainNameserversDataWithDefaults instantiates a new ResultModelDomainNameserversData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultModelDomainNameserversDataWithDefaults() *ResultModelDomainNameserversData {
	this := ResultModelDomainNameserversData{}
	return &this
}

// GetDomain returns the Domain field value
func (o *ResultModelDomainNameserversData) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ResultModelDomainNameserversData) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ResultModelDomainNameserversData) SetDomain(v string) {
	o.Domain = v
}

// GetNameservers returns the Nameservers field value
func (o *ResultModelDomainNameserversData) GetNameservers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value
// and a boolean to check if the value has been set.
func (o *ResultModelDomainNameserversData) GetNameserversOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nameservers, true
}

// SetNameservers sets field value
func (o *ResultModelDomainNameserversData) SetNameservers(v []string) {
	o.Nameservers = v
}

func (o ResultModelDomainNameserversData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultModelDomainNameserversData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	toSerialize["nameservers"] = o.Nameservers
	return toSerialize, nil
}

func (o *ResultModelDomainNameserversData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"nameservers",
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

	varResultModelDomainNameserversData := _ResultModelDomainNameserversData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultModelDomainNameserversData)

	if err != nil {
		return err
	}

	*o = ResultModelDomainNameserversData(varResultModelDomainNameserversData)

	return err
}

type NullableResultModelDomainNameserversData struct {
	value *ResultModelDomainNameserversData
	isSet bool
}

func (v NullableResultModelDomainNameserversData) Get() *ResultModelDomainNameserversData {
	return v.value
}

func (v *NullableResultModelDomainNameserversData) Set(val *ResultModelDomainNameserversData) {
	v.value = val
	v.isSet = true
}

func (v NullableResultModelDomainNameserversData) IsSet() bool {
	return v.isSet
}

func (v *NullableResultModelDomainNameserversData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultModelDomainNameserversData(val *ResultModelDomainNameserversData) *NullableResultModelDomainNameserversData {
	return &NullableResultModelDomainNameserversData{value: val, isSet: true}
}

func (v NullableResultModelDomainNameserversData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultModelDomainNameserversData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


