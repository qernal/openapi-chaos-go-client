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
	"fmt"
)

// FunctionType Type of function, worker types are not exposed to ingress routes
type FunctionType string

// List of FunctionType
const (
	FUNCTIONTYPE_HTTP FunctionType = "http"
	FUNCTIONTYPE_WORKER FunctionType = "worker"
)

// All allowed values of FunctionType enum
var AllowedFunctionTypeEnumValues = []FunctionType{
	"http",
	"worker",
}

func (v *FunctionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FunctionType(value)
	for _, existing := range AllowedFunctionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FunctionType", value)
}

// NewFunctionTypeFromValue returns a pointer to a valid FunctionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFunctionTypeFromValue(v string) (*FunctionType, error) {
	ev := FunctionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FunctionType: valid values are %v", v, AllowedFunctionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FunctionType) IsValid() bool {
	for _, existing := range AllowedFunctionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FunctionType value
func (v FunctionType) Ptr() *FunctionType {
	return &v
}

type NullableFunctionType struct {
	value *FunctionType
	isSet bool
}

func (v NullableFunctionType) Get() *FunctionType {
	return v.value
}

func (v *NullableFunctionType) Set(val *FunctionType) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionType) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionType(val *FunctionType) *NullableFunctionType {
	return &NullableFunctionType{value: val, isSet: true}
}

func (v NullableFunctionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

