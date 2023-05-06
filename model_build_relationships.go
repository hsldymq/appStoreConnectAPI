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

// checks if the BuildRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildRelationships{}

// BuildRelationships struct for BuildRelationships
type BuildRelationships struct {
	PreReleaseVersion *BuildRelationshipsPreReleaseVersion `json:"preReleaseVersion,omitempty"`
	IndividualTesters *BetaGroupRelationshipsBetaTesters `json:"individualTesters,omitempty"`
	BetaGroups *AppRelationshipsBetaGroups `json:"betaGroups,omitempty"`
	BetaBuildLocalizations *BuildRelationshipsBetaBuildLocalizations `json:"betaBuildLocalizations,omitempty"`
	AppEncryptionDeclaration *BuildRelationshipsAppEncryptionDeclaration `json:"appEncryptionDeclaration,omitempty"`
	BetaAppReviewSubmission *BuildRelationshipsBetaAppReviewSubmission `json:"betaAppReviewSubmission,omitempty"`
	App *AppAvailabilityRelationshipsApp `json:"app,omitempty"`
	BuildBetaDetail *BuildRelationshipsBuildBetaDetail `json:"buildBetaDetail,omitempty"`
	AppStoreVersion *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion `json:"appStoreVersion,omitempty"`
	Icons *BuildRelationshipsIcons `json:"icons,omitempty"`
	BuildBundles *BuildRelationshipsBuildBundles `json:"buildBundles,omitempty"`
}

// NewBuildRelationships instantiates a new BuildRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildRelationships() *BuildRelationships {
	this := BuildRelationships{}
	return &this
}

// NewBuildRelationshipsWithDefaults instantiates a new BuildRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildRelationshipsWithDefaults() *BuildRelationships {
	this := BuildRelationships{}
	return &this
}

// GetPreReleaseVersion returns the PreReleaseVersion field value if set, zero value otherwise.
func (o *BuildRelationships) GetPreReleaseVersion() BuildRelationshipsPreReleaseVersion {
	if o == nil || IsNil(o.PreReleaseVersion) {
		var ret BuildRelationshipsPreReleaseVersion
		return ret
	}
	return *o.PreReleaseVersion
}

// GetPreReleaseVersionOk returns a tuple with the PreReleaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetPreReleaseVersionOk() (*BuildRelationshipsPreReleaseVersion, bool) {
	if o == nil || IsNil(o.PreReleaseVersion) {
		return nil, false
	}
	return o.PreReleaseVersion, true
}

// HasPreReleaseVersion returns a boolean if a field has been set.
func (o *BuildRelationships) HasPreReleaseVersion() bool {
	if o != nil && !IsNil(o.PreReleaseVersion) {
		return true
	}

	return false
}

// SetPreReleaseVersion gets a reference to the given BuildRelationshipsPreReleaseVersion and assigns it to the PreReleaseVersion field.
func (o *BuildRelationships) SetPreReleaseVersion(v BuildRelationshipsPreReleaseVersion) {
	o.PreReleaseVersion = &v
}

// GetIndividualTesters returns the IndividualTesters field value if set, zero value otherwise.
func (o *BuildRelationships) GetIndividualTesters() BetaGroupRelationshipsBetaTesters {
	if o == nil || IsNil(o.IndividualTesters) {
		var ret BetaGroupRelationshipsBetaTesters
		return ret
	}
	return *o.IndividualTesters
}

// GetIndividualTestersOk returns a tuple with the IndividualTesters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetIndividualTestersOk() (*BetaGroupRelationshipsBetaTesters, bool) {
	if o == nil || IsNil(o.IndividualTesters) {
		return nil, false
	}
	return o.IndividualTesters, true
}

// HasIndividualTesters returns a boolean if a field has been set.
func (o *BuildRelationships) HasIndividualTesters() bool {
	if o != nil && !IsNil(o.IndividualTesters) {
		return true
	}

	return false
}

