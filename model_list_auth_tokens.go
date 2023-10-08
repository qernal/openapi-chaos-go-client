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

// checks if the ListAuthTokens type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAuthTokens{}

// ListAuthTokens List of auth tokens
type ListAuthTokens struct {
	Meta PaginationMeta `json:"meta"`
	Data []AuthTokenMeta `json:"data"`
}

// NewListAuthTokens instantiates a new ListAuthTokens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAuthTokens(meta PaginationMeta, data []AuthTokenMeta) *ListAuthTokens {
	this := ListAuthTokens{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewListAuthTokensWithDefaults instantiates a new ListAuthTokens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAuthTokensWithDefaults() *ListAuthTokens {
	this := ListAuthTokens{}
	return &this
}

// GetMeta returns the Meta field value
func (o *ListAuthTokens) GetMeta() PaginationMeta {
	if o == nil {
		var ret PaginationMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ListAuthTokens) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ListAuthTokens) SetMeta(v PaginationMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *ListAuthTokens) GetData() []AuthTokenMeta {
	if o == nil {
		var ret []AuthTokenMeta
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListAuthTokens) GetDataOk() ([]AuthTokenMeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListAuthTokens) SetData(v []AuthTokenMeta) {
	o.Data = v
}

func (o ListAuthTokens) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAuthTokens) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListAuthTokens struct {
	value *ListAuthTokens
	isSet bool
}

func (v NullableListAuthTokens) Get() *ListAuthTokens {
	return v.value
}

func (v *NullableListAuthTokens) Set(val *ListAuthTokens) {
	v.value = val
	v.isSet = true
}

func (v NullableListAuthTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableListAuthTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAuthTokens(val *ListAuthTokens) *NullableListAuthTokens {
	return &NullableListAuthTokens{value: val, isSet: true}
}

func (v NullableListAuthTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAuthTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


