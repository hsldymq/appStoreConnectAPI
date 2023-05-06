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

// checks if the SubscriptionGracePeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGracePeriod{}

// SubscriptionGracePeriod struct for SubscriptionGracePeriod
type SubscriptionGracePeriod struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *SubscriptionGracePeriodAttributes `json:"attributes,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewSubscriptionGracePeriod instantiates a new SubscriptionGracePeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGracePeriod(type_ string, id string, links ResourceLinks) *SubscriptionGracePeriod {
	this := SubscriptionGracePeriod{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewSubscriptionGracePeriodWithDefaults instantiates a new SubscriptionGracePeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGracePeriodWithDefaults() *SubscriptionGracePeriod {
	this := SubscriptionGracePeriod{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionGracePeriod) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGracePeriod) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionGracePeriod) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionGracePeriod) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGracePeriod) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionGracePeriod) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionGracePeriod) GetAttributes() SubscriptionGracePeriodAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionGracePeriodAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGracePeriod) GetAttributesOk() (*SubscriptionGracePeriodAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionGracePeriod) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionGracePeriodAttributes and assigns it to the Attributes field.
func (o *SubscriptionGracePeriod) SetAttributes(v SubscriptionGracePeriodAttributes) {
	o.Attributes = &v
}

// GetLinks returns the Links field value
func (o *SubscriptionGracePeriod) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGracePeriod) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionGracePeriod) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o SubscriptionGracePeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGracePeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableSubscriptionGracePeriod struct {
	value *SubscriptionGracePeriod
	isSet bool
}

func (v NullableSubscriptionGracePeriod) Get() *SubscriptionGracePeriod {
	return v.value
}

func (v *NullableSubscriptionGracePeriod) Set(val *SubscriptionGracePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGracePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGracePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGracePeriod(val *SubscriptionGracePeriod) *NullableSubscriptionGracePeriod {
	return &NullableSubscriptionGracePeriod{value: val, isSet: true}
}

func (v NullableSubscriptionGracePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGracePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


