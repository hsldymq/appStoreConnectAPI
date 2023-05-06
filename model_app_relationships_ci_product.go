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

// checks if the AppRelationshipsCiProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsCiProduct{}

// AppRelationshipsCiProduct struct for AppRelationshipsCiProduct
type AppRelationshipsCiProduct struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *AppRelationshipsCiProductData `json:"data,omitempty"`
}

// NewAppRelationshipsCiProduct instantiates a new AppRelationshipsCiProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsCiProduct() *AppRelationshipsCiProduct {
	this := AppRelationshipsCiProduct{}
	return &this
}

// NewAppRelationshipsCiProductWithDefaults instantiates a new AppRelationshipsCiProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsCiProductWithDefaults() *AppRelationshipsCiProduct {
	this := AppRelationshipsCiProduct{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppRelationshipsCiProduct) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsCiProduct) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppRelationshipsCiProduct) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *AppRelationshipsCiProduct) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppRelationshipsCiProduct) GetData() AppRelationshipsCiProductData {
	if o == nil || IsNil(o.Data) {
		var ret AppRelationshipsCiProductData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRelationshipsCiProduct) GetDataOk() (*AppRelationshipsCiProductData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppRelationshipsCiProduct) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppRelationshipsCiProductData and assigns it to the Data field.
func (o *AppRelationshipsCiProduct) SetData(v AppRelationshipsCiProductData) {
	o.Data = &v
}

func (o AppRelationshipsCiProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsCiProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppRelationshipsCiProduct struct {
	value *AppRelationshipsCiProduct
	isSet bool
}

func (v NullableAppRelationshipsCiProduct) Get() *AppRelationshipsCiProduct {
	return v.value
}

func (v *NullableAppRelationshipsCiProduct) Set(val *AppRelationshipsCiProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsCiProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsCiProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsCiProduct(val *AppRelationshipsCiProduct) *NullableAppRelationshipsCiProduct {
	return &NullableAppRelationshipsCiProduct{value: val, isSet: true}
}

func (v NullableAppRelationshipsCiProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsCiProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


