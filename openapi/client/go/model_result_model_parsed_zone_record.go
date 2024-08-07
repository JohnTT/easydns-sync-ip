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

// checks if the ResultModelParsedZoneRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultModelParsedZoneRecord{}

// ResultModelParsedZoneRecord struct for ResultModelParsedZoneRecord
type ResultModelParsedZoneRecord struct {
	// Zone record unique ID. Records that parse to multiple zone records will share the same ID!
	Id int32 `json:"id"`
	// The domain name that contains the zone records
	Domain string `json:"domain"`
	// The hostname of the record (without the domain name). '@' indicates the record is for the root of the zone.
	Host string `json:"host"`
	// (OPTIONAL) The Time-to-live (TTL) of the zone record. If not returned, assume the default.
	Ttl *int32 `json:"ttl,omitempty"`
	// (OPTIONAL) The priority of the record. This is only required for records that include the prio field (such as MX).
	Prio *int32 `json:"prio,omitempty"`
	// (OPTIONAL) The ID of the geolocation this record applies to.
	GeozoneId *int32 `json:"geozone_id,omitempty"`
	// The zone record type of the record. PLEASE NOTE: Some record types are specific to easyDNS and may require translation on other providers.
	Type string `json:"type"`
	// The contents of the rdata for a zone record.
	Rdata string `json:"rdata"`
	// The time this zone record was last updated.
	LastMod *string `json:"last_mod,omitempty"`
	// The original URL that a record uses as a target. This applies to custom records like URL/STEALTH.
	Url *string `json:"url,omitempty"`
	// The data originally stored in the rdata field for custom types and records containing keywords
	OrigRdata *string `json:"orig_rdata,omitempty"`
}

type _ResultModelParsedZoneRecord ResultModelParsedZoneRecord

// NewResultModelParsedZoneRecord instantiates a new ResultModelParsedZoneRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultModelParsedZoneRecord(id int32, domain string, host string, type_ string, rdata string) *ResultModelParsedZoneRecord {
	this := ResultModelParsedZoneRecord{}
	this.Id = id
	this.Domain = domain
	this.Host = host
	var ttl int32 = 600
	this.Ttl = &ttl
	var prio int32 = 0
	this.Prio = &prio
	var geozoneId int32 = 0
	this.GeozoneId = &geozoneId
	this.Type = type_
	this.Rdata = rdata
	return &this
}

// NewResultModelParsedZoneRecordWithDefaults instantiates a new ResultModelParsedZoneRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultModelParsedZoneRecordWithDefaults() *ResultModelParsedZoneRecord {
	this := ResultModelParsedZoneRecord{}
	var ttl int32 = 600
	this.Ttl = &ttl
	var prio int32 = 0
	this.Prio = &prio
	var geozoneId int32 = 0
	this.GeozoneId = &geozoneId
	return &this
}

// GetId returns the Id field value
func (o *ResultModelParsedZoneRecord) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResultModelParsedZoneRecord) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResultModelParsedZoneRecord) SetId(v int32) {
	o.Id = v
}

// GetDomain returns the Domain field value
func (o *ResultModelParsedZoneRecord) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *ResultModelParsedZoneRecord) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *ResultModelParsedZoneRecord) SetDomain(v string) {
	o.Domain = v
}

// GetHost returns the Host field value
func (o *ResultModelParsedZoneRecord) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *ResultModelParsedZoneRecord) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *ResultModelParsedZoneRecord) SetHost(v string) {
	o.Host = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ResultModelParsedZoneRecord) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelParsedZoneRecord) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ResultModelParsedZoneRecord) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *ResultModelParsedZoneRecord) SetTtl(v int32) {
	o.Ttl = &v
}

// GetPrio returns the Prio field value if set, zero value otherwise.
func (o *ResultModelParsedZoneRecord) GetPrio() int32 {
	if o == nil || IsNil(o.Prio) {
		var ret int32
		return ret
	}
	return *o.Prio
}

// GetPrioOk returns a tuple with the Prio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelParsedZoneRecord) GetPrioOk() (*int32, bool) {
	if o == nil || IsNil(o.Prio) {
		return nil, false
	}
	return o.Prio, true
}

// HasPrio returns a boolean if a field has been set.
func (o *ResultModelParsedZoneRecord) HasPrio() bool {
	if o != nil && !IsNil(o.Prio) {
		return true
	}

	return false
}

// SetPrio gets a reference to the given int32 and assigns it to the Prio field.
func (o *ResultModelParsedZoneRecord) SetPrio(v int32) {
	o.Prio = &v
}

// GetGeozoneId returns the GeozoneId field value if set, zero value otherwise.
func (o *ResultModelParsedZoneRecord) GetGeozoneId() int32 {
	if o == nil || IsNil(o.GeozoneId) {
		var ret int32
		return ret
	}
	return *o.GeozoneId
}

// GetGeozoneIdOk returns a tuple with the GeozoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelParsedZoneRecord) GetGeozoneIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GeozoneId) {
		return nil, false
	}
	return o.GeozoneId, true
}

