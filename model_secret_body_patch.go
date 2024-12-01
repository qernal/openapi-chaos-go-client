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

// checks if the SecretBodyPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretBodyPatch{}

// SecretBodyPatch Secret body patch fields
type SecretBodyPatch struct {
	Type SecretCreateType `json:"type"`
	Payload SecretCreatePayload `json:"payload"`
	// Encryption entity
	Encryption string `json:"encryption" validate:"regexp=^keys\\/dek\\/[0-9]+$"`
}

type _SecretBodyPatch SecretBodyPatch

// NewSecretBodyPatch instantiates a new SecretBodyPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretBodyPatch(type_ SecretCreateType, payload SecretCreatePayload, encryption string) *SecretBodyPatch {
	this := SecretBodyPatch{}
	this.Type = type_
	this.Payload = payload
	this.Encryption = encryption
	return &this
}

// NewSecretBodyPatchWithDefaults instantiates a new SecretBodyPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretBodyPatchWithDefaults() *SecretBodyPatch {
	this := SecretBodyPatch{}
	return &this
}

// GetType returns the Type field value
func (o *SecretBodyPatch) GetType() SecretCreateType {
	if o == nil {
		var ret SecretCreateType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecretBodyPatch) GetTypeOk() (*SecretCreateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecretBodyPatch) SetType(v SecretCreateType) {
	o.Type = v
}

// GetPayload returns the Payload field value
func (o *SecretBodyPatch) GetPayload() SecretCreatePayload {
	if o == nil {
		var ret SecretCreatePayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *SecretBodyPatch) GetPayloadOk() (*SecretCreatePayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *SecretBodyPatch) SetPayload(v SecretCreatePayload) {
	o.Payload = v
}

// GetEncryption returns the Encryption field value
func (o *SecretBodyPatch) GetEncryption() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value
// and a boolean to check if the value has been set.
func (o *SecretBodyPatch) GetEncryptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encryption, true
}

// SetEncryption sets field value
func (o *SecretBodyPatch) SetEncryption(v string) {
	o.Encryption = v
}

func (o SecretBodyPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretBodyPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["payload"] = o.Payload
	toSerialize["encryption"] = o.Encryption
	return toSerialize, nil
}

func (o *SecretBodyPatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"payload",
		"encryption",
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

	varSecretBodyPatch := _SecretBodyPatch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSecretBodyPatch)

	if err != nil {
		return err
	}

	*o = SecretBodyPatch(varSecretBodyPatch)

	return err
}

type NullableSecretBodyPatch struct {
	value *SecretBodyPatch
	isSet bool
}

func (v NullableSecretBodyPatch) Get() *SecretBodyPatch {
	return v.value
}

func (v *NullableSecretBodyPatch) Set(val *SecretBodyPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretBodyPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretBodyPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretBodyPatch(val *SecretBodyPatch) *NullableSecretBodyPatch {
	return &NullableSecretBodyPatch{value: val, isSet: true}
}

func (v NullableSecretBodyPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretBodyPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


