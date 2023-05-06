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

// checks if the SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner{}

// SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner struct for SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner
type SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner instantiates a new SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner(type_ string, id string) *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner {
	this := SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInnerWithDefaults instantiates a new SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInnerWithDefaults() *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner {
	this := SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) SetId(v string) {
	o.Id = v
}

func (o SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner struct {
	value *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner
	isSet bool
}

func (v NullableSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) Get() *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner {
	return v.value
}

func (v *NullableSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) Set(val *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner(val *SubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) *NullableSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner {
	return &NullableSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeRelationshipsOneTimeUseCodesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

