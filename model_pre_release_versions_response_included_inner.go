/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PreReleaseVersionsResponseIncludedInner - struct for PreReleaseVersionsResponseIncludedInner
type PreReleaseVersionsResponseIncludedInner struct {
	App *App
	Build *Build
}

// AppAsPreReleaseVersionsResponseIncludedInner is a convenience function that returns App wrapped in PreReleaseVersionsResponseIncludedInner
func AppAsPreReleaseVersionsResponseIncludedInner(v *App) PreReleaseVersionsResponseIncludedInner {
	return PreReleaseVersionsResponseIncludedInner{
		App: v,
	}
}

// BuildAsPreReleaseVersionsResponseIncludedInner is a convenience function that returns Build wrapped in PreReleaseVersionsResponseIncludedInner
func BuildAsPreReleaseVersionsResponseIncludedInner(v *Build) PreReleaseVersionsResponseIncludedInner {
	return PreReleaseVersionsResponseIncludedInner{
		Build: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PreReleaseVersionsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into App
	err = newStrictDecoder(data).Decode(&dst.App)
	if err == nil {
		jsonApp, _ := json.Marshal(dst.App)
		if string(jsonApp) == "{}" { // empty struct
			dst.App = nil
		} else {
			match++
		}
	} else {
		dst.App = nil
	}

	// try to unmarshal data into Build
	err = newStrictDecoder(data).Decode(&dst.Build)
	if err == nil {
		jsonBuild, _ := json.Marshal(dst.Build)
		if string(jsonBuild) == "{}" { // empty struct
			dst.Build = nil
		} else {
			match++
		}
	} else {
		dst.Build = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.App = nil
		dst.Build = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PreReleaseVersionsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PreReleaseVersionsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PreReleaseVersionsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.App != nil {
		return json.Marshal(&src.App)
	}

	if src.Build != nil {
		return json.Marshal(&src.Build)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PreReleaseVersionsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.App != nil {
		return obj.App
	}

	if obj.Build != nil {
		return obj.Build
	}

	// all schemas are nil
	return nil
}

type NullablePreReleaseVersionsResponseIncludedInner struct {
	value *PreReleaseVersionsResponseIncludedInner
	isSet bool
}

func (v NullablePreReleaseVersionsResponseIncludedInner) Get() *PreReleaseVersionsResponseIncludedInner {
	return v.value
}

func (v *NullablePreReleaseVersionsResponseIncludedInner) Set(val *PreReleaseVersionsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePreReleaseVersionsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePreReleaseVersionsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreReleaseVersionsResponseIncludedInner(val *PreReleaseVersionsResponseIncludedInner) *NullablePreReleaseVersionsResponseIncludedInner {
	return &NullablePreReleaseVersionsResponseIncludedInner{value: val, isSet: true}
}

func (v NullablePreReleaseVersionsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreReleaseVersionsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