// SetIndividualTesters gets a reference to the given BetaGroupRelationshipsBetaTesters and assigns it to the IndividualTesters field.
func (o *BuildRelationships) SetIndividualTesters(v BetaGroupRelationshipsBetaTesters) {
	o.IndividualTesters = &v
}

// GetBetaGroups returns the BetaGroups field value if set, zero value otherwise.
func (o *BuildRelationships) GetBetaGroups() AppRelationshipsBetaGroups {
	if o == nil || IsNil(o.BetaGroups) {
		var ret AppRelationshipsBetaGroups
		return ret
	}
	return *o.BetaGroups
}

// GetBetaGroupsOk returns a tuple with the BetaGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetBetaGroupsOk() (*AppRelationshipsBetaGroups, bool) {
	if o == nil || IsNil(o.BetaGroups) {
		return nil, false
	}
	return o.BetaGroups, true
}

// HasBetaGroups returns a boolean if a field has been set.
func (o *BuildRelationships) HasBetaGroups() bool {
	if o != nil && !IsNil(o.BetaGroups) {
		return true
	}

	return false
}

// SetBetaGroups gets a reference to the given AppRelationshipsBetaGroups and assigns it to the BetaGroups field.
func (o *BuildRelationships) SetBetaGroups(v AppRelationshipsBetaGroups) {
	o.BetaGroups = &v
}

// GetBetaBuildLocalizations returns the BetaBuildLocalizations field value if set, zero value otherwise.
func (o *BuildRelationships) GetBetaBuildLocalizations() BuildRelationshipsBetaBuildLocalizations {
	if o == nil || IsNil(o.BetaBuildLocalizations) {
		var ret BuildRelationshipsBetaBuildLocalizations
		return ret
	}
	return *o.BetaBuildLocalizations
}

// GetBetaBuildLocalizationsOk returns a tuple with the BetaBuildLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetBetaBuildLocalizationsOk() (*BuildRelationshipsBetaBuildLocalizations, bool) {
	if o == nil || IsNil(o.BetaBuildLocalizations) {
		return nil, false
	}
	return o.BetaBuildLocalizations, true
}

// HasBetaBuildLocalizations returns a boolean if a field has been set.
func (o *BuildRelationships) HasBetaBuildLocalizations() bool {
	if o != nil && !IsNil(o.BetaBuildLocalizations) {
		return true
	}

	return false
}

// SetBetaBuildLocalizations gets a reference to the given BuildRelationshipsBetaBuildLocalizations and assigns it to the BetaBuildLocalizations field.
func (o *BuildRelationships) SetBetaBuildLocalizations(v BuildRelationshipsBetaBuildLocalizations) {
	o.BetaBuildLocalizations = &v
}

// GetAppEncryptionDeclaration returns the AppEncryptionDeclaration field value if set, zero value otherwise.
func (o *BuildRelationships) GetAppEncryptionDeclaration() BuildRelationshipsAppEncryptionDeclaration {
	if o == nil || IsNil(o.AppEncryptionDeclaration) {
		var ret BuildRelationshipsAppEncryptionDeclaration
		return ret
	}
	return *o.AppEncryptionDeclaration
}

// GetAppEncryptionDeclarationOk returns a tuple with the AppEncryptionDeclaration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetAppEncryptionDeclarationOk() (*BuildRelationshipsAppEncryptionDeclaration, bool) {
	if o == nil || IsNil(o.AppEncryptionDeclaration) {
		return nil, false
	}
	return o.AppEncryptionDeclaration, true
}

// HasAppEncryptionDeclaration returns a boolean if a field has been set.
func (o *BuildRelationships) HasAppEncryptionDeclaration() bool {
	if o != nil && !IsNil(o.AppEncryptionDeclaration) {
		return true
	}

	return false
}

// SetAppEncryptionDeclaration gets a reference to the given BuildRelationshipsAppEncryptionDeclaration and assigns it to the AppEncryptionDeclaration field.
func (o *BuildRelationships) SetAppEncryptionDeclaration(v BuildRelationshipsAppEncryptionDeclaration) {
	o.AppEncryptionDeclaration = &v
}

