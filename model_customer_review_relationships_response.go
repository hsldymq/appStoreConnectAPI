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

// checks if the CustomerReviewRelationshipsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerReviewRelationshipsResponse{}

// CustomerReviewRelationshipsResponse struct for CustomerReviewRelationshipsResponse
type CustomerReviewRelationshipsResponse struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *CustomerReviewRelationshipsResponseData `json:"data,omitempty"`
}

// NewCustomerReviewRelationshipsResponse instantiates a new CustomerReviewRelationshipsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerReviewRelationshipsResponse() *CustomerReviewRelationshipsResponse {
	this := CustomerReviewRelationshipsResponse{}
	return &this
}

// NewCustomerReviewRelationshipsResponseWithDefaults instantiates a new CustomerReviewRelationshipsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerReviewRelationshipsResponseWithDefaults() *CustomerReviewRelationshipsResponse {
	this := CustomerReviewRelationshipsResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CustomerReviewRelationshipsResponse) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerReviewRelationshipsResponse) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CustomerReviewRelationshipsResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *CustomerReviewRelationshipsResponse) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerReviewRelationshipsResponse) GetData() CustomerReviewRelationshipsResponseData {
	if o == nil || IsNil(o.Data) {
		var ret CustomerReviewRelationshipsResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerReviewRelationshipsResponse) GetDataOk() (*CustomerReviewRelationshipsResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerReviewRelationshipsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CustomerReviewRelationshipsResponseData and assigns it to the Data field.
func (o *CustomerReviewRelationshipsResponse) SetData(v CustomerReviewRelationshipsResponseData) {
	o.Data = &v
}

func (o CustomerReviewRelationshipsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerReviewRelationshipsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCustomerReviewRelationshipsResponse struct {
	value *CustomerReviewRelationshipsResponse
	isSet bool
}

func (v NullableCustomerReviewRelationshipsResponse) Get() *CustomerReviewRelationshipsResponse {
	return v.value
}

func (v *NullableCustomerReviewRelationshipsResponse) Set(val *CustomerReviewRelationshipsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerReviewRelationshipsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerReviewRelationshipsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerReviewRelationshipsResponse(val *CustomerReviewRelationshipsResponse) *NullableCustomerReviewRelationshipsResponse {
	return &NullableCustomerReviewRelationshipsResponse{value: val, isSet: true}
}

func (v NullableCustomerReviewRelationshipsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerReviewRelationshipsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


