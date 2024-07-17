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

// checks if the ResultModelMailmapCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultModelMailmapCreate{}

// ResultModelMailmapCreate struct for ResultModelMailmapCreate
type ResultModelMailmapCreate struct {
	Status int32 `json:"status"`
	// The unix timestamp of the request
	Tm int32 `json:"tm"`
	// Processing result message related to response
	Msg string `json:"msg"`
	Data ResultModelMailmapCreateData `json:"data"`
}

type _ResultModelMailmapCreate ResultModelMailmapCreate

// NewResultModelMailmapCreate instantiates a new ResultModelMailmapCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultModelMailmapCreate(status int32, tm int32, msg string, data ResultModelMailmapCreateData) *ResultModelMailmapCreate {
	this := ResultModelMailmapCreate{}
	this.Status = status
	this.Tm = tm
	this.Msg = msg
	this.Data = data
	return &this
}

// NewResultModelMailmapCreateWithDefaults instantiates a new ResultModelMailmapCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultModelMailmapCreateWithDefaults() *ResultModelMailmapCreate {
	this := ResultModelMailmapCreate{}
	return &this
}

// GetStatus returns the Status field value
func (o *ResultModelMailmapCreate) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResultModelMailmapCreate) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResultModelMailmapCreate) SetStatus(v int32) {
	o.Status = v
}

// GetTm returns the Tm field value
func (o *ResultModelMailmapCreate) GetTm() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Tm
}

// GetTmOk returns a tuple with the Tm field value
// and a boolean to check if the value has been set.
func (o *ResultModelMailmapCreate) GetTmOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tm, true
}

// SetTm sets field value
func (o *ResultModelMailmapCreate) SetTm(v int32) {
	o.Tm = v
}

// GetMsg returns the Msg field value
func (o *ResultModelMailmapCreate) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *ResultModelMailmapCreate) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *ResultModelMailmapCreate) SetMsg(v string) {
	o.Msg = v
}

// GetData returns the Data field value
func (o *ResultModelMailmapCreate) GetData() ResultModelMailmapCreateData {
	if o == nil {
		var ret ResultModelMailmapCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResultModelMailmapCreate) GetDataOk() (*ResultModelMailmapCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResultModelMailmapCreate) SetData(v ResultModelMailmapCreateData) {
	o.Data = v
}

func (o ResultModelMailmapCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultModelMailmapCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["tm"] = o.Tm
	toSerialize["msg"] = o.Msg
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResultModelMailmapCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"tm",
		"msg",
		"data",
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

	varResultModelMailmapCreate := _ResultModelMailmapCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultModelMailmapCreate)

	if err != nil {
		return err
	}

	*o = ResultModelMailmapCreate(varResultModelMailmapCreate)

	return err
}

type NullableResultModelMailmapCreate struct {
	value *ResultModelMailmapCreate
	isSet bool
}

func (v NullableResultModelMailmapCreate) Get() *ResultModelMailmapCreate {
	return v.value
}

func (v *NullableResultModelMailmapCreate) Set(val *ResultModelMailmapCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableResultModelMailmapCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableResultModelMailmapCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultModelMailmapCreate(val *ResultModelMailmapCreate) *NullableResultModelMailmapCreate {
	return &NullableResultModelMailmapCreate{value: val, isSet: true}
}

func (v NullableResultModelMailmapCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultModelMailmapCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


