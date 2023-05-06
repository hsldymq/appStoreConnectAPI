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

// checks if the BundleIdRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdRelationships{}

// BundleIdRelationships struct for BundleIdRelationships
type BundleIdRelationships struct {
	Profiles *BundleIdRelationshipsProfiles `json:"profiles,omitempty"`
	BundleIdCapabilities *BundleIdRelationshipsBundleIdCapabilities `json:"bundleIdCapabilities,omitempty"`
	App *AppAvailabilityRelationshipsApp `json:"app,omitempty"`
}

// NewBundleIdRelationships instantiates a new BundleIdRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdRelationships() *BundleIdRelationships {
	this := BundleIdRelationships{}
	return &this
}

// NewBundleIdRelationshipsWithDefaults instantiates a new BundleIdRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdRelationshipsWithDefaults() *BundleIdRelationships {
	this := BundleIdRelationships{}
	return &this
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *BundleIdRelationships) GetProfiles() BundleIdRelationshipsProfiles {
	if o == nil || IsNil(o.Profiles) {
		var ret BundleIdRelationshipsProfiles
		return ret
	}
	return *o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdRelationships) GetProfilesOk() (*BundleIdRelationshipsProfiles, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *BundleIdRelationships) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given BundleIdRelationshipsProfiles and assigns it to the Profiles field.
func (o *BundleIdRelationships) SetProfiles(v BundleIdRelationshipsProfiles) {
	o.Profiles = &v
}

// GetBundleIdCapabilities returns the BundleIdCapabilities field value if set, zero value otherwise.
func (o *BundleIdRelationships) GetBundleIdCapabilities() BundleIdRelationshipsBundleIdCapabilities {
	if o == nil || IsNil(o.BundleIdCapabilities) {
		var ret BundleIdRelationshipsBundleIdCapabilities
		return ret
	}
	return *o.BundleIdCapabilities
}

// GetBundleIdCapabilitiesOk returns a tuple with the BundleIdCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdRelationships) GetBundleIdCapabilitiesOk() (*BundleIdRelationshipsBundleIdCapabilities, bool) {
	if o == nil || IsNil(o.BundleIdCapabilities) {
		return nil, false
	}
	return o.BundleIdCapabilities, true
}

// HasBundleIdCapabilities returns a boolean if a field has been set.
func (o *BundleIdRelationships) HasBundleIdCapabilities() bool {
	if o != nil && !IsNil(o.BundleIdCapabilities) {
		return true
	}

	return false
}

// SetBundleIdCapabilities gets a reference to the given BundleIdRelationshipsBundleIdCapabilities and assigns it to the BundleIdCapabilities field.
func (o *BundleIdRelationships) SetBundleIdCapabilities(v BundleIdRelationshipsBundleIdCapabilities) {
	o.BundleIdCapabilities = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *BundleIdRelationships) GetApp() AppAvailabilityRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret AppAvailabilityRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdRelationships) GetAppOk() (*AppAvailabilityRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *BundleIdRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppAvailabilityRelationshipsApp and assigns it to the App field.
func (o *BundleIdRelationships) SetApp(v AppAvailabilityRelationshipsApp) {
	o.App = &v
}

func (o BundleIdRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	if !IsNil(o.BundleIdCapabilities) {
		toSerialize["bundleIdCapabilities"] = o.BundleIdCapabilities
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	return toSerialize, nil
}

type NullableBundleIdRelationships struct {
	value *BundleIdRelationships
	isSet bool
}

func (v NullableBundleIdRelationships) Get() *BundleIdRelationships {
	return v.value
}

func (v *NullableBundleIdRelationships) Set(val *BundleIdRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdRelationships(val *BundleIdRelationships) *NullableBundleIdRelationships {
	return &NullableBundleIdRelationships{value: val, isSet: true}
}

func (v NullableBundleIdRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


