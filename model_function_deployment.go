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

// checks if the FunctionDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionDeployment{}

// FunctionDeployment struct for FunctionDeployment
type FunctionDeployment struct {
	// ID of the deployment
	Id *string `json:"id,omitempty"`
	Location Location `json:"location"`
	Replicas FunctionReplicas `json:"replicas"`
}

type _FunctionDeployment FunctionDeployment

// NewFunctionDeployment instantiates a new FunctionDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionDeployment(location Location, replicas FunctionReplicas) *FunctionDeployment {
	this := FunctionDeployment{}
	this.Location = location
	this.Replicas = replicas
	return &this
}

// NewFunctionDeploymentWithDefaults instantiates a new FunctionDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionDeploymentWithDefaults() *FunctionDeployment {
	this := FunctionDeployment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FunctionDeployment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunctionDeployment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FunctionDeployment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FunctionDeployment) SetId(v string) {
	o.Id = &v
}

// GetLocation returns the Location field value
func (o *FunctionDeployment) GetLocation() Location {
	if o == nil {
		var ret Location
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *FunctionDeployment) GetLocationOk() (*Location, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *FunctionDeployment) SetLocation(v Location) {
	o.Location = v
}

// GetReplicas returns the Replicas field value
func (o *FunctionDeployment) GetReplicas() FunctionReplicas {
	if o == nil {
		var ret FunctionReplicas
		return ret
	}

	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *FunctionDeployment) GetReplicasOk() (*FunctionReplicas, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value
func (o *FunctionDeployment) SetReplicas(v FunctionReplicas) {
	o.Replicas = v
}

func (o FunctionDeployment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["location"] = o.Location
	toSerialize["replicas"] = o.Replicas
	return toSerialize, nil
}

func (o *FunctionDeployment) UnmarshalJSON(data []byte) (err error) {
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

	varFunctionDeployment := _FunctionDeployment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFunctionDeployment)

	if err != nil {
		return err
	}

	*o = FunctionDeployment(varFunctionDeployment)

	return err
}

type NullableFunctionDeployment struct {
	value *FunctionDeployment
	isSet bool
}

func (v NullableFunctionDeployment) Get() *FunctionDeployment {
	return v.value
}

func (v *NullableFunctionDeployment) Set(val *FunctionDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionDeployment(val *FunctionDeployment) *NullableFunctionDeployment {
	return &NullableFunctionDeployment{value: val, isSet: true}
}

func (v NullableFunctionDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


