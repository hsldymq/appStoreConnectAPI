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

// checks if the AppPreviewSetCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewSetCreateRequest{}

// AppPreviewSetCreateRequest struct for AppPreviewSetCreateRequest
type AppPreviewSetCreateRequest struct {
	Data AppPreviewSetCreateRequestData `json:"data"`
}

// NewAppPreviewSetCreateRequest instantiates a new AppPreviewSetCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewSetCreateRequest(data AppPreviewSetCreateRequestData) *AppPreviewSetCreateRequest {
	this := AppPreviewSetCreateRequest{}
	this.Data = data
	return &this
}

// NewAppPreviewSetCreateRequestWithDefaults instantiates a new AppPreviewSetCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewSetCreateRequestWithDefaults() *AppPreviewSetCreateRequest {
	this := AppPreviewSetCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppPreviewSetCreateRequest) GetData() AppPreviewSetCreateRequestData {
	if o == nil {
		var ret AppPreviewSetCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppPreviewSetCreateRequest) GetDataOk() (*AppPreviewSetCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppPreviewSetCreateRequest) SetData(v AppPreviewSetCreateRequestData) {
	o.Data = v
}

func (o AppPreviewSetCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewSetCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppPreviewSetCreateRequest struct {
	value *AppPreviewSetCreateRequest
	isSet bool
}

func (v NullableAppPreviewSetCreateRequest) Get() *AppPreviewSetCreateRequest {
	return v.value
}

func (v *NullableAppPreviewSetCreateRequest) Set(val *AppPreviewSetCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewSetCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewSetCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewSetCreateRequest(val *AppPreviewSetCreateRequest) *NullableAppPreviewSetCreateRequest {
	return &NullableAppPreviewSetCreateRequest{value: val, isSet: true}
}

func (v NullableAppPreviewSetCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewSetCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


