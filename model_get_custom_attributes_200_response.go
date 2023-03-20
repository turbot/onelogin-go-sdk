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

// checks if the GetCustomAttributes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCustomAttributes200Response{}

// GetCustomAttributes200Response struct for GetCustomAttributes200Response
type GetCustomAttributes200Response struct {
	Status *Error `json:"status,omitempty"`
	// Provides a list of custom attribute fields (also known as custom user fields) that are available for your account. The values returned correspond to the values you provided in the Shortname field when you defined the custom user field. For details about defining custom user fields, see Custom User Fields.
	Data [][]string `json:"data,omitempty"`
}

// NewGetCustomAttributes200Response instantiates a new GetCustomAttributes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomAttributes200Response() *GetCustomAttributes200Response {
	this := GetCustomAttributes200Response{}
	return &this
}

// NewGetCustomAttributes200ResponseWithDefaults instantiates a new GetCustomAttributes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomAttributes200ResponseWithDefaults() *GetCustomAttributes200Response {
	this := GetCustomAttributes200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetCustomAttributes200Response) GetStatus() Error {
	if o == nil || isNil(o.Status) {
		var ret Error
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCustomAttributes200Response) GetStatusOk() (*Error, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetCustomAttributes200Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Error and assigns it to the Status field.
func (o *GetCustomAttributes200Response) SetStatus(v Error) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCustomAttributes200Response) GetData() [][]string {
	if o == nil || isNil(o.Data) {
		var ret [][]string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCustomAttributes200Response) GetDataOk() ([][]string, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCustomAttributes200Response) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given [][]string and assigns it to the Data field.
func (o *GetCustomAttributes200Response) SetData(v [][]string) {
	o.Data = v
}

func (o GetCustomAttributes200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCustomAttributes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetCustomAttributes200Response struct {
	value *GetCustomAttributes200Response
	isSet bool
}

func (v NullableGetCustomAttributes200Response) Get() *GetCustomAttributes200Response {
	return v.value
}

func (v *NullableGetCustomAttributes200Response) Set(val *GetCustomAttributes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomAttributes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomAttributes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomAttributes200Response(val *GetCustomAttributes200Response) *NullableGetCustomAttributes200Response {
	return &NullableGetCustomAttributes200Response{value: val, isSet: true}
}

func (v NullableGetCustomAttributes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomAttributes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

