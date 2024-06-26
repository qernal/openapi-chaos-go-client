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

// checks if the FunctionRoute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionRoute{}

// FunctionRoute struct for FunctionRoute
type FunctionRoute struct {
	// Can be a regular expression
	Path string `json:"path"`
	// HTTP Verb(s) for this function
	Methods []string `json:"methods"`
	// The route weight for consideration
	Weight int32 `json:"weight"`
}

type _FunctionRoute FunctionRoute

// NewFunctionRoute instantiates a new FunctionRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionRoute(path string, methods []string, weight int32) *FunctionRoute {
	this := FunctionRoute{}
	this.Path = path
	this.Methods = methods
	this.Weight = weight
	return &this
}

// NewFunctionRouteWithDefaults instantiates a new FunctionRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionRouteWithDefaults() *FunctionRoute {
	this := FunctionRoute{}
	return &this
}

// GetPath returns the Path field value
func (o *FunctionRoute) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *FunctionRoute) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *FunctionRoute) SetPath(v string) {
	o.Path = v
}

// GetMethods returns the Methods field value
func (o *FunctionRoute) GetMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value
// and a boolean to check if the value has been set.
func (o *FunctionRoute) GetMethodsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Methods, true
}

// SetMethods sets field value
func (o *FunctionRoute) SetMethods(v []string) {
	o.Methods = v
}

// GetWeight returns the Weight field value
func (o *FunctionRoute) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *FunctionRoute) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *FunctionRoute) SetWeight(v int32) {
	o.Weight = v
}

func (o FunctionRoute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["methods"] = o.Methods
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

func (o *FunctionRoute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
		"methods",
		"weight",
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

	varFunctionRoute := _FunctionRoute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFunctionRoute)

	if err != nil {
		return err
	}

	*o = FunctionRoute(varFunctionRoute)

	return err
}

type NullableFunctionRoute struct {
	value *FunctionRoute
	isSet bool
}

func (v NullableFunctionRoute) Get() *FunctionRoute {
	return v.value
}

func (v *NullableFunctionRoute) Set(val *FunctionRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionRoute(val *FunctionRoute) *NullableFunctionRoute {
	return &NullableFunctionRoute{value: val, isSet: true}
}

func (v NullableFunctionRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


