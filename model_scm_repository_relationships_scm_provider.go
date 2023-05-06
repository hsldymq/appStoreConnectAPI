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

// checks if the ScmRepositoryRelationshipsScmProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmRepositoryRelationshipsScmProvider{}

// ScmRepositoryRelationshipsScmProvider struct for ScmRepositoryRelationshipsScmProvider
type ScmRepositoryRelationshipsScmProvider struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *ScmRepositoryRelationshipsScmProviderData `json:"data,omitempty"`
}

// NewScmRepositoryRelationshipsScmProvider instantiates a new ScmRepositoryRelationshipsScmProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmRepositoryRelationshipsScmProvider() *ScmRepositoryRelationshipsScmProvider {
	this := ScmRepositoryRelationshipsScmProvider{}
	return &this
}

// NewScmRepositoryRelationshipsScmProviderWithDefaults instantiates a new ScmRepositoryRelationshipsScmProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmRepositoryRelationshipsScmProviderWithDefaults() *ScmRepositoryRelationshipsScmProvider {
	this := ScmRepositoryRelationshipsScmProvider{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ScmRepositoryRelationshipsScmProvider) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoryRelationshipsScmProvider) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ScmRepositoryRelationshipsScmProvider) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *ScmRepositoryRelationshipsScmProvider) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ScmRepositoryRelationshipsScmProvider) GetData() ScmRepositoryRelationshipsScmProviderData {
	if o == nil || IsNil(o.Data) {
		var ret ScmRepositoryRelationshipsScmProviderData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoryRelationshipsScmProvider) GetDataOk() (*ScmRepositoryRelationshipsScmProviderData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ScmRepositoryRelationshipsScmProvider) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ScmRepositoryRelationshipsScmProviderData and assigns it to the Data field.
func (o *ScmRepositoryRelationshipsScmProvider) SetData(v ScmRepositoryRelationshipsScmProviderData) {
	o.Data = &v
}

func (o ScmRepositoryRelationshipsScmProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmRepositoryRelationshipsScmProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableScmRepositoryRelationshipsScmProvider struct {
	value *ScmRepositoryRelationshipsScmProvider
	isSet bool
}

func (v NullableScmRepositoryRelationshipsScmProvider) Get() *ScmRepositoryRelationshipsScmProvider {
	return v.value
}

func (v *NullableScmRepositoryRelationshipsScmProvider) Set(val *ScmRepositoryRelationshipsScmProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableScmRepositoryRelationshipsScmProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableScmRepositoryRelationshipsScmProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmRepositoryRelationshipsScmProvider(val *ScmRepositoryRelationshipsScmProvider) *NullableScmRepositoryRelationshipsScmProvider {
	return &NullableScmRepositoryRelationshipsScmProvider{value: val, isSet: true}
}

func (v NullableScmRepositoryRelationshipsScmProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmRepositoryRelationshipsScmProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


