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

// checks if the AppEventVideoClipsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventVideoClipsResponse{}

// AppEventVideoClipsResponse struct for AppEventVideoClipsResponse
type AppEventVideoClipsResponse struct {
	Data []AppEventVideoClip `json:"data"`
	Included []AppEventLocalization `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewAppEventVideoClipsResponse instantiates a new AppEventVideoClipsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventVideoClipsResponse(data []AppEventVideoClip, links PagedDocumentLinks) *AppEventVideoClipsResponse {
	this := AppEventVideoClipsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppEventVideoClipsResponseWithDefaults instantiates a new AppEventVideoClipsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventVideoClipsResponseWithDefaults() *AppEventVideoClipsResponse {
	this := AppEventVideoClipsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppEventVideoClipsResponse) GetData() []AppEventVideoClip {
	if o == nil {
		var ret []AppEventVideoClip
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipsResponse) GetDataOk() ([]AppEventVideoClip, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppEventVideoClipsResponse) SetData(v []AppEventVideoClip) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppEventVideoClipsResponse) GetIncluded() []AppEventLocalization {
	if o == nil || IsNil(o.Included) {
		var ret []AppEventLocalization
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipsResponse) GetIncludedOk() ([]AppEventLocalization, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppEventVideoClipsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppEventLocalization and assigns it to the Included field.
func (o *AppEventVideoClipsResponse) SetIncluded(v []AppEventLocalization) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppEventVideoClipsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppEventVideoClipsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppEventVideoClipsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppEventVideoClipsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppEventVideoClipsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppEventVideoClipsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventVideoClipsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAppEventVideoClipsResponse struct {
	value *AppEventVideoClipsResponse
	isSet bool
}

func (v NullableAppEventVideoClipsResponse) Get() *AppEventVideoClipsResponse {
	return v.value
}

func (v *NullableAppEventVideoClipsResponse) Set(val *AppEventVideoClipsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventVideoClipsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventVideoClipsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventVideoClipsResponse(val *AppEventVideoClipsResponse) *NullableAppEventVideoClipsResponse {
	return &NullableAppEventVideoClipsResponse{value: val, isSet: true}
}

func (v NullableAppEventVideoClipsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventVideoClipsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


