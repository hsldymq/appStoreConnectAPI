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

// checks if the BetaGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroup{}

// BetaGroup struct for BetaGroup
type BetaGroup struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *BetaGroupAttributes `json:"attributes,omitempty"`
	Relationships *BetaGroupRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewBetaGroup instantiates a new BetaGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroup(type_ string, id string, links ResourceLinks) *BetaGroup {
	this := BetaGroup{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewBetaGroupWithDefaults instantiates a new BetaGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupWithDefaults() *BetaGroup {
	this := BetaGroup{}
	return &this
}

// GetType returns the Type field value
func (o *BetaGroup) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaGroup) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaGroup) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BetaGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BetaGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BetaGroup) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BetaGroup) GetAttributes() BetaGroupAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BetaGroupAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroup) GetAttributesOk() (*BetaGroupAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BetaGroup) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BetaGroupAttributes and assigns it to the Attributes field.
func (o *BetaGroup) SetAttributes(v BetaGroupAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BetaGroup) GetRelationships() BetaGroupRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BetaGroupRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroup) GetRelationshipsOk() (*BetaGroupRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BetaGroup) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BetaGroupRelationships and assigns it to the Relationships field.
func (o *BetaGroup) SetRelationships(v BetaGroupRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *BetaGroup) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaGroup) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaGroup) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o BetaGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroup) ToMap() (map[string]interface{}, error) {
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

type NullableBetaGroup struct {
	value *BetaGroup
	isSet bool
}

func (v NullableBetaGroup) Get() *BetaGroup {
	return v.value
}

func (v *NullableBetaGroup) Set(val *BetaGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroup(val *BetaGroup) *NullableBetaGroup {
	return &NullableBetaGroup{value: val, isSet: true}
}

func (v NullableBetaGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


