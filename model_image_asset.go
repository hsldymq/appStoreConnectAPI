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

// checks if the ImageAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageAsset{}

// ImageAsset struct for ImageAsset
type ImageAsset struct {
	TemplateUrl *string `json:"templateUrl,omitempty"`
	Width *int32 `json:"width,omitempty"`
	Height *int32 `json:"height,omitempty"`
}

// NewImageAsset instantiates a new ImageAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageAsset() *ImageAsset {
	this := ImageAsset{}
	return &this
}

// NewImageAssetWithDefaults instantiates a new ImageAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageAssetWithDefaults() *ImageAsset {
	this := ImageAsset{}
	return &this
}

// GetTemplateUrl returns the TemplateUrl field value if set, zero value otherwise.
func (o *ImageAsset) GetTemplateUrl() string {
	if o == nil || IsNil(o.TemplateUrl) {
		var ret string
		return ret
	}
	return *o.TemplateUrl
}

// GetTemplateUrlOk returns a tuple with the TemplateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAsset) GetTemplateUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateUrl) {
		return nil, false
	}
	return o.TemplateUrl, true
}

// HasTemplateUrl returns a boolean if a field has been set.
func (o *ImageAsset) HasTemplateUrl() bool {
	if o != nil && !IsNil(o.TemplateUrl) {
		return true
	}

	return false
}

// SetTemplateUrl gets a reference to the given string and assigns it to the TemplateUrl field.
func (o *ImageAsset) SetTemplateUrl(v string) {
	o.TemplateUrl = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ImageAsset) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAsset) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ImageAsset) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *ImageAsset) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ImageAsset) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAsset) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ImageAsset) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ImageAsset) SetHeight(v int32) {
	o.Height = &v
}

func (o ImageAsset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateUrl) {
		toSerialize["templateUrl"] = o.TemplateUrl
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableImageAsset struct {
	value *ImageAsset
	isSet bool
}

func (v NullableImageAsset) Get() *ImageAsset {
	return v.value
}

func (v *NullableImageAsset) Set(val *ImageAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableImageAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableImageAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageAsset(val *ImageAsset) *NullableImageAsset {
	return &NullableImageAsset{value: val, isSet: true}
}

func (v NullableImageAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


