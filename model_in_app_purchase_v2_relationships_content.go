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

// checks if the InAppPurchaseV2RelationshipsContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2RelationshipsContent{}

// InAppPurchaseV2RelationshipsContent struct for InAppPurchaseV2RelationshipsContent
type InAppPurchaseV2RelationshipsContent struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *InAppPurchaseV2RelationshipsContentData `json:"data,omitempty"`
}

// NewInAppPurchaseV2RelationshipsContent instantiates a new InAppPurchaseV2RelationshipsContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2RelationshipsContent() *InAppPurchaseV2RelationshipsContent {
	this := InAppPurchaseV2RelationshipsContent{}
	return &this
}

// NewInAppPurchaseV2RelationshipsContentWithDefaults instantiates a new InAppPurchaseV2RelationshipsContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2RelationshipsContentWithDefaults() *InAppPurchaseV2RelationshipsContent {
	this := InAppPurchaseV2RelationshipsContent{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InAppPurchaseV2RelationshipsContent) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsContent) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InAppPurchaseV2RelationshipsContent) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *InAppPurchaseV2RelationshipsContent) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InAppPurchaseV2RelationshipsContent) GetData() InAppPurchaseV2RelationshipsContentData {
	if o == nil || IsNil(o.Data) {
		var ret InAppPurchaseV2RelationshipsContentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsContent) GetDataOk() (*InAppPurchaseV2RelationshipsContentData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InAppPurchaseV2RelationshipsContent) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given InAppPurchaseV2RelationshipsContentData and assigns it to the Data field.
func (o *InAppPurchaseV2RelationshipsContent) SetData(v InAppPurchaseV2RelationshipsContentData) {
	o.Data = &v
}

func (o InAppPurchaseV2RelationshipsContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2RelationshipsContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableInAppPurchaseV2RelationshipsContent struct {
	value *InAppPurchaseV2RelationshipsContent
	isSet bool
}

func (v NullableInAppPurchaseV2RelationshipsContent) Get() *InAppPurchaseV2RelationshipsContent {
	return v.value
}

func (v *NullableInAppPurchaseV2RelationshipsContent) Set(val *InAppPurchaseV2RelationshipsContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2RelationshipsContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2RelationshipsContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2RelationshipsContent(val *InAppPurchaseV2RelationshipsContent) *NullableInAppPurchaseV2RelationshipsContent {
	return &NullableInAppPurchaseV2RelationshipsContent{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2RelationshipsContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2RelationshipsContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

