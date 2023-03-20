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

// checks if the UpdatePasswordInsecureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePasswordInsecureRequest{}

// UpdatePasswordInsecureRequest struct for UpdatePasswordInsecureRequest
type UpdatePasswordInsecureRequest struct {
	// Set to the password value using cleartext. Hashes are never stored as cleartext. They are stored securely using cryptographic hash. OneLogin continuously upgrades the strength of the hash. Ensure that the value meets the password strength requirements set for the account.
	Password string `json:"password"`
	// Ensure that this value matches the password value exactly.
	PasswordConfirmation string `json:"password_confirmation"`
	// Will passwords validate against the User Policy. Defaults to false.
	ValidatePolicy *bool `json:"validate_policy,omitempty"`
}

// NewUpdatePasswordInsecureRequest instantiates a new UpdatePasswordInsecureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePasswordInsecureRequest(password string, passwordConfirmation string) *UpdatePasswordInsecureRequest {
	this := UpdatePasswordInsecureRequest{}
	this.Password = password
	this.PasswordConfirmation = passwordConfirmation
	var validatePolicy bool = false
	this.ValidatePolicy = &validatePolicy
	return &this
}

// NewUpdatePasswordInsecureRequestWithDefaults instantiates a new UpdatePasswordInsecureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePasswordInsecureRequestWithDefaults() *UpdatePasswordInsecureRequest {
	this := UpdatePasswordInsecureRequest{}
	var validatePolicy bool = false
	this.ValidatePolicy = &validatePolicy
	return &this
}

// GetPassword returns the Password field value
func (o *UpdatePasswordInsecureRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UpdatePasswordInsecureRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UpdatePasswordInsecureRequest) SetPassword(v string) {
	o.Password = v
}

// GetPasswordConfirmation returns the PasswordConfirmation field value
func (o *UpdatePasswordInsecureRequest) GetPasswordConfirmation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordConfirmation
}

// GetPasswordConfirmationOk returns a tuple with the PasswordConfirmation field value
// and a boolean to check if the value has been set.
func (o *UpdatePasswordInsecureRequest) GetPasswordConfirmationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordConfirmation, true
}

// SetPasswordConfirmation sets field value
func (o *UpdatePasswordInsecureRequest) SetPasswordConfirmation(v string) {
	o.PasswordConfirmation = v
}

// GetValidatePolicy returns the ValidatePolicy field value if set, zero value otherwise.
func (o *UpdatePasswordInsecureRequest) GetValidatePolicy() bool {
	if o == nil || isNil(o.ValidatePolicy) {
		var ret bool
		return ret
	}
	return *o.ValidatePolicy
}

// GetValidatePolicyOk returns a tuple with the ValidatePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePasswordInsecureRequest) GetValidatePolicyOk() (*bool, bool) {
	if o == nil || isNil(o.ValidatePolicy) {
		return nil, false
	}
	return o.ValidatePolicy, true
}

// HasValidatePolicy returns a boolean if a field has been set.
func (o *UpdatePasswordInsecureRequest) HasValidatePolicy() bool {
	if o != nil && !isNil(o.ValidatePolicy) {
		return true
	}

	return false
}

// SetValidatePolicy gets a reference to the given bool and assigns it to the ValidatePolicy field.
func (o *UpdatePasswordInsecureRequest) SetValidatePolicy(v bool) {
	o.ValidatePolicy = &v
}

func (o UpdatePasswordInsecureRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePasswordInsecureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	toSerialize["password_confirmation"] = o.PasswordConfirmation
	if !isNil(o.ValidatePolicy) {
		toSerialize["validate_policy"] = o.ValidatePolicy
	}
	return toSerialize, nil
}

type NullableUpdatePasswordInsecureRequest struct {
	value *UpdatePasswordInsecureRequest
	isSet bool
}

func (v NullableUpdatePasswordInsecureRequest) Get() *UpdatePasswordInsecureRequest {
	return v.value
}

func (v *NullableUpdatePasswordInsecureRequest) Set(val *UpdatePasswordInsecureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePasswordInsecureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePasswordInsecureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePasswordInsecureRequest(val *UpdatePasswordInsecureRequest) *NullableUpdatePasswordInsecureRequest {
	return &NullableUpdatePasswordInsecureRequest{value: val, isSet: true}
}

func (v NullableUpdatePasswordInsecureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePasswordInsecureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

