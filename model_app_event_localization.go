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

// checks if the AppEventLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalization{}

// AppEventLocalization struct for AppEventLocalization
type AppEventLocalization struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppEventLocalizationAttributes `json:"attributes,omitempty"`
	Relationships *AppEventLocalizationRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewAppEventLocalization instantiates a new AppEventLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalization(type_ string, id string, links ResourceLinks) *AppEventLocalization {
	this := AppEventLocalization{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppEventLocalizationWithDefaults instantiates a new AppEventLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationWithDefaults() *AppEventLocalization {
	this := AppEventLocalization{}
	return &this
}

// GetType returns the Type field value
func (o *AppEventLocalization) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalization) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEventLocalization) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppEventLocalization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppEventLocalization) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppEventLocalization) GetAttributes() AppEventLocalizationAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppEventLocalizationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalization) GetAttributesOk() (*AppEventLocalizationAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppEventLocalization) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppEventLocalizationAttributes and assigns it to the Attributes field.
func (o *AppEventLocalization) SetAttributes(v AppEventLocalizationAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppEventLocalization) GetRelationships() AppEventLocalizationRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppEventLocalizationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalization) GetRelationshipsOk() (*AppEventLocalizationRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppEventLocalization) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppEventLocalizationRelationships and assigns it to the Relationships field.
func (o *AppEventLocalization) SetRelationships(v AppEventLocalizationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppEventLocalization) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalization) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppEventLocalization) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppEventLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalization) ToMap() (map[string]interface{}, error) {
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

type NullableAppEventLocalization struct {
	value *AppEventLocalization
	isSet bool
}

func (v NullableAppEventLocalization) Get() *AppEventLocalization {
	return v.value
}

func (v *NullableAppEventLocalization) Set(val *AppEventLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalization(val *AppEventLocalization) *NullableAppEventLocalization {
	return &NullableAppEventLocalization{value: val, isSet: true}
}

func (v NullableAppEventLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


