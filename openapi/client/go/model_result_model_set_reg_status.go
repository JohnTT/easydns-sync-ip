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

// checks if the ResultModelSetRegStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultModelSetRegStatus{}

// ResultModelSetRegStatus struct for ResultModelSetRegStatus
type ResultModelSetRegStatus struct {
	Msg *string `json:"msg,omitempty"`
	Tm *int32 `json:"tm,omitempty"`
	Data []ResultModelSetRegStatusItem `json:"data,omitempty"`
	Status *int32 `json:"status,omitempty"`
}

// NewResultModelSetRegStatus instantiates a new ResultModelSetRegStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultModelSetRegStatus() *ResultModelSetRegStatus {
	this := ResultModelSetRegStatus{}
	return &this
}

// NewResultModelSetRegStatusWithDefaults instantiates a new ResultModelSetRegStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultModelSetRegStatusWithDefaults() *ResultModelSetRegStatus {
	this := ResultModelSetRegStatus{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ResultModelSetRegStatus) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelSetRegStatus) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ResultModelSetRegStatus) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ResultModelSetRegStatus) SetMsg(v string) {
	o.Msg = &v
}

// GetTm returns the Tm field value if set, zero value otherwise.
func (o *ResultModelSetRegStatus) GetTm() int32 {
	if o == nil || IsNil(o.Tm) {
		var ret int32
		return ret
	}
	return *o.Tm
}

// GetTmOk returns a tuple with the Tm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelSetRegStatus) GetTmOk() (*int32, bool) {
	if o == nil || IsNil(o.Tm) {
		return nil, false
	}
	return o.Tm, true
}

// HasTm returns a boolean if a field has been set.
func (o *ResultModelSetRegStatus) HasTm() bool {
	if o != nil && !IsNil(o.Tm) {
		return true
	}

	return false
}

// SetTm gets a reference to the given int32 and assigns it to the Tm field.
func (o *ResultModelSetRegStatus) SetTm(v int32) {
	o.Tm = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResultModelSetRegStatus) GetData() []ResultModelSetRegStatusItem {
	if o == nil || IsNil(o.Data) {
		var ret []ResultModelSetRegStatusItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelSetRegStatus) GetDataOk() ([]ResultModelSetRegStatusItem, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResultModelSetRegStatus) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ResultModelSetRegStatusItem and assigns it to the Data field.
func (o *ResultModelSetRegStatus) SetData(v []ResultModelSetRegStatusItem) {
	o.Data = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResultModelSetRegStatus) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelSetRegStatus) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResultModelSetRegStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ResultModelSetRegStatus) SetStatus(v int32) {
	o.Status = &v
}

func (o ResultModelSetRegStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultModelSetRegStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.Tm) {
		toSerialize["tm"] = o.Tm
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableResultModelSetRegStatus struct {
	value *ResultModelSetRegStatus
	isSet bool
}

func (v NullableResultModelSetRegStatus) Get() *ResultModelSetRegStatus {
	return v.value
}

func (v *NullableResultModelSetRegStatus) Set(val *ResultModelSetRegStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableResultModelSetRegStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableResultModelSetRegStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultModelSetRegStatus(val *ResultModelSetRegStatus) *NullableResultModelSetRegStatus {
	return &NullableResultModelSetRegStatus{value: val, isSet: true}
}

func (v NullableResultModelSetRegStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultModelSetRegStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


