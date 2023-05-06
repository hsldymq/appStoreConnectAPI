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

// checks if the AppCategoryRelationshipsSubcategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCategoryRelationshipsSubcategories{}

// AppCategoryRelationshipsSubcategories struct for AppCategoryRelationshipsSubcategories
type AppCategoryRelationshipsSubcategories struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Meta *PagingInformation `json:"meta,omitempty"`
	Data []AppCategoryRelationshipsSubcategoriesDataInner `json:"data,omitempty"`
}

// NewAppCategoryRelationshipsSubcategories instantiates a new AppCategoryRelationshipsSubcategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCategoryRelationshipsSubcategories() *AppCategoryRelationshipsSubcategories {
	this := AppCategoryRelationshipsSubcategories{}
	return &this
}

// NewAppCategoryRelationshipsSubcategoriesWithDefaults instantiates a new AppCategoryRelationshipsSubcategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCategoryRelationshipsSubcategoriesWithDefaults() *AppCategoryRelationshipsSubcategories {
	this := AppCategoryRelationshipsSubcategories{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppCategoryRelationshipsSubcategories) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationshipsSubcategories) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppCategoryRelationshipsSubcategories) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *AppCategoryRelationshipsSubcategories) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppCategoryRelationshipsSubcategories) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationshipsSubcategories) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppCategoryRelationshipsSubcategories) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppCategoryRelationshipsSubcategories) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppCategoryRelationshipsSubcategories) GetData() []AppCategoryRelationshipsSubcategoriesDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []AppCategoryRelationshipsSubcategoriesDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationshipsSubcategories) GetDataOk() ([]AppCategoryRelationshipsSubcategoriesDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppCategoryRelationshipsSubcategories) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppCategoryRelationshipsSubcategoriesDataInner and assigns it to the Data field.
func (o *AppCategoryRelationshipsSubcategories) SetData(v []AppCategoryRelationshipsSubcategoriesDataInner) {
	o.Data = v
}

func (o AppCategoryRelationshipsSubcategories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCategoryRelationshipsSubcategories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppCategoryRelationshipsSubcategories struct {
	value *AppCategoryRelationshipsSubcategories
	isSet bool
}

func (v NullableAppCategoryRelationshipsSubcategories) Get() *AppCategoryRelationshipsSubcategories {
	return v.value
}

func (v *NullableAppCategoryRelationshipsSubcategories) Set(val *AppCategoryRelationshipsSubcategories) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCategoryRelationshipsSubcategories) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCategoryRelationshipsSubcategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCategoryRelationshipsSubcategories(val *AppCategoryRelationshipsSubcategories) *NullableAppCategoryRelationshipsSubcategories {
	return &NullableAppCategoryRelationshipsSubcategories{value: val, isSet: true}
}

func (v NullableAppCategoryRelationshipsSubcategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCategoryRelationshipsSubcategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


