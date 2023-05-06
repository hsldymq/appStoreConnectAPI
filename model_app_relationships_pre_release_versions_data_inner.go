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

// checks if the AppRelationshipsPreReleaseVersionsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsPreReleaseVersionsDataInner{}

// AppRelationshipsPreReleaseVersionsDataInner struct for AppRelationshipsPreReleaseVersionsDataInner
type AppRelationshipsPreReleaseVersionsDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppRelationshipsPreReleaseVersionsDataInner instantiates a new AppRelationshipsPreReleaseVersionsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsPreReleaseVersionsDataInner(type_ string, id string) *AppRelationshipsPreReleaseVersionsDataInner {
	this := AppRelationshipsPreReleaseVersionsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppRelationshipsPreReleaseVersionsDataInnerWithDefaults instantiates a new AppRelationshipsPreReleaseVersionsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsPreReleaseVersionsDataInnerWithDefaults() *AppRelationshipsPreReleaseVersionsDataInner {
	this := AppRelationshipsPreReleaseVersionsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppRelationshipsPreReleaseVersionsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsPreReleaseVersionsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppRelationshipsPreReleaseVersionsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppRelationshipsPreReleaseVersionsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsPreReleaseVersionsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppRelationshipsPreReleaseVersionsDataInner) SetId(v string) {
	o.Id = v
}

func (o AppRelationshipsPreReleaseVersionsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsPreReleaseVersionsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppRelationshipsPreReleaseVersionsDataInner struct {
	value *AppRelationshipsPreReleaseVersionsDataInner
	isSet bool
}

func (v NullableAppRelationshipsPreReleaseVersionsDataInner) Get() *AppRelationshipsPreReleaseVersionsDataInner {
	return v.value
}

func (v *NullableAppRelationshipsPreReleaseVersionsDataInner) Set(val *AppRelationshipsPreReleaseVersionsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsPreReleaseVersionsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsPreReleaseVersionsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsPreReleaseVersionsDataInner(val *AppRelationshipsPreReleaseVersionsDataInner) *NullableAppRelationshipsPreReleaseVersionsDataInner {
	return &NullableAppRelationshipsPreReleaseVersionsDataInner{value: val, isSet: true}
}

func (v NullableAppRelationshipsPreReleaseVersionsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsPreReleaseVersionsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


