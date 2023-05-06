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

// checks if the AppStoreVersionRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionRelationships{}

// AppStoreVersionRelationships struct for AppStoreVersionRelationships
type AppStoreVersionRelationships struct {
	App *AppAvailabilityRelationshipsApp `json:"app,omitempty"`
	// Deprecated
	AgeRatingDeclaration *AppStoreVersionRelationshipsAgeRatingDeclaration `json:"ageRatingDeclaration,omitempty"`
	AppStoreVersionLocalizations *AppStoreVersionRelationshipsAppStoreVersionLocalizations `json:"appStoreVersionLocalizations,omitempty"`
	Build *AppStoreVersionRelationshipsBuild `json:"build,omitempty"`
	AppStoreVersionPhasedRelease *AppStoreVersionRelationshipsAppStoreVersionPhasedRelease `json:"appStoreVersionPhasedRelease,omitempty"`
	RoutingAppCoverage *AppStoreVersionRelationshipsRoutingAppCoverage `json:"routingAppCoverage,omitempty"`
	AppStoreReviewDetail *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail `json:"appStoreReviewDetail,omitempty"`
	AppStoreVersionSubmission *AppStoreVersionRelationshipsAppStoreVersionSubmission `json:"appStoreVersionSubmission,omitempty"`
	AppClipDefaultExperience *AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience `json:"appClipDefaultExperience,omitempty"`
	AppStoreVersionExperiments *AppStoreVersionRelationshipsAppStoreVersionExperiments `json:"appStoreVersionExperiments,omitempty"`
}

// NewAppStoreVersionRelationships instantiates a new AppStoreVersionRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionRelationships() *AppStoreVersionRelationships {
	this := AppStoreVersionRelationships{}
	return &this
}

// NewAppStoreVersionRelationshipsWithDefaults instantiates a new AppStoreVersionRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionRelationshipsWithDefaults() *AppStoreVersionRelationships {
	this := AppStoreVersionRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *AppStoreVersionRelationships) GetApp() AppAvailabilityRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret AppAvailabilityRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationships) GetAppOk() (*AppAvailabilityRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *AppStoreVersionRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppAvailabilityRelationshipsApp and assigns it to the App field.
func (o *AppStoreVersionRelationships) SetApp(v AppAvailabilityRelationshipsApp) {
	o.App = &v
}

// GetAgeRatingDeclaration returns the AgeRatingDeclaration field value if set, zero value otherwise.
// Deprecated
func (o *AppStoreVersionRelationships) GetAgeRatingDeclaration() AppStoreVersionRelationshipsAgeRatingDeclaration {
	if o == nil || IsNil(o.AgeRatingDeclaration) {
		var ret AppStoreVersionRelationshipsAgeRatingDeclaration
		return ret
	}
	return *o.AgeRatingDeclaration
}

// GetAgeRatingDeclarationOk returns a tuple with the AgeRatingDeclaration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AppStoreVersionRelationships) GetAgeRatingDeclarationOk() (*AppStoreVersionRelationshipsAgeRatingDeclaration, bool) {
	if o == nil || IsNil(o.AgeRatingDeclaration) {
		return nil, false
	}
	return o.AgeRatingDeclaration, true
}

// HasAgeRatingDeclaration returns a boolean if a field has been set.
func (o *AppStoreVersionRelationships) HasAgeRatingDeclaration() bool {
	if o != nil && !IsNil(o.AgeRatingDeclaration) {
		return true
	}

	return false
}

// SetAgeRatingDeclaration gets a reference to the given AppStoreVersionRelationshipsAgeRatingDeclaration and assigns it to the AgeRatingDeclaration field.
// Deprecated
func (o *AppStoreVersionRelationships) SetAgeRatingDeclaration(v AppStoreVersionRelationshipsAgeRatingDeclaration) {
	o.AgeRatingDeclaration = &v
}

// GetAppStoreVersionLocalizations returns the AppStoreVersionLocalizations field value if set, zero value otherwise.
func (o *AppStoreVersionRelationships) GetAppStoreVersionLocalizations() AppStoreVersionRelationshipsAppStoreVersionLocalizations {
	if o == nil || IsNil(o.AppStoreVersionLocalizations) {
		var ret AppStoreVersionRelationshipsAppStoreVersionLocalizations
		return ret
	}
	return *o.AppStoreVersionLocalizations
}

