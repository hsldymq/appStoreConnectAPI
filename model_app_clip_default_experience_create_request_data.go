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

// checks if the AppClipDefaultExperienceCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceCreateRequestData{}

// AppClipDefaultExperienceCreateRequestData struct for AppClipDefaultExperienceCreateRequestData
type AppClipDefaultExperienceCreateRequestData struct {
	Type string `json:"type"`
	Attributes *AppClipDefaultExperienceAttributes `json:"attributes,omitempty"`
	Relationships AppClipDefaultExperienceCreateRequestDataRelationships `json:"relationships"`
}

// NewAppClipDefaultExperienceCreateRequestData instantiates a new AppClipDefaultExperienceCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceCreateRequestData(type_ string, relationships AppClipDefaultExperienceCreateRequestDataRelationships) *AppClipDefaultExperienceCreateRequestData {
	this := AppClipDefaultExperienceCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewAppClipDefaultExperienceCreateRequestDataWithDefaults instantiates a new AppClipDefaultExperienceCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceCreateRequestDataWithDefaults() *AppClipDefaultExperienceCreateRequestData {
	this := AppClipDefaultExperienceCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipDefaultExperienceCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipDefaultExperienceCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceCreateRequestData) GetAttributes() AppClipDefaultExperienceAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipDefaultExperienceAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceCreateRequestData) GetAttributesOk() (*AppClipDefaultExperienceAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceCreateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipDefaultExperienceAttributes and assigns it to the Attributes field.
func (o *AppClipDefaultExperienceCreateRequestData) SetAttributes(v AppClipDefaultExperienceAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value
func (o *AppClipDefaultExperienceCreateRequestData) GetRelationships() AppClipDefaultExperienceCreateRequestDataRelationships {
	if o == nil {
		var ret AppClipDefaultExperienceCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceCreateRequestData) GetRelationshipsOk() (*AppClipDefaultExperienceCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppClipDefaultExperienceCreateRequestData) SetRelationships(v AppClipDefaultExperienceCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppClipDefaultExperienceCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceCreateRequestData struct {
	value *AppClipDefaultExperienceCreateRequestData
	isSet bool
}

func (v NullableAppClipDefaultExperienceCreateRequestData) Get() *AppClipDefaultExperienceCreateRequestData {
	return v.value
}

func (v *NullableAppClipDefaultExperienceCreateRequestData) Set(val *AppClipDefaultExperienceCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceCreateRequestData(val *AppClipDefaultExperienceCreateRequestData) *NullableAppClipDefaultExperienceCreateRequestData {
	return &NullableAppClipDefaultExperienceCreateRequestData{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


