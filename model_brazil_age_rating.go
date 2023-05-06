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

// BrazilAgeRating the model 'BrazilAgeRating'
type BrazilAgeRating string

// List of BrazilAgeRating
const (
	L BrazilAgeRating = "L"
	TEN BrazilAgeRating = "TEN"
	TWELVE BrazilAgeRating = "TWELVE"
	FOURTEEN BrazilAgeRating = "FOURTEEN"
	SIXTEEN BrazilAgeRating = "SIXTEEN"
	EIGHTEEN BrazilAgeRating = "EIGHTEEN"
)

// All allowed values of BrazilAgeRating enum
var AllowedBrazilAgeRatingEnumValues = []BrazilAgeRating{
	"L",
	"TEN",
	"TWELVE",
	"FOURTEEN",
	"SIXTEEN",
	"EIGHTEEN",
}

func (v *BrazilAgeRating) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BrazilAgeRating(value)
	for _, existing := range AllowedBrazilAgeRatingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BrazilAgeRating", value)
}

// NewBrazilAgeRatingFromValue returns a pointer to a valid BrazilAgeRating
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrazilAgeRatingFromValue(v string) (*BrazilAgeRating, error) {
	ev := BrazilAgeRating(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrazilAgeRating: valid values are %v", v, AllowedBrazilAgeRatingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrazilAgeRating) IsValid() bool {
	for _, existing := range AllowedBrazilAgeRatingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BrazilAgeRating value
func (v BrazilAgeRating) Ptr() *BrazilAgeRating {
	return &v
}

type NullableBrazilAgeRating struct {
	value *BrazilAgeRating
	isSet bool
}

func (v NullableBrazilAgeRating) Get() *BrazilAgeRating {
	return v.value
}

func (v *NullableBrazilAgeRating) Set(val *BrazilAgeRating) {
	v.value = val
	v.isSet = true
}

func (v NullableBrazilAgeRating) IsSet() bool {
	return v.isSet
}

func (v *NullableBrazilAgeRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrazilAgeRating(val *BrazilAgeRating) *NullableBrazilAgeRating {
	return &NullableBrazilAgeRating{value: val, isSet: true}
}

func (v NullableBrazilAgeRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrazilAgeRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

