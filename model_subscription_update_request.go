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

// checks if the SubscriptionUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionUpdateRequest{}

// SubscriptionUpdateRequest struct for SubscriptionUpdateRequest
type SubscriptionUpdateRequest struct {
	Data SubscriptionUpdateRequestData `json:"data"`
	Included []SubscriptionUpdateRequestIncludedInner `json:"included,omitempty"`
}

// NewSubscriptionUpdateRequest instantiates a new SubscriptionUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionUpdateRequest(data SubscriptionUpdateRequestData) *SubscriptionUpdateRequest {
	this := SubscriptionUpdateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionUpdateRequestWithDefaults instantiates a new SubscriptionUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionUpdateRequestWithDefaults() *SubscriptionUpdateRequest {
	this := SubscriptionUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionUpdateRequest) GetData() SubscriptionUpdateRequestData {
	if o == nil {
		var ret SubscriptionUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionUpdateRequest) GetDataOk() (*SubscriptionUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionUpdateRequest) SetData(v SubscriptionUpdateRequestData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionUpdateRequest) GetIncluded() []SubscriptionUpdateRequestIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionUpdateRequestIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionUpdateRequest) GetIncludedOk() ([]SubscriptionUpdateRequestIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionUpdateRequest) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionUpdateRequestIncludedInner and assigns it to the Included field.
func (o *SubscriptionUpdateRequest) SetIncluded(v []SubscriptionUpdateRequestIncludedInner) {
	o.Included = v
}

func (o SubscriptionUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

type NullableSubscriptionUpdateRequest struct {
	value *SubscriptionUpdateRequest
	isSet bool
}

func (v NullableSubscriptionUpdateRequest) Get() *SubscriptionUpdateRequest {
	return v.value
}

func (v *NullableSubscriptionUpdateRequest) Set(val *SubscriptionUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionUpdateRequest(val *SubscriptionUpdateRequest) *NullableSubscriptionUpdateRequest {
	return &NullableSubscriptionUpdateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


