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

// checks if the SubscriptionOfferCodeCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCreateRequestDataAttributes{}

// SubscriptionOfferCodeCreateRequestDataAttributes struct for SubscriptionOfferCodeCreateRequestDataAttributes
type SubscriptionOfferCodeCreateRequestDataAttributes struct {
	Name string `json:"name"`
	CustomerEligibilities []SubscriptionCustomerEligibility `json:"customerEligibilities"`
	OfferEligibility SubscriptionOfferEligibility `json:"offerEligibility"`
	Duration SubscriptionOfferDuration `json:"duration"`
	OfferMode SubscriptionOfferMode `json:"offerMode"`
	NumberOfPeriods int32 `json:"numberOfPeriods"`
}

// NewSubscriptionOfferCodeCreateRequestDataAttributes instantiates a new SubscriptionOfferCodeCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCreateRequestDataAttributes(name string, customerEligibilities []SubscriptionCustomerEligibility, offerEligibility SubscriptionOfferEligibility, duration SubscriptionOfferDuration, offerMode SubscriptionOfferMode, numberOfPeriods int32) *SubscriptionOfferCodeCreateRequestDataAttributes {
	this := SubscriptionOfferCodeCreateRequestDataAttributes{}
	this.Name = name
	this.CustomerEligibilities = customerEligibilities
	this.OfferEligibility = offerEligibility
	this.Duration = duration
	this.OfferMode = offerMode
	this.NumberOfPeriods = numberOfPeriods
	return &this
}

// NewSubscriptionOfferCodeCreateRequestDataAttributesWithDefaults instantiates a new SubscriptionOfferCodeCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCreateRequestDataAttributesWithDefaults() *SubscriptionOfferCodeCreateRequestDataAttributes {
	this := SubscriptionOfferCodeCreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetCustomerEligibilities returns the CustomerEligibilities field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetCustomerEligibilities() []SubscriptionCustomerEligibility {
	if o == nil {
		var ret []SubscriptionCustomerEligibility
		return ret
	}

	return o.CustomerEligibilities
}

// GetCustomerEligibilitiesOk returns a tuple with the CustomerEligibilities field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetCustomerEligibilitiesOk() ([]SubscriptionCustomerEligibility, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerEligibilities, true
}

// SetCustomerEligibilities sets field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) SetCustomerEligibilities(v []SubscriptionCustomerEligibility) {
	o.CustomerEligibilities = v
}

// GetOfferEligibility returns the OfferEligibility field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetOfferEligibility() SubscriptionOfferEligibility {
	if o == nil {
		var ret SubscriptionOfferEligibility
		return ret
	}

	return o.OfferEligibility
}

// GetOfferEligibilityOk returns a tuple with the OfferEligibility field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetOfferEligibilityOk() (*SubscriptionOfferEligibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferEligibility, true
}

// SetOfferEligibility sets field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) SetOfferEligibility(v SubscriptionOfferEligibility) {
	o.OfferEligibility = v
}

// GetDuration returns the Duration field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetDuration() SubscriptionOfferDuration {
	if o == nil {
		var ret SubscriptionOfferDuration
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetDurationOk() (*SubscriptionOfferDuration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) SetDuration(v SubscriptionOfferDuration) {
	o.Duration = v
}

// GetOfferMode returns the OfferMode field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetOfferMode() SubscriptionOfferMode {
	if o == nil {
		var ret SubscriptionOfferMode
		return ret
	}

	return o.OfferMode
}

// GetOfferModeOk returns a tuple with the OfferMode field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetOfferModeOk() (*SubscriptionOfferMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferMode, true
}

// SetOfferMode sets field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) SetOfferMode(v SubscriptionOfferMode) {
	o.OfferMode = v
}

// GetNumberOfPeriods returns the NumberOfPeriods field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetNumberOfPeriods() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfPeriods
}

// GetNumberOfPeriodsOk returns a tuple with the NumberOfPeriods field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) GetNumberOfPeriodsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfPeriods, true
}

// SetNumberOfPeriods sets field value
func (o *SubscriptionOfferCodeCreateRequestDataAttributes) SetNumberOfPeriods(v int32) {
	o.NumberOfPeriods = v
}

func (o SubscriptionOfferCodeCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["customerEligibilities"] = o.CustomerEligibilities
	toSerialize["offerEligibility"] = o.OfferEligibility
	toSerialize["duration"] = o.Duration
	toSerialize["offerMode"] = o.OfferMode
	toSerialize["numberOfPeriods"] = o.NumberOfPeriods
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeCreateRequestDataAttributes struct {
	value *SubscriptionOfferCodeCreateRequestDataAttributes
	isSet bool
}

func (v NullableSubscriptionOfferCodeCreateRequestDataAttributes) Get() *SubscriptionOfferCodeCreateRequestDataAttributes {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCreateRequestDataAttributes) Set(val *SubscriptionOfferCodeCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCreateRequestDataAttributes(val *SubscriptionOfferCodeCreateRequestDataAttributes) *NullableSubscriptionOfferCodeCreateRequestDataAttributes {
	return &NullableSubscriptionOfferCodeCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

