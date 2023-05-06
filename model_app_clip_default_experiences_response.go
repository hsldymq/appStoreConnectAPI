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

// checks if the AppClipDefaultExperiencesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperiencesResponse{}

// AppClipDefaultExperiencesResponse struct for AppClipDefaultExperiencesResponse
type AppClipDefaultExperiencesResponse struct {
	Data []AppClipDefaultExperience `json:"data"`
	Included []AppClipDefaultExperiencesResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewAppClipDefaultExperiencesResponse instantiates a new AppClipDefaultExperiencesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperiencesResponse(data []AppClipDefaultExperience, links PagedDocumentLinks) *AppClipDefaultExperiencesResponse {
	this := AppClipDefaultExperiencesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppClipDefaultExperiencesResponseWithDefaults instantiates a new AppClipDefaultExperiencesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperiencesResponseWithDefaults() *AppClipDefaultExperiencesResponse {
	this := AppClipDefaultExperiencesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipDefaultExperiencesResponse) GetData() []AppClipDefaultExperience {
	if o == nil {
		var ret []AppClipDefaultExperience
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperiencesResponse) GetDataOk() ([]AppClipDefaultExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppClipDefaultExperiencesResponse) SetData(v []AppClipDefaultExperience) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppClipDefaultExperiencesResponse) GetIncluded() []AppClipDefaultExperiencesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppClipDefaultExperiencesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperiencesResponse) GetIncludedOk() ([]AppClipDefaultExperiencesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppClipDefaultExperiencesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppClipDefaultExperiencesResponseIncludedInner and assigns it to the Included field.
func (o *AppClipDefaultExperiencesResponse) SetIncluded(v []AppClipDefaultExperiencesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppClipDefaultExperiencesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperiencesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClipDefaultExperiencesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppClipDefaultExperiencesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperiencesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppClipDefaultExperiencesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppClipDefaultExperiencesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppClipDefaultExperiencesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperiencesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableAppClipDefaultExperiencesResponse struct {
	value *AppClipDefaultExperiencesResponse
	isSet bool
}

func (v NullableAppClipDefaultExperiencesResponse) Get() *AppClipDefaultExperiencesResponse {
	return v.value
}

func (v *NullableAppClipDefaultExperiencesResponse) Set(val *AppClipDefaultExperiencesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperiencesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperiencesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperiencesResponse(val *AppClipDefaultExperiencesResponse) *NullableAppClipDefaultExperiencesResponse {
	return &NullableAppClipDefaultExperiencesResponse{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperiencesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperiencesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


