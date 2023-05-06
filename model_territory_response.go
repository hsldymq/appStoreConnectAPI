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

// checks if the TerritoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerritoryResponse{}

// TerritoryResponse struct for TerritoryResponse
type TerritoryResponse struct {
	Data Territory `json:"data"`
	Links DocumentLinks `json:"links"`
}

// NewTerritoryResponse instantiates a new TerritoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerritoryResponse(data Territory, links DocumentLinks) *TerritoryResponse {
	this := TerritoryResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewTerritoryResponseWithDefaults instantiates a new TerritoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerritoryResponseWithDefaults() *TerritoryResponse {
	this := TerritoryResponse{}
	return &this
}

// GetData returns the Data field value
func (o *TerritoryResponse) GetData() Territory {
	if o == nil {
		var ret Territory
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TerritoryResponse) GetDataOk() (*Territory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TerritoryResponse) SetData(v Territory) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *TerritoryResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TerritoryResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *TerritoryResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o TerritoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerritoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableTerritoryResponse struct {
	value *TerritoryResponse
	isSet bool
}

func (v NullableTerritoryResponse) Get() *TerritoryResponse {
	return v.value
}

func (v *NullableTerritoryResponse) Set(val *TerritoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerritoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerritoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerritoryResponse(val *TerritoryResponse) *NullableTerritoryResponse {
	return &NullableTerritoryResponse{value: val, isSet: true}
}

func (v NullableTerritoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerritoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


