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

// checks if the BetaLicenseAgreementUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaLicenseAgreementUpdateRequest{}

// BetaLicenseAgreementUpdateRequest struct for BetaLicenseAgreementUpdateRequest
type BetaLicenseAgreementUpdateRequest struct {
	Data BetaLicenseAgreementUpdateRequestData `json:"data"`
}

// NewBetaLicenseAgreementUpdateRequest instantiates a new BetaLicenseAgreementUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaLicenseAgreementUpdateRequest(data BetaLicenseAgreementUpdateRequestData) *BetaLicenseAgreementUpdateRequest {
	this := BetaLicenseAgreementUpdateRequest{}
	this.Data = data
	return &this
}

// NewBetaLicenseAgreementUpdateRequestWithDefaults instantiates a new BetaLicenseAgreementUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaLicenseAgreementUpdateRequestWithDefaults() *BetaLicenseAgreementUpdateRequest {
	this := BetaLicenseAgreementUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaLicenseAgreementUpdateRequest) GetData() BetaLicenseAgreementUpdateRequestData {
	if o == nil {
		var ret BetaLicenseAgreementUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreementUpdateRequest) GetDataOk() (*BetaLicenseAgreementUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaLicenseAgreementUpdateRequest) SetData(v BetaLicenseAgreementUpdateRequestData) {
	o.Data = v
}

func (o BetaLicenseAgreementUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaLicenseAgreementUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBetaLicenseAgreementUpdateRequest struct {
	value *BetaLicenseAgreementUpdateRequest
	isSet bool
}

func (v NullableBetaLicenseAgreementUpdateRequest) Get() *BetaLicenseAgreementUpdateRequest {
	return v.value
}

func (v *NullableBetaLicenseAgreementUpdateRequest) Set(val *BetaLicenseAgreementUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaLicenseAgreementUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaLicenseAgreementUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaLicenseAgreementUpdateRequest(val *BetaLicenseAgreementUpdateRequest) *NullableBetaLicenseAgreementUpdateRequest {
	return &NullableBetaLicenseAgreementUpdateRequest{value: val, isSet: true}
}

func (v NullableBetaLicenseAgreementUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaLicenseAgreementUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


