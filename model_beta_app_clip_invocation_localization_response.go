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

// checks if the BetaAppClipInvocationLocalizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationLocalizationResponse{}

// BetaAppClipInvocationLocalizationResponse struct for BetaAppClipInvocationLocalizationResponse
type BetaAppClipInvocationLocalizationResponse struct {
	Data BetaAppClipInvocationLocalization `json:"data"`
	Links DocumentLinks `json:"links"`
}

// NewBetaAppClipInvocationLocalizationResponse instantiates a new BetaAppClipInvocationLocalizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationLocalizationResponse(data BetaAppClipInvocationLocalization, links DocumentLinks) *BetaAppClipInvocationLocalizationResponse {
	this := BetaAppClipInvocationLocalizationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaAppClipInvocationLocalizationResponseWithDefaults instantiates a new BetaAppClipInvocationLocalizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationLocalizationResponseWithDefaults() *BetaAppClipInvocationLocalizationResponse {
	this := BetaAppClipInvocationLocalizationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppClipInvocationLocalizationResponse) GetData() BetaAppClipInvocationLocalization {
	if o == nil {
		var ret BetaAppClipInvocationLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationResponse) GetDataOk() (*BetaAppClipInvocationLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaAppClipInvocationLocalizationResponse) SetData(v BetaAppClipInvocationLocalization) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BetaAppClipInvocationLocalizationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaAppClipInvocationLocalizationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BetaAppClipInvocationLocalizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationLocalizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBetaAppClipInvocationLocalizationResponse struct {
	value *BetaAppClipInvocationLocalizationResponse
	isSet bool
}

func (v NullableBetaAppClipInvocationLocalizationResponse) Get() *BetaAppClipInvocationLocalizationResponse {
	return v.value
}

func (v *NullableBetaAppClipInvocationLocalizationResponse) Set(val *BetaAppClipInvocationLocalizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationLocalizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationLocalizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationLocalizationResponse(val *BetaAppClipInvocationLocalizationResponse) *NullableBetaAppClipInvocationLocalizationResponse {
	return &NullableBetaAppClipInvocationLocalizationResponse{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationLocalizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationLocalizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


