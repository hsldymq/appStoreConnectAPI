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

// checks if the AppEventScreenshotCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventScreenshotCreateRequestData{}

// AppEventScreenshotCreateRequestData struct for AppEventScreenshotCreateRequestData
type AppEventScreenshotCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppEventScreenshotCreateRequestDataAttributes `json:"attributes"`
	Relationships AppEventScreenshotCreateRequestDataRelationships `json:"relationships"`
}

// NewAppEventScreenshotCreateRequestData instantiates a new AppEventScreenshotCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventScreenshotCreateRequestData(type_ string, attributes AppEventScreenshotCreateRequestDataAttributes, relationships AppEventScreenshotCreateRequestDataRelationships) *AppEventScreenshotCreateRequestData {
	this := AppEventScreenshotCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppEventScreenshotCreateRequestDataWithDefaults instantiates a new AppEventScreenshotCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventScreenshotCreateRequestDataWithDefaults() *AppEventScreenshotCreateRequestData {
	this := AppEventScreenshotCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppEventScreenshotCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEventScreenshotCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppEventScreenshotCreateRequestData) GetAttributes() AppEventScreenshotCreateRequestDataAttributes {
	if o == nil {
		var ret AppEventScreenshotCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotCreateRequestData) GetAttributesOk() (*AppEventScreenshotCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppEventScreenshotCreateRequestData) SetAttributes(v AppEventScreenshotCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppEventScreenshotCreateRequestData) GetRelationships() AppEventScreenshotCreateRequestDataRelationships {
	if o == nil {
		var ret AppEventScreenshotCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotCreateRequestData) GetRelationshipsOk() (*AppEventScreenshotCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppEventScreenshotCreateRequestData) SetRelationships(v AppEventScreenshotCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppEventScreenshotCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventScreenshotCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppEventScreenshotCreateRequestData struct {
	value *AppEventScreenshotCreateRequestData
	isSet bool
}

func (v NullableAppEventScreenshotCreateRequestData) Get() *AppEventScreenshotCreateRequestData {
	return v.value
}

func (v *NullableAppEventScreenshotCreateRequestData) Set(val *AppEventScreenshotCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventScreenshotCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventScreenshotCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventScreenshotCreateRequestData(val *AppEventScreenshotCreateRequestData) *NullableAppEventScreenshotCreateRequestData {
	return &NullableAppEventScreenshotCreateRequestData{value: val, isSet: true}
}

func (v NullableAppEventScreenshotCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventScreenshotCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

