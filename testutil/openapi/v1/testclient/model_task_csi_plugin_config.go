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

// TaskCSIPluginConfig struct for TaskCSIPluginConfig
type TaskCSIPluginConfig struct {
	ID       *string `json:"ID,omitempty"`
	MountDir *string `json:"MountDir,omitempty"`
	Type     *string `json:"Type,omitempty"`
}

// NewTaskCSIPluginConfig instantiates a new TaskCSIPluginConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskCSIPluginConfig() *TaskCSIPluginConfig {
	this := TaskCSIPluginConfig{}
	return &this
}

// NewTaskCSIPluginConfigWithDefaults instantiates a new TaskCSIPluginConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskCSIPluginConfigWithDefaults() *TaskCSIPluginConfig {
	this := TaskCSIPluginConfig{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TaskCSIPluginConfig) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCSIPluginConfig) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TaskCSIPluginConfig) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TaskCSIPluginConfig) SetID(v string) {
	o.ID = &v
}

// GetMountDir returns the MountDir field value if set, zero value otherwise.
func (o *TaskCSIPluginConfig) GetMountDir() string {
	if o == nil || o.MountDir == nil {
		var ret string
		return ret
	}
	return *o.MountDir
}

// GetMountDirOk returns a tuple with the MountDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCSIPluginConfig) GetMountDirOk() (*string, bool) {
	if o == nil || o.MountDir == nil {
		return nil, false
	}
	return o.MountDir, true
}

// HasMountDir returns a boolean if a field has been set.
func (o *TaskCSIPluginConfig) HasMountDir() bool {
	if o != nil && o.MountDir != nil {
		return true
	}

	return false
}

// SetMountDir gets a reference to the given string and assigns it to the MountDir field.
func (o *TaskCSIPluginConfig) SetMountDir(v string) {
	o.MountDir = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaskCSIPluginConfig) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCSIPluginConfig) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaskCSIPluginConfig) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaskCSIPluginConfig) SetType(v string) {
	o.Type = &v
}

func (o TaskCSIPluginConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.MountDir != nil {
		toSerialize["MountDir"] = o.MountDir
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTaskCSIPluginConfig struct {
	value *TaskCSIPluginConfig
	isSet bool
}

func (v NullableTaskCSIPluginConfig) Get() *TaskCSIPluginConfig {
	return v.value
}

func (v *NullableTaskCSIPluginConfig) Set(val *TaskCSIPluginConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskCSIPluginConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskCSIPluginConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskCSIPluginConfig(val *TaskCSIPluginConfig) *NullableTaskCSIPluginConfig {
	return &NullableTaskCSIPluginConfig{value: val, isSet: true}
}

func (v NullableTaskCSIPluginConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskCSIPluginConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
