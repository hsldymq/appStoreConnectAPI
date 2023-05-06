/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SandboxTesterV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SandboxTesterV2Response{}

// SandboxTesterV2Response struct for SandboxTesterV2Response
type SandboxTesterV2Response struct {
	Data SandboxTesterV2 `json:"data"`
	Links DocumentLinks `json:"links"`
}

// NewSandboxTesterV2Response instantiates a new SandboxTesterV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTesterV2Response(data SandboxTesterV2, links DocumentLinks) *SandboxTesterV2Response {
	this := SandboxTesterV2Response{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSandboxTesterV2ResponseWithDefaults instantiates a new SandboxTesterV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTesterV2ResponseWithDefaults() *SandboxTesterV2Response {
	this := SandboxTesterV2Response{}
	return &this
}

// GetData returns the Data field value
func (o *SandboxTesterV2Response) GetData() SandboxTesterV2 {
	if o == nil {
		var ret SandboxTesterV2
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SandboxTesterV2Response) GetDataOk() (*SandboxTesterV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SandboxTesterV2Response) SetData(v SandboxTesterV2) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *SandboxTesterV2Response) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SandboxTesterV2Response) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SandboxTesterV2Response) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o SandboxTesterV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SandboxTesterV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableSandboxTesterV2Response struct {
	value *SandboxTesterV2Response
	isSet bool
}

func (v NullableSandboxTesterV2Response) Get() *SandboxTesterV2Response {
	return v.value
}

func (v *NullableSandboxTesterV2Response) Set(val *SandboxTesterV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTesterV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTesterV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTesterV2Response(val *SandboxTesterV2Response) *NullableSandboxTesterV2Response {
	return &NullableSandboxTesterV2Response{value: val, isSet: true}
}

func (v NullableSandboxTesterV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTesterV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


