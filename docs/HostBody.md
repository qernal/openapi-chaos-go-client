# HostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname | 
**Certificate** | **string** | Reference to secret certificate path | 
**Disabled** | **bool** | If the host is disabled, then this host won&#39;t be accessible and so the deployments will not be routable on this host | 

## Methods

### NewHostBody

`func NewHostBody(host string, certificate string, disabled bool, ) *HostBody`

NewHostBody instantiates a new HostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostBodyWithDefaults

`func NewHostBodyWithDefaults() *HostBody`

NewHostBodyWithDefaults instantiates a new HostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HostBody) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HostBody) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HostBody) SetHost(v string)`

SetHost sets Host field to given value.


### GetCertificate

`func (o *HostBody) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *HostBody) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *HostBody) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetDisabled

`func (o *HostBody) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *HostBody) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *HostBody) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


