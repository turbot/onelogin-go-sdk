/*
OneLogin API

OpenAPI Specification for OneLogin

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onelogin

import (
	"encoding/json"
)

// checks if the HookOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HookOptions{}

// HookOptions A set of attributes allow control over the information that is included in the hook context.
type HookOptions struct {
	RiskEnabled *bool `json:"risk_enabled,omitempty"`
	LocationEnabled *bool `json:"location_enabled,omitempty"`
	MfaDeviceInfoEnabled *bool `json:"mfa_device_info_enabled,omitempty"`
}

// NewHookOptions instantiates a new HookOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHookOptions() *HookOptions {
	this := HookOptions{}
	return &this
}

// NewHookOptionsWithDefaults instantiates a new HookOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHookOptionsWithDefaults() *HookOptions {
	this := HookOptions{}
	return &this
}

// GetRiskEnabled returns the RiskEnabled field value if set, zero value otherwise.
func (o *HookOptions) GetRiskEnabled() bool {
	if o == nil || isNil(o.RiskEnabled) {
		var ret bool
		return ret
	}
	return *o.RiskEnabled
}

// GetRiskEnabledOk returns a tuple with the RiskEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookOptions) GetRiskEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.RiskEnabled) {
		return nil, false
	}
	return o.RiskEnabled, true
}

// HasRiskEnabled returns a boolean if a field has been set.
func (o *HookOptions) HasRiskEnabled() bool {
	if o != nil && !isNil(o.RiskEnabled) {
		return true
	}

	return false
}

// SetRiskEnabled gets a reference to the given bool and assigns it to the RiskEnabled field.
func (o *HookOptions) SetRiskEnabled(v bool) {
	o.RiskEnabled = &v
}

// GetLocationEnabled returns the LocationEnabled field value if set, zero value otherwise.
func (o *HookOptions) GetLocationEnabled() bool {
	if o == nil || isNil(o.LocationEnabled) {
		var ret bool
		return ret
	}
	return *o.LocationEnabled
}

// GetLocationEnabledOk returns a tuple with the LocationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookOptions) GetLocationEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.LocationEnabled) {
		return nil, false
	}
	return o.LocationEnabled, true
}

// HasLocationEnabled returns a boolean if a field has been set.
func (o *HookOptions) HasLocationEnabled() bool {
	if o != nil && !isNil(o.LocationEnabled) {
		return true
	}

	return false
}

// SetLocationEnabled gets a reference to the given bool and assigns it to the LocationEnabled field.
func (o *HookOptions) SetLocationEnabled(v bool) {
	o.LocationEnabled = &v
}

// GetMfaDeviceInfoEnabled returns the MfaDeviceInfoEnabled field value if set, zero value otherwise.
func (o *HookOptions) GetMfaDeviceInfoEnabled() bool {
	if o == nil || isNil(o.MfaDeviceInfoEnabled) {
		var ret bool
		return ret
	}
	return *o.MfaDeviceInfoEnabled
}

// GetMfaDeviceInfoEnabledOk returns a tuple with the MfaDeviceInfoEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookOptions) GetMfaDeviceInfoEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.MfaDeviceInfoEnabled) {
		return nil, false
	}
	return o.MfaDeviceInfoEnabled, true
}

// HasMfaDeviceInfoEnabled returns a boolean if a field has been set.
func (o *HookOptions) HasMfaDeviceInfoEnabled() bool {
	if o != nil && !isNil(o.MfaDeviceInfoEnabled) {
		return true
	}

	return false
}

// SetMfaDeviceInfoEnabled gets a reference to the given bool and assigns it to the MfaDeviceInfoEnabled field.
func (o *HookOptions) SetMfaDeviceInfoEnabled(v bool) {
	o.MfaDeviceInfoEnabled = &v
}

func (o HookOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HookOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RiskEnabled) {
		toSerialize["risk_enabled"] = o.RiskEnabled
	}
	if !isNil(o.LocationEnabled) {
		toSerialize["location_enabled"] = o.LocationEnabled
	}
	if !isNil(o.MfaDeviceInfoEnabled) {
		toSerialize["mfa_device_info_enabled"] = o.MfaDeviceInfoEnabled
	}
	return toSerialize, nil
}

type NullableHookOptions struct {
	value *HookOptions
	isSet bool
}

func (v NullableHookOptions) Get() *HookOptions {
	return v.value
}

func (v *NullableHookOptions) Set(val *HookOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableHookOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableHookOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHookOptions(val *HookOptions) *NullableHookOptions {
	return &NullableHookOptions{value: val, isSet: true}
}

func (v NullableHookOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHookOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


