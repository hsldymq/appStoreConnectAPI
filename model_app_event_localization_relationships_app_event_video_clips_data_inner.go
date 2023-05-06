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

// checks if the AppEventLocalizationRelationshipsAppEventVideoClipsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationRelationshipsAppEventVideoClipsDataInner{}

// AppEventLocalizationRelationshipsAppEventVideoClipsDataInner struct for AppEventLocalizationRelationshipsAppEventVideoClipsDataInner
type AppEventLocalizationRelationshipsAppEventVideoClipsDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppEventLocalizationRelationshipsAppEventVideoClipsDataInner instantiates a new AppEventLocalizationRelationshipsAppEventVideoClipsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationRelationshipsAppEventVideoClipsDataInner(type_ string, id string) *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner {
	this := AppEventLocalizationRelationshipsAppEventVideoClipsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppEventLocalizationRelationshipsAppEventVideoClipsDataInnerWithDefaults instantiates a new AppEventLocalizationRelationshipsAppEventVideoClipsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationRelationshipsAppEventVideoClipsDataInnerWithDefaults() *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner {
	this := AppEventLocalizationRelationshipsAppEventVideoClipsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner) SetId(v string) {
	o.Id = v
}

func (o AppEventLocalizationRelationshipsAppEventVideoClipsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationRelationshipsAppEventVideoClipsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppEventLocalizationRelationshipsAppEventVideoClipsDataInner struct {
	value *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner
	isSet bool
}

func (v NullableAppEventLocalizationRelationshipsAppEventVideoClipsDataInner) Get() *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner {
	return v.value
}

func (v *NullableAppEventLocalizationRelationshipsAppEventVideoClipsDataInner) Set(val *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationRelationshipsAppEventVideoClipsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationRelationshipsAppEventVideoClipsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationRelationshipsAppEventVideoClipsDataInner(val *AppEventLocalizationRelationshipsAppEventVideoClipsDataInner) *NullableAppEventLocalizationRelationshipsAppEventVideoClipsDataInner {
	return &NullableAppEventLocalizationRelationshipsAppEventVideoClipsDataInner{value: val, isSet: true}
}

func (v NullableAppEventLocalizationRelationshipsAppEventVideoClipsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationRelationshipsAppEventVideoClipsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


