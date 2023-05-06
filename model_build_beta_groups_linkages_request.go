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

// checks if the BuildBetaGroupsLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBetaGroupsLinkagesRequest{}

// BuildBetaGroupsLinkagesRequest struct for BuildBetaGroupsLinkagesRequest
type BuildBetaGroupsLinkagesRequest struct {
	Data []AppRelationshipsBetaGroupsDataInner `json:"data"`
}

// NewBuildBetaGroupsLinkagesRequest instantiates a new BuildBetaGroupsLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaGroupsLinkagesRequest(data []AppRelationshipsBetaGroupsDataInner) *BuildBetaGroupsLinkagesRequest {
	this := BuildBetaGroupsLinkagesRequest{}
	this.Data = data
	return &this
}

// NewBuildBetaGroupsLinkagesRequestWithDefaults instantiates a new BuildBetaGroupsLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaGroupsLinkagesRequestWithDefaults() *BuildBetaGroupsLinkagesRequest {
	this := BuildBetaGroupsLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BuildBetaGroupsLinkagesRequest) GetData() []AppRelationshipsBetaGroupsDataInner {
	if o == nil {
		var ret []AppRelationshipsBetaGroupsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildBetaGroupsLinkagesRequest) GetDataOk() ([]AppRelationshipsBetaGroupsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BuildBetaGroupsLinkagesRequest) SetData(v []AppRelationshipsBetaGroupsDataInner) {
	o.Data = v
}

func (o BuildBetaGroupsLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBetaGroupsLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBuildBetaGroupsLinkagesRequest struct {
	value *BuildBetaGroupsLinkagesRequest
	isSet bool
}

func (v NullableBuildBetaGroupsLinkagesRequest) Get() *BuildBetaGroupsLinkagesRequest {
	return v.value
}

func (v *NullableBuildBetaGroupsLinkagesRequest) Set(val *BuildBetaGroupsLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaGroupsLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaGroupsLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaGroupsLinkagesRequest(val *BuildBetaGroupsLinkagesRequest) *NullableBuildBetaGroupsLinkagesRequest {
	return &NullableBuildBetaGroupsLinkagesRequest{value: val, isSet: true}
}

func (v NullableBuildBetaGroupsLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaGroupsLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


