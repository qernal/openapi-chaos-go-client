/*
Chaos

Central Management API - publicly exposed set of APIs for managing deployments

API version: 1.0.0
Contact: support@qernal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_chaos_client

import (
	"encoding/json"
)

// checks if the Date type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Date{}

// Date Object date information
type Date struct {
	// UTC creation datetime (ISO 8601 date format)
	CreatedAt string `json:"created_at"`
	// UTC update datetime (ISO 8601 date format)
	UpdatedAt string `json:"updated_at"`
}

// NewDate instantiates a new Date object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDate(createdAt string, updatedAt string) *Date {
	this := Date{}
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewDateWithDefaults instantiates a new Date object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateWithDefaults() *Date {
	this := Date{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Date) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Date) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Date) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Date) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Date) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Date) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o Date) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Date) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableDate struct {
	value *Date
	isSet bool
}

func (v NullableDate) Get() *Date {
	return v.value
}

func (v *NullableDate) Set(val *Date) {
	v.value = val
	v.isSet = true
}

func (v NullableDate) IsSet() bool {
	return v.isSet
}

func (v *NullableDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDate(val *Date) *NullableDate {
	return &NullableDate{value: val, isSet: true}
}

func (v NullableDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


