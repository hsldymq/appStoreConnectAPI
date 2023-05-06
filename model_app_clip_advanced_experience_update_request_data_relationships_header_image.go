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

// checks if the AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage{}

// AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage struct for AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage
type AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage struct {
	Data *AppClipAdvancedExperienceRelationshipsHeaderImageData `json:"data,omitempty"`
}

// NewAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage instantiates a new AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage() *AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage {
	this := AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage{}
	return &this
}

// NewAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImageWithDefaults instantiates a new AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImageWithDefaults() *AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage {
	this := AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) GetData() AppClipAdvancedExperienceRelationshipsHeaderImageData {
	if o == nil || IsNil(o.Data) {
		var ret AppClipAdvancedExperienceRelationshipsHeaderImageData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) GetDataOk() (*AppClipAdvancedExperienceRelationshipsHeaderImageData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppClipAdvancedExperienceRelationshipsHeaderImageData and assigns it to the Data field.
func (o *AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) SetData(v AppClipAdvancedExperienceRelationshipsHeaderImageData) {
	o.Data = &v
}

func (o AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage struct {
	value *AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage
	isSet bool
}

func (v NullableAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) Get() *AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) Set(val *AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage(val *AppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) *NullableAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage {
	return &NullableAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceUpdateRequestDataRelationshipsHeaderImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


