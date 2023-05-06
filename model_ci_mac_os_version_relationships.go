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

// checks if the CiMacOsVersionRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiMacOsVersionRelationships{}

// CiMacOsVersionRelationships struct for CiMacOsVersionRelationships
type CiMacOsVersionRelationships struct {
	XcodeVersions *CiMacOsVersionRelationshipsXcodeVersions `json:"xcodeVersions,omitempty"`
}

// NewCiMacOsVersionRelationships instantiates a new CiMacOsVersionRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiMacOsVersionRelationships() *CiMacOsVersionRelationships {
	this := CiMacOsVersionRelationships{}
	return &this
}

// NewCiMacOsVersionRelationshipsWithDefaults instantiates a new CiMacOsVersionRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiMacOsVersionRelationshipsWithDefaults() *CiMacOsVersionRelationships {
	this := CiMacOsVersionRelationships{}
	return &this
}

// GetXcodeVersions returns the XcodeVersions field value if set, zero value otherwise.
func (o *CiMacOsVersionRelationships) GetXcodeVersions() CiMacOsVersionRelationshipsXcodeVersions {
	if o == nil || IsNil(o.XcodeVersions) {
		var ret CiMacOsVersionRelationshipsXcodeVersions
		return ret
	}
	return *o.XcodeVersions
}

// GetXcodeVersionsOk returns a tuple with the XcodeVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionRelationships) GetXcodeVersionsOk() (*CiMacOsVersionRelationshipsXcodeVersions, bool) {
	if o == nil || IsNil(o.XcodeVersions) {
		return nil, false
	}
	return o.XcodeVersions, true
}

// HasXcodeVersions returns a boolean if a field has been set.
func (o *CiMacOsVersionRelationships) HasXcodeVersions() bool {
	if o != nil && !IsNil(o.XcodeVersions) {
		return true
	}

	return false
}

// SetXcodeVersions gets a reference to the given CiMacOsVersionRelationshipsXcodeVersions and assigns it to the XcodeVersions field.
func (o *CiMacOsVersionRelationships) SetXcodeVersions(v CiMacOsVersionRelationshipsXcodeVersions) {
	o.XcodeVersions = &v
}

func (o CiMacOsVersionRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiMacOsVersionRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.XcodeVersions) {
		toSerialize["xcodeVersions"] = o.XcodeVersions
	}
	return toSerialize, nil
}

type NullableCiMacOsVersionRelationships struct {
	value *CiMacOsVersionRelationships
	isSet bool
}

func (v NullableCiMacOsVersionRelationships) Get() *CiMacOsVersionRelationships {
	return v.value
}

func (v *NullableCiMacOsVersionRelationships) Set(val *CiMacOsVersionRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCiMacOsVersionRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCiMacOsVersionRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiMacOsVersionRelationships(val *CiMacOsVersionRelationships) *NullableCiMacOsVersionRelationships {
	return &NullableCiMacOsVersionRelationships{value: val, isSet: true}
}

func (v NullableCiMacOsVersionRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiMacOsVersionRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

