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

// checks if the PrivilegePrivilege type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivilegePrivilege{}

// PrivilegePrivilege struct for PrivilegePrivilege
type PrivilegePrivilege struct {
	Version *string `json:"Version,omitempty"`
	Statement []PrivilegePrivilegeStatementInner `json:"Statement,omitempty"`
}

// NewPrivilegePrivilege instantiates a new PrivilegePrivilege object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegePrivilege() *PrivilegePrivilege {
	this := PrivilegePrivilege{}
	return &this
}

// NewPrivilegePrivilegeWithDefaults instantiates a new PrivilegePrivilege object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegePrivilegeWithDefaults() *PrivilegePrivilege {
	this := PrivilegePrivilege{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PrivilegePrivilege) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegePrivilege) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PrivilegePrivilege) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PrivilegePrivilege) SetVersion(v string) {
	o.Version = &v
}

// GetStatement returns the Statement field value if set, zero value otherwise.
func (o *PrivilegePrivilege) GetStatement() []PrivilegePrivilegeStatementInner {
	if o == nil || isNil(o.Statement) {
		var ret []PrivilegePrivilegeStatementInner
		return ret
	}
	return o.Statement
}

// GetStatementOk returns a tuple with the Statement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegePrivilege) GetStatementOk() ([]PrivilegePrivilegeStatementInner, bool) {
	if o == nil || isNil(o.Statement) {
		return nil, false
	}
	return o.Statement, true
}

// HasStatement returns a boolean if a field has been set.
func (o *PrivilegePrivilege) HasStatement() bool {
	if o != nil && !isNil(o.Statement) {
		return true
	}

	return false
}

// SetStatement gets a reference to the given []PrivilegePrivilegeStatementInner and assigns it to the Statement field.
func (o *PrivilegePrivilege) SetStatement(v []PrivilegePrivilegeStatementInner) {
	o.Statement = v
}

func (o PrivilegePrivilege) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivilegePrivilege) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !isNil(o.Statement) {
		toSerialize["Statement"] = o.Statement
	}
	return toSerialize, nil
}

type NullablePrivilegePrivilege struct {
	value *PrivilegePrivilege
	isSet bool
}

func (v NullablePrivilegePrivilege) Get() *PrivilegePrivilege {
	return v.value
}

func (v *NullablePrivilegePrivilege) Set(val *PrivilegePrivilege) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegePrivilege) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegePrivilege) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegePrivilege(val *PrivilegePrivilege) *NullablePrivilegePrivilege {
	return &NullablePrivilegePrivilege{value: val, isSet: true}
}

func (v NullablePrivilegePrivilege) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegePrivilege) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


