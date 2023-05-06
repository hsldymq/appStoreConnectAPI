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

// checks if the AppStoreReviewDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewDetail{}

// AppStoreReviewDetail struct for AppStoreReviewDetail
type AppStoreReviewDetail struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppStoreReviewDetailAttributes `json:"attributes,omitempty"`
	Relationships *AppStoreReviewDetailRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewAppStoreReviewDetail instantiates a new AppStoreReviewDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewDetail(type_ string, id string, links ResourceLinks) *AppStoreReviewDetail {
	this := AppStoreReviewDetail{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppStoreReviewDetailWithDefaults instantiates a new AppStoreReviewDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewDetailWithDefaults() *AppStoreReviewDetail {
	this := AppStoreReviewDetail{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreReviewDetail) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetail) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreReviewDetail) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreReviewDetail) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetail) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreReviewDetail) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppStoreReviewDetail) GetAttributes() AppStoreReviewDetailAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppStoreReviewDetailAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetail) GetAttributesOk() (*AppStoreReviewDetailAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppStoreReviewDetail) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppStoreReviewDetailAttributes and assigns it to the Attributes field.
func (o *AppStoreReviewDetail) SetAttributes(v AppStoreReviewDetailAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppStoreReviewDetail) GetRelationships() AppStoreReviewDetailRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppStoreReviewDetailRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetail) GetRelationshipsOk() (*AppStoreReviewDetailRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppStoreReviewDetail) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppStoreReviewDetailRelationships and assigns it to the Relationships field.
func (o *AppStoreReviewDetail) SetRelationships(v AppStoreReviewDetailRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppStoreReviewDetail) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetail) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreReviewDetail) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppStoreReviewDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppStoreReviewDetail struct {
	value *AppStoreReviewDetail
	isSet bool
}

func (v NullableAppStoreReviewDetail) Get() *AppStoreReviewDetail {
	return v.value
}

func (v *NullableAppStoreReviewDetail) Set(val *AppStoreReviewDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetail(val *AppStoreReviewDetail) *NullableAppStoreReviewDetail {
	return &NullableAppStoreReviewDetail{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


