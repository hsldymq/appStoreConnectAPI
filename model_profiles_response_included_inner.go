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

// ProfilesResponseIncludedInner - struct for ProfilesResponseIncludedInner
type ProfilesResponseIncludedInner struct {
	BundleId *BundleId
	Certificate *Certificate
	Device *Device
}

// BundleIdAsProfilesResponseIncludedInner is a convenience function that returns BundleId wrapped in ProfilesResponseIncludedInner
func BundleIdAsProfilesResponseIncludedInner(v *BundleId) ProfilesResponseIncludedInner {
	return ProfilesResponseIncludedInner{
		BundleId: v,
	}
}

// CertificateAsProfilesResponseIncludedInner is a convenience function that returns Certificate wrapped in ProfilesResponseIncludedInner
func CertificateAsProfilesResponseIncludedInner(v *Certificate) ProfilesResponseIncludedInner {
	return ProfilesResponseIncludedInner{
		Certificate: v,
	}
}

// DeviceAsProfilesResponseIncludedInner is a convenience function that returns Device wrapped in ProfilesResponseIncludedInner
func DeviceAsProfilesResponseIncludedInner(v *Device) ProfilesResponseIncludedInner {
	return ProfilesResponseIncludedInner{
		Device: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ProfilesResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BundleId
	err = newStrictDecoder(data).Decode(&dst.BundleId)
	if err == nil {
		jsonBundleId, _ := json.Marshal(dst.BundleId)
		if string(jsonBundleId) == "{}" { // empty struct
			dst.BundleId = nil
		} else {
			match++
		}
	} else {
		dst.BundleId = nil
	}

	// try to unmarshal data into Certificate
	err = newStrictDecoder(data).Decode(&dst.Certificate)
	if err == nil {
		jsonCertificate, _ := json.Marshal(dst.Certificate)
		if string(jsonCertificate) == "{}" { // empty struct
			dst.Certificate = nil
		} else {
			match++
		}
	} else {
		dst.Certificate = nil
	}

	// try to unmarshal data into Device
	err = newStrictDecoder(data).Decode(&dst.Device)
	if err == nil {
		jsonDevice, _ := json.Marshal(dst.Device)
		if string(jsonDevice) == "{}" { // empty struct
			dst.Device = nil
		} else {
			match++
		}
	} else {
		dst.Device = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BundleId = nil
		dst.Certificate = nil
		dst.Device = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ProfilesResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ProfilesResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProfilesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.BundleId != nil {
		return json.Marshal(&src.BundleId)
	}

	if src.Certificate != nil {
		return json.Marshal(&src.Certificate)
	}

	if src.Device != nil {
		return json.Marshal(&src.Device)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ProfilesResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BundleId != nil {
		return obj.BundleId
	}

	if obj.Certificate != nil {
		return obj.Certificate
	}

	if obj.Device != nil {
		return obj.Device
	}

	// all schemas are nil
	return nil
}

type NullableProfilesResponseIncludedInner struct {
	value *ProfilesResponseIncludedInner
	isSet bool
}

func (v NullableProfilesResponseIncludedInner) Get() *ProfilesResponseIncludedInner {
	return v.value
}

func (v *NullableProfilesResponseIncludedInner) Set(val *ProfilesResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProfilesResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProfilesResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfilesResponseIncludedInner(val *ProfilesResponseIncludedInner) *NullableProfilesResponseIncludedInner {
	return &NullableProfilesResponseIncludedInner{value: val, isSet: true}
}

func (v NullableProfilesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfilesResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


