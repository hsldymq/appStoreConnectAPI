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

// checks if the SandboxTesterV2UpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SandboxTesterV2UpdateRequestDataAttributes{}

// SandboxTesterV2UpdateRequestDataAttributes struct for SandboxTesterV2UpdateRequestDataAttributes
type SandboxTesterV2UpdateRequestDataAttributes struct {
	Territory *TerritoryCode `json:"territory,omitempty"`
	InterruptPurchases *bool `json:"interruptPurchases,omitempty"`
	SubscriptionRenewalRate *string `json:"subscriptionRenewalRate,omitempty"`
}

// NewSandboxTesterV2UpdateRequestDataAttributes instantiates a new SandboxTesterV2UpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxTesterV2UpdateRequestDataAttributes() *SandboxTesterV2UpdateRequestDataAttributes {
	this := SandboxTesterV2UpdateRequestDataAttributes{}
	return &this
}

// NewSandboxTesterV2UpdateRequestDataAttributesWithDefaults instantiates a new SandboxTesterV2UpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxTesterV2UpdateRequestDataAttributesWithDefaults() *SandboxTesterV2UpdateRequestDataAttributes {
	this := SandboxTesterV2UpdateRequestDataAttributes{}
	return &this
}

// GetTerritory returns the Territory field value if set, zero value otherwise.
func (o *SandboxTesterV2UpdateRequestDataAttributes) GetTerritory() TerritoryCode {
	if o == nil || IsNil(o.Territory) {
		var ret TerritoryCode
		return ret
	}
	return *o.Territory
}

// GetTerritoryOk returns a tuple with the Territory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTesterV2UpdateRequestDataAttributes) GetTerritoryOk() (*TerritoryCode, bool) {
	if o == nil || IsNil(o.Territory) {
		return nil, false
	}
	return o.Territory, true
}

// HasTerritory returns a boolean if a field has been set.
func (o *SandboxTesterV2UpdateRequestDataAttributes) HasTerritory() bool {
	if o != nil && !IsNil(o.Territory) {
		return true
	}

	return false
}

// SetTerritory gets a reference to the given TerritoryCode and assigns it to the Territory field.
func (o *SandboxTesterV2UpdateRequestDataAttributes) SetTerritory(v TerritoryCode) {
	o.Territory = &v
}

// GetInterruptPurchases returns the InterruptPurchases field value if set, zero value otherwise.
func (o *SandboxTesterV2UpdateRequestDataAttributes) GetInterruptPurchases() bool {
	if o == nil || IsNil(o.InterruptPurchases) {
		var ret bool
		return ret
	}
	return *o.InterruptPurchases
}

// GetInterruptPurchasesOk returns a tuple with the InterruptPurchases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTesterV2UpdateRequestDataAttributes) GetInterruptPurchasesOk() (*bool, bool) {
	if o == nil || IsNil(o.InterruptPurchases) {
		return nil, false
	}
	return o.InterruptPurchases, true
}

// HasInterruptPurchases returns a boolean if a field has been set.
func (o *SandboxTesterV2UpdateRequestDataAttributes) HasInterruptPurchases() bool {
	if o != nil && !IsNil(o.InterruptPurchases) {
		return true
	}

	return false
}

// SetInterruptPurchases gets a reference to the given bool and assigns it to the InterruptPurchases field.
func (o *SandboxTesterV2UpdateRequestDataAttributes) SetInterruptPurchases(v bool) {
	o.InterruptPurchases = &v
}

// GetSubscriptionRenewalRate returns the SubscriptionRenewalRate field value if set, zero value otherwise.
func (o *SandboxTesterV2UpdateRequestDataAttributes) GetSubscriptionRenewalRate() string {
	if o == nil || IsNil(o.SubscriptionRenewalRate) {
		var ret string
		return ret
	}
	return *o.SubscriptionRenewalRate
}

// GetSubscriptionRenewalRateOk returns a tuple with the SubscriptionRenewalRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxTesterV2UpdateRequestDataAttributes) GetSubscriptionRenewalRateOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionRenewalRate) {
		return nil, false
	}
	return o.SubscriptionRenewalRate, true
}

// HasSubscriptionRenewalRate returns a boolean if a field has been set.
func (o *SandboxTesterV2UpdateRequestDataAttributes) HasSubscriptionRenewalRate() bool {
	if o != nil && !IsNil(o.SubscriptionRenewalRate) {
		return true
	}

	return false
}

// SetSubscriptionRenewalRate gets a reference to the given string and assigns it to the SubscriptionRenewalRate field.
func (o *SandboxTesterV2UpdateRequestDataAttributes) SetSubscriptionRenewalRate(v string) {
	o.SubscriptionRenewalRate = &v
}

func (o SandboxTesterV2UpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SandboxTesterV2UpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Territory) {
		toSerialize["territory"] = o.Territory
	}
	if !IsNil(o.InterruptPurchases) {
		toSerialize["interruptPurchases"] = o.InterruptPurchases
	}
	if !IsNil(o.SubscriptionRenewalRate) {
		toSerialize["subscriptionRenewalRate"] = o.SubscriptionRenewalRate
	}
	return toSerialize, nil
}

type NullableSandboxTesterV2UpdateRequestDataAttributes struct {
	value *SandboxTesterV2UpdateRequestDataAttributes
	isSet bool
}

func (v NullableSandboxTesterV2UpdateRequestDataAttributes) Get() *SandboxTesterV2UpdateRequestDataAttributes {
	return v.value
}

func (v *NullableSandboxTesterV2UpdateRequestDataAttributes) Set(val *SandboxTesterV2UpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxTesterV2UpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxTesterV2UpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxTesterV2UpdateRequestDataAttributes(val *SandboxTesterV2UpdateRequestDataAttributes) *NullableSandboxTesterV2UpdateRequestDataAttributes {
	return &NullableSandboxTesterV2UpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableSandboxTesterV2UpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxTesterV2UpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


