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

// checks if the BetaTesterBuildsLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterBuildsLinkagesRequest{}

// BetaTesterBuildsLinkagesRequest struct for BetaTesterBuildsLinkagesRequest
type BetaTesterBuildsLinkagesRequest struct {
	Data []AppEncryptionDeclarationRelationshipsBuildsDataInner `json:"data"`
}

// NewBetaTesterBuildsLinkagesRequest instantiates a new BetaTesterBuildsLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterBuildsLinkagesRequest(data []AppEncryptionDeclarationRelationshipsBuildsDataInner) *BetaTesterBuildsLinkagesRequest {
	this := BetaTesterBuildsLinkagesRequest{}
	this.Data = data
	return &this
}

// NewBetaTesterBuildsLinkagesRequestWithDefaults instantiates a new BetaTesterBuildsLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterBuildsLinkagesRequestWithDefaults() *BetaTesterBuildsLinkagesRequest {
	this := BetaTesterBuildsLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaTesterBuildsLinkagesRequest) GetData() []AppEncryptionDeclarationRelationshipsBuildsDataInner {
	if o == nil {
		var ret []AppEncryptionDeclarationRelationshipsBuildsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaTesterBuildsLinkagesRequest) GetDataOk() ([]AppEncryptionDeclarationRelationshipsBuildsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaTesterBuildsLinkagesRequest) SetData(v []AppEncryptionDeclarationRelationshipsBuildsDataInner) {
	o.Data = v
}

func (o BetaTesterBuildsLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterBuildsLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBetaTesterBuildsLinkagesRequest struct {
	value *BetaTesterBuildsLinkagesRequest
	isSet bool
}

func (v NullableBetaTesterBuildsLinkagesRequest) Get() *BetaTesterBuildsLinkagesRequest {
	return v.value
}

func (v *NullableBetaTesterBuildsLinkagesRequest) Set(val *BetaTesterBuildsLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterBuildsLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterBuildsLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterBuildsLinkagesRequest(val *BetaTesterBuildsLinkagesRequest) *NullableBetaTesterBuildsLinkagesRequest {
	return &NullableBetaTesterBuildsLinkagesRequest{value: val, isSet: true}
}

func (v NullableBetaTesterBuildsLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterBuildsLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


