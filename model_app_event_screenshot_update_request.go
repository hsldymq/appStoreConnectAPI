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

// checks if the AppEventScreenshotUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventScreenshotUpdateRequest{}

// AppEventScreenshotUpdateRequest struct for AppEventScreenshotUpdateRequest
type AppEventScreenshotUpdateRequest struct {
	Data AppEventScreenshotUpdateRequestData `json:"data"`
}

// NewAppEventScreenshotUpdateRequest instantiates a new AppEventScreenshotUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventScreenshotUpdateRequest(data AppEventScreenshotUpdateRequestData) *AppEventScreenshotUpdateRequest {
	this := AppEventScreenshotUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppEventScreenshotUpdateRequestWithDefaults instantiates a new AppEventScreenshotUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventScreenshotUpdateRequestWithDefaults() *AppEventScreenshotUpdateRequest {
	this := AppEventScreenshotUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppEventScreenshotUpdateRequest) GetData() AppEventScreenshotUpdateRequestData {
	if o == nil {
		var ret AppEventScreenshotUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotUpdateRequest) GetDataOk() (*AppEventScreenshotUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppEventScreenshotUpdateRequest) SetData(v AppEventScreenshotUpdateRequestData) {
	o.Data = v
}

func (o AppEventScreenshotUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventScreenshotUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppEventScreenshotUpdateRequest struct {
	value *AppEventScreenshotUpdateRequest
	isSet bool
}

func (v NullableAppEventScreenshotUpdateRequest) Get() *AppEventScreenshotUpdateRequest {
	return v.value
}

func (v *NullableAppEventScreenshotUpdateRequest) Set(val *AppEventScreenshotUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventScreenshotUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventScreenshotUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventScreenshotUpdateRequest(val *AppEventScreenshotUpdateRequest) *NullableAppEventScreenshotUpdateRequest {
	return &NullableAppEventScreenshotUpdateRequest{value: val, isSet: true}
}

func (v NullableAppEventScreenshotUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventScreenshotUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


