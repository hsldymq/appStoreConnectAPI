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

// checks if the AppPriceInlineCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPriceInlineCreate{}

// AppPriceInlineCreate struct for AppPriceInlineCreate
type AppPriceInlineCreate struct {
	Type string `json:"type"`
	Id *string `json:"id,omitempty"`
}

// NewAppPriceInlineCreate instantiates a new AppPriceInlineCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceInlineCreate(type_ string) *AppPriceInlineCreate {
	this := AppPriceInlineCreate{}
	this.Type = type_
	return &this
}

// NewAppPriceInlineCreateWithDefaults instantiates a new AppPriceInlineCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceInlineCreateWithDefaults() *AppPriceInlineCreate {
	this := AppPriceInlineCreate{}
	return &this
}

// GetType returns the Type field value
func (o *AppPriceInlineCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPriceInlineCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPriceInlineCreate) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppPriceInlineCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceInlineCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppPriceInlineCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppPriceInlineCreate) SetId(v string) {
	o.Id = &v
}

func (o AppPriceInlineCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPriceInlineCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAppPriceInlineCreate struct {
	value *AppPriceInlineCreate
	isSet bool
}

func (v NullableAppPriceInlineCreate) Get() *AppPriceInlineCreate {
	return v.value
}

func (v *NullableAppPriceInlineCreate) Set(val *AppPriceInlineCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceInlineCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceInlineCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceInlineCreate(val *AppPriceInlineCreate) *NullableAppPriceInlineCreate {
	return &NullableAppPriceInlineCreate{value: val, isSet: true}
}

func (v NullableAppPriceInlineCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceInlineCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


