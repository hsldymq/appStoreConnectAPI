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

// checks if the AppClipAdvancedExperienceRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceRelationships{}

// AppClipAdvancedExperienceRelationships struct for AppClipAdvancedExperienceRelationships
type AppClipAdvancedExperienceRelationships struct {
	AppClip *AppClipAdvancedExperienceRelationshipsAppClip `json:"appClip,omitempty"`
	HeaderImage *AppClipAdvancedExperienceRelationshipsHeaderImage `json:"headerImage,omitempty"`
	Localizations *AppClipAdvancedExperienceRelationshipsLocalizations `json:"localizations,omitempty"`
}

// NewAppClipAdvancedExperienceRelationships instantiates a new AppClipAdvancedExperienceRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceRelationships() *AppClipAdvancedExperienceRelationships {
	this := AppClipAdvancedExperienceRelationships{}
	return &this
}

// NewAppClipAdvancedExperienceRelationshipsWithDefaults instantiates a new AppClipAdvancedExperienceRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceRelationshipsWithDefaults() *AppClipAdvancedExperienceRelationships {
	this := AppClipAdvancedExperienceRelationships{}
	return &this
}

// GetAppClip returns the AppClip field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceRelationships) GetAppClip() AppClipAdvancedExperienceRelationshipsAppClip {
	if o == nil || IsNil(o.AppClip) {
		var ret AppClipAdvancedExperienceRelationshipsAppClip
		return ret
	}
	return *o.AppClip
}

// GetAppClipOk returns a tuple with the AppClip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceRelationships) GetAppClipOk() (*AppClipAdvancedExperienceRelationshipsAppClip, bool) {
	if o == nil || IsNil(o.AppClip) {
		return nil, false
	}
	return o.AppClip, true
}

// HasAppClip returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceRelationships) HasAppClip() bool {
	if o != nil && !IsNil(o.AppClip) {
		return true
	}

	return false
}

// SetAppClip gets a reference to the given AppClipAdvancedExperienceRelationshipsAppClip and assigns it to the AppClip field.
func (o *AppClipAdvancedExperienceRelationships) SetAppClip(v AppClipAdvancedExperienceRelationshipsAppClip) {
	o.AppClip = &v
}

// GetHeaderImage returns the HeaderImage field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceRelationships) GetHeaderImage() AppClipAdvancedExperienceRelationshipsHeaderImage {
	if o == nil || IsNil(o.HeaderImage) {
		var ret AppClipAdvancedExperienceRelationshipsHeaderImage
		return ret
	}
	return *o.HeaderImage
}

// GetHeaderImageOk returns a tuple with the HeaderImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceRelationships) GetHeaderImageOk() (*AppClipAdvancedExperienceRelationshipsHeaderImage, bool) {
	if o == nil || IsNil(o.HeaderImage) {
		return nil, false
	}
	return o.HeaderImage, true
}

// HasHeaderImage returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceRelationships) HasHeaderImage() bool {
	if o != nil && !IsNil(o.HeaderImage) {
		return true
	}

	return false
}

// SetHeaderImage gets a reference to the given AppClipAdvancedExperienceRelationshipsHeaderImage and assigns it to the HeaderImage field.
func (o *AppClipAdvancedExperienceRelationships) SetHeaderImage(v AppClipAdvancedExperienceRelationshipsHeaderImage) {
	o.HeaderImage = &v
}

// GetLocalizations returns the Localizations field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceRelationships) GetLocalizations() AppClipAdvancedExperienceRelationshipsLocalizations {
	if o == nil || IsNil(o.Localizations) {
		var ret AppClipAdvancedExperienceRelationshipsLocalizations
		return ret
	}
	return *o.Localizations
}

// GetLocalizationsOk returns a tuple with the Localizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceRelationships) GetLocalizationsOk() (*AppClipAdvancedExperienceRelationshipsLocalizations, bool) {
	if o == nil || IsNil(o.Localizations) {
		return nil, false
	}
	return o.Localizations, true
}

// HasLocalizations returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceRelationships) HasLocalizations() bool {
	if o != nil && !IsNil(o.Localizations) {
		return true
	}

	return false
}

// SetLocalizations gets a reference to the given AppClipAdvancedExperienceRelationshipsLocalizations and assigns it to the Localizations field.
func (o *AppClipAdvancedExperienceRelationships) SetLocalizations(v AppClipAdvancedExperienceRelationshipsLocalizations) {
	o.Localizations = &v
}

func (o AppClipAdvancedExperienceRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppClip) {
		toSerialize["appClip"] = o.AppClip
	}
	if !IsNil(o.HeaderImage) {
		toSerialize["headerImage"] = o.HeaderImage
	}
	if !IsNil(o.Localizations) {
		toSerialize["localizations"] = o.Localizations
	}
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceRelationships struct {
	value *AppClipAdvancedExperienceRelationships
	isSet bool
}

func (v NullableAppClipAdvancedExperienceRelationships) Get() *AppClipAdvancedExperienceRelationships {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceRelationships) Set(val *AppClipAdvancedExperienceRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceRelationships(val *AppClipAdvancedExperienceRelationships) *NullableAppClipAdvancedExperienceRelationships {
	return &NullableAppClipAdvancedExperienceRelationships{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


