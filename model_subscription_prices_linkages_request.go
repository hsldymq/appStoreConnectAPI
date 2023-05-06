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

// checks if the SubscriptionPricesLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPricesLinkagesRequest{}

// SubscriptionPricesLinkagesRequest struct for SubscriptionPricesLinkagesRequest
type SubscriptionPricesLinkagesRequest struct {
	Data []SubscriptionRelationshipsPricesDataInner `json:"data"`
}

// NewSubscriptionPricesLinkagesRequest instantiates a new SubscriptionPricesLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPricesLinkagesRequest(data []SubscriptionRelationshipsPricesDataInner) *SubscriptionPricesLinkagesRequest {
	this := SubscriptionPricesLinkagesRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionPricesLinkagesRequestWithDefaults instantiates a new SubscriptionPricesLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPricesLinkagesRequestWithDefaults() *SubscriptionPricesLinkagesRequest {
	this := SubscriptionPricesLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionPricesLinkagesRequest) GetData() []SubscriptionRelationshipsPricesDataInner {
	if o == nil {
		var ret []SubscriptionRelationshipsPricesDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPricesLinkagesRequest) GetDataOk() ([]SubscriptionRelationshipsPricesDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscriptionPricesLinkagesRequest) SetData(v []SubscriptionRelationshipsPricesDataInner) {
	o.Data = v
}

func (o SubscriptionPricesLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPricesLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionPricesLinkagesRequest struct {
	value *SubscriptionPricesLinkagesRequest
	isSet bool
}

func (v NullableSubscriptionPricesLinkagesRequest) Get() *SubscriptionPricesLinkagesRequest {
	return v.value
}

func (v *NullableSubscriptionPricesLinkagesRequest) Set(val *SubscriptionPricesLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPricesLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPricesLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPricesLinkagesRequest(val *SubscriptionPricesLinkagesRequest) *NullableSubscriptionPricesLinkagesRequest {
	return &NullableSubscriptionPricesLinkagesRequest{value: val, isSet: true}
}

func (v NullableSubscriptionPricesLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPricesLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


