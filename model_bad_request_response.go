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

// checks if the BadRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BadRequestResponse{}

// BadRequestResponse Bad request
type BadRequestResponse struct {
	Message string `json:"message"`
	Fields BadRequestResponseFields `json:"fields"`
}

// NewBadRequestResponse instantiates a new BadRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestResponse(message string, fields BadRequestResponseFields) *BadRequestResponse {
	this := BadRequestResponse{}
	this.Message = message
	this.Fields = fields
	return &this
}

// NewBadRequestResponseWithDefaults instantiates a new BadRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestResponseWithDefaults() *BadRequestResponse {
	this := BadRequestResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *BadRequestResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BadRequestResponse) SetMessage(v string) {
	o.Message = v
}

// GetFields returns the Fields field value
func (o *BadRequestResponse) GetFields() BadRequestResponseFields {
	if o == nil {
		var ret BadRequestResponseFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetFieldsOk() (*BadRequestResponseFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *BadRequestResponse) SetFields(v BadRequestResponseFields) {
	o.Fields = v
}

func (o BadRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BadRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

type NullableBadRequestResponse struct {
	value *BadRequestResponse
	isSet bool
}

func (v NullableBadRequestResponse) Get() *BadRequestResponse {
	return v.value
}

func (v *NullableBadRequestResponse) Set(val *BadRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestResponse(val *BadRequestResponse) *NullableBadRequestResponse {
	return &NullableBadRequestResponse{value: val, isSet: true}
}

func (v NullableBadRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


