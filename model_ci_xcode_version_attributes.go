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

// checks if the CiXcodeVersionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiXcodeVersionAttributes{}

// CiXcodeVersionAttributes struct for CiXcodeVersionAttributes
type CiXcodeVersionAttributes struct {
	Version *string `json:"version,omitempty"`
	Name *string `json:"name,omitempty"`
	TestDestinations []CiXcodeVersionAttributesTestDestinationsInner `json:"testDestinations,omitempty"`
}

// NewCiXcodeVersionAttributes instantiates a new CiXcodeVersionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiXcodeVersionAttributes() *CiXcodeVersionAttributes {
	this := CiXcodeVersionAttributes{}
	return &this
}

// NewCiXcodeVersionAttributesWithDefaults instantiates a new CiXcodeVersionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiXcodeVersionAttributesWithDefaults() *CiXcodeVersionAttributes {
	this := CiXcodeVersionAttributes{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CiXcodeVersionAttributes) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiXcodeVersionAttributes) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CiXcodeVersionAttributes) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CiXcodeVersionAttributes) SetVersion(v string) {
	o.Version = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CiXcodeVersionAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiXcodeVersionAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CiXcodeVersionAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CiXcodeVersionAttributes) SetName(v string) {
	o.Name = &v
}

// GetTestDestinations returns the TestDestinations field value if set, zero value otherwise.
func (o *CiXcodeVersionAttributes) GetTestDestinations() []CiXcodeVersionAttributesTestDestinationsInner {
	if o == nil || IsNil(o.TestDestinations) {
		var ret []CiXcodeVersionAttributesTestDestinationsInner
		return ret
	}
	return o.TestDestinations
}

// GetTestDestinationsOk returns a tuple with the TestDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiXcodeVersionAttributes) GetTestDestinationsOk() ([]CiXcodeVersionAttributesTestDestinationsInner, bool) {
	if o == nil || IsNil(o.TestDestinations) {
		return nil, false
	}
	return o.TestDestinations, true
}

// HasTestDestinations returns a boolean if a field has been set.
func (o *CiXcodeVersionAttributes) HasTestDestinations() bool {
	if o != nil && !IsNil(o.TestDestinations) {
		return true
	}

	return false
}

// SetTestDestinations gets a reference to the given []CiXcodeVersionAttributesTestDestinationsInner and assigns it to the TestDestinations field.
func (o *CiXcodeVersionAttributes) SetTestDestinations(v []CiXcodeVersionAttributesTestDestinationsInner) {
	o.TestDestinations = v
}

func (o CiXcodeVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiXcodeVersionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TestDestinations) {
		toSerialize["testDestinations"] = o.TestDestinations
	}
	return toSerialize, nil
}

type NullableCiXcodeVersionAttributes struct {
	value *CiXcodeVersionAttributes
	isSet bool
}

func (v NullableCiXcodeVersionAttributes) Get() *CiXcodeVersionAttributes {
	return v.value
}

func (v *NullableCiXcodeVersionAttributes) Set(val *CiXcodeVersionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCiXcodeVersionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCiXcodeVersionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiXcodeVersionAttributes(val *CiXcodeVersionAttributes) *NullableCiXcodeVersionAttributes {
	return &NullableCiXcodeVersionAttributes{value: val, isSet: true}
}

func (v NullableCiXcodeVersionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiXcodeVersionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


