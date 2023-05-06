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

// checks if the AppRelationshipsGameCenterEnabledVersions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsGameCenterEnabledVersions{}

// AppRelationshipsGameCenterEnabledVersions struct for AppRelationshipsGameCenterEnabledVersions
type AppRelationshipsGameCenterEnabledVersions struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []AppRelationshipsGameCenterEnabledVersionsDataInner `json:"data,omitempty"`
}

// NewAppRelationshipsGameCenterEnabledVersions instantiates a new AppRelationshipsGameCenterEnabledVersions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsGameCenterEnabledVersions() *AppRelationshipsGameCenterEnabledVersions {
	this := AppRelationshipsGameCenterEnabledVersions{}
	return &this
}

// NewAppRelationshipsGameCenterEnabledVersionsWithDefaults instantiates a new AppRelationshipsGameCenterEnabledVersions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsGameCenterEnabledVersionsWithDefaults() *AppRelationshipsGameCenterEnabledVersions {
	this := AppRelationshipsGameCenterEnabledVersions{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppRelationshipsGameCenterEnabledVersions) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsGameCenterEnabledVersions) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppRelationshipsGameCenterEnabledVersions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *AppRelationshipsGameCenterEnabledVersions) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppRelationshipsGameCenterEnabledVersions) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsGameCenterEnabledVersions) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppRelationshipsGameCenterEnabledVersions) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppRelationshipsGameCenterEnabledVersions) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppRelationshipsGameCenterEnabledVersions) GetData() []AppRelationshipsGameCenterEnabledVersionsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []AppRelationshipsGameCenterEnabledVersionsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsGameCenterEnabledVersions) GetDataOk() ([]AppRelationshipsGameCenterEnabledVersionsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppRelationshipsGameCenterEnabledVersions) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppRelationshipsGameCenterEnabledVersionsDataInner and assigns it to the Data field.
func (o *AppRelationshipsGameCenterEnabledVersions) SetData(v []AppRelationshipsGameCenterEnabledVersionsDataInner) {
	o.Data = v
}

func (o AppRelationshipsGameCenterEnabledVersions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsGameCenterEnabledVersions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppRelationshipsGameCenterEnabledVersions struct {
	value *AppRelationshipsGameCenterEnabledVersions
	isSet bool
}

func (v NullableAppRelationshipsGameCenterEnabledVersions) Get() *AppRelationshipsGameCenterEnabledVersions {
	return v.value
}

func (v *NullableAppRelationshipsGameCenterEnabledVersions) Set(val *AppRelationshipsGameCenterEnabledVersions) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsGameCenterEnabledVersions) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsGameCenterEnabledVersions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsGameCenterEnabledVersions(val *AppRelationshipsGameCenterEnabledVersions) *NullableAppRelationshipsGameCenterEnabledVersions {
	return &NullableAppRelationshipsGameCenterEnabledVersions{value: val, isSet: true}
}

func (v NullableAppRelationshipsGameCenterEnabledVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsGameCenterEnabledVersions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


