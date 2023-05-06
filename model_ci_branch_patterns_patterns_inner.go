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

// checks if the CiBranchPatternsPatternsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBranchPatternsPatternsInner{}

// CiBranchPatternsPatternsInner struct for CiBranchPatternsPatternsInner
type CiBranchPatternsPatternsInner struct {
	Pattern *string `json:"pattern,omitempty"`
	IsPrefix *bool `json:"isPrefix,omitempty"`
}

// NewCiBranchPatternsPatternsInner instantiates a new CiBranchPatternsPatternsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBranchPatternsPatternsInner() *CiBranchPatternsPatternsInner {
	this := CiBranchPatternsPatternsInner{}
	return &this
}

// NewCiBranchPatternsPatternsInnerWithDefaults instantiates a new CiBranchPatternsPatternsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBranchPatternsPatternsInnerWithDefaults() *CiBranchPatternsPatternsInner {
	this := CiBranchPatternsPatternsInner{}
	return &this
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *CiBranchPatternsPatternsInner) GetPattern() string {
	if o == nil || IsNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBranchPatternsPatternsInner) GetPatternOk() (*string, bool) {
	if o == nil || IsNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *CiBranchPatternsPatternsInner) HasPattern() bool {
	if o != nil && !IsNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *CiBranchPatternsPatternsInner) SetPattern(v string) {
	o.Pattern = &v
}

// GetIsPrefix returns the IsPrefix field value if set, zero value otherwise.
func (o *CiBranchPatternsPatternsInner) GetIsPrefix() bool {
	if o == nil || IsNil(o.IsPrefix) {
		var ret bool
		return ret
	}
	return *o.IsPrefix
}

// GetIsPrefixOk returns a tuple with the IsPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBranchPatternsPatternsInner) GetIsPrefixOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrefix) {
		return nil, false
	}
	return o.IsPrefix, true
}

// HasIsPrefix returns a boolean if a field has been set.
func (o *CiBranchPatternsPatternsInner) HasIsPrefix() bool {
	if o != nil && !IsNil(o.IsPrefix) {
		return true
	}

	return false
}

// SetIsPrefix gets a reference to the given bool and assigns it to the IsPrefix field.
func (o *CiBranchPatternsPatternsInner) SetIsPrefix(v bool) {
	o.IsPrefix = &v
}

func (o CiBranchPatternsPatternsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBranchPatternsPatternsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	if !IsNil(o.IsPrefix) {
		toSerialize["isPrefix"] = o.IsPrefix
	}
	return toSerialize, nil
}

type NullableCiBranchPatternsPatternsInner struct {
	value *CiBranchPatternsPatternsInner
	isSet bool
}

func (v NullableCiBranchPatternsPatternsInner) Get() *CiBranchPatternsPatternsInner {
	return v.value
}

func (v *NullableCiBranchPatternsPatternsInner) Set(val *CiBranchPatternsPatternsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBranchPatternsPatternsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBranchPatternsPatternsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBranchPatternsPatternsInner(val *CiBranchPatternsPatternsInner) *NullableCiBranchPatternsPatternsInner {
	return &NullableCiBranchPatternsPatternsInner{value: val, isSet: true}
}

func (v NullableCiBranchPatternsPatternsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBranchPatternsPatternsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


