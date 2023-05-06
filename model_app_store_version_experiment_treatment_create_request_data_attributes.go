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

// checks if the AppStoreVersionExperimentTreatmentCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentCreateRequestDataAttributes{}

// AppStoreVersionExperimentTreatmentCreateRequestDataAttributes struct for AppStoreVersionExperimentTreatmentCreateRequestDataAttributes
type AppStoreVersionExperimentTreatmentCreateRequestDataAttributes struct {
	Name string `json:"name"`
	AppIconName *string `json:"appIconName,omitempty"`
}

// NewAppStoreVersionExperimentTreatmentCreateRequestDataAttributes instantiates a new AppStoreVersionExperimentTreatmentCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentCreateRequestDataAttributes(name string) *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes {
	this := AppStoreVersionExperimentTreatmentCreateRequestDataAttributes{}
	this.Name = name
	return &this
}

// NewAppStoreVersionExperimentTreatmentCreateRequestDataAttributesWithDefaults instantiates a new AppStoreVersionExperimentTreatmentCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentCreateRequestDataAttributesWithDefaults() *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes {
	this := AppStoreVersionExperimentTreatmentCreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetAppIconName returns the AppIconName field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) GetAppIconName() string {
	if o == nil || IsNil(o.AppIconName) {
		var ret string
		return ret
	}
	return *o.AppIconName
}

// GetAppIconNameOk returns a tuple with the AppIconName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) GetAppIconNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppIconName) {
		return nil, false
	}
	return o.AppIconName, true
}

// HasAppIconName returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) HasAppIconName() bool {
	if o != nil && !IsNil(o.AppIconName) {
		return true
	}

	return false
}

// SetAppIconName gets a reference to the given string and assigns it to the AppIconName field.
func (o *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) SetAppIconName(v string) {
	o.AppIconName = &v
}

func (o AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.AppIconName) {
		toSerialize["appIconName"] = o.AppIconName
	}
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentCreateRequestDataAttributes struct {
	value *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequestDataAttributes) Get() *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequestDataAttributes) Set(val *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentCreateRequestDataAttributes(val *AppStoreVersionExperimentTreatmentCreateRequestDataAttributes) *NullableAppStoreVersionExperimentTreatmentCreateRequestDataAttributes {
	return &NullableAppStoreVersionExperimentTreatmentCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


