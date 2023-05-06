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

// checks if the SubscriptionAvailabilityCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionAvailabilityCreateRequest{}

// SubscriptionAvailabilityCreateRequest struct for SubscriptionAvailabilityCreateRequest
type SubscriptionAvailabilityCreateRequest struct {
	Data SubscriptionAvailabilityCreateRequestData `json:"data"`
}

// NewSubscriptionAvailabilityCreateRequest instantiates a new SubscriptionAvailabilityCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionAvailabilityCreateRequest(data SubscriptionAvailabilityCreateRequestData) *SubscriptionAvailabilityCreateRequest {
	this := SubscriptionAvailabilityCreateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionAvailabilityCreateRequestWithDefaults instantiates a new SubscriptionAvailabilityCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionAvailabilityCreateRequestWithDefaults() *SubscriptionAvailabilityCreateRequest {
	this := SubscriptionAvailabilityCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionAvailabilityCreateRequest) GetData() SubscriptionAvailabilityCreateRequestData {
	if o == nil {
		var ret SubscriptionAvailabilityCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionAvailabilityCreateRequest) GetDataOk() (*SubscriptionAvailabilityCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionAvailabilityCreateRequest) SetData(v SubscriptionAvailabilityCreateRequestData) {
	o.Data = v
}

func (o SubscriptionAvailabilityCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionAvailabilityCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionAvailabilityCreateRequest struct {
	value *SubscriptionAvailabilityCreateRequest
	isSet bool
}

func (v NullableSubscriptionAvailabilityCreateRequest) Get() *SubscriptionAvailabilityCreateRequest {
	return v.value
}

func (v *NullableSubscriptionAvailabilityCreateRequest) Set(val *SubscriptionAvailabilityCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAvailabilityCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAvailabilityCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAvailabilityCreateRequest(val *SubscriptionAvailabilityCreateRequest) *NullableSubscriptionAvailabilityCreateRequest {
	return &NullableSubscriptionAvailabilityCreateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionAvailabilityCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAvailabilityCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


