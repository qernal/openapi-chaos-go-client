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

// checks if the SecretEnvironment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretEnvironment{}

// SecretEnvironment Encrypted ENV secret, `type: environment`
type SecretEnvironment struct {
	// Encrypted environment value
	EnvironmentValue string `json:"environment_value"`
}

type _SecretEnvironment SecretEnvironment

// NewSecretEnvironment instantiates a new SecretEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretEnvironment(environmentValue string) *SecretEnvironment {
	this := SecretEnvironment{}
	this.EnvironmentValue = environmentValue
	return &this
}

// NewSecretEnvironmentWithDefaults instantiates a new SecretEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretEnvironmentWithDefaults() *SecretEnvironment {
	this := SecretEnvironment{}
	return &this
}

// GetEnvironmentValue returns the EnvironmentValue field value
func (o *SecretEnvironment) GetEnvironmentValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentValue
}

// GetEnvironmentValueOk returns a tuple with the EnvironmentValue field value
// and a boolean to check if the value has been set.
func (o *SecretEnvironment) GetEnvironmentValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentValue, true
}

// SetEnvironmentValue sets field value
func (o *SecretEnvironment) SetEnvironmentValue(v string) {
	o.EnvironmentValue = v
}

func (o SecretEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretEnvironment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment_value"] = o.EnvironmentValue
	return toSerialize, nil
}

func (o *SecretEnvironment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environment_value",
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

	varSecretEnvironment := _SecretEnvironment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSecretEnvironment)

	if err != nil {
		return err
	}

	*o = SecretEnvironment(varSecretEnvironment)

	return err
}

type NullableSecretEnvironment struct {
	value *SecretEnvironment
	isSet bool
}

func (v NullableSecretEnvironment) Get() *SecretEnvironment {
	return v.value
}

func (v *NullableSecretEnvironment) Set(val *SecretEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretEnvironment(val *SecretEnvironment) *NullableSecretEnvironment {
	return &NullableSecretEnvironment{value: val, isSet: true}
}

func (v NullableSecretEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


