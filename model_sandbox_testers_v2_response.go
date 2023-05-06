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

// checks if the SandboxTestersV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SandboxTestersV2Response{}

// SandboxTestersV2Response struct for SandboxTestersV2Response
type SandboxTestersV2Response struct {
	Data []SandboxTesterV2 `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewSandboxTestersV2Response instantiates a new SandboxTestersV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTestersV2Response(data []SandboxTesterV2, links PagedDocumentLinks) *SandboxTestersV2Response {
	this := SandboxTestersV2Response{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSandboxTestersV2ResponseWithDefaults instantiates a new SandboxTestersV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTestersV2ResponseWithDefaults() *SandboxTestersV2Response {
	this := SandboxTestersV2Response{}
	return &this
}

// GetData returns the Data field value
func (o *SandboxTestersV2Response) GetData() []SandboxTesterV2 {
	if o == nil {
		var ret []SandboxTesterV2
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SandboxTestersV2Response) GetDataOk() ([]SandboxTesterV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SandboxTestersV2Response) SetData(v []SandboxTesterV2) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *SandboxTestersV2Response) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SandboxTestersV2Response) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SandboxTestersV2Response) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SandboxTestersV2Response) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTestersV2Response) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SandboxTestersV2Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SandboxTestersV2Response) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o SandboxTestersV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SandboxTestersV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableSandboxTestersV2Response struct {
	value *SandboxTestersV2Response
	isSet bool
}

func (v NullableSandboxTestersV2Response) Get() *SandboxTestersV2Response {
	return v.value
}

func (v *NullableSandboxTestersV2Response) Set(val *SandboxTestersV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTestersV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTestersV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTestersV2Response(val *SandboxTestersV2Response) *NullableSandboxTestersV2Response {
	return &NullableSandboxTestersV2Response{value: val, isSet: true}
}

func (v NullableSandboxTestersV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTestersV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


