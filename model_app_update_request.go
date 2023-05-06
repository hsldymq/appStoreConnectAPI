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

// checks if the AppUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUpdateRequest{}

// AppUpdateRequest struct for AppUpdateRequest
type AppUpdateRequest struct {
	Data AppUpdateRequestData `json:"data"`
	Included []AppPriceInlineCreate `json:"included,omitempty"`
}

// NewAppUpdateRequest instantiates a new AppUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUpdateRequest(data AppUpdateRequestData) *AppUpdateRequest {
	this := AppUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppUpdateRequestWithDefaults instantiates a new AppUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUpdateRequestWithDefaults() *AppUpdateRequest {
	this := AppUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppUpdateRequest) GetData() AppUpdateRequestData {
	if o == nil {
		var ret AppUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppUpdateRequest) GetDataOk() (*AppUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppUpdateRequest) SetData(v AppUpdateRequestData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppUpdateRequest) GetIncluded() []AppPriceInlineCreate {
	if o == nil || IsNil(o.Included) {
		var ret []AppPriceInlineCreate
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUpdateRequest) GetIncludedOk() ([]AppPriceInlineCreate, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppUpdateRequest) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppPriceInlineCreate and assigns it to the Included field.
func (o *AppUpdateRequest) SetIncluded(v []AppPriceInlineCreate) {
	o.Included = v
}

func (o AppUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

type NullableAppUpdateRequest struct {
	value *AppUpdateRequest
	isSet bool
}

func (v NullableAppUpdateRequest) Get() *AppUpdateRequest {
	return v.value
}

func (v *NullableAppUpdateRequest) Set(val *AppUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUpdateRequest(val *AppUpdateRequest) *NullableAppUpdateRequest {
	return &NullableAppUpdateRequest{value: val, isSet: true}
}

func (v NullableAppUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


