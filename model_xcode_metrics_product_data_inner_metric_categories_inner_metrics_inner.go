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

// checks if the XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner{}

// XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner struct for XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner
type XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner struct {
	Identifier *string `json:"identifier,omitempty"`
	GoalKeys []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner `json:"goalKeys,omitempty"`
	Unit *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerUnit `json:"unit,omitempty"`
	Datasets []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInner `json:"datasets,omitempty"`
}

// NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner instantiates a new XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner() *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner {
	this := XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner{}
	return &this
}

// NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerWithDefaults instantiates a new XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerWithDefaults() *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner {
	this := XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetGoalKeys returns the GoalKeys field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) GetGoalKeys() []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner {
	if o == nil || IsNil(o.GoalKeys) {
		var ret []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner
		return ret
	}
	return o.GoalKeys
}

// GetGoalKeysOk returns a tuple with the GoalKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) GetGoalKeysOk() ([]XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner, bool) {
	if o == nil || IsNil(o.GoalKeys) {
		return nil, false
	}
	return o.GoalKeys, true
}

// HasGoalKeys returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) HasGoalKeys() bool {
	if o != nil && !IsNil(o.GoalKeys) {
		return true
	}

	return false
}

// SetGoalKeys gets a reference to the given []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner and assigns it to the GoalKeys field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) SetGoalKeys(v []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerGoalKeysInner) {
	o.GoalKeys = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) GetUnit() XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerUnit {
	if o == nil || IsNil(o.Unit) {
		var ret XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerUnit
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) GetUnitOk() (*XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerUnit, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerUnit and assigns it to the Unit field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) SetUnit(v XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerUnit) {
	o.Unit = &v
}

// GetDatasets returns the Datasets field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) GetDatasets() []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInner {
	if o == nil || IsNil(o.Datasets) {
		var ret []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInner
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) GetDatasetsOk() ([]XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInner, bool) {
	if o == nil || IsNil(o.Datasets) {
		return nil, false
	}
	return o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) HasDatasets() bool {
	if o != nil && !IsNil(o.Datasets) {
		return true
	}

	return false
}

// SetDatasets gets a reference to the given []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInner and assigns it to the Datasets field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) SetDatasets(v []XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInner) {
	o.Datasets = v
}

func (o XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.GoalKeys) {
		toSerialize["goalKeys"] = o.GoalKeys
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Datasets) {
		toSerialize["datasets"] = o.Datasets
	}
	return toSerialize, nil
}

type NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner struct {
	value *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner
	isSet bool
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) Get() *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner {
	return v.value
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) Set(val *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner(val *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner {
	return &NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner{value: val, isSet: true}
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


