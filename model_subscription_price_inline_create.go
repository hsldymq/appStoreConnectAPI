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

// checks if the SubscriptionPriceInlineCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPriceInlineCreate{}

// SubscriptionPriceInlineCreate struct for SubscriptionPriceInlineCreate
type SubscriptionPriceInlineCreate struct {
	Type string `json:"type"`
	Id *string `json:"id,omitempty"`
	Attributes *SubscriptionPriceInlineCreateAttributes `json:"attributes,omitempty"`
	Relationships *SubscriptionIntroductoryOfferInlineCreateRelationships `json:"relationships,omitempty"`
}

// NewSubscriptionPriceInlineCreate instantiates a new SubscriptionPriceInlineCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPriceInlineCreate(type_ string) *SubscriptionPriceInlineCreate {
	this := SubscriptionPriceInlineCreate{}
	this.Type = type_
	return &this
}

// NewSubscriptionPriceInlineCreateWithDefaults instantiates a new SubscriptionPriceInlineCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPriceInlineCreateWithDefaults() *SubscriptionPriceInlineCreate {
	this := SubscriptionPriceInlineCreate{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionPriceInlineCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceInlineCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionPriceInlineCreate) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionPriceInlineCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceInlineCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionPriceInlineCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionPriceInlineCreate) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionPriceInlineCreate) GetAttributes() SubscriptionPriceInlineCreateAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionPriceInlineCreateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceInlineCreate) GetAttributesOk() (*SubscriptionPriceInlineCreateAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionPriceInlineCreate) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionPriceInlineCreateAttributes and assigns it to the Attributes field.
func (o *SubscriptionPriceInlineCreate) SetAttributes(v SubscriptionPriceInlineCreateAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SubscriptionPriceInlineCreate) GetRelationships() SubscriptionIntroductoryOfferInlineCreateRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SubscriptionIntroductoryOfferInlineCreateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceInlineCreate) GetRelationshipsOk() (*SubscriptionIntroductoryOfferInlineCreateRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionPriceInlineCreate) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SubscriptionIntroductoryOfferInlineCreateRelationships and assigns it to the Relationships field.
func (o *SubscriptionPriceInlineCreate) SetRelationships(v SubscriptionIntroductoryOfferInlineCreateRelationships) {
	o.Relationships = &v
}

func (o SubscriptionPriceInlineCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPriceInlineCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableSubscriptionPriceInlineCreate struct {
	value *SubscriptionPriceInlineCreate
	isSet bool
}

func (v NullableSubscriptionPriceInlineCreate) Get() *SubscriptionPriceInlineCreate {
	return v.value
}

func (v *NullableSubscriptionPriceInlineCreate) Set(val *SubscriptionPriceInlineCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPriceInlineCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPriceInlineCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPriceInlineCreate(val *SubscriptionPriceInlineCreate) *NullableSubscriptionPriceInlineCreate {
	return &NullableSubscriptionPriceInlineCreate{value: val, isSet: true}
}

func (v NullableSubscriptionPriceInlineCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPriceInlineCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


