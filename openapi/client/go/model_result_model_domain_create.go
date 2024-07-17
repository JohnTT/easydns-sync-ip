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

// checks if the ResultModelDomainCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultModelDomainCreate{}

// ResultModelDomainCreate struct for ResultModelDomainCreate
type ResultModelDomainCreate struct {
	Msg string `json:"msg"`
	Data []ResultModelDomainCreateData `json:"data,omitempty"`
	Status string `json:"status"`
}

type _ResultModelDomainCreate ResultModelDomainCreate

// NewResultModelDomainCreate instantiates a new ResultModelDomainCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultModelDomainCreate(msg string, status string) *ResultModelDomainCreate {
	this := ResultModelDomainCreate{}
	this.Msg = msg
	this.Status = status
	return &this
}

// NewResultModelDomainCreateWithDefaults instantiates a new ResultModelDomainCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultModelDomainCreateWithDefaults() *ResultModelDomainCreate {
	this := ResultModelDomainCreate{}
	return &this
}

// GetMsg returns the Msg field value
func (o *ResultModelDomainCreate) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *ResultModelDomainCreate) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *ResultModelDomainCreate) SetMsg(v string) {
	o.Msg = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResultModelDomainCreate) GetData() []ResultModelDomainCreateData {
	if o == nil || IsNil(o.Data) {
		var ret []ResultModelDomainCreateData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelDomainCreate) GetDataOk() ([]ResultModelDomainCreateData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResultModelDomainCreate) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ResultModelDomainCreateData and assigns it to the Data field.
func (o *ResultModelDomainCreate) SetData(v []ResultModelDomainCreateData) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *ResultModelDomainCreate) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResultModelDomainCreate) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResultModelDomainCreate) SetStatus(v string) {
	o.Status = v
}

func (o ResultModelDomainCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultModelDomainCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msg"] = o.Msg
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ResultModelDomainCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"msg",
		"status",
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

	varResultModelDomainCreate := _ResultModelDomainCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultModelDomainCreate)

	if err != nil {
		return err
	}

	*o = ResultModelDomainCreate(varResultModelDomainCreate)

	return err
}

type NullableResultModelDomainCreate struct {
	value *ResultModelDomainCreate
	isSet bool
}

func (v NullableResultModelDomainCreate) Get() *ResultModelDomainCreate {
	return v.value
}

func (v *NullableResultModelDomainCreate) Set(val *ResultModelDomainCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableResultModelDomainCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableResultModelDomainCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultModelDomainCreate(val *ResultModelDomainCreate) *NullableResultModelDomainCreate {
	return &NullableResultModelDomainCreate{value: val, isSet: true}
}

func (v NullableResultModelDomainCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultModelDomainCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

