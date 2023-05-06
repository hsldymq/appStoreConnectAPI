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

// checks if the AppEventLocalizationCreateRequestDataRelationshipsAppEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationCreateRequestDataRelationshipsAppEvent{}

// AppEventLocalizationCreateRequestDataRelationshipsAppEvent struct for AppEventLocalizationCreateRequestDataRelationshipsAppEvent
type AppEventLocalizationCreateRequestDataRelationshipsAppEvent struct {
	Data AppEventLocalizationRelationshipsAppEventData `json:"data"`
}

// NewAppEventLocalizationCreateRequestDataRelationshipsAppEvent instantiates a new AppEventLocalizationCreateRequestDataRelationshipsAppEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationCreateRequestDataRelationshipsAppEvent(data AppEventLocalizationRelationshipsAppEventData) *AppEventLocalizationCreateRequestDataRelationshipsAppEvent {
	this := AppEventLocalizationCreateRequestDataRelationshipsAppEvent{}
	this.Data = data
	return &this
}

// NewAppEventLocalizationCreateRequestDataRelationshipsAppEventWithDefaults instantiates a new AppEventLocalizationCreateRequestDataRelationshipsAppEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationCreateRequestDataRelationshipsAppEventWithDefaults() *AppEventLocalizationCreateRequestDataRelationshipsAppEvent {
	this := AppEventLocalizationCreateRequestDataRelationshipsAppEvent{}
	return &this
}

// GetData returns the Data field value
func (o *AppEventLocalizationCreateRequestDataRelationshipsAppEvent) GetData() AppEventLocalizationRelationshipsAppEventData {
	if o == nil {
		var ret AppEventLocalizationRelationshipsAppEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationCreateRequestDataRelationshipsAppEvent) GetDataOk() (*AppEventLocalizationRelationshipsAppEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppEventLocalizationCreateRequestDataRelationshipsAppEvent) SetData(v AppEventLocalizationRelationshipsAppEventData) {
	o.Data = v
}

func (o AppEventLocalizationCreateRequestDataRelationshipsAppEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationCreateRequestDataRelationshipsAppEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppEventLocalizationCreateRequestDataRelationshipsAppEvent struct {
	value *AppEventLocalizationCreateRequestDataRelationshipsAppEvent
	isSet bool
}

func (v NullableAppEventLocalizationCreateRequestDataRelationshipsAppEvent) Get() *AppEventLocalizationCreateRequestDataRelationshipsAppEvent {
	return v.value
}

func (v *NullableAppEventLocalizationCreateRequestDataRelationshipsAppEvent) Set(val *AppEventLocalizationCreateRequestDataRelationshipsAppEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationCreateRequestDataRelationshipsAppEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationCreateRequestDataRelationshipsAppEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationCreateRequestDataRelationshipsAppEvent(val *AppEventLocalizationCreateRequestDataRelationshipsAppEvent) *NullableAppEventLocalizationCreateRequestDataRelationshipsAppEvent {
	return &NullableAppEventLocalizationCreateRequestDataRelationshipsAppEvent{value: val, isSet: true}
}

func (v NullableAppEventLocalizationCreateRequestDataRelationshipsAppEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationCreateRequestDataRelationshipsAppEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


