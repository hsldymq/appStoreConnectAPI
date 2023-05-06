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

// checks if the InAppPurchaseAvailabilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseAvailabilityResponse{}

// InAppPurchaseAvailabilityResponse struct for InAppPurchaseAvailabilityResponse
type InAppPurchaseAvailabilityResponse struct {
	Data InAppPurchaseAvailability `json:"data"`
	Included []Territory `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewInAppPurchaseAvailabilityResponse instantiates a new InAppPurchaseAvailabilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseAvailabilityResponse(data InAppPurchaseAvailability, links DocumentLinks) *InAppPurchaseAvailabilityResponse {
	this := InAppPurchaseAvailabilityResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewInAppPurchaseAvailabilityResponseWithDefaults instantiates a new InAppPurchaseAvailabilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseAvailabilityResponseWithDefaults() *InAppPurchaseAvailabilityResponse {
	this := InAppPurchaseAvailabilityResponse{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchaseAvailabilityResponse) GetData() InAppPurchaseAvailability {
	if o == nil {
		var ret InAppPurchaseAvailability
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailabilityResponse) GetDataOk() (*InAppPurchaseAvailability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InAppPurchaseAvailabilityResponse) SetData(v InAppPurchaseAvailability) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *InAppPurchaseAvailabilityResponse) GetIncluded() []Territory {
	if o == nil || IsNil(o.Included) {
		var ret []Territory
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailabilityResponse) GetIncludedOk() ([]Territory, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *InAppPurchaseAvailabilityResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []Territory and assigns it to the Included field.
func (o *InAppPurchaseAvailabilityResponse) SetIncluded(v []Territory) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *InAppPurchaseAvailabilityResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailabilityResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *InAppPurchaseAvailabilityResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o InAppPurchaseAvailabilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseAvailabilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableInAppPurchaseAvailabilityResponse struct {
	value *InAppPurchaseAvailabilityResponse
	isSet bool
}

func (v NullableInAppPurchaseAvailabilityResponse) Get() *InAppPurchaseAvailabilityResponse {
	return v.value
}

func (v *NullableInAppPurchaseAvailabilityResponse) Set(val *InAppPurchaseAvailabilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseAvailabilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseAvailabilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseAvailabilityResponse(val *InAppPurchaseAvailabilityResponse) *NullableInAppPurchaseAvailabilityResponse {
	return &NullableInAppPurchaseAvailabilityResponse{value: val, isSet: true}
}

func (v NullableInAppPurchaseAvailabilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseAvailabilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