// GetBetaAppReviewSubmission returns the BetaAppReviewSubmission field value if set, zero value otherwise.
func (o *BuildRelationships) GetBetaAppReviewSubmission() BuildRelationshipsBetaAppReviewSubmission {
	if o == nil || IsNil(o.BetaAppReviewSubmission) {
		var ret BuildRelationshipsBetaAppReviewSubmission
		return ret
	}
	return *o.BetaAppReviewSubmission
}

// GetBetaAppReviewSubmissionOk returns a tuple with the BetaAppReviewSubmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetBetaAppReviewSubmissionOk() (*BuildRelationshipsBetaAppReviewSubmission, bool) {
	if o == nil || IsNil(o.BetaAppReviewSubmission) {
		return nil, false
	}
	return o.BetaAppReviewSubmission, true
}

// HasBetaAppReviewSubmission returns a boolean if a field has been set.
func (o *BuildRelationships) HasBetaAppReviewSubmission() bool {
	if o != nil && !IsNil(o.BetaAppReviewSubmission) {
		return true
	}

	return false
}

// SetBetaAppReviewSubmission gets a reference to the given BuildRelationshipsBetaAppReviewSubmission and assigns it to the BetaAppReviewSubmission field.
func (o *BuildRelationships) SetBetaAppReviewSubmission(v BuildRelationshipsBetaAppReviewSubmission) {
	o.BetaAppReviewSubmission = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *BuildRelationships) GetApp() AppAvailabilityRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret AppAvailabilityRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetAppOk() (*AppAvailabilityRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *BuildRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppAvailabilityRelationshipsApp and assigns it to the App field.
func (o *BuildRelationships) SetApp(v AppAvailabilityRelationshipsApp) {
	o.App = &v
}

// GetBuildBetaDetail returns the BuildBetaDetail field value if set, zero value otherwise.
func (o *BuildRelationships) GetBuildBetaDetail() BuildRelationshipsBuildBetaDetail {
	if o == nil || IsNil(o.BuildBetaDetail) {
		var ret BuildRelationshipsBuildBetaDetail
		return ret
	}
	return *o.BuildBetaDetail
}

// GetBuildBetaDetailOk returns a tuple with the BuildBetaDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetBuildBetaDetailOk() (*BuildRelationshipsBuildBetaDetail, bool) {
	if o == nil || IsNil(o.BuildBetaDetail) {
		return nil, false
	}
	return o.BuildBetaDetail, true
}

// HasBuildBetaDetail returns a boolean if a field has been set.
func (o *BuildRelationships) HasBuildBetaDetail() bool {
	if o != nil && !IsNil(o.BuildBetaDetail) {
		return true
	}

	return false
}

// SetBuildBetaDetail gets a reference to the given BuildRelationshipsBuildBetaDetail and assigns it to the BuildBetaDetail field.
func (o *BuildRelationships) SetBuildBetaDetail(v BuildRelationshipsBuildBetaDetail) {
	o.BuildBetaDetail = &v
}

// GetAppStoreVersion returns the AppStoreVersion field value if set, zero value otherwise.
func (o *BuildRelationships) GetAppStoreVersion() AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion {
	if o == nil || IsNil(o.AppStoreVersion) {
		var ret AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion
		return ret
	}
	return *o.AppStoreVersion
}

// GetAppStoreVersionOk returns a tuple with the AppStoreVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetAppStoreVersionOk() (*AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion, bool) {
	if o == nil || IsNil(o.AppStoreVersion) {
		return nil, false
	}
	return o.AppStoreVersion, true
}

// HasAppStoreVersion returns a boolean if a field has been set.
func (o *BuildRelationships) HasAppStoreVersion() bool {
	if o != nil && !IsNil(o.AppStoreVersion) {
		return true
	}

	return false
}

