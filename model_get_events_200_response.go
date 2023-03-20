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

// checks if the GetEvents200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEvents200Response{}

// GetEvents200Response struct for GetEvents200Response
type GetEvents200Response struct {
	Status *Error `json:"status,omitempty"`
	Pagination *GetEvents200ResponsePagination `json:"pagination,omitempty"`
	Data []Event `json:"data,omitempty"`
}

// NewGetEvents200Response instantiates a new GetEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEvents200Response() *GetEvents200Response {
	this := GetEvents200Response{}
	return &this
}

// NewGetEvents200ResponseWithDefaults instantiates a new GetEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEvents200ResponseWithDefaults() *GetEvents200Response {
	this := GetEvents200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetEvents200Response) GetStatus() Error {
	if o == nil || isNil(o.Status) {
		var ret Error
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200Response) GetStatusOk() (*Error, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetEvents200Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Error and assigns it to the Status field.
func (o *GetEvents200Response) SetStatus(v Error) {
	o.Status = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetEvents200Response) GetPagination() GetEvents200ResponsePagination {
	if o == nil || isNil(o.Pagination) {
		var ret GetEvents200ResponsePagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200Response) GetPaginationOk() (*GetEvents200ResponsePagination, bool) {
	if o == nil || isNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetEvents200Response) HasPagination() bool {
	if o != nil && !isNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given GetEvents200ResponsePagination and assigns it to the Pagination field.
func (o *GetEvents200Response) SetPagination(v GetEvents200ResponsePagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetEvents200Response) GetData() []Event {
	if o == nil || isNil(o.Data) {
		var ret []Event
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200Response) GetDataOk() ([]Event, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetEvents200Response) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Event and assigns it to the Data field.
func (o *GetEvents200Response) SetData(v []Event) {
	o.Data = v
}

func (o GetEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEvents200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetEvents200Response struct {
	value *GetEvents200Response
	isSet bool
}

func (v NullableGetEvents200Response) Get() *GetEvents200Response {
	return v.value
}

func (v *NullableGetEvents200Response) Set(val *GetEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEvents200Response(val *GetEvents200Response) *NullableGetEvents200Response {
	return &NullableGetEvents200Response{value: val, isSet: true}
}

func (v NullableGetEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


