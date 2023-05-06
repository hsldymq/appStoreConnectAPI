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

// checks if the SubscriptionGroupLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupLocalizationCreateRequest{}

// SubscriptionGroupLocalizationCreateRequest struct for SubscriptionGroupLocalizationCreateRequest
type SubscriptionGroupLocalizationCreateRequest struct {
	Data SubscriptionGroupLocalizationCreateRequestData `json:"data"`
}

// NewSubscriptionGroupLocalizationCreateRequest instantiates a new SubscriptionGroupLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupLocalizationCreateRequest(data SubscriptionGroupLocalizationCreateRequestData) *SubscriptionGroupLocalizationCreateRequest {
	this := SubscriptionGroupLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionGroupLocalizationCreateRequestWithDefaults instantiates a new SubscriptionGroupLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupLocalizationCreateRequestWithDefaults() *SubscriptionGroupLocalizationCreateRequest {
	this := SubscriptionGroupLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionGroupLocalizationCreateRequest) GetData() SubscriptionGroupLocalizationCreateRequestData {
	if o == nil {
		var ret SubscriptionGroupLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationCreateRequest) GetDataOk() (*SubscriptionGroupLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionGroupLocalizationCreateRequest) SetData(v SubscriptionGroupLocalizationCreateRequestData) {
	o.Data = v
}

func (o SubscriptionGroupLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionGroupLocalizationCreateRequest struct {
	value *SubscriptionGroupLocalizationCreateRequest
	isSet bool
}

func (v NullableSubscriptionGroupLocalizationCreateRequest) Get() *SubscriptionGroupLocalizationCreateRequest {
	return v.value
}

func (v *NullableSubscriptionGroupLocalizationCreateRequest) Set(val *SubscriptionGroupLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupLocalizationCreateRequest(val *SubscriptionGroupLocalizationCreateRequest) *NullableSubscriptionGroupLocalizationCreateRequest {
	return &NullableSubscriptionGroupLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionGroupLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