// GetAppStoreVersionLocalizationsOk returns a tuple with the AppStoreVersionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationships) GetAppStoreVersionLocalizationsOk() (*AppStoreVersionRelationshipsAppStoreVersionLocalizations, bool) {
	if o == nil || IsNil(o.AppStoreVersionLocalizations) {
		return nil, false
	}
	return o.AppStoreVersionLocalizations, true
}

// HasAppStoreVersionLocalizations returns a boolean if a field has been set.
func (o *AppStoreVersionRelationships) HasAppStoreVersionLocalizations() bool {
	if o != nil && !IsNil(o.AppStoreVersionLocalizations) {
		return true
	}

	return false
}

// SetAppStoreVersionLocalizations gets a reference to the given AppStoreVersionRelationshipsAppStoreVersionLocalizations and assigns it to the AppStoreVersionLocalizations field.
func (o *AppStoreVersionRelationships) SetAppStoreVersionLocalizations(v AppStoreVersionRelationshipsAppStoreVersionLocalizations) {
	o.AppStoreVersionLocalizations = &v
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *AppStoreVersionRelationships) GetBuild() AppStoreVersionRelationshipsBuild {
	if o == nil || IsNil(o.Build) {
		var ret AppStoreVersionRelationshipsBuild
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationships) GetBuildOk() (*AppStoreVersionRelationshipsBuild, bool) {
	if o == nil || IsNil(o.Build) {
		return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *AppStoreVersionRelationships) HasBuild() bool {
	if o != nil && !IsNil(o.Build) {
		return true
	}

	return false
}

// SetBuild gets a reference to the given AppStoreVersionRelationshipsBuild and assigns it to the Build field.
func (o *AppStoreVersionRelationships) SetBuild(v AppStoreVersionRelationshipsBuild) {
	o.Build = &v
}

// GetAppStoreVersionPhasedRelease returns the AppStoreVersionPhasedRelease field value if set, zero value otherwise.
func (o *AppStoreVersionRelationships) GetAppStoreVersionPhasedRelease() AppStoreVersionRelationshipsAppStoreVersionPhasedRelease {
	if o == nil || IsNil(o.AppStoreVersionPhasedRelease) {
		var ret AppStoreVersionRelationshipsAppStoreVersionPhasedRelease
		return ret
	}
	return *o.AppStoreVersionPhasedRelease
}

// GetAppStoreVersionPhasedReleaseOk returns a tuple with the AppStoreVersionPhasedRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationships) GetAppStoreVersionPhasedReleaseOk() (*AppStoreVersionRelationshipsAppStoreVersionPhasedRelease, bool) {
	if o == nil || IsNil(o.AppStoreVersionPhasedRelease) {
		return nil, false
	}
	return o.AppStoreVersionPhasedRelease, true
}

// HasAppStoreVersionPhasedRelease returns a boolean if a field has been set.
func (o *AppStoreVersionRelationships) HasAppStoreVersionPhasedRelease() bool {
	if o != nil && !IsNil(o.AppStoreVersionPhasedRelease) {
		return true
	}

	return false
}

// SetAppStoreVersionPhasedRelease gets a reference to the given AppStoreVersionRelationshipsAppStoreVersionPhasedRelease and assigns it to the AppStoreVersionPhasedRelease field.
func (o *AppStoreVersionRelationships) SetAppStoreVersionPhasedRelease(v AppStoreVersionRelationshipsAppStoreVersionPhasedRelease) {
	o.AppStoreVersionPhasedRelease = &v
}

// GetRoutingAppCoverage returns the RoutingAppCoverage field value if set, zero value otherwise.
func (o *AppStoreVersionRelationships) GetRoutingAppCoverage() AppStoreVersionRelationshipsRoutingAppCoverage {
	if o == nil || IsNil(o.RoutingAppCoverage) {
		var ret AppStoreVersionRelationshipsRoutingAppCoverage
		return ret
	}
	return *o.RoutingAppCoverage
}

// GetRoutingAppCoverageOk returns a tuple with the RoutingAppCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationships) GetRoutingAppCoverageOk() (*AppStoreVersionRelationshipsRoutingAppCoverage, bool) {
	if o == nil || IsNil(o.RoutingAppCoverage) {
		return nil, false
	}
	return o.RoutingAppCoverage, true
}

