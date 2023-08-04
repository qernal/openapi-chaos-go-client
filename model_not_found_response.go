/*
Chaos

Central Management API - publicly exposed set of APIs for managing deployments

API version: 1.0.0
Contact: support@qernal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-chaos-client

import (
	"encoding/json"
)

// checks if the NotFoundResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotFoundResponse{}

// NotFoundResponse Resource not found
type NotFoundResponse struct {
	Message string `json:"message"`
}

// NewNotFoundResponse instantiates a new NotFoundResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotFoundResponse(message string) *NotFoundResponse {
	this := NotFoundResponse{}
	this.Message = message
	return &this
}

// NewNotFoundResponseWithDefaults instantiates a new NotFoundResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotFoundResponseWithDefaults() *NotFoundResponse {
	this := NotFoundResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *NotFoundResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NotFoundResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NotFoundResponse) SetMessage(v string) {
	o.Message = v
}

func (o NotFoundResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotFoundResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableNotFoundResponse struct {
	value *NotFoundResponse
	isSet bool
}

func (v NullableNotFoundResponse) Get() *NotFoundResponse {
	return v.value
}

func (v *NullableNotFoundResponse) Set(val *NotFoundResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundResponse(val *NotFoundResponse) *NullableNotFoundResponse {
	return &NullableNotFoundResponse{value: val, isSet: true}
}

func (v NullableNotFoundResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFoundResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


