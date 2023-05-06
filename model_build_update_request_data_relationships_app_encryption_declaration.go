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

// checks if the BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration{}

// BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration struct for BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration
type BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration struct {
	Data *AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData `json:"data,omitempty"`
}

// NewBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration instantiates a new BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration() *BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration {
	this := BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration{}
	return &this
}

// NewBuildUpdateRequestDataRelationshipsAppEncryptionDeclarationWithDefaults instantiates a new BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildUpdateRequestDataRelationshipsAppEncryptionDeclarationWithDefaults() *BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration {
	this := BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) GetData() AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData {
	if o == nil || IsNil(o.Data) {
		var ret AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) GetDataOk() (*AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData and assigns it to the Data field.
func (o *BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) SetData(v AppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData) {
	o.Data = &v
}

func (o BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration struct {
	value *BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration
	isSet bool
}

func (v NullableBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) Get() *BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration {
	return v.value
}

func (v *NullableBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) Set(val *BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration(val *BuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) *NullableBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration {
	return &NullableBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration{value: val, isSet: true}
}

func (v NullableBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildUpdateRequestDataRelationshipsAppEncryptionDeclaration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


