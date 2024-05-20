# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Host id | 
**Host** | **string** | Hostname, this can be the root of a domain or a subdomain | 
**Certificate** | Pointer to **string** | The secret reference to the certificate | [optional] 
**ProjectId** | **string** | Project ID this is attached to | 
**ReadOnly** | **bool** | If the host is read only and cannot be removed, primarily used for *.qrnl.app domains | 
**Disabled** | **bool** | If the host is disabled, then this host won&#39;t be accessible and so the deployments will not be routable | 
**TxtVerification** | **string** | TXT record of host to verify ownership - if this record is removed, it may become unverified as this is checked periodically to continually verify ownership | 
**VerifiedAt** | Pointer to **string** | UTC datetime when the host was verified (ISO 8601 date format). | [optional] 
**Date** | [**Date**](Date.md) |  | 
**VerificationStatus** | [**HostVerificationStatus**](HostVerificationStatus.md) |  | 

## Methods

### NewHost

`func NewHost(id string, host string, projectId string, readOnly bool, disabled bool, txtVerification string, date Date, verificationStatus HostVerificationStatus, ) *Host`

NewHost instantiates a new Host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostWithDefaults

`func NewHostWithDefaults() *Host`

NewHostWithDefaults instantiates a new Host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Host) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Host) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Host) SetId(v string)`

SetId sets Id field to given value.


### GetHost

`func (o *Host) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Host) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Host) SetHost(v string)`

SetHost sets Host field to given value.


### GetCertificate

`func (o *Host) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Host) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Host) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *Host) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetProjectId

`func (o *Host) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Host) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Host) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetReadOnly

`func (o *Host) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *Host) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *Host) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.


### GetDisabled

`func (o *Host) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Host) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Host) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetTxtVerification

`func (o *Host) GetTxtVerification() string`

GetTxtVerification returns the TxtVerification field if non-nil, zero value otherwise.

### GetTxtVerificationOk

`func (o *Host) GetTxtVerificationOk() (*string, bool)`

GetTxtVerificationOk returns a tuple with the TxtVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtVerification

`func (o *Host) SetTxtVerification(v string)`

SetTxtVerification sets TxtVerification field to given value.


### GetVerifiedAt

`func (o *Host) GetVerifiedAt() string`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *Host) GetVerifiedAtOk() (*string, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *Host) SetVerifiedAt(v string)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *Host) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### GetDate

`func (o *Host) GetDate() Date`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Host) GetDateOk() (*Date, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Host) SetDate(v Date)`

SetDate sets Date field to given value.


### GetVerificationStatus

`func (o *Host) GetVerificationStatus() HostVerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *Host) GetVerificationStatusOk() (*HostVerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *Host) SetVerificationStatus(v HostVerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


