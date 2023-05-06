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

// checks if the AppClip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClip{}

// AppClip struct for AppClip
type AppClip struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppClipAttributes `json:"attributes,omitempty"`
	Relationships *AppClipRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewAppClip instantiates a new AppClip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClip(type_ string, id string, links ResourceLinks) *AppClip {
	this := AppClip{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppClipWithDefaults instantiates a new AppClip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipWithDefaults() *AppClip {
	this := AppClip{}
	return &this
}

// GetType returns the Type field value
func (o *AppClip) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClip) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClip) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClip) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClip) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClip) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClip) GetAttributes() AppClipAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClip) GetAttributesOk() (*AppClipAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClip) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipAttributes and assigns it to the Attributes field.
func (o *AppClip) SetAttributes(v AppClipAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppClip) GetRelationships() AppClipRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppClipRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClip) GetRelationshipsOk() (*AppClipRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppClip) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppClipRelationships and assigns it to the Relationships field.
func (o *AppClip) SetRelationships(v AppClipRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppClip) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClip) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClip) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppClip) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClip) ToMap() (map[string]interface{}, error) {
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

type NullableAppClip struct {
	value *AppClip
	isSet bool
}

func (v NullableAppClip) Get() *AppClip {
	return v.value
}

func (v *NullableAppClip) Set(val *AppClip) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClip) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClip) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClip(val *AppClip) *NullableAppClip {
	return &NullableAppClip{value: val, isSet: true}
}

func (v NullableAppClip) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


