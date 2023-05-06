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

// checks if the InAppPurchaseLocalizationCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseLocalizationCreateRequestDataAttributes{}

// InAppPurchaseLocalizationCreateRequestDataAttributes struct for InAppPurchaseLocalizationCreateRequestDataAttributes
type InAppPurchaseLocalizationCreateRequestDataAttributes struct {
	Name string `json:"name"`
	Locale string `json:"locale"`
	Description *string `json:"description,omitempty"`
}

// NewInAppPurchaseLocalizationCreateRequestDataAttributes instantiates a new InAppPurchaseLocalizationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseLocalizationCreateRequestDataAttributes(name string, locale string) *InAppPurchaseLocalizationCreateRequestDataAttributes {
	this := InAppPurchaseLocalizationCreateRequestDataAttributes{}
	this.Name = name
	this.Locale = locale
	return &this
}

// NewInAppPurchaseLocalizationCreateRequestDataAttributesWithDefaults instantiates a new InAppPurchaseLocalizationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseLocalizationCreateRequestDataAttributesWithDefaults() *InAppPurchaseLocalizationCreateRequestDataAttributes {
	this := InAppPurchaseLocalizationCreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *InAppPurchaseLocalizationCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InAppPurchaseLocalizationCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetLocale returns the Locale field value
func (o *InAppPurchaseLocalizationCreateRequestDataAttributes) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationCreateRequestDataAttributes) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *InAppPurchaseLocalizationCreateRequestDataAttributes) SetLocale(v string) {
	o.Locale = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InAppPurchaseLocalizationCreateRequestDataAttributes) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationCreateRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InAppPurchaseLocalizationCreateRequestDataAttributes) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InAppPurchaseLocalizationCreateRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

func (o InAppPurchaseLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseLocalizationCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["locale"] = o.Locale
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableInAppPurchaseLocalizationCreateRequestDataAttributes struct {
	value *InAppPurchaseLocalizationCreateRequestDataAttributes
	isSet bool
}

func (v NullableInAppPurchaseLocalizationCreateRequestDataAttributes) Get() *InAppPurchaseLocalizationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableInAppPurchaseLocalizationCreateRequestDataAttributes) Set(val *InAppPurchaseLocalizationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseLocalizationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseLocalizationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseLocalizationCreateRequestDataAttributes(val *InAppPurchaseLocalizationCreateRequestDataAttributes) *NullableInAppPurchaseLocalizationCreateRequestDataAttributes {
	return &NullableInAppPurchaseLocalizationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableInAppPurchaseLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseLocalizationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


