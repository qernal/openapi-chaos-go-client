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

// checks if the OrganisationBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganisationBody{}

// OrganisationBody Organisation body
type OrganisationBody struct {
	// Organisation name
	Name string `json:"name"`
}

// NewOrganisationBody instantiates a new OrganisationBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganisationBody(name string) *OrganisationBody {
	this := OrganisationBody{}
	this.Name = name
	return &this
}

// NewOrganisationBodyWithDefaults instantiates a new OrganisationBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganisationBodyWithDefaults() *OrganisationBody {
	this := OrganisationBody{}
	return &this
}

// GetName returns the Name field value
func (o *OrganisationBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganisationBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganisationBody) SetName(v string) {
	o.Name = v
}

func (o OrganisationBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganisationBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableOrganisationBody struct {
	value *OrganisationBody
	isSet bool
}

func (v NullableOrganisationBody) Get() *OrganisationBody {
	return v.value
}

func (v *NullableOrganisationBody) Set(val *OrganisationBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisationBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisationBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisationBody(val *OrganisationBody) *NullableOrganisationBody {
	return &NullableOrganisationBody{value: val, isSet: true}
}

func (v NullableOrganisationBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisationBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


