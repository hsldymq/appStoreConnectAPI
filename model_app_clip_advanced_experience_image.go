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

// checks if the AppClipAdvancedExperienceImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceImage{}

// AppClipAdvancedExperienceImage struct for AppClipAdvancedExperienceImage
type AppClipAdvancedExperienceImage struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppClipAdvancedExperienceImageAttributes `json:"attributes,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewAppClipAdvancedExperienceImage instantiates a new AppClipAdvancedExperienceImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceImage(type_ string, id string, links ResourceLinks) *AppClipAdvancedExperienceImage {
	this := AppClipAdvancedExperienceImage{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppClipAdvancedExperienceImageWithDefaults instantiates a new AppClipAdvancedExperienceImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceImageWithDefaults() *AppClipAdvancedExperienceImage {
	this := AppClipAdvancedExperienceImage{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipAdvancedExperienceImage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipAdvancedExperienceImage) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipAdvancedExperienceImage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipAdvancedExperienceImage) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceImage) GetAttributes() AppClipAdvancedExperienceImageAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipAdvancedExperienceImageAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImage) GetAttributesOk() (*AppClipAdvancedExperienceImageAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceImage) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipAdvancedExperienceImageAttributes and assigns it to the Attributes field.
func (o *AppClipAdvancedExperienceImage) SetAttributes(v AppClipAdvancedExperienceImageAttributes) {
	o.Attributes = &v
}

// GetLinks returns the Links field value
func (o *AppClipAdvancedExperienceImage) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImage) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClipAdvancedExperienceImage) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppClipAdvancedExperienceImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceImage struct {
	value *AppClipAdvancedExperienceImage
	isSet bool
}

func (v NullableAppClipAdvancedExperienceImage) Get() *AppClipAdvancedExperienceImage {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceImage) Set(val *AppClipAdvancedExperienceImage) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceImage) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceImage(val *AppClipAdvancedExperienceImage) *NullableAppClipAdvancedExperienceImage {
	return &NullableAppClipAdvancedExperienceImage{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


