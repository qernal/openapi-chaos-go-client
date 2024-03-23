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
	"bytes"
	"fmt"
)

// checks if the ListProjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListProjectResponse{}

// ListProjectResponse List of projects
type ListProjectResponse struct {
	Meta PaginationMeta `json:"meta"`
	Data []ProjectResponse `json:"data"`
}

type _ListProjectResponse ListProjectResponse

// NewListProjectResponse instantiates a new ListProjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListProjectResponse(meta PaginationMeta, data []ProjectResponse) *ListProjectResponse {
	this := ListProjectResponse{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewListProjectResponseWithDefaults instantiates a new ListProjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListProjectResponseWithDefaults() *ListProjectResponse {
	this := ListProjectResponse{}
	return &this
}

// GetMeta returns the Meta field value
func (o *ListProjectResponse) GetMeta() PaginationMeta {
	if o == nil {
		var ret PaginationMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ListProjectResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ListProjectResponse) SetMeta(v PaginationMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *ListProjectResponse) GetData() []ProjectResponse {
	if o == nil {
		var ret []ProjectResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListProjectResponse) GetDataOk() ([]ProjectResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListProjectResponse) SetData(v []ProjectResponse) {
	o.Data = v
}

func (o ListProjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListProjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ListProjectResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListProjectResponse := _ListProjectResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListProjectResponse)

	if err != nil {
		return err
	}

	*o = ListProjectResponse(varListProjectResponse)

	return err
}

type NullableListProjectResponse struct {
	value *ListProjectResponse
	isSet bool
}

func (v NullableListProjectResponse) Get() *ListProjectResponse {
	return v.value
}

func (v *NullableListProjectResponse) Set(val *ListProjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListProjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListProjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListProjectResponse(val *ListProjectResponse) *NullableListProjectResponse {
	return &NullableListProjectResponse{value: val, isSet: true}
}

func (v NullableListProjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListProjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


