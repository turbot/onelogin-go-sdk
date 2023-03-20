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

// checks if the SsoSamlCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoSamlCertificate{}

// SsoSamlCertificate The Certificate used for signing
type SsoSamlCertificate struct {
	// SAML Certificate ID
	Id *int32 `json:"id,omitempty"`
	// SAML Certificate Name
	Name *string `json:"name,omitempty"`
	// SAML Certificate Value
	Value *string `json:"value,omitempty"`
}

// NewSsoSamlCertificate instantiates a new SsoSamlCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoSamlCertificate() *SsoSamlCertificate {
	this := SsoSamlCertificate{}
	return &this
}

// NewSsoSamlCertificateWithDefaults instantiates a new SsoSamlCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoSamlCertificateWithDefaults() *SsoSamlCertificate {
	this := SsoSamlCertificate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SsoSamlCertificate) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSamlCertificate) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SsoSamlCertificate) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SsoSamlCertificate) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SsoSamlCertificate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSamlCertificate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SsoSamlCertificate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SsoSamlCertificate) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SsoSamlCertificate) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoSamlCertificate) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SsoSamlCertificate) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SsoSamlCertificate) SetValue(v string) {
	o.Value = &v
}

func (o SsoSamlCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoSamlCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSsoSamlCertificate struct {
	value *SsoSamlCertificate
	isSet bool
}

func (v NullableSsoSamlCertificate) Get() *SsoSamlCertificate {
	return v.value
}

func (v *NullableSsoSamlCertificate) Set(val *SsoSamlCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoSamlCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoSamlCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoSamlCertificate(val *SsoSamlCertificate) *NullableSsoSamlCertificate {
	return &NullableSsoSamlCertificate{value: val, isSet: true}
}

func (v NullableSsoSamlCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoSamlCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


