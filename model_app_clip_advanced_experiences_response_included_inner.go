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

// AppClipAdvancedExperiencesResponseIncludedInner - struct for AppClipAdvancedExperiencesResponseIncludedInner
type AppClipAdvancedExperiencesResponseIncludedInner struct {
	AppClip *AppClip
	AppClipAdvancedExperienceImage *AppClipAdvancedExperienceImage
	AppClipAdvancedExperienceLocalization *AppClipAdvancedExperienceLocalization
}

// AppClipAsAppClipAdvancedExperiencesResponseIncludedInner is a convenience function that returns AppClip wrapped in AppClipAdvancedExperiencesResponseIncludedInner
func AppClipAsAppClipAdvancedExperiencesResponseIncludedInner(v *AppClip) AppClipAdvancedExperiencesResponseIncludedInner {
	return AppClipAdvancedExperiencesResponseIncludedInner{
		AppClip: v,
	}
}

// AppClipAdvancedExperienceImageAsAppClipAdvancedExperiencesResponseIncludedInner is a convenience function that returns AppClipAdvancedExperienceImage wrapped in AppClipAdvancedExperiencesResponseIncludedInner
func AppClipAdvancedExperienceImageAsAppClipAdvancedExperiencesResponseIncludedInner(v *AppClipAdvancedExperienceImage) AppClipAdvancedExperiencesResponseIncludedInner {
	return AppClipAdvancedExperiencesResponseIncludedInner{
		AppClipAdvancedExperienceImage: v,
	}
}

// AppClipAdvancedExperienceLocalizationAsAppClipAdvancedExperiencesResponseIncludedInner is a convenience function that returns AppClipAdvancedExperienceLocalization wrapped in AppClipAdvancedExperiencesResponseIncludedInner
func AppClipAdvancedExperienceLocalizationAsAppClipAdvancedExperiencesResponseIncludedInner(v *AppClipAdvancedExperienceLocalization) AppClipAdvancedExperiencesResponseIncludedInner {
	return AppClipAdvancedExperiencesResponseIncludedInner{
		AppClipAdvancedExperienceLocalization: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppClipAdvancedExperiencesResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppClip
	err = newStrictDecoder(data).Decode(&dst.AppClip)
	if err == nil {
		jsonAppClip, _ := json.Marshal(dst.AppClip)
		if string(jsonAppClip) == "{}" { // empty struct
			dst.AppClip = nil
		} else {
			match++
		}
	} else {
		dst.AppClip = nil
	}

	// try to unmarshal data into AppClipAdvancedExperienceImage
	err = newStrictDecoder(data).Decode(&dst.AppClipAdvancedExperienceImage)
	if err == nil {
		jsonAppClipAdvancedExperienceImage, _ := json.Marshal(dst.AppClipAdvancedExperienceImage)
		if string(jsonAppClipAdvancedExperienceImage) == "{}" { // empty struct
			dst.AppClipAdvancedExperienceImage = nil
		} else {
			match++
		}
	} else {
		dst.AppClipAdvancedExperienceImage = nil
	}

	// try to unmarshal data into AppClipAdvancedExperienceLocalization
	err = newStrictDecoder(data).Decode(&dst.AppClipAdvancedExperienceLocalization)
	if err == nil {
		jsonAppClipAdvancedExperienceLocalization, _ := json.Marshal(dst.AppClipAdvancedExperienceLocalization)
		if string(jsonAppClipAdvancedExperienceLocalization) == "{}" { // empty struct
			dst.AppClipAdvancedExperienceLocalization = nil
		} else {
			match++
		}
	} else {
		dst.AppClipAdvancedExperienceLocalization = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppClip = nil
		dst.AppClipAdvancedExperienceImage = nil
		dst.AppClipAdvancedExperienceLocalization = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppClipAdvancedExperiencesResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppClipAdvancedExperiencesResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppClipAdvancedExperiencesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppClip != nil {
		return json.Marshal(&src.AppClip)
	}

	if src.AppClipAdvancedExperienceImage != nil {
		return json.Marshal(&src.AppClipAdvancedExperienceImage)
	}

	if src.AppClipAdvancedExperienceLocalization != nil {
		return json.Marshal(&src.AppClipAdvancedExperienceLocalization)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppClipAdvancedExperiencesResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppClip != nil {
		return obj.AppClip
	}

	if obj.AppClipAdvancedExperienceImage != nil {
		return obj.AppClipAdvancedExperienceImage
	}

	if obj.AppClipAdvancedExperienceLocalization != nil {
		return obj.AppClipAdvancedExperienceLocalization
	}

	// all schemas are nil
	return nil
}

type NullableAppClipAdvancedExperiencesResponseIncludedInner struct {
	value *AppClipAdvancedExperiencesResponseIncludedInner
	isSet bool
}

func (v NullableAppClipAdvancedExperiencesResponseIncludedInner) Get() *AppClipAdvancedExperiencesResponseIncludedInner {
	return v.value
}

func (v *NullableAppClipAdvancedExperiencesResponseIncludedInner) Set(val *AppClipAdvancedExperiencesResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperiencesResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperiencesResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperiencesResponseIncludedInner(val *AppClipAdvancedExperiencesResponseIncludedInner) *NullableAppClipAdvancedExperiencesResponseIncludedInner {
	return &NullableAppClipAdvancedExperiencesResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperiencesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperiencesResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


