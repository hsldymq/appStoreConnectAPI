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

// CiCompletionStatus the model 'CiCompletionStatus'
type CiCompletionStatus string

// List of CiCompletionStatus
const (
	SUCCEEDED CiCompletionStatus = "SUCCEEDED"
	FAILED CiCompletionStatus = "FAILED"
	ERRORED CiCompletionStatus = "ERRORED"
	CANCELED CiCompletionStatus = "CANCELED"
	SKIPPED CiCompletionStatus = "SKIPPED"
)

// All allowed values of CiCompletionStatus enum
var AllowedCiCompletionStatusEnumValues = []CiCompletionStatus{
	"SUCCEEDED",
	"FAILED",
	"ERRORED",
	"CANCELED",
	"SKIPPED",
}

func (v *CiCompletionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CiCompletionStatus(value)
	for _, existing := range AllowedCiCompletionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CiCompletionStatus", value)
}

// NewCiCompletionStatusFromValue returns a pointer to a valid CiCompletionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCiCompletionStatusFromValue(v string) (*CiCompletionStatus, error) {
	ev := CiCompletionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CiCompletionStatus: valid values are %v", v, AllowedCiCompletionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CiCompletionStatus) IsValid() bool {
	for _, existing := range AllowedCiCompletionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CiCompletionStatus value
func (v CiCompletionStatus) Ptr() *CiCompletionStatus {
	return &v
}

type NullableCiCompletionStatus struct {
	value *CiCompletionStatus
	isSet bool
}

func (v NullableCiCompletionStatus) Get() *CiCompletionStatus {
	return v.value
}

func (v *NullableCiCompletionStatus) Set(val *CiCompletionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCiCompletionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCiCompletionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiCompletionStatus(val *CiCompletionStatus) *NullableCiCompletionStatus {
	return &NullableCiCompletionStatus{value: val, isSet: true}
}

func (v NullableCiCompletionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiCompletionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

