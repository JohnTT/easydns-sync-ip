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

// checks if the ResultModelCreateDomainGlueData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultModelCreateDomainGlueData{}

// ResultModelCreateDomainGlueData Response data related to request
type ResultModelCreateDomainGlueData struct {
	// The domain the glue records are provided by
	Domain string `json:"domain"`
	// The hostname that the glue record is defined for
	NameserverName string `json:"nameserver_name"`
	// The IPv4 address defined for this glue record if one is defined
	Ipaddress *string `json:"ipaddress,omitempty"`
	// The IPv6 address defined for this glue record if one is defined
	Ipv6 *string `json:"ipv6,omitempty"`
}

type _ResultModelCreateDomainGlueData ResultModelCreateDomainGlueData

// NewResultModelCreateDomainGlueData instantiates a new ResultModelCreateDomainGlueData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultModelCreateDomainGlueData(domain string, nameserverName string) *ResultModelCreateDomainGlueData {
	this := ResultModelCreateDomainGlueData{}
	this.Domain = domain
	this.NameserverName = nameserverName
	return &this
}

// NewResultModelCreateDomainGlueDataWithDefaults instantiates a new ResultModelCreateDomainGlueData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultModelCreateDomainGlueDataWithDefaults() *ResultModelCreateDomainGlueData {
	this := ResultModelCreateDomainGlueData{}
	return &this
}

// GetDomain returns the Domain field value
func (o *ResultModelCreateDomainGlueData) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ResultModelCreateDomainGlueData) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ResultModelCreateDomainGlueData) SetDomain(v string) {
	o.Domain = v
}

// GetNameserverName returns the NameserverName field value
func (o *ResultModelCreateDomainGlueData) GetNameserverName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameserverName
}

// GetNameserverNameOk returns a tuple with the NameserverName field value
// and a boolean to check if the value has been set.
func (o *ResultModelCreateDomainGlueData) GetNameserverNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameserverName, true
}

// SetNameserverName sets field value
func (o *ResultModelCreateDomainGlueData) SetNameserverName(v string) {
	o.NameserverName = v
}

// GetIpaddress returns the Ipaddress field value if set, zero value otherwise.
func (o *ResultModelCreateDomainGlueData) GetIpaddress() string {
	if o == nil || IsNil(o.Ipaddress) {
		var ret string
		return ret
	}
	return *o.Ipaddress
}

// GetIpaddressOk returns a tuple with the Ipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelCreateDomainGlueData) GetIpaddressOk() (*string, bool) {
	if o == nil || IsNil(o.Ipaddress) {
		return nil, false
	}
	return o.Ipaddress, true
}

// HasIpaddress returns a boolean if a field has been set.
func (o *ResultModelCreateDomainGlueData) HasIpaddress() bool {
	if o != nil && !IsNil(o.Ipaddress) {
		return true
	}

	return false
}

// SetIpaddress gets a reference to the given string and assigns it to the Ipaddress field.
func (o *ResultModelCreateDomainGlueData) SetIpaddress(v string) {
	o.Ipaddress = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *ResultModelCreateDomainGlueData) GetIpv6() string {
	if o == nil || IsNil(o.Ipv6) {
		var ret string
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelCreateDomainGlueData) GetIpv6Ok() (*string, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *ResultModelCreateDomainGlueData) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given string and assigns it to the Ipv6 field.
func (o *ResultModelCreateDomainGlueData) SetIpv6(v string) {
	o.Ipv6 = &v
}

func (o ResultModelCreateDomainGlueData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultModelCreateDomainGlueData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	toSerialize["nameserver_name"] = o.NameserverName
	if !IsNil(o.Ipaddress) {
		toSerialize["ipaddress"] = o.Ipaddress
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	return toSerialize, nil
}

func (o *ResultModelCreateDomainGlueData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"nameserver_name",
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

	varResultModelCreateDomainGlueData := _ResultModelCreateDomainGlueData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultModelCreateDomainGlueData)

	if err != nil {
		return err
	}

	*o = ResultModelCreateDomainGlueData(varResultModelCreateDomainGlueData)

	return err
}

type NullableResultModelCreateDomainGlueData struct {
	value *ResultModelCreateDomainGlueData
	isSet bool
}

func (v NullableResultModelCreateDomainGlueData) Get() *ResultModelCreateDomainGlueData {
	return v.value
}

func (v *NullableResultModelCreateDomainGlueData) Set(val *ResultModelCreateDomainGlueData) {
	v.value = val
	v.isSet = true
}

func (v NullableResultModelCreateDomainGlueData) IsSet() bool {
	return v.isSet
}

func (v *NullableResultModelCreateDomainGlueData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultModelCreateDomainGlueData(val *ResultModelCreateDomainGlueData) *NullableResultModelCreateDomainGlueData {
	return &NullableResultModelCreateDomainGlueData{value: val, isSet: true}
}

func (v NullableResultModelCreateDomainGlueData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultModelCreateDomainGlueData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


