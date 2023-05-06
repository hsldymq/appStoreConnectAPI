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

// checks if the BetaTestersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTestersResponse{}

// BetaTestersResponse struct for BetaTestersResponse
type BetaTestersResponse struct {
	Data []BetaTester `json:"data"`
	Included []BetaTestersResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewBetaTestersResponse instantiates a new BetaTestersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTestersResponse(data []BetaTester, links PagedDocumentLinks) *BetaTestersResponse {
	this := BetaTestersResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaTestersResponseWithDefaults instantiates a new BetaTestersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTestersResponseWithDefaults() *BetaTestersResponse {
	this := BetaTestersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaTestersResponse) GetData() []BetaTester {
	if o == nil {
		var ret []BetaTester
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaTestersResponse) GetDataOk() ([]BetaTester, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaTestersResponse) SetData(v []BetaTester) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BetaTestersResponse) GetIncluded() []BetaTestersResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []BetaTestersResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTestersResponse) GetIncludedOk() ([]BetaTestersResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BetaTestersResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []BetaTestersResponseIncludedInner and assigns it to the Included field.
func (o *BetaTestersResponse) SetIncluded(v []BetaTestersResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *BetaTestersResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaTestersResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaTestersResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BetaTestersResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTestersResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BetaTestersResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BetaTestersResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BetaTestersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTestersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableBetaTestersResponse struct {
	value *BetaTestersResponse
	isSet bool
}

func (v NullableBetaTestersResponse) Get() *BetaTestersResponse {
	return v.value
}

func (v *NullableBetaTestersResponse) Set(val *BetaTestersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTestersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTestersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTestersResponse(val *BetaTestersResponse) *NullableBetaTestersResponse {
	return &NullableBetaTestersResponse{value: val, isSet: true}
}

func (v NullableBetaTestersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTestersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


