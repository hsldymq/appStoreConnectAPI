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

// checks if the CiBuildRunCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildRunCreateRequestData{}

// CiBuildRunCreateRequestData struct for CiBuildRunCreateRequestData
type CiBuildRunCreateRequestData struct {
	Type string `json:"type"`
	Attributes *CiBuildRunCreateRequestDataAttributes `json:"attributes,omitempty"`
	Relationships *CiBuildRunCreateRequestDataRelationships `json:"relationships,omitempty"`
}

// NewCiBuildRunCreateRequestData instantiates a new CiBuildRunCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildRunCreateRequestData(type_ string) *CiBuildRunCreateRequestData {
	this := CiBuildRunCreateRequestData{}
	this.Type = type_
	return &this
}

// NewCiBuildRunCreateRequestDataWithDefaults instantiates a new CiBuildRunCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildRunCreateRequestDataWithDefaults() *CiBuildRunCreateRequestData {
	this := CiBuildRunCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *CiBuildRunCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CiBuildRunCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CiBuildRunCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CiBuildRunCreateRequestData) GetAttributes() CiBuildRunCreateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CiBuildRunCreateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunCreateRequestData) GetAttributesOk() (*CiBuildRunCreateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CiBuildRunCreateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CiBuildRunCreateRequestDataAttributes and assigns it to the Attributes field.
func (o *CiBuildRunCreateRequestData) SetAttributes(v CiBuildRunCreateRequestDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CiBuildRunCreateRequestData) GetRelationships() CiBuildRunCreateRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CiBuildRunCreateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunCreateRequestData) GetRelationshipsOk() (*CiBuildRunCreateRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CiBuildRunCreateRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CiBuildRunCreateRequestDataRelationships and assigns it to the Relationships field.
func (o *CiBuildRunCreateRequestData) SetRelationships(v CiBuildRunCreateRequestDataRelationships) {
	o.Relationships = &v
}

func (o CiBuildRunCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildRunCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableCiBuildRunCreateRequestData struct {
	value *CiBuildRunCreateRequestData
	isSet bool
}

func (v NullableCiBuildRunCreateRequestData) Get() *CiBuildRunCreateRequestData {
	return v.value
}

func (v *NullableCiBuildRunCreateRequestData) Set(val *CiBuildRunCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildRunCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildRunCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildRunCreateRequestData(val *CiBuildRunCreateRequestData) *NullableCiBuildRunCreateRequestData {
	return &NullableCiBuildRunCreateRequestData{value: val, isSet: true}
}

func (v NullableCiBuildRunCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildRunCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


