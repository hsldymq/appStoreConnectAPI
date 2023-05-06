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

// checks if the AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion{}

// AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion struct for AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion
type AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData `json:"data,omitempty"`
}

// NewAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion instantiates a new AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion() *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion {
	this := AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion{}
	return &this
}

// NewAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionWithDefaults instantiates a new AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionWithDefaults() *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion {
	this := AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) GetData() AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData {
	if o == nil || IsNil(o.Data) {
		var ret AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) GetDataOk() (*AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData and assigns it to the Data field.
func (o *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) SetData(v AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData) {
	o.Data = &v
}

func (o AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion struct {
	value *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion
	isSet bool
}

func (v NullableAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) Get() *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion {
	return v.value
}

func (v *NullableAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) Set(val *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion(val *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) *NullableAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion {
	return &NullableAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


