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

// checks if the CustomerReviewResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerReviewResponseV1{}

// CustomerReviewResponseV1 struct for CustomerReviewResponseV1
type CustomerReviewResponseV1 struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *CustomerReviewResponseV1Attributes `json:"attributes,omitempty"`
	Relationships *CustomerReviewResponseV1Relationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewCustomerReviewResponseV1 instantiates a new CustomerReviewResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerReviewResponseV1(type_ string, id string, links ResourceLinks) *CustomerReviewResponseV1 {
	this := CustomerReviewResponseV1{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewCustomerReviewResponseV1WithDefaults instantiates a new CustomerReviewResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerReviewResponseV1WithDefaults() *CustomerReviewResponseV1 {
	this := CustomerReviewResponseV1{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerReviewResponseV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerReviewResponseV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerReviewResponseV1) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CustomerReviewResponseV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerReviewResponseV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerReviewResponseV1) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerReviewResponseV1) GetAttributes() CustomerReviewResponseV1Attributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CustomerReviewResponseV1Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerReviewResponseV1) GetAttributesOk() (*CustomerReviewResponseV1Attributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerReviewResponseV1) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CustomerReviewResponseV1Attributes and assigns it to the Attributes field.
func (o *CustomerReviewResponseV1) SetAttributes(v CustomerReviewResponseV1Attributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerReviewResponseV1) GetRelationships() CustomerReviewResponseV1Relationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CustomerReviewResponseV1Relationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerReviewResponseV1) GetRelationshipsOk() (*CustomerReviewResponseV1Relationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerReviewResponseV1) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerReviewResponseV1Relationships and assigns it to the Relationships field.
func (o *CustomerReviewResponseV1) SetRelationships(v CustomerReviewResponseV1Relationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *CustomerReviewResponseV1) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CustomerReviewResponseV1) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CustomerReviewResponseV1) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o CustomerReviewResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerReviewResponseV1) ToMap() (map[string]interface{}, error) {
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

type NullableCustomerReviewResponseV1 struct {
	value *CustomerReviewResponseV1
	isSet bool
}

func (v NullableCustomerReviewResponseV1) Get() *CustomerReviewResponseV1 {
	return v.value
}

func (v *NullableCustomerReviewResponseV1) Set(val *CustomerReviewResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerReviewResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerReviewResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerReviewResponseV1(val *CustomerReviewResponseV1) *NullableCustomerReviewResponseV1 {
	return &NullableCustomerReviewResponseV1{value: val, isSet: true}
}

func (v NullableCustomerReviewResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerReviewResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


