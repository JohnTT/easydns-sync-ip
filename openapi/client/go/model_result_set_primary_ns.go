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

// checks if the ResultSetPrimaryNS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultSetPrimaryNS{}

// ResultSetPrimaryNS struct for ResultSetPrimaryNS
type ResultSetPrimaryNS struct {
	// Result message of request
	Msg string `json:"msg"`
	Data ResultSetPrimaryNSData `json:"data"`
	// HTTP status code of result
	Status int32 `json:"status"`
}

type _ResultSetPrimaryNS ResultSetPrimaryNS

// NewResultSetPrimaryNS instantiates a new ResultSetPrimaryNS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultSetPrimaryNS(msg string, data ResultSetPrimaryNSData, status int32) *ResultSetPrimaryNS {
	this := ResultSetPrimaryNS{}
	this.Msg = msg
	this.Data = data
	this.Status = status
	return &this
}

// NewResultSetPrimaryNSWithDefaults instantiates a new ResultSetPrimaryNS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultSetPrimaryNSWithDefaults() *ResultSetPrimaryNS {
	this := ResultSetPrimaryNS{}
	var status int32 = 200
	this.Status = status
	return &this
}

// GetMsg returns the Msg field value
func (o *ResultSetPrimaryNS) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *ResultSetPrimaryNS) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *ResultSetPrimaryNS) SetMsg(v string) {
	o.Msg = v
}

// GetData returns the Data field value
func (o *ResultSetPrimaryNS) GetData() ResultSetPrimaryNSData {
	if o == nil {
		var ret ResultSetPrimaryNSData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResultSetPrimaryNS) GetDataOk() (*ResultSetPrimaryNSData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResultSetPrimaryNS) SetData(v ResultSetPrimaryNSData) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *ResultSetPrimaryNS) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResultSetPrimaryNS) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResultSetPrimaryNS) SetStatus(v int32) {
	o.Status = v
}

func (o ResultSetPrimaryNS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultSetPrimaryNS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msg"] = o.Msg
	toSerialize["data"] = o.Data
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ResultSetPrimaryNS) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"msg",
		"data",
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

	varResultSetPrimaryNS := _ResultSetPrimaryNS{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultSetPrimaryNS)

	if err != nil {
		return err
	}

	*o = ResultSetPrimaryNS(varResultSetPrimaryNS)

	return err
}

type NullableResultSetPrimaryNS struct {
	value *ResultSetPrimaryNS
	isSet bool
}

func (v NullableResultSetPrimaryNS) Get() *ResultSetPrimaryNS {
	return v.value
}

func (v *NullableResultSetPrimaryNS) Set(val *ResultSetPrimaryNS) {
	v.value = val
	v.isSet = true
}

func (v NullableResultSetPrimaryNS) IsSet() bool {
	return v.isSet
}

func (v *NullableResultSetPrimaryNS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultSetPrimaryNS(val *ResultSetPrimaryNS) *NullableResultSetPrimaryNS {
	return &NullableResultSetPrimaryNS{value: val, isSet: true}
}

func (v NullableResultSetPrimaryNS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultSetPrimaryNS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


