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

// checks if the SubscriptionOfferCodeCustomCodeUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCustomCodeUpdateRequestData{}

// SubscriptionOfferCodeCustomCodeUpdateRequestData struct for SubscriptionOfferCodeCustomCodeUpdateRequestData
type SubscriptionOfferCodeCustomCodeUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewSubscriptionOfferCodeCustomCodeUpdateRequestData instantiates a new SubscriptionOfferCodeCustomCodeUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCustomCodeUpdateRequestData(type_ string, id string) *SubscriptionOfferCodeCustomCodeUpdateRequestData {
	this := SubscriptionOfferCodeCustomCodeUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionOfferCodeCustomCodeUpdateRequestDataWithDefaults instantiates a new SubscriptionOfferCodeCustomCodeUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCustomCodeUpdateRequestDataWithDefaults() *SubscriptionOfferCodeCustomCodeUpdateRequestData {
	this := SubscriptionOfferCodeCustomCodeUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestData) GetAttributes() SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestData) GetAttributesOk() (*SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestData) SetAttributes(v SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o SubscriptionOfferCodeCustomCodeUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCustomCodeUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeCustomCodeUpdateRequestData struct {
	value *SubscriptionOfferCodeCustomCodeUpdateRequestData
	isSet bool
}

func (v NullableSubscriptionOfferCodeCustomCodeUpdateRequestData) Get() *SubscriptionOfferCodeCustomCodeUpdateRequestData {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCustomCodeUpdateRequestData) Set(val *SubscriptionOfferCodeCustomCodeUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCustomCodeUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCustomCodeUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCustomCodeUpdateRequestData(val *SubscriptionOfferCodeCustomCodeUpdateRequestData) *NullableSubscriptionOfferCodeCustomCodeUpdateRequestData {
	return &NullableSubscriptionOfferCodeCustomCodeUpdateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCustomCodeUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCustomCodeUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


