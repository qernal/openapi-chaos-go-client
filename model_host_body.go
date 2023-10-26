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

// checks if the HostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostBody{}

// HostBody Host body
type HostBody struct {
	// Hostname
	Host string `json:"host"`
	// Reference to secret certificate path
	Certificate string `json:"certificate"`
	// If the host is disabled, then this host won't be accessible and so the deployments will not be routable on this host
	Disabled bool `json:"disabled"`
}

// NewHostBody instantiates a new HostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostBody(host string, certificate string, disabled bool) *HostBody {
	this := HostBody{}
	this.Host = host
	this.Certificate = certificate
	this.Disabled = disabled
	return &this
}

// NewHostBodyWithDefaults instantiates a new HostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostBodyWithDefaults() *HostBody {
	this := HostBody{}
	return &this
}

// GetHost returns the Host field value
func (o *HostBody) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *HostBody) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *HostBody) SetHost(v string) {
	o.Host = v
}

// GetCertificate returns the Certificate field value
func (o *HostBody) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *HostBody) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *HostBody) SetCertificate(v string) {
	o.Certificate = v
}

// GetDisabled returns the Disabled field value
func (o *HostBody) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *HostBody) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *HostBody) SetDisabled(v bool) {
	o.Disabled = v
}

func (o HostBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["certificate"] = o.Certificate
	toSerialize["disabled"] = o.Disabled
	return toSerialize, nil
}

type NullableHostBody struct {
	value *HostBody
	isSet bool
}

func (v NullableHostBody) Get() *HostBody {
	return v.value
}

func (v *NullableHostBody) Set(val *HostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableHostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableHostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostBody(val *HostBody) *NullableHostBody {
	return &NullableHostBody{value: val, isSet: true}
}

func (v NullableHostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

