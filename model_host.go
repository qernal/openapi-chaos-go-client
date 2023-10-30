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

// checks if the Host type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Host{}

// Host Host response
type Host struct {
	// Host id
	Id string `json:"id"`
	// Hostname, this can be a root or a subdomain
	Host string `json:"host"`
	// The secret reference to the certificate
	Certificate *string `json:"certificate,omitempty"`
	// Project ID this is attached to
	ProjectId string `json:"project_id"`
	// If the host is read only and cannot be removed, primarily used for *.qrnl.app domains
	ReadOnly bool `json:"read_only"`
	// If the host is disabled and so won't be routable
	Disabled bool `json:"disabled"`
	// TXT record of host to verify ownership - if this record is removed, it may become unverified as this is checked periodically to continually verify ownership
	TxtVerification string `json:"txt_verification"`
	// UTC datetime when the host was verified (ISO 8601 date format).
	VerifiedAt *string `json:"verified_at,omitempty"`
	Date Date `json:"date"`
	VerificationStatus HostVerificationStatus `json:"verification_status"`
}

// NewHost instantiates a new Host object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHost(id string, host string, projectId string, readOnly bool, disabled bool, txtVerification string, date Date, verificationStatus HostVerificationStatus) *Host {
	this := Host{}
	this.Id = id
	this.Host = host
	this.ProjectId = projectId
	this.ReadOnly = readOnly
	this.Disabled = disabled
	this.TxtVerification = txtVerification
	this.Date = date
	this.VerificationStatus = verificationStatus
	return &this
}

// NewHostWithDefaults instantiates a new Host object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostWithDefaults() *Host {
	this := Host{}
	return &this
}

// GetId returns the Id field value
func (o *Host) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Host) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Host) SetId(v string) {
	o.Id = v
}

// GetHost returns the Host field value
func (o *Host) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *Host) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *Host) SetHost(v string) {
	o.Host = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *Host) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *Host) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *Host) SetCertificate(v string) {
	o.Certificate = &v
}

// GetProjectId returns the ProjectId field value
func (o *Host) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Host) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Host) SetProjectId(v string) {
	o.ProjectId = v
}

// GetReadOnly returns the ReadOnly field value
func (o *Host) GetReadOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value
// and a boolean to check if the value has been set.
func (o *Host) GetReadOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadOnly, true
}

// SetReadOnly sets field value
func (o *Host) SetReadOnly(v bool) {
	o.ReadOnly = v
}

// GetDisabled returns the Disabled field value
func (o *Host) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *Host) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *Host) SetDisabled(v bool) {
	o.Disabled = v
}

// GetTxtVerification returns the TxtVerification field value
func (o *Host) GetTxtVerification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxtVerification
}

// GetTxtVerificationOk returns a tuple with the TxtVerification field value
// and a boolean to check if the value has been set.
func (o *Host) GetTxtVerificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxtVerification, true
}

// SetTxtVerification sets field value
func (o *Host) SetTxtVerification(v string) {
	o.TxtVerification = v
}

// GetVerifiedAt returns the VerifiedAt field value if set, zero value otherwise.
func (o *Host) GetVerifiedAt() string {
	if o == nil || IsNil(o.VerifiedAt) {
		var ret string
		return ret
	}
	return *o.VerifiedAt
}

// GetVerifiedAtOk returns a tuple with the VerifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Host) GetVerifiedAtOk() (*string, bool) {
	if o == nil || IsNil(o.VerifiedAt) {
		return nil, false
	}
	return o.VerifiedAt, true
}

// HasVerifiedAt returns a boolean if a field has been set.
func (o *Host) HasVerifiedAt() bool {
	if o != nil && !IsNil(o.VerifiedAt) {
		return true
	}

	return false
}

// SetVerifiedAt gets a reference to the given string and assigns it to the VerifiedAt field.
func (o *Host) SetVerifiedAt(v string) {
	o.VerifiedAt = &v
}

// GetDate returns the Date field value
func (o *Host) GetDate() Date {
	if o == nil {
		var ret Date
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Host) GetDateOk() (*Date, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Host) SetDate(v Date) {
	o.Date = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *Host) GetVerificationStatus() HostVerificationStatus {
	if o == nil {
		var ret HostVerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *Host) GetVerificationStatusOk() (*HostVerificationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *Host) SetVerificationStatus(v HostVerificationStatus) {
	o.VerificationStatus = v
}

func (o Host) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Host) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["host"] = o.Host
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	toSerialize["project_id"] = o.ProjectId
	toSerialize["read_only"] = o.ReadOnly
	toSerialize["disabled"] = o.Disabled
	toSerialize["txt_verification"] = o.TxtVerification
	if !IsNil(o.VerifiedAt) {
		toSerialize["verified_at"] = o.VerifiedAt
	}
	toSerialize["date"] = o.Date
	toSerialize["verification_status"] = o.VerificationStatus
	return toSerialize, nil
}

type NullableHost struct {
	value *Host
	isSet bool
}

func (v NullableHost) Get() *Host {
	return v.value
}

func (v *NullableHost) Set(val *Host) {
	v.value = val
	v.isSet = true
}

func (v NullableHost) IsSet() bool {
	return v.isSet
}

func (v *NullableHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHost(val *Host) *NullableHost {
	return &NullableHost{value: val, isSet: true}
}

func (v NullableHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