// SetAppStoreVersion gets a reference to the given AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion and assigns it to the AppStoreVersion field.
func (o *BuildRelationships) SetAppStoreVersion(v AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) {
	o.AppStoreVersion = &v
}

// GetIcons returns the Icons field value if set, zero value otherwise.
func (o *BuildRelationships) GetIcons() BuildRelationshipsIcons {
	if o == nil || IsNil(o.Icons) {
		var ret BuildRelationshipsIcons
		return ret
	}
	return *o.Icons
}

// GetIconsOk returns a tuple with the Icons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetIconsOk() (*BuildRelationshipsIcons, bool) {
	if o == nil || IsNil(o.Icons) {
		return nil, false
	}
	return o.Icons, true
}

// HasIcons returns a boolean if a field has been set.
func (o *BuildRelationships) HasIcons() bool {
	if o != nil && !IsNil(o.Icons) {
		return true
	}

	return false
}

// SetIcons gets a reference to the given BuildRelationshipsIcons and assigns it to the Icons field.
func (o *BuildRelationships) SetIcons(v BuildRelationshipsIcons) {
	o.Icons = &v
}

// GetBuildBundles returns the BuildBundles field value if set, zero value otherwise.
func (o *BuildRelationships) GetBuildBundles() BuildRelationshipsBuildBundles {
	if o == nil || IsNil(o.BuildBundles) {
		var ret BuildRelationshipsBuildBundles
		return ret
	}
	return *o.BuildBundles
}

// GetBuildBundlesOk returns a tuple with the BuildBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationships) GetBuildBundlesOk() (*BuildRelationshipsBuildBundles, bool) {
	if o == nil || IsNil(o.BuildBundles) {
		return nil, false
	}
	return o.BuildBundles, true
}

// HasBuildBundles returns a boolean if a field has been set.
func (o *BuildRelationships) HasBuildBundles() bool {
	if o != nil && !IsNil(o.BuildBundles) {
		return true
	}

	return false
}

// SetBuildBundles gets a reference to the given BuildRelationshipsBuildBundles and assigns it to the BuildBundles field.
func (o *BuildRelationships) SetBuildBundles(v BuildRelationshipsBuildBundles) {
	o.BuildBundles = &v
}

func (o BuildRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreReleaseVersion) {
		toSerialize["preReleaseVersion"] = o.PreReleaseVersion
	}
	if !IsNil(o.IndividualTesters) {
		toSerialize["individualTesters"] = o.IndividualTesters
	}
	if !IsNil(o.BetaGroups) {
		toSerialize["betaGroups"] = o.BetaGroups
	}
	if !IsNil(o.BetaBuildLocalizations) {
		toSerialize["betaBuildLocalizations"] = o.BetaBuildLocalizations
	}
	if !IsNil(o.AppEncryptionDeclaration) {
		toSerialize["appEncryptionDeclaration"] = o.AppEncryptionDeclaration
	}
	if !IsNil(o.BetaAppReviewSubmission) {
		toSerialize["betaAppReviewSubmission"] = o.BetaAppReviewSubmission
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.BuildBetaDetail) {
		toSerialize["buildBetaDetail"] = o.BuildBetaDetail
	}
	if !IsNil(o.AppStoreVersion) {
		toSerialize["appStoreVersion"] = o.AppStoreVersion
	}
	if !IsNil(o.Icons) {
		toSerialize["icons"] = o.Icons
	}
	if !IsNil(o.BuildBundles) {
		toSerialize["buildBundles"] = o.BuildBundles
	}
	return toSerialize, nil
}

type NullableBuildRelationships struct {
	value *BuildRelationships
	isSet bool
}

func (v NullableBuildRelationships) Get() *BuildRelationships {
	return v.value
}

func (v *NullableBuildRelationships) Set(val *BuildRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildRelationships(val *BuildRelationships) *NullableBuildRelationships {
	return &NullableBuildRelationships{value: val, isSet: true}
}

func (v NullableBuildRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


