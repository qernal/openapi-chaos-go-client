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

// checks if the LogLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogLog{}

// LogLog Log item
type LogLog struct {
	// Which log stream
	Stream *string `json:"stream,omitempty"`
	// If this was an event on the function or a log line
	Kind *string `json:"kind,omitempty"`
	// An array of labels
	Labels []string `json:"labels,omitempty"`
	// Log line type
	Type *string `json:"type,omitempty"`
	// Log line
	Line *string `json:"line,omitempty"`
	// The date/time that this log was generated
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewLogLog instantiates a new LogLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogLog() *LogLog {
	this := LogLog{}
	return &this
}

// NewLogLogWithDefaults instantiates a new LogLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogLogWithDefaults() *LogLog {
	this := LogLog{}
	return &this
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *LogLog) GetStream() string {
	if o == nil || IsNil(o.Stream) {
		var ret string
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogLog) GetStreamOk() (*string, bool) {
	if o == nil || IsNil(o.Stream) {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *LogLog) HasStream() bool {
	if o != nil && !IsNil(o.Stream) {
		return true
	}

	return false
}

// SetStream gets a reference to the given string and assigns it to the Stream field.
func (o *LogLog) SetStream(v string) {
	o.Stream = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *LogLog) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogLog) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *LogLog) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *LogLog) SetKind(v string) {
	o.Kind = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *LogLog) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogLog) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LogLog) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *LogLog) SetLabels(v []string) {
	o.Labels = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogLog) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogLog) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogLog) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogLog) SetType(v string) {
	o.Type = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *LogLog) GetLine() string {
	if o == nil || IsNil(o.Line) {
		var ret string
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogLog) GetLineOk() (*string, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *LogLog) HasLine() bool {
	if o != nil && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given string and assigns it to the Line field.
func (o *LogLog) SetLine(v string) {
	o.Line = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LogLog) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogLog) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LogLog) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *LogLog) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o LogLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Stream) {
		toSerialize["stream"] = o.Stream
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Line) {
		toSerialize["line"] = o.Line
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableLogLog struct {
	value *LogLog
	isSet bool
}

func (v NullableLogLog) Get() *LogLog {
	return v.value
}

func (v *NullableLogLog) Set(val *LogLog) {
	v.value = val
	v.isSet = true
}

func (v NullableLogLog) IsSet() bool {
	return v.isSet
}

func (v *NullableLogLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogLog(val *LogLog) *NullableLogLog {
	return &NullableLogLog{value: val, isSet: true}
}

func (v NullableLogLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


