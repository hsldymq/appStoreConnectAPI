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

// checks if the AppCustomProductPagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPagesResponse{}

// AppCustomProductPagesResponse struct for AppCustomProductPagesResponse
type AppCustomProductPagesResponse struct {
	Data []AppCustomProductPage `json:"data"`
	Included []AppCustomProductPagesResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewAppCustomProductPagesResponse instantiates a new AppCustomProductPagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPagesResponse(data []AppCustomProductPage, links PagedDocumentLinks) *AppCustomProductPagesResponse {
	this := AppCustomProductPagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppCustomProductPagesResponseWithDefaults instantiates a new AppCustomProductPagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPagesResponseWithDefaults() *AppCustomProductPagesResponse {
	this := AppCustomProductPagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppCustomProductPagesResponse) GetData() []AppCustomProductPage {
	if o == nil {
		var ret []AppCustomProductPage
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPagesResponse) GetDataOk() ([]AppCustomProductPage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppCustomProductPagesResponse) SetData(v []AppCustomProductPage) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppCustomProductPagesResponse) GetIncluded() []AppCustomProductPagesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppCustomProductPagesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPagesResponse) GetIncludedOk() ([]AppCustomProductPagesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppCustomProductPagesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppCustomProductPagesResponseIncludedInner and assigns it to the Included field.
func (o *AppCustomProductPagesResponse) SetIncluded(v []AppCustomProductPagesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppCustomProductPagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppCustomProductPagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppCustomProductPagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppCustomProductPagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppCustomProductPagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppCustomProductPagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPagesResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAppCustomProductPagesResponse struct {
	value *AppCustomProductPagesResponse
	isSet bool
}

func (v NullableAppCustomProductPagesResponse) Get() *AppCustomProductPagesResponse {
	return v.value
}

func (v *NullableAppCustomProductPagesResponse) Set(val *AppCustomProductPagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPagesResponse(val *AppCustomProductPagesResponse) *NullableAppCustomProductPagesResponse {
	return &NullableAppCustomProductPagesResponse{value: val, isSet: true}
}

func (v NullableAppCustomProductPagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


