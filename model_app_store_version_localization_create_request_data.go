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

// checks if the AppStoreVersionLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionLocalizationCreateRequestData{}

// AppStoreVersionLocalizationCreateRequestData struct for AppStoreVersionLocalizationCreateRequestData
type AppStoreVersionLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppStoreVersionLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships AppStoreReviewDetailCreateRequestDataRelationships `json:"relationships"`
}

// NewAppStoreVersionLocalizationCreateRequestData instantiates a new AppStoreVersionLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationCreateRequestData(type_ string, attributes AppStoreVersionLocalizationCreateRequestDataAttributes, relationships AppStoreReviewDetailCreateRequestDataRelationships) *AppStoreVersionLocalizationCreateRequestData {
	this := AppStoreVersionLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppStoreVersionLocalizationCreateRequestDataWithDefaults instantiates a new AppStoreVersionLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationCreateRequestDataWithDefaults() *AppStoreVersionLocalizationCreateRequestData {
	this := AppStoreVersionLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppStoreVersionLocalizationCreateRequestData) GetAttributes() AppStoreVersionLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret AppStoreVersionLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestData) GetAttributesOk() (*AppStoreVersionLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppStoreVersionLocalizationCreateRequestData) SetAttributes(v AppStoreVersionLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppStoreVersionLocalizationCreateRequestData) GetRelationships() AppStoreReviewDetailCreateRequestDataRelationships {
	if o == nil {
		var ret AppStoreReviewDetailCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationCreateRequestData) GetRelationshipsOk() (*AppStoreReviewDetailCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppStoreVersionLocalizationCreateRequestData) SetRelationships(v AppStoreReviewDetailCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppStoreVersionLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppStoreVersionLocalizationCreateRequestData struct {
	value *AppStoreVersionLocalizationCreateRequestData
	isSet bool
}

func (v NullableAppStoreVersionLocalizationCreateRequestData) Get() *AppStoreVersionLocalizationCreateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationCreateRequestData) Set(val *AppStoreVersionLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationCreateRequestData(val *AppStoreVersionLocalizationCreateRequestData) *NullableAppStoreVersionLocalizationCreateRequestData {
	return &NullableAppStoreVersionLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