// HasRoutingAppCoverage returns a boolean if a field has been set.
func (o *AppStoreVersionRelationships) HasRoutingAppCoverage() bool {
	if o != nil && !IsNil(o.RoutingAppCoverage) {
		return true
	}

	return false
}

// SetRoutingAppCoverage gets a reference to the given AppStoreVersionRelationshipsRoutingAppCoverage and assigns it to the RoutingAppCoverage field.
func (o *AppStoreVersionRelationships) SetRoutingAppCoverage(v AppStoreVersionRelationshipsRoutingAppCoverage) {
	o.RoutingAppCoverage = &v
}

// GetAppStoreReviewDetail returns the AppStoreReviewDetail field value if set, zero value otherwise.
func (o *AppStoreVersionRelationships) GetAppStoreReviewDetail() AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail {
	if o == nil || IsNil(o.AppStoreReviewDetail) {
		var ret AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail
		return ret
	}
	return *o.AppStoreReviewDetail
}

// GetAppStoreReviewDetailOk returns a tuple with the AppStoreReviewDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationships) GetAppStoreReviewDetailOk() (*AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail, bool) {
	if o == nil || IsNil(o.AppStoreReviewDetail) {
		return nil, false
	}
	return o.AppStoreReviewDetail, true
}

// HasAppStoreReviewDetail returns a boolean if a field has been set.
func (o *AppStoreVersionRelationships) HasAppStoreReviewDetail() bool {
	if o != nil && !IsNil(o.AppStoreReviewDetail) {
		return true
	}

	return false
}

// SetAppStoreReviewDetail gets a reference to the given AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail and assigns it to the AppStoreReviewDetail field.
func (o *AppStoreVersionRelationships) SetAppStoreReviewDetail(v AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) {
	o.AppStoreReviewDetail = &v
}

// GetAppStoreVersionSubmission returns the AppStoreVersionSubmission field value if set, zero value otherwise.
func (o *AppStoreVersionRelationships) GetAppStoreVersionSubmission() AppStoreVersionRelationshipsAppStoreVersionSubmission {
	if o == nil || IsNil(o.AppStoreVersionSubmission) {
		var ret AppStoreVersionRelationshipsAppStoreVersionSubmission
		return ret
	}
	return *o.AppStoreVersionSubmission
}

// GetAppStoreVersionSubmissionOk returns a tuple with the AppStoreVersionSubmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationships) GetAppStoreVersionSubmissionOk() (*AppStoreVersionRelationshipsAppStoreVersionSubmission, bool) {
	if o == nil || IsNil(o.AppStoreVersionSubmission) {
		return nil, false
	}
	return o.AppStoreVersionSubmission, true
}

// HasAppStoreVersionSubmission returns a boolean if a field has been set.
func (o *AppStoreVersionRelationships) HasAppStoreVersionSubmission() bool {
	if o != nil && !IsNil(o.AppStoreVersionSubmission) {
		return true
	}

	return false
}

// SetAppStoreVersionSubmission gets a reference to the given AppStoreVersionRelationshipsAppStoreVersionSubmission and assigns it to the AppStoreVersionSubmission field.
func (o *AppStoreVersionRelationships) SetAppStoreVersionSubmission(v AppStoreVersionRelationshipsAppStoreVersionSubmission) {
	o.AppStoreVersionSubmission = &v
}

// GetAppClipDefaultExperience returns the AppClipDefaultExperience field value if set, zero value otherwise.
func (o *AppStoreVersionRelationships) GetAppClipDefaultExperience() AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience {
	if o == nil || IsNil(o.AppClipDefaultExperience) {
		var ret AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience
		return ret
	}
	return *o.AppClipDefaultExperience
}

