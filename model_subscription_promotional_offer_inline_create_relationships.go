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

// checks if the SubscriptionPromotionalOfferInlineCreateRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPromotionalOfferInlineCreateRelationships{}

// SubscriptionPromotionalOfferInlineCreateRelationships struct for SubscriptionPromotionalOfferInlineCreateRelationships
type SubscriptionPromotionalOfferInlineCreateRelationships struct {
	Subscription *PromotedPurchaseCreateRequestDataRelationshipsSubscription `json:"subscription,omitempty"`
	Prices *SubscriptionPromotionalOfferInlineCreateRelationshipsPrices `json:"prices,omitempty"`
}

// NewSubscriptionPromotionalOfferInlineCreateRelationships instantiates a new SubscriptionPromotionalOfferInlineCreateRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPromotionalOfferInlineCreateRelationships() *SubscriptionPromotionalOfferInlineCreateRelationships {
	this := SubscriptionPromotionalOfferInlineCreateRelationships{}
	return &this
}

// NewSubscriptionPromotionalOfferInlineCreateRelationshipsWithDefaults instantiates a new SubscriptionPromotionalOfferInlineCreateRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPromotionalOfferInlineCreateRelationshipsWithDefaults() *SubscriptionPromotionalOfferInlineCreateRelationships {
	this := SubscriptionPromotionalOfferInlineCreateRelationships{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferInlineCreateRelationships) GetSubscription() PromotedPurchaseCreateRequestDataRelationshipsSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret PromotedPurchaseCreateRequestDataRelationshipsSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferInlineCreateRelationships) GetSubscriptionOk() (*PromotedPurchaseCreateRequestDataRelationshipsSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferInlineCreateRelationships) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given PromotedPurchaseCreateRequestDataRelationshipsSubscription and assigns it to the Subscription field.
func (o *SubscriptionPromotionalOfferInlineCreateRelationships) SetSubscription(v PromotedPurchaseCreateRequestDataRelationshipsSubscription) {
	o.Subscription = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferInlineCreateRelationships) GetPrices() SubscriptionPromotionalOfferInlineCreateRelationshipsPrices {
	if o == nil || IsNil(o.Prices) {
		var ret SubscriptionPromotionalOfferInlineCreateRelationshipsPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferInlineCreateRelationships) GetPricesOk() (*SubscriptionPromotionalOfferInlineCreateRelationshipsPrices, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferInlineCreateRelationships) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given SubscriptionPromotionalOfferInlineCreateRelationshipsPrices and assigns it to the Prices field.
func (o *SubscriptionPromotionalOfferInlineCreateRelationships) SetPrices(v SubscriptionPromotionalOfferInlineCreateRelationshipsPrices) {
	o.Prices = &v
}

func (o SubscriptionPromotionalOfferInlineCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPromotionalOfferInlineCreateRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	return toSerialize, nil
}

type NullableSubscriptionPromotionalOfferInlineCreateRelationships struct {
	value *SubscriptionPromotionalOfferInlineCreateRelationships
	isSet bool
}

func (v NullableSubscriptionPromotionalOfferInlineCreateRelationships) Get() *SubscriptionPromotionalOfferInlineCreateRelationships {
	return v.value
}

func (v *NullableSubscriptionPromotionalOfferInlineCreateRelationships) Set(val *SubscriptionPromotionalOfferInlineCreateRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOfferInlineCreateRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOfferInlineCreateRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOfferInlineCreateRelationships(val *SubscriptionPromotionalOfferInlineCreateRelationships) *NullableSubscriptionPromotionalOfferInlineCreateRelationships {
	return &NullableSubscriptionPromotionalOfferInlineCreateRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOfferInlineCreateRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOfferInlineCreateRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


