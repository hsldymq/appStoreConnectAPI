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

// checks if the InAppPurchaseLocalizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseLocalizationsResponse{}

// InAppPurchaseLocalizationsResponse struct for InAppPurchaseLocalizationsResponse
type InAppPurchaseLocalizationsResponse struct {
	Data []InAppPurchaseLocalization `json:"data"`
	Included []InAppPurchaseV2 `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewInAppPurchaseLocalizationsResponse instantiates a new InAppPurchaseLocalizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseLocalizationsResponse(data []InAppPurchaseLocalization, links PagedDocumentLinks) *InAppPurchaseLocalizationsResponse {
	this := InAppPurchaseLocalizationsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewInAppPurchaseLocalizationsResponseWithDefaults instantiates a new InAppPurchaseLocalizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseLocalizationsResponseWithDefaults() *InAppPurchaseLocalizationsResponse {
	this := InAppPurchaseLocalizationsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchaseLocalizationsResponse) GetData() []InAppPurchaseLocalization {
	if o == nil {
		var ret []InAppPurchaseLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationsResponse) GetDataOk() ([]InAppPurchaseLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *InAppPurchaseLocalizationsResponse) SetData(v []InAppPurchaseLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *InAppPurchaseLocalizationsResponse) GetIncluded() []InAppPurchaseV2 {
	if o == nil || IsNil(o.Included) {
		var ret []InAppPurchaseV2
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationsResponse) GetIncludedOk() ([]InAppPurchaseV2, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *InAppPurchaseLocalizationsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []InAppPurchaseV2 and assigns it to the Included field.
func (o *InAppPurchaseLocalizationsResponse) SetIncluded(v []InAppPurchaseV2) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *InAppPurchaseLocalizationsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *InAppPurchaseLocalizationsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InAppPurchaseLocalizationsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InAppPurchaseLocalizationsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *InAppPurchaseLocalizationsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o InAppPurchaseLocalizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseLocalizationsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableInAppPurchaseLocalizationsResponse struct {
	value *InAppPurchaseLocalizationsResponse
	isSet bool
}

func (v NullableInAppPurchaseLocalizationsResponse) Get() *InAppPurchaseLocalizationsResponse {
	return v.value
}

func (v *NullableInAppPurchaseLocalizationsResponse) Set(val *InAppPurchaseLocalizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseLocalizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseLocalizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseLocalizationsResponse(val *InAppPurchaseLocalizationsResponse) *NullableInAppPurchaseLocalizationsResponse {
	return &NullableInAppPurchaseLocalizationsResponse{value: val, isSet: true}
}

func (v NullableInAppPurchaseLocalizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseLocalizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

