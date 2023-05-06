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

// checks if the SubscriptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionResponse{}

// SubscriptionResponse struct for SubscriptionResponse
type SubscriptionResponse struct {
	Data Subscription `json:"data"`
	Included []SubscriptionsResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewSubscriptionResponse instantiates a new SubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionResponse(data Subscription, links DocumentLinks) *SubscriptionResponse {
	this := SubscriptionResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionResponseWithDefaults instantiates a new SubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionResponseWithDefaults() *SubscriptionResponse {
	this := SubscriptionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionResponse) GetData() Subscription {
	if o == nil {
		var ret Subscription
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetDataOk() (*Subscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionResponse) SetData(v Subscription) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetIncluded() []SubscriptionsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetIncludedOk() ([]SubscriptionsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionsResponseIncludedInner and assigns it to the Included field.
func (o *SubscriptionResponse) SetIncluded(v []SubscriptionsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o SubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableSubscriptionResponse struct {
	value *SubscriptionResponse
	isSet bool
}

func (v NullableSubscriptionResponse) Get() *SubscriptionResponse {
	return v.value
}

func (v *NullableSubscriptionResponse) Set(val *SubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionResponse(val *SubscriptionResponse) *NullableSubscriptionResponse {
	return &NullableSubscriptionResponse{value: val, isSet: true}
}

func (v NullableSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


