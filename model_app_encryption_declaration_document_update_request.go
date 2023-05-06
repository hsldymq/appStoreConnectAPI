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

// checks if the AppEncryptionDeclarationDocumentUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEncryptionDeclarationDocumentUpdateRequest{}

// AppEncryptionDeclarationDocumentUpdateRequest struct for AppEncryptionDeclarationDocumentUpdateRequest
type AppEncryptionDeclarationDocumentUpdateRequest struct {
	Data AppEncryptionDeclarationDocumentUpdateRequestData `json:"data"`
}

// NewAppEncryptionDeclarationDocumentUpdateRequest instantiates a new AppEncryptionDeclarationDocumentUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationDocumentUpdateRequest(data AppEncryptionDeclarationDocumentUpdateRequestData) *AppEncryptionDeclarationDocumentUpdateRequest {
	this := AppEncryptionDeclarationDocumentUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppEncryptionDeclarationDocumentUpdateRequestWithDefaults instantiates a new AppEncryptionDeclarationDocumentUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationDocumentUpdateRequestWithDefaults() *AppEncryptionDeclarationDocumentUpdateRequest {
	this := AppEncryptionDeclarationDocumentUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppEncryptionDeclarationDocumentUpdateRequest) GetData() AppEncryptionDeclarationDocumentUpdateRequestData {
	if o == nil {
		var ret AppEncryptionDeclarationDocumentUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationDocumentUpdateRequest) GetDataOk() (*AppEncryptionDeclarationDocumentUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppEncryptionDeclarationDocumentUpdateRequest) SetData(v AppEncryptionDeclarationDocumentUpdateRequestData) {
	o.Data = v
}

func (o AppEncryptionDeclarationDocumentUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEncryptionDeclarationDocumentUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppEncryptionDeclarationDocumentUpdateRequest struct {
	value *AppEncryptionDeclarationDocumentUpdateRequest
	isSet bool
}

func (v NullableAppEncryptionDeclarationDocumentUpdateRequest) Get() *AppEncryptionDeclarationDocumentUpdateRequest {
	return v.value
}

func (v *NullableAppEncryptionDeclarationDocumentUpdateRequest) Set(val *AppEncryptionDeclarationDocumentUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationDocumentUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationDocumentUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationDocumentUpdateRequest(val *AppEncryptionDeclarationDocumentUpdateRequest) *NullableAppEncryptionDeclarationDocumentUpdateRequest {
	return &NullableAppEncryptionDeclarationDocumentUpdateRequest{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationDocumentUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationDocumentUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

