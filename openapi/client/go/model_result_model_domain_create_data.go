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

// checks if the ResultModelDomainCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultModelDomainCreateData{}

// ResultModelDomainCreateData struct for ResultModelDomainCreateData
type ResultModelDomainCreateData struct {
	// The domain name processed as part of the request
	Domain string `json:"domain"`
	// The period (in years) to create the domain for on the system/registry (if applicable)
	Term int32 `json:"term"`
	// The service ID the domain was created as
	Service int32 `json:"service"`
	Tld string `json:"tld"`
	// The ID of the invoice that was created and processed to fullfill the request
	InvId int32 `json:"inv_id"`
	// The currency used to process the invoice
	Currency *string `json:"currency,omitempty"`
	User *string `json:"user,omitempty"`
}

type _ResultModelDomainCreateData ResultModelDomainCreateData

// NewResultModelDomainCreateData instantiates a new ResultModelDomainCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultModelDomainCreateData(domain string, term int32, service int32, tld string, invId int32) *ResultModelDomainCreateData {
	this := ResultModelDomainCreateData{}
	this.Domain = domain
	this.Term = term
	this.Service = service
	this.Tld = tld
	this.InvId = invId
	return &this
}

// NewResultModelDomainCreateDataWithDefaults instantiates a new ResultModelDomainCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultModelDomainCreateDataWithDefaults() *ResultModelDomainCreateData {
	this := ResultModelDomainCreateData{}
	return &this
}

// GetDomain returns the Domain field value
func (o *ResultModelDomainCreateData) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ResultModelDomainCreateData) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ResultModelDomainCreateData) SetDomain(v string) {
	o.Domain = v
}

// GetTerm returns the Term field value
func (o *ResultModelDomainCreateData) GetTerm() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Term
}

// GetTermOk returns a tuple with the Term field value
// and a boolean to check if the value has been set.
func (o *ResultModelDomainCreateData) GetTermOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Term, true
}

// SetTerm sets field value
func (o *ResultModelDomainCreateData) SetTerm(v int32) {
	o.Term = v
}

// GetService returns the Service field value
func (o *ResultModelDomainCreateData) GetService() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ResultModelDomainCreateData) GetServiceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ResultModelDomainCreateData) SetService(v int32) {
	o.Service = v
}

// GetTld returns the Tld field value
func (o *ResultModelDomainCreateData) GetTld() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tld
}

// GetTldOk returns a tuple with the Tld field value
// and a boolean to check if the value has been set.
func (o *ResultModelDomainCreateData) GetTldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tld, true
}

// SetTld sets field value
func (o *ResultModelDomainCreateData) SetTld(v string) {
	o.Tld = v
}

// GetInvId returns the InvId field value
func (o *ResultModelDomainCreateData) GetInvId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InvId
}

// GetInvIdOk returns a tuple with the InvId field value
// and a boolean to check if the value has been set.
func (o *ResultModelDomainCreateData) GetInvIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvId, true
}

// SetInvId sets field value
func (o *ResultModelDomainCreateData) SetInvId(v int32) {
	o.InvId = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ResultModelDomainCreateData) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelDomainCreateData) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ResultModelDomainCreateData) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ResultModelDomainCreateData) SetCurrency(v string) {
	o.Currency = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ResultModelDomainCreateData) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelDomainCreateData) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ResultModelDomainCreateData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ResultModelDomainCreateData) SetUser(v string) {
	o.User = &v
}

func (o ResultModelDomainCreateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultModelDomainCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	toSerialize["term"] = o.Term
	toSerialize["service"] = o.Service
	toSerialize["tld"] = o.Tld
	toSerialize["inv_id"] = o.InvId
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

func (o *ResultModelDomainCreateData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"term",
		"service",
		"tld",
		"inv_id",
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

	varResultModelDomainCreateData := _ResultModelDomainCreateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultModelDomainCreateData)

	if err != nil {
		return err
	}

	*o = ResultModelDomainCreateData(varResultModelDomainCreateData)

	return err
}

type NullableResultModelDomainCreateData struct {
	value *ResultModelDomainCreateData
	isSet bool
}

func (v NullableResultModelDomainCreateData) Get() *ResultModelDomainCreateData {
	return v.value
}

func (v *NullableResultModelDomainCreateData) Set(val *ResultModelDomainCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableResultModelDomainCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableResultModelDomainCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultModelDomainCreateData(val *ResultModelDomainCreateData) *NullableResultModelDomainCreateData {
	return &NullableResultModelDomainCreateData{value: val, isSet: true}
}

func (v NullableResultModelDomainCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultModelDomainCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


