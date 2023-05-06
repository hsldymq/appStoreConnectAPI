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

// checks if the AppStoreVersionExperimentTreatmentLocalizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentLocalizationsResponse{}

// AppStoreVersionExperimentTreatmentLocalizationsResponse struct for AppStoreVersionExperimentTreatmentLocalizationsResponse
type AppStoreVersionExperimentTreatmentLocalizationsResponse struct {
	Data []AppStoreVersionExperimentTreatmentLocalization `json:"data"`
	Included []AppStoreVersionExperimentTreatmentLocalizationsResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewAppStoreVersionExperimentTreatmentLocalizationsResponse instantiates a new AppStoreVersionExperimentTreatmentLocalizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentLocalizationsResponse(data []AppStoreVersionExperimentTreatmentLocalization, links PagedDocumentLinks) *AppStoreVersionExperimentTreatmentLocalizationsResponse {
	this := AppStoreVersionExperimentTreatmentLocalizationsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionExperimentTreatmentLocalizationsResponseWithDefaults instantiates a new AppStoreVersionExperimentTreatmentLocalizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentLocalizationsResponseWithDefaults() *AppStoreVersionExperimentTreatmentLocalizationsResponse {
	this := AppStoreVersionExperimentTreatmentLocalizationsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) GetData() []AppStoreVersionExperimentTreatmentLocalization {
	if o == nil {
		var ret []AppStoreVersionExperimentTreatmentLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) GetDataOk() ([]AppStoreVersionExperimentTreatmentLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) SetData(v []AppStoreVersionExperimentTreatmentLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) GetIncluded() []AppStoreVersionExperimentTreatmentLocalizationsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppStoreVersionExperimentTreatmentLocalizationsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) GetIncludedOk() ([]AppStoreVersionExperimentTreatmentLocalizationsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppStoreVersionExperimentTreatmentLocalizationsResponseIncludedInner and assigns it to the Included field.
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) SetIncluded(v []AppStoreVersionExperimentTreatmentLocalizationsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppStoreVersionExperimentTreatmentLocalizationsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppStoreVersionExperimentTreatmentLocalizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentLocalizationsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAppStoreVersionExperimentTreatmentLocalizationsResponse struct {
	value *AppStoreVersionExperimentTreatmentLocalizationsResponse
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationsResponse) Get() *AppStoreVersionExperimentTreatmentLocalizationsResponse {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationsResponse) Set(val *AppStoreVersionExperimentTreatmentLocalizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentLocalizationsResponse(val *AppStoreVersionExperimentTreatmentLocalizationsResponse) *NullableAppStoreVersionExperimentTreatmentLocalizationsResponse {
	return &NullableAppStoreVersionExperimentTreatmentLocalizationsResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentLocalizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentLocalizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


