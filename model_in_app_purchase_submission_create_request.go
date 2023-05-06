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

// checks if the InAppPurchaseSubmissionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseSubmissionCreateRequest{}

// InAppPurchaseSubmissionCreateRequest struct for InAppPurchaseSubmissionCreateRequest
type InAppPurchaseSubmissionCreateRequest struct {
	Data InAppPurchaseSubmissionCreateRequestData `json:"data"`
}

// NewInAppPurchaseSubmissionCreateRequest instantiates a new InAppPurchaseSubmissionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseSubmissionCreateRequest(data InAppPurchaseSubmissionCreateRequestData) *InAppPurchaseSubmissionCreateRequest {
	this := InAppPurchaseSubmissionCreateRequest{}
	this.Data = data
	return &this
}

// NewInAppPurchaseSubmissionCreateRequestWithDefaults instantiates a new InAppPurchaseSubmissionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseSubmissionCreateRequestWithDefaults() *InAppPurchaseSubmissionCreateRequest {
	this := InAppPurchaseSubmissionCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchaseSubmissionCreateRequest) GetData() InAppPurchaseSubmissionCreateRequestData {
	if o == nil {
		var ret InAppPurchaseSubmissionCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseSubmissionCreateRequest) GetDataOk() (*InAppPurchaseSubmissionCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InAppPurchaseSubmissionCreateRequest) SetData(v InAppPurchaseSubmissionCreateRequestData) {
	o.Data = v
}

func (o InAppPurchaseSubmissionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseSubmissionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableInAppPurchaseSubmissionCreateRequest struct {
	value *InAppPurchaseSubmissionCreateRequest
	isSet bool
}

func (v NullableInAppPurchaseSubmissionCreateRequest) Get() *InAppPurchaseSubmissionCreateRequest {
	return v.value
}

func (v *NullableInAppPurchaseSubmissionCreateRequest) Set(val *InAppPurchaseSubmissionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseSubmissionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseSubmissionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseSubmissionCreateRequest(val *InAppPurchaseSubmissionCreateRequest) *NullableInAppPurchaseSubmissionCreateRequest {
	return &NullableInAppPurchaseSubmissionCreateRequest{value: val, isSet: true}
}

func (v NullableInAppPurchaseSubmissionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseSubmissionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


