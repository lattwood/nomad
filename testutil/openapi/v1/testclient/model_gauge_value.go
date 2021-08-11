/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.3
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testclient

import (
	"encoding/json"
)

// GaugeValue struct for GaugeValue
type GaugeValue struct {
	Labels *map[string]string `json:"Labels,omitempty"`
	Name   *string            `json:"Name,omitempty"`
	Value  *float32           `json:"Value,omitempty"`
}

// NewGaugeValue instantiates a new GaugeValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGaugeValue() *GaugeValue {
	this := GaugeValue{}
	return &this
}

// NewGaugeValueWithDefaults instantiates a new GaugeValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGaugeValueWithDefaults() *GaugeValue {
	this := GaugeValue{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *GaugeValue) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GaugeValue) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *GaugeValue) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *GaugeValue) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GaugeValue) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GaugeValue) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GaugeValue) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GaugeValue) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GaugeValue) GetValue() float32 {
	if o == nil || o.Value == nil {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GaugeValue) GetValueOk() (*float32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GaugeValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *GaugeValue) SetValue(v float32) {
	o.Value = &v
}

func (o GaugeValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableGaugeValue struct {
	value *GaugeValue
	isSet bool
}

func (v NullableGaugeValue) Get() *GaugeValue {
	return v.value
}

func (v *NullableGaugeValue) Set(val *GaugeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGaugeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGaugeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGaugeValue(val *GaugeValue) *NullableGaugeValue {
	return &NullableGaugeValue{value: val, isSet: true}
}

func (v NullableGaugeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGaugeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