// GetAppClipDefaultExperienceOk returns a tuple with the AppClipDefaultExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationships) GetAppClipDefaultExperienceOk() (*AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience, bool) {
	if o == nil || IsNil(o.AppClipDefaultExperience) {
		return nil, false
	}
	return o.AppClipDefaultExperience, true
}

// HasAppClipDefaultExperience returns a boolean if a field has been set.
func (o *AppStoreVersionRelationships) HasAppClipDefaultExperience() bool {
	if o != nil && !IsNil(o.AppClipDefaultExperience) {
		return true
	}

	return false
}

// SetAppClipDefaultExperience gets a reference to the given AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience and assigns it to the AppClipDefaultExperience field.
func (o *AppStoreVersionRelationships) SetAppClipDefaultExperience(v AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperience) {
	o.AppClipDefaultExperience = &v
}

// GetAppStoreVersionExperiments returns the AppStoreVersionExperiments field value if set, zero value otherwise.
func (o *AppStoreVersionRelationships) GetAppStoreVersionExperiments() AppStoreVersionRelationshipsAppStoreVersionExperiments {
	if o == nil || IsNil(o.AppStoreVersionExperiments) {
		var ret AppStoreVersionRelationshipsAppStoreVersionExperiments
		return ret
	}
	return *o.AppStoreVersionExperiments
}

// GetAppStoreVersionExperimentsOk returns a tuple with the AppStoreVersionExperiments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationships) GetAppStoreVersionExperimentsOk() (*AppStoreVersionRelationshipsAppStoreVersionExperiments, bool) {
	if o == nil || IsNil(o.AppStoreVersionExperiments) {
		return nil, false
	}
	return o.AppStoreVersionExperiments, true
}

// HasAppStoreVersionExperiments returns a boolean if a field has been set.
func (o *AppStoreVersionRelationships) HasAppStoreVersionExperiments() bool {
	if o != nil && !IsNil(o.AppStoreVersionExperiments) {
		return true
	}

	return false
}

// SetAppStoreVersionExperiments gets a reference to the given AppStoreVersionRelationshipsAppStoreVersionExperiments and assigns it to the AppStoreVersionExperiments field.
func (o *AppStoreVersionRelationships) SetAppStoreVersionExperiments(v AppStoreVersionRelationshipsAppStoreVersionExperiments) {
	o.AppStoreVersionExperiments = &v
}

func (o AppStoreVersionRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.AgeRatingDeclaration) {
		toSerialize["ageRatingDeclaration"] = o.AgeRatingDeclaration
	}
	if !IsNil(o.AppStoreVersionLocalizations) {
		toSerialize["appStoreVersionLocalizations"] = o.AppStoreVersionLocalizations
	}
	if !IsNil(o.Build) {
		toSerialize["build"] = o.Build
	}
	if !IsNil(o.AppStoreVersionPhasedRelease) {
		toSerialize["appStoreVersionPhasedRelease"] = o.AppStoreVersionPhasedRelease
	}
	if !IsNil(o.RoutingAppCoverage) {
		toSerialize["routingAppCoverage"] = o.RoutingAppCoverage
	}
	if !IsNil(o.AppStoreReviewDetail) {
		toSerialize["appStoreReviewDetail"] = o.AppStoreReviewDetail
	}
	if !IsNil(o.AppStoreVersionSubmission) {
		toSerialize["appStoreVersionSubmission"] = o.AppStoreVersionSubmission
	}
	if !IsNil(o.AppClipDefaultExperience) {
		toSerialize["appClipDefaultExperience"] = o.AppClipDefaultExperience
	}
	if !IsNil(o.AppStoreVersionExperiments) {
		toSerialize["appStoreVersionExperiments"] = o.AppStoreVersionExperiments
	}
	return toSerialize, nil
}

type NullableAppStoreVersionRelationships struct {
	value *AppStoreVersionRelationships
	isSet bool
}

func (v NullableAppStoreVersionRelationships) Get() *AppStoreVersionRelationships {
	return v.value
}

func (v *NullableAppStoreVersionRelationships) Set(val *AppStoreVersionRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionRelationships(val *AppStoreVersionRelationships) *NullableAppStoreVersionRelationships {
	return &NullableAppStoreVersionRelationships{value: val, isSet: true}
}

func (v NullableAppStoreVersionRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


