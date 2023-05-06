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

// checks if the BetaTesterCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterCreateRequestData{}

// BetaTesterCreateRequestData struct for BetaTesterCreateRequestData
type BetaTesterCreateRequestData struct {
	Type string `json:"type"`
	Attributes BetaTesterCreateRequestDataAttributes `json:"attributes"`
	Relationships *BetaTesterCreateRequestDataRelationships `json:"relationships,omitempty"`
}

// NewBetaTesterCreateRequestData instantiates a new BetaTesterCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterCreateRequestData(type_ string, attributes BetaTesterCreateRequestDataAttributes) *BetaTesterCreateRequestData {
	this := BetaTesterCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBetaTesterCreateRequestDataWithDefaults instantiates a new BetaTesterCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterCreateRequestDataWithDefaults() *BetaTesterCreateRequestData {
	this := BetaTesterCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BetaTesterCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaTesterCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaTesterCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BetaTesterCreateRequestData) GetAttributes() BetaTesterCreateRequestDataAttributes {
	if o == nil {
		var ret BetaTesterCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BetaTesterCreateRequestData) GetAttributesOk() (*BetaTesterCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BetaTesterCreateRequestData) SetAttributes(v BetaTesterCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BetaTesterCreateRequestData) GetRelationships() BetaTesterCreateRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BetaTesterCreateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterCreateRequestData) GetRelationshipsOk() (*BetaTesterCreateRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BetaTesterCreateRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BetaTesterCreateRequestDataRelationships and assigns it to the Relationships field.
func (o *BetaTesterCreateRequestData) SetRelationships(v BetaTesterCreateRequestDataRelationships) {
	o.Relationships = &v
}

func (o BetaTesterCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableBetaTesterCreateRequestData struct {
	value *BetaTesterCreateRequestData
	isSet bool
}

func (v NullableBetaTesterCreateRequestData) Get() *BetaTesterCreateRequestData {
	return v.value
}

func (v *NullableBetaTesterCreateRequestData) Set(val *BetaTesterCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterCreateRequestData(val *BetaTesterCreateRequestData) *NullableBetaTesterCreateRequestData {
	return &NullableBetaTesterCreateRequestData{value: val, isSet: true}
}

func (v NullableBetaTesterCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


