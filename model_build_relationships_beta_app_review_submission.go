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

// checks if the BuildRelationshipsBetaAppReviewSubmission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildRelationshipsBetaAppReviewSubmission{}

// BuildRelationshipsBetaAppReviewSubmission struct for BuildRelationshipsBetaAppReviewSubmission
type BuildRelationshipsBetaAppReviewSubmission struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *BuildRelationshipsBetaAppReviewSubmissionData `json:"data,omitempty"`
}

// NewBuildRelationshipsBetaAppReviewSubmission instantiates a new BuildRelationshipsBetaAppReviewSubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildRelationshipsBetaAppReviewSubmission() *BuildRelationshipsBetaAppReviewSubmission {
	this := BuildRelationshipsBetaAppReviewSubmission{}
	return &this
}

// NewBuildRelationshipsBetaAppReviewSubmissionWithDefaults instantiates a new BuildRelationshipsBetaAppReviewSubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildRelationshipsBetaAppReviewSubmissionWithDefaults() *BuildRelationshipsBetaAppReviewSubmission {
	this := BuildRelationshipsBetaAppReviewSubmission{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BuildRelationshipsBetaAppReviewSubmission) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationshipsBetaAppReviewSubmission) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BuildRelationshipsBetaAppReviewSubmission) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *BuildRelationshipsBetaAppReviewSubmission) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BuildRelationshipsBetaAppReviewSubmission) GetData() BuildRelationshipsBetaAppReviewSubmissionData {
	if o == nil || IsNil(o.Data) {
		var ret BuildRelationshipsBetaAppReviewSubmissionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildRelationshipsBetaAppReviewSubmission) GetDataOk() (*BuildRelationshipsBetaAppReviewSubmissionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BuildRelationshipsBetaAppReviewSubmission) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BuildRelationshipsBetaAppReviewSubmissionData and assigns it to the Data field.
func (o *BuildRelationshipsBetaAppReviewSubmission) SetData(v BuildRelationshipsBetaAppReviewSubmissionData) {
	o.Data = &v
}

func (o BuildRelationshipsBetaAppReviewSubmission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildRelationshipsBetaAppReviewSubmission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBuildRelationshipsBetaAppReviewSubmission struct {
	value *BuildRelationshipsBetaAppReviewSubmission
	isSet bool
}

func (v NullableBuildRelationshipsBetaAppReviewSubmission) Get() *BuildRelationshipsBetaAppReviewSubmission {
	return v.value
}

func (v *NullableBuildRelationshipsBetaAppReviewSubmission) Set(val *BuildRelationshipsBetaAppReviewSubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildRelationshipsBetaAppReviewSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildRelationshipsBetaAppReviewSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildRelationshipsBetaAppReviewSubmission(val *BuildRelationshipsBetaAppReviewSubmission) *NullableBuildRelationshipsBetaAppReviewSubmission {
	return &NullableBuildRelationshipsBetaAppReviewSubmission{value: val, isSet: true}
}

func (v NullableBuildRelationshipsBetaAppReviewSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildRelationshipsBetaAppReviewSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


