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

// ScmRepositoriesResponseIncludedInner - struct for ScmRepositoriesResponseIncludedInner
type ScmRepositoriesResponseIncludedInner struct {
	ScmGitReference *ScmGitReference
	ScmProvider *ScmProvider
}

// ScmGitReferenceAsScmRepositoriesResponseIncludedInner is a convenience function that returns ScmGitReference wrapped in ScmRepositoriesResponseIncludedInner
func ScmGitReferenceAsScmRepositoriesResponseIncludedInner(v *ScmGitReference) ScmRepositoriesResponseIncludedInner {
	return ScmRepositoriesResponseIncludedInner{
		ScmGitReference: v,
	}
}

// ScmProviderAsScmRepositoriesResponseIncludedInner is a convenience function that returns ScmProvider wrapped in ScmRepositoriesResponseIncludedInner
func ScmProviderAsScmRepositoriesResponseIncludedInner(v *ScmProvider) ScmRepositoriesResponseIncludedInner {
	return ScmRepositoriesResponseIncludedInner{
		ScmProvider: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ScmRepositoriesResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ScmGitReference
	err = newStrictDecoder(data).Decode(&dst.ScmGitReference)
	if err == nil {
		jsonScmGitReference, _ := json.Marshal(dst.ScmGitReference)
		if string(jsonScmGitReference) == "{}" { // empty struct
			dst.ScmGitReference = nil
		} else {
			match++
		}
	} else {
		dst.ScmGitReference = nil
	}

	// try to unmarshal data into ScmProvider
	err = newStrictDecoder(data).Decode(&dst.ScmProvider)
	if err == nil {
		jsonScmProvider, _ := json.Marshal(dst.ScmProvider)
		if string(jsonScmProvider) == "{}" { // empty struct
			dst.ScmProvider = nil
		} else {
			match++
		}
	} else {
		dst.ScmProvider = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ScmGitReference = nil
		dst.ScmProvider = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ScmRepositoriesResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ScmRepositoriesResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ScmRepositoriesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.ScmGitReference != nil {
		return json.Marshal(&src.ScmGitReference)
	}

	if src.ScmProvider != nil {
		return json.Marshal(&src.ScmProvider)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ScmRepositoriesResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ScmGitReference != nil {
		return obj.ScmGitReference
	}

	if obj.ScmProvider != nil {
		return obj.ScmProvider
	}

	// all schemas are nil
	return nil
}

type NullableScmRepositoriesResponseIncludedInner struct {
	value *ScmRepositoriesResponseIncludedInner
	isSet bool
}

func (v NullableScmRepositoriesResponseIncludedInner) Get() *ScmRepositoriesResponseIncludedInner {
	return v.value
}

func (v *NullableScmRepositoriesResponseIncludedInner) Set(val *ScmRepositoriesResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableScmRepositoriesResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableScmRepositoriesResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmRepositoriesResponseIncludedInner(val *ScmRepositoriesResponseIncludedInner) *NullableScmRepositoriesResponseIncludedInner {
	return &NullableScmRepositoriesResponseIncludedInner{value: val, isSet: true}
}

func (v NullableScmRepositoriesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmRepositoriesResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

