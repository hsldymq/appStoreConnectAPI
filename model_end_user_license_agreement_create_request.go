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

// checks if the EndUserLicenseAgreementCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndUserLicenseAgreementCreateRequest{}

// EndUserLicenseAgreementCreateRequest struct for EndUserLicenseAgreementCreateRequest
type EndUserLicenseAgreementCreateRequest struct {
	Data EndUserLicenseAgreementCreateRequestData `json:"data"`
}

// NewEndUserLicenseAgreementCreateRequest instantiates a new EndUserLicenseAgreementCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserLicenseAgreementCreateRequest(data EndUserLicenseAgreementCreateRequestData) *EndUserLicenseAgreementCreateRequest {
	this := EndUserLicenseAgreementCreateRequest{}
	this.Data = data
	return &this
}

// NewEndUserLicenseAgreementCreateRequestWithDefaults instantiates a new EndUserLicenseAgreementCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserLicenseAgreementCreateRequestWithDefaults() *EndUserLicenseAgreementCreateRequest {
	this := EndUserLicenseAgreementCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *EndUserLicenseAgreementCreateRequest) GetData() EndUserLicenseAgreementCreateRequestData {
	if o == nil {
		var ret EndUserLicenseAgreementCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementCreateRequest) GetDataOk() (*EndUserLicenseAgreementCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EndUserLicenseAgreementCreateRequest) SetData(v EndUserLicenseAgreementCreateRequestData) {
	o.Data = v
}

func (o EndUserLicenseAgreementCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndUserLicenseAgreementCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableEndUserLicenseAgreementCreateRequest struct {
	value *EndUserLicenseAgreementCreateRequest
	isSet bool
}

func (v NullableEndUserLicenseAgreementCreateRequest) Get() *EndUserLicenseAgreementCreateRequest {
	return v.value
}

func (v *NullableEndUserLicenseAgreementCreateRequest) Set(val *EndUserLicenseAgreementCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserLicenseAgreementCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserLicenseAgreementCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserLicenseAgreementCreateRequest(val *EndUserLicenseAgreementCreateRequest) *NullableEndUserLicenseAgreementCreateRequest {
	return &NullableEndUserLicenseAgreementCreateRequest{value: val, isSet: true}
}

func (v NullableEndUserLicenseAgreementCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserLicenseAgreementCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


