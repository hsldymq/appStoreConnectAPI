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

// checks if the CiWorkflowRelationshipsMacOsVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiWorkflowRelationshipsMacOsVersion{}

// CiWorkflowRelationshipsMacOsVersion struct for CiWorkflowRelationshipsMacOsVersion
type CiWorkflowRelationshipsMacOsVersion struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *CiWorkflowRelationshipsMacOsVersionData `json:"data,omitempty"`
}

// NewCiWorkflowRelationshipsMacOsVersion instantiates a new CiWorkflowRelationshipsMacOsVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiWorkflowRelationshipsMacOsVersion() *CiWorkflowRelationshipsMacOsVersion {
	this := CiWorkflowRelationshipsMacOsVersion{}
	return &this
}

// NewCiWorkflowRelationshipsMacOsVersionWithDefaults instantiates a new CiWorkflowRelationshipsMacOsVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiWorkflowRelationshipsMacOsVersionWithDefaults() *CiWorkflowRelationshipsMacOsVersion {
	this := CiWorkflowRelationshipsMacOsVersion{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CiWorkflowRelationshipsMacOsVersion) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowRelationshipsMacOsVersion) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CiWorkflowRelationshipsMacOsVersion) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *CiWorkflowRelationshipsMacOsVersion) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CiWorkflowRelationshipsMacOsVersion) GetData() CiWorkflowRelationshipsMacOsVersionData {
	if o == nil || IsNil(o.Data) {
		var ret CiWorkflowRelationshipsMacOsVersionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowRelationshipsMacOsVersion) GetDataOk() (*CiWorkflowRelationshipsMacOsVersionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CiWorkflowRelationshipsMacOsVersion) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CiWorkflowRelationshipsMacOsVersionData and assigns it to the Data field.
func (o *CiWorkflowRelationshipsMacOsVersion) SetData(v CiWorkflowRelationshipsMacOsVersionData) {
	o.Data = &v
}

func (o CiWorkflowRelationshipsMacOsVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiWorkflowRelationshipsMacOsVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCiWorkflowRelationshipsMacOsVersion struct {
	value *CiWorkflowRelationshipsMacOsVersion
	isSet bool
}

func (v NullableCiWorkflowRelationshipsMacOsVersion) Get() *CiWorkflowRelationshipsMacOsVersion {
	return v.value
}

func (v *NullableCiWorkflowRelationshipsMacOsVersion) Set(val *CiWorkflowRelationshipsMacOsVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflowRelationshipsMacOsVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflowRelationshipsMacOsVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflowRelationshipsMacOsVersion(val *CiWorkflowRelationshipsMacOsVersion) *NullableCiWorkflowRelationshipsMacOsVersion {
	return &NullableCiWorkflowRelationshipsMacOsVersion{value: val, isSet: true}
}

func (v NullableCiWorkflowRelationshipsMacOsVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflowRelationshipsMacOsVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


