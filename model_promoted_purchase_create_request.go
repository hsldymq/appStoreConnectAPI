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

// checks if the PromotedPurchaseCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseCreateRequest{}

// PromotedPurchaseCreateRequest struct for PromotedPurchaseCreateRequest
type PromotedPurchaseCreateRequest struct {
	Data PromotedPurchaseCreateRequestData `json:"data"`
}

// NewPromotedPurchaseCreateRequest instantiates a new PromotedPurchaseCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseCreateRequest(data PromotedPurchaseCreateRequestData) *PromotedPurchaseCreateRequest {
	this := PromotedPurchaseCreateRequest{}
	this.Data = data
	return &this
}

// NewPromotedPurchaseCreateRequestWithDefaults instantiates a new PromotedPurchaseCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseCreateRequestWithDefaults() *PromotedPurchaseCreateRequest {
	this := PromotedPurchaseCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PromotedPurchaseCreateRequest) GetData() PromotedPurchaseCreateRequestData {
	if o == nil {
		var ret PromotedPurchaseCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseCreateRequest) GetDataOk() (*PromotedPurchaseCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PromotedPurchaseCreateRequest) SetData(v PromotedPurchaseCreateRequestData) {
	o.Data = v
}

func (o PromotedPurchaseCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePromotedPurchaseCreateRequest struct {
	value *PromotedPurchaseCreateRequest
	isSet bool
}

func (v NullablePromotedPurchaseCreateRequest) Get() *PromotedPurchaseCreateRequest {
	return v.value
}

func (v *NullablePromotedPurchaseCreateRequest) Set(val *PromotedPurchaseCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseCreateRequest(val *PromotedPurchaseCreateRequest) *NullablePromotedPurchaseCreateRequest {
	return &NullablePromotedPurchaseCreateRequest{value: val, isSet: true}
}

func (v NullablePromotedPurchaseCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


