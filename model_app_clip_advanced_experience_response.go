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

// checks if the AppClipAdvancedExperienceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceResponse{}

// AppClipAdvancedExperienceResponse struct for AppClipAdvancedExperienceResponse
type AppClipAdvancedExperienceResponse struct {
	Data AppClipAdvancedExperience `json:"data"`
	Included []AppClipAdvancedExperiencesResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewAppClipAdvancedExperienceResponse instantiates a new AppClipAdvancedExperienceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceResponse(data AppClipAdvancedExperience, links DocumentLinks) *AppClipAdvancedExperienceResponse {
	this := AppClipAdvancedExperienceResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppClipAdvancedExperienceResponseWithDefaults instantiates a new AppClipAdvancedExperienceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceResponseWithDefaults() *AppClipAdvancedExperienceResponse {
	this := AppClipAdvancedExperienceResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipAdvancedExperienceResponse) GetData() AppClipAdvancedExperience {
	if o == nil {
		var ret AppClipAdvancedExperience
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceResponse) GetDataOk() (*AppClipAdvancedExperience, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipAdvancedExperienceResponse) SetData(v AppClipAdvancedExperience) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceResponse) GetIncluded() []AppClipAdvancedExperiencesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppClipAdvancedExperiencesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceResponse) GetIncludedOk() ([]AppClipAdvancedExperiencesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppClipAdvancedExperiencesResponseIncludedInner and assigns it to the Included field.
func (o *AppClipAdvancedExperienceResponse) SetIncluded(v []AppClipAdvancedExperiencesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppClipAdvancedExperienceResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClipAdvancedExperienceResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppClipAdvancedExperienceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceResponse struct {
	value *AppClipAdvancedExperienceResponse
	isSet bool
}

func (v NullableAppClipAdvancedExperienceResponse) Get() *AppClipAdvancedExperienceResponse {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceResponse) Set(val *AppClipAdvancedExperienceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceResponse(val *AppClipAdvancedExperienceResponse) *NullableAppClipAdvancedExperienceResponse {
	return &NullableAppClipAdvancedExperienceResponse{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


