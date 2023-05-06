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

// checks if the SubscriptionPromotionalOfferRelationshipsPrices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPromotionalOfferRelationshipsPrices{}

// SubscriptionPromotionalOfferRelationshipsPrices struct for SubscriptionPromotionalOfferRelationshipsPrices
type SubscriptionPromotionalOfferRelationshipsPrices struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []SubscriptionPromotionalOfferRelationshipsPricesDataInner `json:"data,omitempty"`
}

// NewSubscriptionPromotionalOfferRelationshipsPrices instantiates a new SubscriptionPromotionalOfferRelationshipsPrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPromotionalOfferRelationshipsPrices() *SubscriptionPromotionalOfferRelationshipsPrices {
	this := SubscriptionPromotionalOfferRelationshipsPrices{}
	return &this
}

// NewSubscriptionPromotionalOfferRelationshipsPricesWithDefaults instantiates a new SubscriptionPromotionalOfferRelationshipsPrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPromotionalOfferRelationshipsPricesWithDefaults() *SubscriptionPromotionalOfferRelationshipsPrices {
	this := SubscriptionPromotionalOfferRelationshipsPrices{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) GetData() []SubscriptionPromotionalOfferRelationshipsPricesDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []SubscriptionPromotionalOfferRelationshipsPricesDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) GetDataOk() ([]SubscriptionPromotionalOfferRelationshipsPricesDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SubscriptionPromotionalOfferRelationshipsPricesDataInner and assigns it to the Data field.
func (o *SubscriptionPromotionalOfferRelationshipsPrices) SetData(v []SubscriptionPromotionalOfferRelationshipsPricesDataInner) {
	o.Data = v
}

func (o SubscriptionPromotionalOfferRelationshipsPrices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPromotionalOfferRelationshipsPrices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSubscriptionPromotionalOfferRelationshipsPrices struct {
	value *SubscriptionPromotionalOfferRelationshipsPrices
	isSet bool
}

func (v NullableSubscriptionPromotionalOfferRelationshipsPrices) Get() *SubscriptionPromotionalOfferRelationshipsPrices {
	return v.value
}

func (v *NullableSubscriptionPromotionalOfferRelationshipsPrices) Set(val *SubscriptionPromotionalOfferRelationshipsPrices) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOfferRelationshipsPrices) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOfferRelationshipsPrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOfferRelationshipsPrices(val *SubscriptionPromotionalOfferRelationshipsPrices) *NullableSubscriptionPromotionalOfferRelationshipsPrices {
	return &NullableSubscriptionPromotionalOfferRelationshipsPrices{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOfferRelationshipsPrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOfferRelationshipsPrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


