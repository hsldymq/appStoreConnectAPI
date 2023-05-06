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

// checks if the AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations{}

// AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations struct for AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations
type AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations struct {
	Data []AppCustomProductPageVersionRelationshipsAppCustomProductPageLocalizationsDataInner `json:"data,omitempty"`
}

// NewAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations instantiates a new AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations() *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations {
	this := AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations{}
	return &this
}

// NewAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizationsWithDefaults instantiates a new AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizationsWithDefaults() *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations {
	this := AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) GetData() []AppCustomProductPageVersionRelationshipsAppCustomProductPageLocalizationsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []AppCustomProductPageVersionRelationshipsAppCustomProductPageLocalizationsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) GetDataOk() ([]AppCustomProductPageVersionRelationshipsAppCustomProductPageLocalizationsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppCustomProductPageVersionRelationshipsAppCustomProductPageLocalizationsDataInner and assigns it to the Data field.
func (o *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) SetData(v []AppCustomProductPageVersionRelationshipsAppCustomProductPageLocalizationsDataInner) {
	o.Data = v
}

func (o AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations struct {
	value *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations
	isSet bool
}

func (v NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) Get() *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations {
	return v.value
}

func (v *NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) Set(val *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations(val *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) *NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations {
	return &NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations{value: val, isSet: true}
}

func (v NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageLocalizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


