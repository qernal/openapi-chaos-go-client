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

// checks if the ProjectBodyPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectBodyPatch{}

// ProjectBodyPatch Project patch fields
type ProjectBodyPatch struct {
	// Organisation id
	OrgId *string `json:"org_id,omitempty"`
	// Project name
	Name *string `json:"name,omitempty" validate:"regexp=^[A-z-]+$"`
}

// NewProjectBodyPatch instantiates a new ProjectBodyPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectBodyPatch() *ProjectBodyPatch {
	this := ProjectBodyPatch{}
	return &this
}

// NewProjectBodyPatchWithDefaults instantiates a new ProjectBodyPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectBodyPatchWithDefaults() *ProjectBodyPatch {
	this := ProjectBodyPatch{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ProjectBodyPatch) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectBodyPatch) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ProjectBodyPatch) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ProjectBodyPatch) SetOrgId(v string) {
	o.OrgId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectBodyPatch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectBodyPatch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectBodyPatch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectBodyPatch) SetName(v string) {
	o.Name = &v
}

func (o ProjectBodyPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectBodyPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableProjectBodyPatch struct {
	value *ProjectBodyPatch
	isSet bool
}

func (v NullableProjectBodyPatch) Get() *ProjectBodyPatch {
	return v.value
}

func (v *NullableProjectBodyPatch) Set(val *ProjectBodyPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectBodyPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectBodyPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectBodyPatch(val *ProjectBodyPatch) *NullableProjectBodyPatch {
	return &NullableProjectBodyPatch{value: val, isSet: true}
}

func (v NullableProjectBodyPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectBodyPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


