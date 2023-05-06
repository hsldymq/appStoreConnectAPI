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

// checks if the BundleIdCapabilityUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdCapabilityUpdateRequest{}

// BundleIdCapabilityUpdateRequest struct for BundleIdCapabilityUpdateRequest
type BundleIdCapabilityUpdateRequest struct {
	Data BundleIdCapabilityUpdateRequestData `json:"data"`
}

// NewBundleIdCapabilityUpdateRequest instantiates a new BundleIdCapabilityUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdCapabilityUpdateRequest(data BundleIdCapabilityUpdateRequestData) *BundleIdCapabilityUpdateRequest {
	this := BundleIdCapabilityUpdateRequest{}
	this.Data = data
	return &this
}

// NewBundleIdCapabilityUpdateRequestWithDefaults instantiates a new BundleIdCapabilityUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdCapabilityUpdateRequestWithDefaults() *BundleIdCapabilityUpdateRequest {
	this := BundleIdCapabilityUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BundleIdCapabilityUpdateRequest) GetData() BundleIdCapabilityUpdateRequestData {
	if o == nil {
		var ret BundleIdCapabilityUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleIdCapabilityUpdateRequest) GetDataOk() (*BundleIdCapabilityUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleIdCapabilityUpdateRequest) SetData(v BundleIdCapabilityUpdateRequestData) {
	o.Data = v
}

func (o BundleIdCapabilityUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdCapabilityUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBundleIdCapabilityUpdateRequest struct {
	value *BundleIdCapabilityUpdateRequest
	isSet bool
}

func (v NullableBundleIdCapabilityUpdateRequest) Get() *BundleIdCapabilityUpdateRequest {
	return v.value
}

func (v *NullableBundleIdCapabilityUpdateRequest) Set(val *BundleIdCapabilityUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdCapabilityUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdCapabilityUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdCapabilityUpdateRequest(val *BundleIdCapabilityUpdateRequest) *NullableBundleIdCapabilityUpdateRequest {
	return &NullableBundleIdCapabilityUpdateRequest{value: val, isSet: true}
}

func (v NullableBundleIdCapabilityUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdCapabilityUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


