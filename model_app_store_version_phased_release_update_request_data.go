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

// checks if the AppStoreVersionPhasedReleaseUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionPhasedReleaseUpdateRequestData{}

// AppStoreVersionPhasedReleaseUpdateRequestData struct for AppStoreVersionPhasedReleaseUpdateRequestData
type AppStoreVersionPhasedReleaseUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppStoreVersionPhasedReleaseCreateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewAppStoreVersionPhasedReleaseUpdateRequestData instantiates a new AppStoreVersionPhasedReleaseUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionPhasedReleaseUpdateRequestData(type_ string, id string) *AppStoreVersionPhasedReleaseUpdateRequestData {
	this := AppStoreVersionPhasedReleaseUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreVersionPhasedReleaseUpdateRequestDataWithDefaults instantiates a new AppStoreVersionPhasedReleaseUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionPhasedReleaseUpdateRequestDataWithDefaults() *AppStoreVersionPhasedReleaseUpdateRequestData {
	this := AppStoreVersionPhasedReleaseUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionPhasedReleaseUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionPhasedReleaseUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersionPhasedReleaseUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersionPhasedReleaseUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppStoreVersionPhasedReleaseUpdateRequestData) GetAttributes() AppStoreVersionPhasedReleaseCreateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppStoreVersionPhasedReleaseCreateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseUpdateRequestData) GetAttributesOk() (*AppStoreVersionPhasedReleaseCreateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppStoreVersionPhasedReleaseUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppStoreVersionPhasedReleaseCreateRequestDataAttributes and assigns it to the Attributes field.
func (o *AppStoreVersionPhasedReleaseUpdateRequestData) SetAttributes(v AppStoreVersionPhasedReleaseCreateRequestDataAttributes) {
	o.Attributes = &v
}

func (o AppStoreVersionPhasedReleaseUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionPhasedReleaseUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableAppStoreVersionPhasedReleaseUpdateRequestData struct {
	value *AppStoreVersionPhasedReleaseUpdateRequestData
	isSet bool
}

func (v NullableAppStoreVersionPhasedReleaseUpdateRequestData) Get() *AppStoreVersionPhasedReleaseUpdateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionPhasedReleaseUpdateRequestData) Set(val *AppStoreVersionPhasedReleaseUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionPhasedReleaseUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionPhasedReleaseUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionPhasedReleaseUpdateRequestData(val *AppStoreVersionPhasedReleaseUpdateRequestData) *NullableAppStoreVersionPhasedReleaseUpdateRequestData {
	return &NullableAppStoreVersionPhasedReleaseUpdateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionPhasedReleaseUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionPhasedReleaseUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


