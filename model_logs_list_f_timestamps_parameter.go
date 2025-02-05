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

// checks if the LogsListFTimestampsParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogsListFTimestampsParameter{}

// LogsListFTimestampsParameter struct for LogsListFTimestampsParameter
type LogsListFTimestampsParameter struct {
	// Restrict to after this timestamp
	After *string `json:"after,omitempty"`
	// Resitrct to before this timestamp
	Before *string `json:"before,omitempty"`
}

// NewLogsListFTimestampsParameter instantiates a new LogsListFTimestampsParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsListFTimestampsParameter() *LogsListFTimestampsParameter {
	this := LogsListFTimestampsParameter{}
	return &this
}

// NewLogsListFTimestampsParameterWithDefaults instantiates a new LogsListFTimestampsParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsListFTimestampsParameterWithDefaults() *LogsListFTimestampsParameter {
	this := LogsListFTimestampsParameter{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *LogsListFTimestampsParameter) GetAfter() string {
	if o == nil || IsNil(o.After) {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsListFTimestampsParameter) GetAfterOk() (*string, bool) {
	if o == nil || IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *LogsListFTimestampsParameter) HasAfter() bool {
	if o != nil && !IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *LogsListFTimestampsParameter) SetAfter(v string) {
	o.After = &v
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *LogsListFTimestampsParameter) GetBefore() string {
	if o == nil || IsNil(o.Before) {
		var ret string
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsListFTimestampsParameter) GetBeforeOk() (*string, bool) {
	if o == nil || IsNil(o.Before) {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *LogsListFTimestampsParameter) HasBefore() bool {
	if o != nil && !IsNil(o.Before) {
		return true
	}

	return false
}

// SetBefore gets a reference to the given string and assigns it to the Before field.
func (o *LogsListFTimestampsParameter) SetBefore(v string) {
	o.Before = &v
}

func (o LogsListFTimestampsParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogsListFTimestampsParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.After) {
		toSerialize["after"] = o.After
	}
	if !IsNil(o.Before) {
		toSerialize["before"] = o.Before
	}
	return toSerialize, nil
}

type NullableLogsListFTimestampsParameter struct {
	value *LogsListFTimestampsParameter
	isSet bool
}

func (v NullableLogsListFTimestampsParameter) Get() *LogsListFTimestampsParameter {
	return v.value
}

func (v *NullableLogsListFTimestampsParameter) Set(val *LogsListFTimestampsParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsListFTimestampsParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsListFTimestampsParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsListFTimestampsParameter(val *LogsListFTimestampsParameter) *NullableLogsListFTimestampsParameter {
	return &NullableLogsListFTimestampsParameter{value: val, isSet: true}
}

func (v NullableLogsListFTimestampsParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsListFTimestampsParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


