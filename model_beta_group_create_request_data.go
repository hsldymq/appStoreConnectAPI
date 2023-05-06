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

// checks if the BetaGroupCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupCreateRequestData{}

// BetaGroupCreateRequestData struct for BetaGroupCreateRequestData
type BetaGroupCreateRequestData struct {
	Type string `json:"type"`
	Attributes BetaGroupCreateRequestDataAttributes `json:"attributes"`
	Relationships BetaGroupCreateRequestDataRelationships `json:"relationships"`
}

// NewBetaGroupCreateRequestData instantiates a new BetaGroupCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupCreateRequestData(type_ string, attributes BetaGroupCreateRequestDataAttributes, relationships BetaGroupCreateRequestDataRelationships) *BetaGroupCreateRequestData {
	this := BetaGroupCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewBetaGroupCreateRequestDataWithDefaults instantiates a new BetaGroupCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupCreateRequestDataWithDefaults() *BetaGroupCreateRequestData {
	this := BetaGroupCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BetaGroupCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaGroupCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BetaGroupCreateRequestData) GetAttributes() BetaGroupCreateRequestDataAttributes {
	if o == nil {
		var ret BetaGroupCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestData) GetAttributesOk() (*BetaGroupCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BetaGroupCreateRequestData) SetAttributes(v BetaGroupCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *BetaGroupCreateRequestData) GetRelationships() BetaGroupCreateRequestDataRelationships {
	if o == nil {
		var ret BetaGroupCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestData) GetRelationshipsOk() (*BetaGroupCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *BetaGroupCreateRequestData) SetRelationships(v BetaGroupCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o BetaGroupCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableBetaGroupCreateRequestData struct {
	value *BetaGroupCreateRequestData
	isSet bool
}

func (v NullableBetaGroupCreateRequestData) Get() *BetaGroupCreateRequestData {
	return v.value
}

func (v *NullableBetaGroupCreateRequestData) Set(val *BetaGroupCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupCreateRequestData(val *BetaGroupCreateRequestData) *NullableBetaGroupCreateRequestData {
	return &NullableBetaGroupCreateRequestData{value: val, isSet: true}
}

func (v NullableBetaGroupCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


