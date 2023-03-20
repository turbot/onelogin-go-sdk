/*
OneLogin API

OpenAPI Specification for OneLogin

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onelogin

import (
	"encoding/json"
	"fmt"
)

// AuthMethod An ID indicating the type of app:   - 0: Password   - 1: OpenId   - 2: SAML   - 3: API   - 4: Google   - 6: Forms Based App   - 7: WSFED   - 8: OpenId Connect
type AuthMethod int32

// List of auth_method
const (
	_0 AuthMethod = 0
	_1 AuthMethod = 1
	_2 AuthMethod = 2
	_3 AuthMethod = 3
	_4 AuthMethod = 4
	_5 AuthMethod = 5
	_6 AuthMethod = 6
	_7 AuthMethod = 7
	_8 AuthMethod = 8
)

// All allowed values of AuthMethod enum
var AllowedAuthMethodEnumValues = []AuthMethod{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
}

func (v *AuthMethod) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthMethod(value)
	for _, existing := range AllowedAuthMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthMethod", value)
}

// NewAuthMethodFromValue returns a pointer to a valid AuthMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthMethodFromValue(v int32) (*AuthMethod, error) {
	ev := AuthMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthMethod: valid values are %v", v, AllowedAuthMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthMethod) IsValid() bool {
	for _, existing := range AllowedAuthMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to auth_method value
func (v AuthMethod) Ptr() *AuthMethod {
	return &v
}

type NullableAuthMethod struct {
	value *AuthMethod
	isSet bool
}

func (v NullableAuthMethod) Get() *AuthMethod {
	return v.value
}

func (v *NullableAuthMethod) Set(val *AuthMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthMethod(val *AuthMethod) *NullableAuthMethod {
	return &NullableAuthMethod{value: val, isSet: true}
}

func (v NullableAuthMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

