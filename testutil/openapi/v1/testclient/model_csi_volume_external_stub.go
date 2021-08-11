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

// CSIVolumeExternalStub struct for CSIVolumeExternalStub
type CSIVolumeExternalStub struct {
	CapacityBytes            *int64             `json:"CapacityBytes,omitempty"`
	CloneID                  *string            `json:"CloneID,omitempty"`
	ExternalID               *string            `json:"ExternalID,omitempty"`
	IsAbnormal               *bool              `json:"IsAbnormal,omitempty"`
	PublishedExternalNodeIDs *[]string          `json:"PublishedExternalNodeIDs,omitempty"`
	SnapshotID               *string            `json:"SnapshotID,omitempty"`
	Status                   *string            `json:"Status,omitempty"`
	VolumeContext            *map[string]string `json:"VolumeContext,omitempty"`
}

// NewCSIVolumeExternalStub instantiates a new CSIVolumeExternalStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSIVolumeExternalStub() *CSIVolumeExternalStub {
	this := CSIVolumeExternalStub{}
	return &this
}

// NewCSIVolumeExternalStubWithDefaults instantiates a new CSIVolumeExternalStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSIVolumeExternalStubWithDefaults() *CSIVolumeExternalStub {
	this := CSIVolumeExternalStub{}
	return &this
}

// GetCapacityBytes returns the CapacityBytes field value if set, zero value otherwise.
func (o *CSIVolumeExternalStub) GetCapacityBytes() int64 {
	if o == nil || o.CapacityBytes == nil {
		var ret int64
		return ret
	}
	return *o.CapacityBytes
}

// GetCapacityBytesOk returns a tuple with the CapacityBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeExternalStub) GetCapacityBytesOk() (*int64, bool) {
	if o == nil || o.CapacityBytes == nil {
		return nil, false
	}
	return o.CapacityBytes, true
}

// HasCapacityBytes returns a boolean if a field has been set.
func (o *CSIVolumeExternalStub) HasCapacityBytes() bool {
	if o != nil && o.CapacityBytes != nil {
		return true
	}

	return false
}

// SetCapacityBytes gets a reference to the given int64 and assigns it to the CapacityBytes field.
func (o *CSIVolumeExternalStub) SetCapacityBytes(v int64) {
	o.CapacityBytes = &v
}

// GetCloneID returns the CloneID field value if set, zero value otherwise.
func (o *CSIVolumeExternalStub) GetCloneID() string {
	if o == nil || o.CloneID == nil {
		var ret string
		return ret
	}
	return *o.CloneID
}

// GetCloneIDOk returns a tuple with the CloneID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeExternalStub) GetCloneIDOk() (*string, bool) {
	if o == nil || o.CloneID == nil {
		return nil, false
	}
	return o.CloneID, true
}

// HasCloneID returns a boolean if a field has been set.
func (o *CSIVolumeExternalStub) HasCloneID() bool {
	if o != nil && o.CloneID != nil {
		return true
	}

	return false
}

// SetCloneID gets a reference to the given string and assigns it to the CloneID field.
func (o *CSIVolumeExternalStub) SetCloneID(v string) {
	o.CloneID = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *CSIVolumeExternalStub) GetExternalID() string {
	if o == nil || o.ExternalID == nil {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeExternalStub) GetExternalIDOk() (*string, bool) {
	if o == nil || o.ExternalID == nil {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *CSIVolumeExternalStub) HasExternalID() bool {
	if o != nil && o.ExternalID != nil {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *CSIVolumeExternalStub) SetExternalID(v string) {
	o.ExternalID = &v
}

// GetIsAbnormal returns the IsAbnormal field value if set, zero value otherwise.
func (o *CSIVolumeExternalStub) GetIsAbnormal() bool {
	if o == nil || o.IsAbnormal == nil {
		var ret bool
		return ret
	}
	return *o.IsAbnormal
}

// GetIsAbnormalOk returns a tuple with the IsAbnormal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeExternalStub) GetIsAbnormalOk() (*bool, bool) {
	if o == nil || o.IsAbnormal == nil {
		return nil, false
	}
	return o.IsAbnormal, true
}

// HasIsAbnormal returns a boolean if a field has been set.
func (o *CSIVolumeExternalStub) HasIsAbnormal() bool {
	if o != nil && o.IsAbnormal != nil {
		return true
	}

	return false
}

// SetIsAbnormal gets a reference to the given bool and assigns it to the IsAbnormal field.
func (o *CSIVolumeExternalStub) SetIsAbnormal(v bool) {
	o.IsAbnormal = &v
}

// GetPublishedExternalNodeIDs returns the PublishedExternalNodeIDs field value if set, zero value otherwise.
func (o *CSIVolumeExternalStub) GetPublishedExternalNodeIDs() []string {
	if o == nil || o.PublishedExternalNodeIDs == nil {
		var ret []string
		return ret
	}
	return *o.PublishedExternalNodeIDs
}

// GetPublishedExternalNodeIDsOk returns a tuple with the PublishedExternalNodeIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeExternalStub) GetPublishedExternalNodeIDsOk() (*[]string, bool) {
	if o == nil || o.PublishedExternalNodeIDs == nil {
		return nil, false
	}
	return o.PublishedExternalNodeIDs, true
}

// HasPublishedExternalNodeIDs returns a boolean if a field has been set.
func (o *CSIVolumeExternalStub) HasPublishedExternalNodeIDs() bool {
	if o != nil && o.PublishedExternalNodeIDs != nil {
		return true
	}

	return false
}

// SetPublishedExternalNodeIDs gets a reference to the given []string and assigns it to the PublishedExternalNodeIDs field.
func (o *CSIVolumeExternalStub) SetPublishedExternalNodeIDs(v []string) {
	o.PublishedExternalNodeIDs = &v
}

// GetSnapshotID returns the SnapshotID field value if set, zero value otherwise.
func (o *CSIVolumeExternalStub) GetSnapshotID() string {
	if o == nil || o.SnapshotID == nil {
		var ret string
		return ret
	}
	return *o.SnapshotID
}

// GetSnapshotIDOk returns a tuple with the SnapshotID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeExternalStub) GetSnapshotIDOk() (*string, bool) {
	if o == nil || o.SnapshotID == nil {
		return nil, false
	}
	return o.SnapshotID, true
}

