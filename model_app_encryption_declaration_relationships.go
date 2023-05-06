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

// checks if the AppEncryptionDeclarationRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEncryptionDeclarationRelationships{}

// AppEncryptionDeclarationRelationships struct for AppEncryptionDeclarationRelationships
type AppEncryptionDeclarationRelationships struct {
	App *AppAvailabilityRelationshipsApp `json:"app,omitempty"`
	Builds *AppEncryptionDeclarationRelationshipsBuilds `json:"builds,omitempty"`
	AppEncryptionDeclarationDocument *AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocument `json:"appEncryptionDeclarationDocument,omitempty"`
}

// NewAppEncryptionDeclarationRelationships instantiates a new AppEncryptionDeclarationRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEncryptionDeclarationRelationships() *AppEncryptionDeclarationRelationships {
	this := AppEncryptionDeclarationRelationships{}
	return &this
}

// NewAppEncryptionDeclarationRelationshipsWithDefaults instantiates a new AppEncryptionDeclarationRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEncryptionDeclarationRelationshipsWithDefaults() *AppEncryptionDeclarationRelationships {
	this := AppEncryptionDeclarationRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationRelationships) GetApp() AppAvailabilityRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret AppAvailabilityRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationRelationships) GetAppOk() (*AppAvailabilityRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppAvailabilityRelationshipsApp and assigns it to the App field.
func (o *AppEncryptionDeclarationRelationships) SetApp(v AppAvailabilityRelationshipsApp) {
	o.App = &v
}

// GetBuilds returns the Builds field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationRelationships) GetBuilds() AppEncryptionDeclarationRelationshipsBuilds {
	if o == nil || IsNil(o.Builds) {
		var ret AppEncryptionDeclarationRelationshipsBuilds
		return ret
	}
	return *o.Builds
}

// GetBuildsOk returns a tuple with the Builds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationRelationships) GetBuildsOk() (*AppEncryptionDeclarationRelationshipsBuilds, bool) {
	if o == nil || IsNil(o.Builds) {
		return nil, false
	}
	return o.Builds, true
}

// HasBuilds returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationRelationships) HasBuilds() bool {
	if o != nil && !IsNil(o.Builds) {
		return true
	}

	return false
}

// SetBuilds gets a reference to the given AppEncryptionDeclarationRelationshipsBuilds and assigns it to the Builds field.
func (o *AppEncryptionDeclarationRelationships) SetBuilds(v AppEncryptionDeclarationRelationshipsBuilds) {
	o.Builds = &v
}

// GetAppEncryptionDeclarationDocument returns the AppEncryptionDeclarationDocument field value if set, zero value otherwise.
func (o *AppEncryptionDeclarationRelationships) GetAppEncryptionDeclarationDocument() AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocument {
	if o == nil || IsNil(o.AppEncryptionDeclarationDocument) {
		var ret AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocument
		return ret
	}
	return *o.AppEncryptionDeclarationDocument
}

// GetAppEncryptionDeclarationDocumentOk returns a tuple with the AppEncryptionDeclarationDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEncryptionDeclarationRelationships) GetAppEncryptionDeclarationDocumentOk() (*AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocument, bool) {
	if o == nil || IsNil(o.AppEncryptionDeclarationDocument) {
		return nil, false
	}
	return o.AppEncryptionDeclarationDocument, true
}

// HasAppEncryptionDeclarationDocument returns a boolean if a field has been set.
func (o *AppEncryptionDeclarationRelationships) HasAppEncryptionDeclarationDocument() bool {
	if o != nil && !IsNil(o.AppEncryptionDeclarationDocument) {
		return true
	}

	return false
}

// SetAppEncryptionDeclarationDocument gets a reference to the given AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocument and assigns it to the AppEncryptionDeclarationDocument field.
func (o *AppEncryptionDeclarationRelationships) SetAppEncryptionDeclarationDocument(v AppEncryptionDeclarationRelationshipsAppEncryptionDeclarationDocument) {
	o.AppEncryptionDeclarationDocument = &v
}

func (o AppEncryptionDeclarationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEncryptionDeclarationRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Builds) {
		toSerialize["builds"] = o.Builds
	}
	if !IsNil(o.AppEncryptionDeclarationDocument) {
		toSerialize["appEncryptionDeclarationDocument"] = o.AppEncryptionDeclarationDocument
	}
	return toSerialize, nil
}

type NullableAppEncryptionDeclarationRelationships struct {
	value *AppEncryptionDeclarationRelationships
	isSet bool
}

func (v NullableAppEncryptionDeclarationRelationships) Get() *AppEncryptionDeclarationRelationships {
	return v.value
}

func (v *NullableAppEncryptionDeclarationRelationships) Set(val *AppEncryptionDeclarationRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEncryptionDeclarationRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEncryptionDeclarationRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEncryptionDeclarationRelationships(val *AppEncryptionDeclarationRelationships) *NullableAppEncryptionDeclarationRelationships {
	return &NullableAppEncryptionDeclarationRelationships{value: val, isSet: true}
}

func (v NullableAppEncryptionDeclarationRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEncryptionDeclarationRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


