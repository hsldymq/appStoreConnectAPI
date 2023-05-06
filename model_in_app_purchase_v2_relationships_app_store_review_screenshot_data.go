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

// checks if the InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData{}

// InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData struct for InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData
type InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData instantiates a new InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData(type_ string, id string) *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData {
	this := InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewInAppPurchaseV2RelationshipsAppStoreReviewScreenshotDataWithDefaults instantiates a new InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2RelationshipsAppStoreReviewScreenshotDataWithDefaults() *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData {
	this := InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData{}
	return &this
}

// GetType returns the Type field value
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) SetId(v string) {
	o.Id = v
}

func (o InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData struct {
	value *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData
	isSet bool
}

func (v NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) Get() *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData {
	return v.value
}

func (v *NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) Set(val *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData(val *InAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) *NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData {
	return &NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2RelationshipsAppStoreReviewScreenshotData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


