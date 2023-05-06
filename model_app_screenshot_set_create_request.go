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

// checks if the AppScreenshotSetCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppScreenshotSetCreateRequest{}

// AppScreenshotSetCreateRequest struct for AppScreenshotSetCreateRequest
type AppScreenshotSetCreateRequest struct {
	Data AppScreenshotSetCreateRequestData `json:"data"`
}

// NewAppScreenshotSetCreateRequest instantiates a new AppScreenshotSetCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotSetCreateRequest(data AppScreenshotSetCreateRequestData) *AppScreenshotSetCreateRequest {
	this := AppScreenshotSetCreateRequest{}
	this.Data = data
	return &this
}

// NewAppScreenshotSetCreateRequestWithDefaults instantiates a new AppScreenshotSetCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotSetCreateRequestWithDefaults() *AppScreenshotSetCreateRequest {
	this := AppScreenshotSetCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppScreenshotSetCreateRequest) GetData() AppScreenshotSetCreateRequestData {
	if o == nil {
		var ret AppScreenshotSetCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetCreateRequest) GetDataOk() (*AppScreenshotSetCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppScreenshotSetCreateRequest) SetData(v AppScreenshotSetCreateRequestData) {
	o.Data = v
}

func (o AppScreenshotSetCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppScreenshotSetCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppScreenshotSetCreateRequest struct {
	value *AppScreenshotSetCreateRequest
	isSet bool
}

func (v NullableAppScreenshotSetCreateRequest) Get() *AppScreenshotSetCreateRequest {
	return v.value
}

func (v *NullableAppScreenshotSetCreateRequest) Set(val *AppScreenshotSetCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotSetCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotSetCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotSetCreateRequest(val *AppScreenshotSetCreateRequest) *NullableAppScreenshotSetCreateRequest {
	return &NullableAppScreenshotSetCreateRequest{value: val, isSet: true}
}

func (v NullableAppScreenshotSetCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotSetCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


