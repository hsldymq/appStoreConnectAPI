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

// checks if the PromotedPurchaseCreateRequestDataRelationshipsSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseCreateRequestDataRelationshipsSubscription{}

// PromotedPurchaseCreateRequestDataRelationshipsSubscription struct for PromotedPurchaseCreateRequestDataRelationshipsSubscription
type PromotedPurchaseCreateRequestDataRelationshipsSubscription struct {
	Data *PromotedPurchaseRelationshipsSubscriptionData `json:"data,omitempty"`
}

// NewPromotedPurchaseCreateRequestDataRelationshipsSubscription instantiates a new PromotedPurchaseCreateRequestDataRelationshipsSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseCreateRequestDataRelationshipsSubscription() *PromotedPurchaseCreateRequestDataRelationshipsSubscription {
	this := PromotedPurchaseCreateRequestDataRelationshipsSubscription{}
	return &this
}

// NewPromotedPurchaseCreateRequestDataRelationshipsSubscriptionWithDefaults instantiates a new PromotedPurchaseCreateRequestDataRelationshipsSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseCreateRequestDataRelationshipsSubscriptionWithDefaults() *PromotedPurchaseCreateRequestDataRelationshipsSubscription {
	this := PromotedPurchaseCreateRequestDataRelationshipsSubscription{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PromotedPurchaseCreateRequestDataRelationshipsSubscription) GetData() PromotedPurchaseRelationshipsSubscriptionData {
	if o == nil || IsNil(o.Data) {
		var ret PromotedPurchaseRelationshipsSubscriptionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseCreateRequestDataRelationshipsSubscription) GetDataOk() (*PromotedPurchaseRelationshipsSubscriptionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PromotedPurchaseCreateRequestDataRelationshipsSubscription) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PromotedPurchaseRelationshipsSubscriptionData and assigns it to the Data field.
func (o *PromotedPurchaseCreateRequestDataRelationshipsSubscription) SetData(v PromotedPurchaseRelationshipsSubscriptionData) {
	o.Data = &v
}

func (o PromotedPurchaseCreateRequestDataRelationshipsSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseCreateRequestDataRelationshipsSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePromotedPurchaseCreateRequestDataRelationshipsSubscription struct {
	value *PromotedPurchaseCreateRequestDataRelationshipsSubscription
	isSet bool
}

func (v NullablePromotedPurchaseCreateRequestDataRelationshipsSubscription) Get() *PromotedPurchaseCreateRequestDataRelationshipsSubscription {
	return v.value
}

func (v *NullablePromotedPurchaseCreateRequestDataRelationshipsSubscription) Set(val *PromotedPurchaseCreateRequestDataRelationshipsSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseCreateRequestDataRelationshipsSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseCreateRequestDataRelationshipsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseCreateRequestDataRelationshipsSubscription(val *PromotedPurchaseCreateRequestDataRelationshipsSubscription) *NullablePromotedPurchaseCreateRequestDataRelationshipsSubscription {
	return &NullablePromotedPurchaseCreateRequestDataRelationshipsSubscription{value: val, isSet: true}
}

func (v NullablePromotedPurchaseCreateRequestDataRelationshipsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseCreateRequestDataRelationshipsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


