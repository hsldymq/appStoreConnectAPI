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

// checks if the AppStoreVersionPromotionCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionPromotionCreateRequestDataRelationships{}

// AppStoreVersionPromotionCreateRequestDataRelationships struct for AppStoreVersionPromotionCreateRequestDataRelationships
type AppStoreVersionPromotionCreateRequestDataRelationships struct {
	AppStoreVersion AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion `json:"appStoreVersion"`
	AppStoreVersionExperimentTreatment AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationshipsAppStoreVersionExperimentTreatment `json:"appStoreVersionExperimentTreatment"`
}

// NewAppStoreVersionPromotionCreateRequestDataRelationships instantiates a new AppStoreVersionPromotionCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionPromotionCreateRequestDataRelationships(appStoreVersion AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion, appStoreVersionExperimentTreatment AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationshipsAppStoreVersionExperimentTreatment) *AppStoreVersionPromotionCreateRequestDataRelationships {
	this := AppStoreVersionPromotionCreateRequestDataRelationships{}
	this.AppStoreVersion = appStoreVersion
	this.AppStoreVersionExperimentTreatment = appStoreVersionExperimentTreatment
	return &this
}

// NewAppStoreVersionPromotionCreateRequestDataRelationshipsWithDefaults instantiates a new AppStoreVersionPromotionCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionPromotionCreateRequestDataRelationshipsWithDefaults() *AppStoreVersionPromotionCreateRequestDataRelationships {
	this := AppStoreVersionPromotionCreateRequestDataRelationships{}
	return &this
}

// GetAppStoreVersion returns the AppStoreVersion field value
func (o *AppStoreVersionPromotionCreateRequestDataRelationships) GetAppStoreVersion() AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion {
	if o == nil {
		var ret AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion
		return ret
	}

	return o.AppStoreVersion
}

// GetAppStoreVersionOk returns a tuple with the AppStoreVersion field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPromotionCreateRequestDataRelationships) GetAppStoreVersionOk() (*AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppStoreVersion, true
}

// SetAppStoreVersion sets field value
func (o *AppStoreVersionPromotionCreateRequestDataRelationships) SetAppStoreVersion(v AppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion) {
	o.AppStoreVersion = v
}

// GetAppStoreVersionExperimentTreatment returns the AppStoreVersionExperimentTreatment field value
func (o *AppStoreVersionPromotionCreateRequestDataRelationships) GetAppStoreVersionExperimentTreatment() AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationshipsAppStoreVersionExperimentTreatment {
	if o == nil {
		var ret AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationshipsAppStoreVersionExperimentTreatment
		return ret
	}

	return o.AppStoreVersionExperimentTreatment
}

// GetAppStoreVersionExperimentTreatmentOk returns a tuple with the AppStoreVersionExperimentTreatment field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPromotionCreateRequestDataRelationships) GetAppStoreVersionExperimentTreatmentOk() (*AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationshipsAppStoreVersionExperimentTreatment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppStoreVersionExperimentTreatment, true
}

// SetAppStoreVersionExperimentTreatment sets field value
func (o *AppStoreVersionPromotionCreateRequestDataRelationships) SetAppStoreVersionExperimentTreatment(v AppStoreVersionExperimentTreatmentLocalizationCreateRequestDataRelationshipsAppStoreVersionExperimentTreatment) {
	o.AppStoreVersionExperimentTreatment = v
}

func (o AppStoreVersionPromotionCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionPromotionCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appStoreVersion"] = o.AppStoreVersion
	toSerialize["appStoreVersionExperimentTreatment"] = o.AppStoreVersionExperimentTreatment
	return toSerialize, nil
}

type NullableAppStoreVersionPromotionCreateRequestDataRelationships struct {
	value *AppStoreVersionPromotionCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppStoreVersionPromotionCreateRequestDataRelationships) Get() *AppStoreVersionPromotionCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppStoreVersionPromotionCreateRequestDataRelationships) Set(val *AppStoreVersionPromotionCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionPromotionCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionPromotionCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionPromotionCreateRequestDataRelationships(val *AppStoreVersionPromotionCreateRequestDataRelationships) *NullableAppStoreVersionPromotionCreateRequestDataRelationships {
	return &NullableAppStoreVersionPromotionCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppStoreVersionPromotionCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionPromotionCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


