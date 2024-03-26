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

// checks if the FunctionDeploymentBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionDeploymentBody{}

// FunctionDeploymentBody struct for FunctionDeploymentBody
type FunctionDeploymentBody struct {
	Location Location `json:"location"`
	Replicas FunctionReplicas `json:"replicas"`
}

type _FunctionDeploymentBody FunctionDeploymentBody

// NewFunctionDeploymentBody instantiates a new FunctionDeploymentBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionDeploymentBody(location Location, replicas FunctionReplicas) *FunctionDeploymentBody {
	this := FunctionDeploymentBody{}
	this.Location = location
	this.Replicas = replicas
	return &this
}

// NewFunctionDeploymentBodyWithDefaults instantiates a new FunctionDeploymentBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionDeploymentBodyWithDefaults() *FunctionDeploymentBody {
	this := FunctionDeploymentBody{}
	return &this
}

// GetLocation returns the Location field value
func (o *FunctionDeploymentBody) GetLocation() Location {
	if o == nil {
		var ret Location
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *FunctionDeploymentBody) GetLocationOk() (*Location, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *FunctionDeploymentBody) SetLocation(v Location) {
	o.Location = v
}

// GetReplicas returns the Replicas field value
func (o *FunctionDeploymentBody) GetReplicas() FunctionReplicas {
	if o == nil {
		var ret FunctionReplicas
		return ret
	}

	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *FunctionDeploymentBody) GetReplicasOk() (*FunctionReplicas, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value
func (o *FunctionDeploymentBody) SetReplicas(v FunctionReplicas) {
	o.Replicas = v
}

func (o FunctionDeploymentBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionDeploymentBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	toSerialize["replicas"] = o.Replicas
	return toSerialize, nil
}

func (o *FunctionDeploymentBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"location",
		"replicas",
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

	varFunctionDeploymentBody := _FunctionDeploymentBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFunctionDeploymentBody)

	if err != nil {
		return err
	}

	*o = FunctionDeploymentBody(varFunctionDeploymentBody)

	return err
}

type NullableFunctionDeploymentBody struct {
	value *FunctionDeploymentBody
	isSet bool
}

func (v NullableFunctionDeploymentBody) Get() *FunctionDeploymentBody {
	return v.value
}

func (v *NullableFunctionDeploymentBody) Set(val *FunctionDeploymentBody) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionDeploymentBody) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionDeploymentBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionDeploymentBody(val *FunctionDeploymentBody) *NullableFunctionDeploymentBody {
	return &NullableFunctionDeploymentBody{value: val, isSet: true}
}

func (v NullableFunctionDeploymentBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionDeploymentBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


