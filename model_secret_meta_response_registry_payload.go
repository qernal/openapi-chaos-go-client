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

// checks if the SecretMetaResponseRegistryPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretMetaResponseRegistryPayload{}

// SecretMetaResponseRegistryPayload Secret metadata registry payload
type SecretMetaResponseRegistryPayload struct {
	// Registry domain
	Registry string `json:"registry"`
}

// NewSecretMetaResponseRegistryPayload instantiates a new SecretMetaResponseRegistryPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretMetaResponseRegistryPayload(registry string) *SecretMetaResponseRegistryPayload {
	this := SecretMetaResponseRegistryPayload{}
	this.Registry = registry
	return &this
}

// NewSecretMetaResponseRegistryPayloadWithDefaults instantiates a new SecretMetaResponseRegistryPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretMetaResponseRegistryPayloadWithDefaults() *SecretMetaResponseRegistryPayload {
	this := SecretMetaResponseRegistryPayload{}
	return &this
}

// GetRegistry returns the Registry field value
func (o *SecretMetaResponseRegistryPayload) GetRegistry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value
// and a boolean to check if the value has been set.
func (o *SecretMetaResponseRegistryPayload) GetRegistryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Registry, true
}

// SetRegistry sets field value
func (o *SecretMetaResponseRegistryPayload) SetRegistry(v string) {
	o.Registry = v
}

func (o SecretMetaResponseRegistryPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretMetaResponseRegistryPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["registry"] = o.Registry
	return toSerialize, nil
}

type NullableSecretMetaResponseRegistryPayload struct {
	value *SecretMetaResponseRegistryPayload
	isSet bool
}

func (v NullableSecretMetaResponseRegistryPayload) Get() *SecretMetaResponseRegistryPayload {
	return v.value
}

func (v *NullableSecretMetaResponseRegistryPayload) Set(val *SecretMetaResponseRegistryPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretMetaResponseRegistryPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretMetaResponseRegistryPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretMetaResponseRegistryPayload(val *SecretMetaResponseRegistryPayload) *NullableSecretMetaResponseRegistryPayload {
	return &NullableSecretMetaResponseRegistryPayload{value: val, isSet: true}
}

func (v NullableSecretMetaResponseRegistryPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretMetaResponseRegistryPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