// HasSnapshotID returns a boolean if a field has been set.
func (o *CSIVolumeExternalStub) HasSnapshotID() bool {
	if o != nil && o.SnapshotID != nil {
		return true
	}

	return false
}

// SetSnapshotID gets a reference to the given string and assigns it to the SnapshotID field.
func (o *CSIVolumeExternalStub) SetSnapshotID(v string) {
	o.SnapshotID = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CSIVolumeExternalStub) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeExternalStub) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CSIVolumeExternalStub) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CSIVolumeExternalStub) SetStatus(v string) {
	o.Status = &v
}

// GetVolumeContext returns the VolumeContext field value if set, zero value otherwise.
func (o *CSIVolumeExternalStub) GetVolumeContext() map[string]string {
	if o == nil || o.VolumeContext == nil {
		var ret map[string]string
		return ret
	}
	return *o.VolumeContext
}

// GetVolumeContextOk returns a tuple with the VolumeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeExternalStub) GetVolumeContextOk() (*map[string]string, bool) {
	if o == nil || o.VolumeContext == nil {
		return nil, false
	}
	return o.VolumeContext, true
}

// HasVolumeContext returns a boolean if a field has been set.
func (o *CSIVolumeExternalStub) HasVolumeContext() bool {
	if o != nil && o.VolumeContext != nil {
		return true
	}

	return false
}

// SetVolumeContext gets a reference to the given map[string]string and assigns it to the VolumeContext field.
func (o *CSIVolumeExternalStub) SetVolumeContext(v map[string]string) {
	o.VolumeContext = &v
}

func (o CSIVolumeExternalStub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CapacityBytes != nil {
		toSerialize["CapacityBytes"] = o.CapacityBytes
	}
	if o.CloneID != nil {
		toSerialize["CloneID"] = o.CloneID
	}
	if o.ExternalID != nil {
		toSerialize["ExternalID"] = o.ExternalID
	}
	if o.IsAbnormal != nil {
		toSerialize["IsAbnormal"] = o.IsAbnormal
	}
	if o.PublishedExternalNodeIDs != nil {
		toSerialize["PublishedExternalNodeIDs"] = o.PublishedExternalNodeIDs
	}
	if o.SnapshotID != nil {
		toSerialize["SnapshotID"] = o.SnapshotID
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.VolumeContext != nil {
		toSerialize["VolumeContext"] = o.VolumeContext
	}
	return json.Marshal(toSerialize)
}

type NullableCSIVolumeExternalStub struct {
	value *CSIVolumeExternalStub
	isSet bool
}

func (v NullableCSIVolumeExternalStub) Get() *CSIVolumeExternalStub {
	return v.value
}

func (v *NullableCSIVolumeExternalStub) Set(val *CSIVolumeExternalStub) {
	v.value = val
	v.isSet = true
}

func (v NullableCSIVolumeExternalStub) IsSet() bool {
	return v.isSet
}

func (v *NullableCSIVolumeExternalStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSIVolumeExternalStub(val *CSIVolumeExternalStub) *NullableCSIVolumeExternalStub {
	return &NullableCSIVolumeExternalStub{value: val, isSet: true}
}

func (v NullableCSIVolumeExternalStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSIVolumeExternalStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
