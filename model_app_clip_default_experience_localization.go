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

// checks if the AppClipDefaultExperienceLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceLocalization{}

// AppClipDefaultExperienceLocalization struct for AppClipDefaultExperienceLocalization
type AppClipDefaultExperienceLocalization struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppClipDefaultExperienceLocalizationAttributes `json:"attributes,omitempty"`
	Relationships *AppClipDefaultExperienceLocalizationRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewAppClipDefaultExperienceLocalization instantiates a new AppClipDefaultExperienceLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceLocalization(type_ string, id string, links ResourceLinks) *AppClipDefaultExperienceLocalization {
	this := AppClipDefaultExperienceLocalization{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppClipDefaultExperienceLocalizationWithDefaults instantiates a new AppClipDefaultExperienceLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceLocalizationWithDefaults() *AppClipDefaultExperienceLocalization {
	this := AppClipDefaultExperienceLocalization{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipDefaultExperienceLocalization) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalization) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipDefaultExperienceLocalization) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipDefaultExperienceLocalization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipDefaultExperienceLocalization) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceLocalization) GetAttributes() AppClipDefaultExperienceLocalizationAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipDefaultExperienceLocalizationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalization) GetAttributesOk() (*AppClipDefaultExperienceLocalizationAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceLocalization) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipDefaultExperienceLocalizationAttributes and assigns it to the Attributes field.
func (o *AppClipDefaultExperienceLocalization) SetAttributes(v AppClipDefaultExperienceLocalizationAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceLocalization) GetRelationships() AppClipDefaultExperienceLocalizationRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppClipDefaultExperienceLocalizationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalization) GetRelationshipsOk() (*AppClipDefaultExperienceLocalizationRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceLocalization) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppClipDefaultExperienceLocalizationRelationships and assigns it to the Relationships field.
func (o *AppClipDefaultExperienceLocalization) SetRelationships(v AppClipDefaultExperienceLocalizationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppClipDefaultExperienceLocalization) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceLocalization) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClipDefaultExperienceLocalization) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppClipDefaultExperienceLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceLocalization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceLocalization struct {
	value *AppClipDefaultExperienceLocalization
	isSet bool
}

func (v NullableAppClipDefaultExperienceLocalization) Get() *AppClipDefaultExperienceLocalization {
	return v.value
}

func (v *NullableAppClipDefaultExperienceLocalization) Set(val *AppClipDefaultExperienceLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceLocalization(val *AppClipDefaultExperienceLocalization) *NullableAppClipDefaultExperienceLocalization {
	return &NullableAppClipDefaultExperienceLocalization{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


