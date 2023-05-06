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

// checks if the AppStoreVersionAppClipDefaultExperienceLinkageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionAppClipDefaultExperienceLinkageRequest{}

// AppStoreVersionAppClipDefaultExperienceLinkageRequest struct for AppStoreVersionAppClipDefaultExperienceLinkageRequest
type AppStoreVersionAppClipDefaultExperienceLinkageRequest struct {
	Data AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData `json:"data"`
}

// NewAppStoreVersionAppClipDefaultExperienceLinkageRequest instantiates a new AppStoreVersionAppClipDefaultExperienceLinkageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionAppClipDefaultExperienceLinkageRequest(data AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData) *AppStoreVersionAppClipDefaultExperienceLinkageRequest {
	this := AppStoreVersionAppClipDefaultExperienceLinkageRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionAppClipDefaultExperienceLinkageRequestWithDefaults instantiates a new AppStoreVersionAppClipDefaultExperienceLinkageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionAppClipDefaultExperienceLinkageRequestWithDefaults() *AppStoreVersionAppClipDefaultExperienceLinkageRequest {
	this := AppStoreVersionAppClipDefaultExperienceLinkageRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionAppClipDefaultExperienceLinkageRequest) GetData() AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData {
	if o == nil {
		var ret AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionAppClipDefaultExperienceLinkageRequest) GetDataOk() (*AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionAppClipDefaultExperienceLinkageRequest) SetData(v AppClipAppStoreReviewDetailRelationshipsAppClipDefaultExperienceData) {
	o.Data = v
}

func (o AppStoreVersionAppClipDefaultExperienceLinkageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionAppClipDefaultExperienceLinkageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppStoreVersionAppClipDefaultExperienceLinkageRequest struct {
	value *AppStoreVersionAppClipDefaultExperienceLinkageRequest
	isSet bool
}

func (v NullableAppStoreVersionAppClipDefaultExperienceLinkageRequest) Get() *AppStoreVersionAppClipDefaultExperienceLinkageRequest {
	return v.value
}

func (v *NullableAppStoreVersionAppClipDefaultExperienceLinkageRequest) Set(val *AppStoreVersionAppClipDefaultExperienceLinkageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionAppClipDefaultExperienceLinkageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionAppClipDefaultExperienceLinkageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionAppClipDefaultExperienceLinkageRequest(val *AppStoreVersionAppClipDefaultExperienceLinkageRequest) *NullableAppStoreVersionAppClipDefaultExperienceLinkageRequest {
	return &NullableAppStoreVersionAppClipDefaultExperienceLinkageRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionAppClipDefaultExperienceLinkageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionAppClipDefaultExperienceLinkageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


