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

// checks if the SubscriptionIntroductoryOfferInlineCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionIntroductoryOfferInlineCreate{}

// SubscriptionIntroductoryOfferInlineCreate struct for SubscriptionIntroductoryOfferInlineCreate
type SubscriptionIntroductoryOfferInlineCreate struct {
	Type string `json:"type"`
	Id *string `json:"id,omitempty"`
	Attributes SubscriptionIntroductoryOfferInlineCreateAttributes `json:"attributes"`
	Relationships *SubscriptionIntroductoryOfferInlineCreateRelationships `json:"relationships,omitempty"`
}

// NewSubscriptionIntroductoryOfferInlineCreate instantiates a new SubscriptionIntroductoryOfferInlineCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionIntroductoryOfferInlineCreate(type_ string, attributes SubscriptionIntroductoryOfferInlineCreateAttributes) *SubscriptionIntroductoryOfferInlineCreate {
	this := SubscriptionIntroductoryOfferInlineCreate{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSubscriptionIntroductoryOfferInlineCreateWithDefaults instantiates a new SubscriptionIntroductoryOfferInlineCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionIntroductoryOfferInlineCreateWithDefaults() *SubscriptionIntroductoryOfferInlineCreate {
	this := SubscriptionIntroductoryOfferInlineCreate{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionIntroductoryOfferInlineCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferInlineCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionIntroductoryOfferInlineCreate) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferInlineCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferInlineCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferInlineCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionIntroductoryOfferInlineCreate) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value
func (o *SubscriptionIntroductoryOfferInlineCreate) GetAttributes() SubscriptionIntroductoryOfferInlineCreateAttributes {
	if o == nil {
		var ret SubscriptionIntroductoryOfferInlineCreateAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferInlineCreate) GetAttributesOk() (*SubscriptionIntroductoryOfferInlineCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SubscriptionIntroductoryOfferInlineCreate) SetAttributes(v SubscriptionIntroductoryOfferInlineCreateAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferInlineCreate) GetRelationships() SubscriptionIntroductoryOfferInlineCreateRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SubscriptionIntroductoryOfferInlineCreateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferInlineCreate) GetRelationshipsOk() (*SubscriptionIntroductoryOfferInlineCreateRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferInlineCreate) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SubscriptionIntroductoryOfferInlineCreateRelationships and assigns it to the Relationships field.
func (o *SubscriptionIntroductoryOfferInlineCreate) SetRelationships(v SubscriptionIntroductoryOfferInlineCreateRelationships) {
	o.Relationships = &v
}

func (o SubscriptionIntroductoryOfferInlineCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionIntroductoryOfferInlineCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableSubscriptionIntroductoryOfferInlineCreate struct {
	value *SubscriptionIntroductoryOfferInlineCreate
	isSet bool
}

func (v NullableSubscriptionIntroductoryOfferInlineCreate) Get() *SubscriptionIntroductoryOfferInlineCreate {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOfferInlineCreate) Set(val *SubscriptionIntroductoryOfferInlineCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOfferInlineCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOfferInlineCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOfferInlineCreate(val *SubscriptionIntroductoryOfferInlineCreate) *NullableSubscriptionIntroductoryOfferInlineCreate {
	return &NullableSubscriptionIntroductoryOfferInlineCreate{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOfferInlineCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOfferInlineCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


