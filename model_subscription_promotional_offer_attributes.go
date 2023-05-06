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

// checks if the SubscriptionPromotionalOfferAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPromotionalOfferAttributes{}

// SubscriptionPromotionalOfferAttributes struct for SubscriptionPromotionalOfferAttributes
type SubscriptionPromotionalOfferAttributes struct {
	Name *string `json:"name,omitempty"`
	OfferCode *string `json:"offerCode,omitempty"`
	Duration *SubscriptionOfferDuration `json:"duration,omitempty"`
	OfferMode *SubscriptionOfferMode `json:"offerMode,omitempty"`
	NumberOfPeriods *int32 `json:"numberOfPeriods,omitempty"`
}

// NewSubscriptionPromotionalOfferAttributes instantiates a new SubscriptionPromotionalOfferAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPromotionalOfferAttributes() *SubscriptionPromotionalOfferAttributes {
	this := SubscriptionPromotionalOfferAttributes{}
	return &this
}

// NewSubscriptionPromotionalOfferAttributesWithDefaults instantiates a new SubscriptionPromotionalOfferAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPromotionalOfferAttributesWithDefaults() *SubscriptionPromotionalOfferAttributes {
	this := SubscriptionPromotionalOfferAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubscriptionPromotionalOfferAttributes) SetName(v string) {
	o.Name = &v
}

// GetOfferCode returns the OfferCode field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferAttributes) GetOfferCode() string {
	if o == nil || IsNil(o.OfferCode) {
		var ret string
		return ret
	}
	return *o.OfferCode
}

// GetOfferCodeOk returns a tuple with the OfferCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferAttributes) GetOfferCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OfferCode) {
		return nil, false
	}
	return o.OfferCode, true
}

// HasOfferCode returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferAttributes) HasOfferCode() bool {
	if o != nil && !IsNil(o.OfferCode) {
		return true
	}

	return false
}

// SetOfferCode gets a reference to the given string and assigns it to the OfferCode field.
func (o *SubscriptionPromotionalOfferAttributes) SetOfferCode(v string) {
	o.OfferCode = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferAttributes) GetDuration() SubscriptionOfferDuration {
	if o == nil || IsNil(o.Duration) {
		var ret SubscriptionOfferDuration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferAttributes) GetDurationOk() (*SubscriptionOfferDuration, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferAttributes) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given SubscriptionOfferDuration and assigns it to the Duration field.
func (o *SubscriptionPromotionalOfferAttributes) SetDuration(v SubscriptionOfferDuration) {
	o.Duration = &v
}

// GetOfferMode returns the OfferMode field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferAttributes) GetOfferMode() SubscriptionOfferMode {
	if o == nil || IsNil(o.OfferMode) {
		var ret SubscriptionOfferMode
		return ret
	}
	return *o.OfferMode
}

// GetOfferModeOk returns a tuple with the OfferMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferAttributes) GetOfferModeOk() (*SubscriptionOfferMode, bool) {
	if o == nil || IsNil(o.OfferMode) {
		return nil, false
	}
	return o.OfferMode, true
}

// HasOfferMode returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferAttributes) HasOfferMode() bool {
	if o != nil && !IsNil(o.OfferMode) {
		return true
	}

	return false
}

// SetOfferMode gets a reference to the given SubscriptionOfferMode and assigns it to the OfferMode field.
func (o *SubscriptionPromotionalOfferAttributes) SetOfferMode(v SubscriptionOfferMode) {
	o.OfferMode = &v
}

// GetNumberOfPeriods returns the NumberOfPeriods field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferAttributes) GetNumberOfPeriods() int32 {
	if o == nil || IsNil(o.NumberOfPeriods) {
		var ret int32
		return ret
	}
	return *o.NumberOfPeriods
}

// GetNumberOfPeriodsOk returns a tuple with the NumberOfPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferAttributes) GetNumberOfPeriodsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfPeriods) {
		return nil, false
	}
	return o.NumberOfPeriods, true
}

// HasNumberOfPeriods returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferAttributes) HasNumberOfPeriods() bool {
	if o != nil && !IsNil(o.NumberOfPeriods) {
		return true
	}

	return false
}

// SetNumberOfPeriods gets a reference to the given int32 and assigns it to the NumberOfPeriods field.
func (o *SubscriptionPromotionalOfferAttributes) SetNumberOfPeriods(v int32) {
	o.NumberOfPeriods = &v
}

func (o SubscriptionPromotionalOfferAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPromotionalOfferAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OfferCode) {
		toSerialize["offerCode"] = o.OfferCode
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.OfferMode) {
		toSerialize["offerMode"] = o.OfferMode
	}
	if !IsNil(o.NumberOfPeriods) {
		toSerialize["numberOfPeriods"] = o.NumberOfPeriods
	}
	return toSerialize, nil
}

type NullableSubscriptionPromotionalOfferAttributes struct {
	value *SubscriptionPromotionalOfferAttributes
	isSet bool
}

func (v NullableSubscriptionPromotionalOfferAttributes) Get() *SubscriptionPromotionalOfferAttributes {
	return v.value
}

func (v *NullableSubscriptionPromotionalOfferAttributes) Set(val *SubscriptionPromotionalOfferAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOfferAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOfferAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOfferAttributes(val *SubscriptionPromotionalOfferAttributes) *NullableSubscriptionPromotionalOfferAttributes {
	return &NullableSubscriptionPromotionalOfferAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOfferAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOfferAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

