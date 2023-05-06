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

// checks if the AppStoreVersionExperimentTreatmentCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentCreateRequestData{}

// AppStoreVersionExperimentTreatmentCreateRequestData struct for AppStoreVersionExperimentTreatmentCreateRequestData
type AppStoreVersionExperimentTreatmentCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppStoreVersionExperimentTreatmentCreateRequestDataAttributes `json:"attributes"`
	Relationships AppStoreVersionExperimentTreatmentCreateRequestDataRelationships `json:"relationships"`
}

// NewAppStoreVersionExperimentTreatmentCreateRequestData instantiates a new AppStoreVersionExperimentTreatmentCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentCreateRequestData(type_ string, attributes AppStoreVersionExperimentTreatmentCreateRequestDataAttributes, relationships AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) *AppStoreVersionExperimentTreatmentCreateRequestData {
	this := AppStoreVersionExperimentTreatmentCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppStoreVersionExperimentTreatmentCreateRequestDataWithDefaults instantiates a new AppStoreVersionExperimentTreatmentCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentCreateRequestDataWithDefaults() *AppStoreVersionExperimentTreatmentCreateRequestData {
	this := AppStoreVersionExperimentTreatmentCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionExperimentTreatmentCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionExperimentTreatmentCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppStoreVersionExperimentTreatmentCreateRequestData) GetAttributes() AppStoreVersionExperimentTreatmentCreateRequestDataAttributes {
	if o == nil {
		var ret AppStoreVersionExperimentTreatmentCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequestData) GetAttributesOk() (*AppStoreVersionExperimentTreatmentCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppStoreVersionExperimentTreatmentCreateRequestData) SetAttributes(v AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppStoreVersionExperimentTreatmentCreateRequestData) GetRelationships() AppStoreVersionExperimentTreatmentCreateRequestDataRelationships {
	if o == nil {
		var ret AppStoreVersionExperimentTreatmentCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequestData) GetRelationshipsOk() (*AppStoreVersionExperimentTreatmentCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppStoreVersionExperimentTreatmentCreateRequestData) SetRelationships(v AppStoreVersionExperimentTreatmentCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppStoreVersionExperimentTreatmentCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentCreateRequestData struct {
	value *AppStoreVersionExperimentTreatmentCreateRequestData
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequestData) Get() *AppStoreVersionExperimentTreatmentCreateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequestData) Set(val *AppStoreVersionExperimentTreatmentCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentCreateRequestData(val *AppStoreVersionExperimentTreatmentCreateRequestData) *NullableAppStoreVersionExperimentTreatmentCreateRequestData {
	return &NullableAppStoreVersionExperimentTreatmentCreateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

