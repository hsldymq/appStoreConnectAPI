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

// checks if the BetaAppClipInvocationLocalizationInlineCreateRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationLocalizationInlineCreateRelationships{}

// BetaAppClipInvocationLocalizationInlineCreateRelationships struct for BetaAppClipInvocationLocalizationInlineCreateRelationships
type BetaAppClipInvocationLocalizationInlineCreateRelationships struct {
	BetaAppClipInvocation *BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocation `json:"betaAppClipInvocation,omitempty"`
}

// NewBetaAppClipInvocationLocalizationInlineCreateRelationships instantiates a new BetaAppClipInvocationLocalizationInlineCreateRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationLocalizationInlineCreateRelationships() *BetaAppClipInvocationLocalizationInlineCreateRelationships {
	this := BetaAppClipInvocationLocalizationInlineCreateRelationships{}
	return &this
}

// NewBetaAppClipInvocationLocalizationInlineCreateRelationshipsWithDefaults instantiates a new BetaAppClipInvocationLocalizationInlineCreateRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationLocalizationInlineCreateRelationshipsWithDefaults() *BetaAppClipInvocationLocalizationInlineCreateRelationships {
	this := BetaAppClipInvocationLocalizationInlineCreateRelationships{}
	return &this
}

// GetBetaAppClipInvocation returns the BetaAppClipInvocation field value if set, zero value otherwise.
func (o *BetaAppClipInvocationLocalizationInlineCreateRelationships) GetBetaAppClipInvocation() BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocation {
	if o == nil || IsNil(o.BetaAppClipInvocation) {
		var ret BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocation
		return ret
	}
	return *o.BetaAppClipInvocation
}

// GetBetaAppClipInvocationOk returns a tuple with the BetaAppClipInvocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationInlineCreateRelationships) GetBetaAppClipInvocationOk() (*BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocation, bool) {
	if o == nil || IsNil(o.BetaAppClipInvocation) {
		return nil, false
	}
	return o.BetaAppClipInvocation, true
}

// HasBetaAppClipInvocation returns a boolean if a field has been set.
func (o *BetaAppClipInvocationLocalizationInlineCreateRelationships) HasBetaAppClipInvocation() bool {
	if o != nil && !IsNil(o.BetaAppClipInvocation) {
		return true
	}

	return false
}

// SetBetaAppClipInvocation gets a reference to the given BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocation and assigns it to the BetaAppClipInvocation field.
func (o *BetaAppClipInvocationLocalizationInlineCreateRelationships) SetBetaAppClipInvocation(v BetaAppClipInvocationLocalizationInlineCreateRelationshipsBetaAppClipInvocation) {
	o.BetaAppClipInvocation = &v
}

func (o BetaAppClipInvocationLocalizationInlineCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationLocalizationInlineCreateRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BetaAppClipInvocation) {
		toSerialize["betaAppClipInvocation"] = o.BetaAppClipInvocation
	}
	return toSerialize, nil
}

type NullableBetaAppClipInvocationLocalizationInlineCreateRelationships struct {
	value *BetaAppClipInvocationLocalizationInlineCreateRelationships
	isSet bool
}

func (v NullableBetaAppClipInvocationLocalizationInlineCreateRelationships) Get() *BetaAppClipInvocationLocalizationInlineCreateRelationships {
	return v.value
}

func (v *NullableBetaAppClipInvocationLocalizationInlineCreateRelationships) Set(val *BetaAppClipInvocationLocalizationInlineCreateRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationLocalizationInlineCreateRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationLocalizationInlineCreateRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationLocalizationInlineCreateRelationships(val *BetaAppClipInvocationLocalizationInlineCreateRelationships) *NullableBetaAppClipInvocationLocalizationInlineCreateRelationships {
	return &NullableBetaAppClipInvocationLocalizationInlineCreateRelationships{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationLocalizationInlineCreateRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationLocalizationInlineCreateRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