// HasGeozoneId returns a boolean if a field has been set.
func (o *ResultModelParsedZoneRecord) HasGeozoneId() bool {
	if o != nil && !IsNil(o.GeozoneId) {
		return true
	}

	return false
}

// SetGeozoneId gets a reference to the given int32 and assigns it to the GeozoneId field.
func (o *ResultModelParsedZoneRecord) SetGeozoneId(v int32) {
	o.GeozoneId = &v
}

// GetType returns the Type field value
func (o *ResultModelParsedZoneRecord) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResultModelParsedZoneRecord) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResultModelParsedZoneRecord) SetType(v string) {
	o.Type = v
}

// GetRdata returns the Rdata field value
func (o *ResultModelParsedZoneRecord) GetRdata() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rdata
}

// GetRdataOk returns a tuple with the Rdata field value
// and a boolean to check if the value has been set.
func (o *ResultModelParsedZoneRecord) GetRdataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rdata, true
}

// SetRdata sets field value
func (o *ResultModelParsedZoneRecord) SetRdata(v string) {
	o.Rdata = v
}

// GetLastMod returns the LastMod field value if set, zero value otherwise.
func (o *ResultModelParsedZoneRecord) GetLastMod() string {
	if o == nil || IsNil(o.LastMod) {
		var ret string
		return ret
	}
	return *o.LastMod
}

// GetLastModOk returns a tuple with the LastMod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelParsedZoneRecord) GetLastModOk() (*string, bool) {
	if o == nil || IsNil(o.LastMod) {
		return nil, false
	}
	return o.LastMod, true
}

// HasLastMod returns a boolean if a field has been set.
func (o *ResultModelParsedZoneRecord) HasLastMod() bool {
	if o != nil && !IsNil(o.LastMod) {
		return true
	}

	return false
}

// SetLastMod gets a reference to the given string and assigns it to the LastMod field.
func (o *ResultModelParsedZoneRecord) SetLastMod(v string) {
	o.LastMod = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ResultModelParsedZoneRecord) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelParsedZoneRecord) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ResultModelParsedZoneRecord) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ResultModelParsedZoneRecord) SetUrl(v string) {
	o.Url = &v
}

// GetOrigRdata returns the OrigRdata field value if set, zero value otherwise.
func (o *ResultModelParsedZoneRecord) GetOrigRdata() string {
	if o == nil || IsNil(o.OrigRdata) {
		var ret string
		return ret
	}
	return *o.OrigRdata
}

// GetOrigRdataOk returns a tuple with the OrigRdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultModelParsedZoneRecord) GetOrigRdataOk() (*string, bool) {
	if o == nil || IsNil(o.OrigRdata) {
		return nil, false
	}
	return o.OrigRdata, true
}

// HasOrigRdata returns a boolean if a field has been set.
func (o *ResultModelParsedZoneRecord) HasOrigRdata() bool {
	if o != nil && !IsNil(o.OrigRdata) {
		return true
	}

	return false
}

// SetOrigRdata gets a reference to the given string and assigns it to the OrigRdata field.
func (o *ResultModelParsedZoneRecord) SetOrigRdata(v string) {
	o.OrigRdata = &v
}

func (o ResultModelParsedZoneRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultModelParsedZoneRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["domain"] = o.Domain
	toSerialize["host"] = o.Host
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.Prio) {
		toSerialize["prio"] = o.Prio
	}
	if !IsNil(o.GeozoneId) {
		toSerialize["geozone_id"] = o.GeozoneId
	}
	toSerialize["type"] = o.Type
	toSerialize["rdata"] = o.Rdata
	if !IsNil(o.LastMod) {
		toSerialize["last_mod"] = o.LastMod
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.OrigRdata) {
		toSerialize["orig_rdata"] = o.OrigRdata
	}
	return toSerialize, nil
}

func (o *ResultModelParsedZoneRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"domain",
		"host",
		"type",
		"rdata",
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

	varResultModelParsedZoneRecord := _ResultModelParsedZoneRecord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultModelParsedZoneRecord)

	if err != nil {
		return err
	}

	*o = ResultModelParsedZoneRecord(varResultModelParsedZoneRecord)

	return err
}

type NullableResultModelParsedZoneRecord struct {
	value *ResultModelParsedZoneRecord
	isSet bool
}

func (v NullableResultModelParsedZoneRecord) Get() *ResultModelParsedZoneRecord {
	return v.value
}

func (v *NullableResultModelParsedZoneRecord) Set(val *ResultModelParsedZoneRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableResultModelParsedZoneRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableResultModelParsedZoneRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultModelParsedZoneRecord(val *ResultModelParsedZoneRecord) *NullableResultModelParsedZoneRecord {
	return &NullableResultModelParsedZoneRecord{value: val, isSet: true}
}

func (v NullableResultModelParsedZoneRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultModelParsedZoneRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


