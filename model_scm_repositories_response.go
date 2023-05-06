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

// checks if the ScmRepositoriesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmRepositoriesResponse{}

// ScmRepositoriesResponse struct for ScmRepositoriesResponse
type ScmRepositoriesResponse struct {
	Data []ScmRepository `json:"data"`
	Included []ScmRepositoriesResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewScmRepositoriesResponse instantiates a new ScmRepositoriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmRepositoriesResponse(data []ScmRepository, links PagedDocumentLinks) *ScmRepositoriesResponse {
	this := ScmRepositoriesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewScmRepositoriesResponseWithDefaults instantiates a new ScmRepositoriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmRepositoriesResponseWithDefaults() *ScmRepositoriesResponse {
	this := ScmRepositoriesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ScmRepositoriesResponse) GetData() []ScmRepository {
	if o == nil {
		var ret []ScmRepository
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ScmRepositoriesResponse) GetDataOk() ([]ScmRepository, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ScmRepositoriesResponse) SetData(v []ScmRepository) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ScmRepositoriesResponse) GetIncluded() []ScmRepositoriesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []ScmRepositoriesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoriesResponse) GetIncludedOk() ([]ScmRepositoriesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ScmRepositoriesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []ScmRepositoriesResponseIncludedInner and assigns it to the Included field.
func (o *ScmRepositoriesResponse) SetIncluded(v []ScmRepositoriesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *ScmRepositoriesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ScmRepositoriesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ScmRepositoriesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScmRepositoriesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoriesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScmRepositoriesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *ScmRepositoriesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o ScmRepositoriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmRepositoriesResponse) ToMap() (map[string]interface{}, error) {
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

type NullableScmRepositoriesResponse struct {
	value *ScmRepositoriesResponse
	isSet bool
}

func (v NullableScmRepositoriesResponse) Get() *ScmRepositoriesResponse {
	return v.value
}

func (v *NullableScmRepositoriesResponse) Set(val *ScmRepositoriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScmRepositoriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScmRepositoriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmRepositoriesResponse(val *ScmRepositoriesResponse) *NullableScmRepositoriesResponse {
	return &NullableScmRepositoriesResponse{value: val, isSet: true}
}

func (v NullableScmRepositoriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmRepositoriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

