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

// checks if the AppPriceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPriceResponse{}

// AppPriceResponse struct for AppPriceResponse
type AppPriceResponse struct {
	// Deprecated
	Data AppPrice `json:"data"`
	Included []AppPricesResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewAppPriceResponse instantiates a new AppPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceResponse(data AppPrice, links DocumentLinks) *AppPriceResponse {
	this := AppPriceResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppPriceResponseWithDefaults instantiates a new AppPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceResponseWithDefaults() *AppPriceResponse {
	this := AppPriceResponse{}
	return &this
}

// GetData returns the Data field value
// Deprecated
func (o *AppPriceResponse) GetData() AppPrice {
	if o == nil {
		var ret AppPrice
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *AppPriceResponse) GetDataOk() (*AppPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
// Deprecated
func (o *AppPriceResponse) SetData(v AppPrice) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppPriceResponse) GetIncluded() []AppPricesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppPricesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceResponse) GetIncludedOk() ([]AppPricesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppPriceResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppPricesResponseIncludedInner and assigns it to the Included field.
func (o *AppPriceResponse) SetIncluded(v []AppPricesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppPriceResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPriceResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPriceResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppPriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppPriceResponse struct {
	value *AppPriceResponse
	isSet bool
}

func (v NullableAppPriceResponse) Get() *AppPriceResponse {
	return v.value
}

func (v *NullableAppPriceResponse) Set(val *AppPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceResponse(val *AppPriceResponse) *NullableAppPriceResponse {
	return &NullableAppPriceResponse{value: val, isSet: true}
}

func (v NullableAppPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

