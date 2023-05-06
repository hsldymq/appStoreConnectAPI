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

// checks if the CiWorkflowRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiWorkflowRelationships{}

// CiWorkflowRelationships struct for CiWorkflowRelationships
type CiWorkflowRelationships struct {
	Product *AppRelationshipsCiProduct `json:"product,omitempty"`
	Repository *CiWorkflowRelationshipsRepository `json:"repository,omitempty"`
	XcodeVersion *CiWorkflowRelationshipsXcodeVersion `json:"xcodeVersion,omitempty"`
	MacOsVersion *CiWorkflowRelationshipsMacOsVersion `json:"macOsVersion,omitempty"`
}

// NewCiWorkflowRelationships instantiates a new CiWorkflowRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiWorkflowRelationships() *CiWorkflowRelationships {
	this := CiWorkflowRelationships{}
	return &this
}

// NewCiWorkflowRelationshipsWithDefaults instantiates a new CiWorkflowRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiWorkflowRelationshipsWithDefaults() *CiWorkflowRelationships {
	this := CiWorkflowRelationships{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *CiWorkflowRelationships) GetProduct() AppRelationshipsCiProduct {
	if o == nil || IsNil(o.Product) {
		var ret AppRelationshipsCiProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowRelationships) GetProductOk() (*AppRelationshipsCiProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *CiWorkflowRelationships) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given AppRelationshipsCiProduct and assigns it to the Product field.
func (o *CiWorkflowRelationships) SetProduct(v AppRelationshipsCiProduct) {
	o.Product = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *CiWorkflowRelationships) GetRepository() CiWorkflowRelationshipsRepository {
	if o == nil || IsNil(o.Repository) {
		var ret CiWorkflowRelationshipsRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowRelationships) GetRepositoryOk() (*CiWorkflowRelationshipsRepository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *CiWorkflowRelationships) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given CiWorkflowRelationshipsRepository and assigns it to the Repository field.
func (o *CiWorkflowRelationships) SetRepository(v CiWorkflowRelationshipsRepository) {
	o.Repository = &v
}

// GetXcodeVersion returns the XcodeVersion field value if set, zero value otherwise.
func (o *CiWorkflowRelationships) GetXcodeVersion() CiWorkflowRelationshipsXcodeVersion {
	if o == nil || IsNil(o.XcodeVersion) {
		var ret CiWorkflowRelationshipsXcodeVersion
		return ret
	}
	return *o.XcodeVersion
}

// GetXcodeVersionOk returns a tuple with the XcodeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowRelationships) GetXcodeVersionOk() (*CiWorkflowRelationshipsXcodeVersion, bool) {
	if o == nil || IsNil(o.XcodeVersion) {
		return nil, false
	}
	return o.XcodeVersion, true
}

// HasXcodeVersion returns a boolean if a field has been set.
func (o *CiWorkflowRelationships) HasXcodeVersion() bool {
	if o != nil && !IsNil(o.XcodeVersion) {
		return true
	}

	return false
}

// SetXcodeVersion gets a reference to the given CiWorkflowRelationshipsXcodeVersion and assigns it to the XcodeVersion field.
func (o *CiWorkflowRelationships) SetXcodeVersion(v CiWorkflowRelationshipsXcodeVersion) {
	o.XcodeVersion = &v
}

// GetMacOsVersion returns the MacOsVersion field value if set, zero value otherwise.
func (o *CiWorkflowRelationships) GetMacOsVersion() CiWorkflowRelationshipsMacOsVersion {
	if o == nil || IsNil(o.MacOsVersion) {
		var ret CiWorkflowRelationshipsMacOsVersion
		return ret
	}
	return *o.MacOsVersion
}

// GetMacOsVersionOk returns a tuple with the MacOsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowRelationships) GetMacOsVersionOk() (*CiWorkflowRelationshipsMacOsVersion, bool) {
	if o == nil || IsNil(o.MacOsVersion) {
		return nil, false
	}
	return o.MacOsVersion, true
}

// HasMacOsVersion returns a boolean if a field has been set.
func (o *CiWorkflowRelationships) HasMacOsVersion() bool {
	if o != nil && !IsNil(o.MacOsVersion) {
		return true
	}

	return false
}

// SetMacOsVersion gets a reference to the given CiWorkflowRelationshipsMacOsVersion and assigns it to the MacOsVersion field.
func (o *CiWorkflowRelationships) SetMacOsVersion(v CiWorkflowRelationshipsMacOsVersion) {
	o.MacOsVersion = &v
}

func (o CiWorkflowRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiWorkflowRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.XcodeVersion) {
		toSerialize["xcodeVersion"] = o.XcodeVersion
	}
	if !IsNil(o.MacOsVersion) {
		toSerialize["macOsVersion"] = o.MacOsVersion
	}
	return toSerialize, nil
}

type NullableCiWorkflowRelationships struct {
	value *CiWorkflowRelationships
	isSet bool
}

func (v NullableCiWorkflowRelationships) Get() *CiWorkflowRelationships {
	return v.value
}

func (v *NullableCiWorkflowRelationships) Set(val *CiWorkflowRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflowRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflowRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflowRelationships(val *CiWorkflowRelationships) *NullableCiWorkflowRelationships {
	return &NullableCiWorkflowRelationships{value: val, isSet: true}
}

func (v NullableCiWorkflowRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflowRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

