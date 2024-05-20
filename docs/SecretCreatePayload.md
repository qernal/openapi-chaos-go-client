# SecretCreatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registry** | **string** | Url to private container repository (for docker registry use docker.io) | 
**RegistryValue** | **string** | Token used for auth to the registry | 
**EnvironmentValue** | **string** | Encrypted environment pairs (key - env key, value - env value) | 
**Certificate** | **string** | Public certificate | 
**CertificateValue** | **string** | Encrypted certificate private key | 

## Methods

### NewSecretCreatePayload

`func NewSecretCreatePayload(registry string, registryValue string, environmentValue string, certificate string, certificateValue string, ) *SecretCreatePayload`

NewSecretCreatePayload instantiates a new SecretCreatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretCreatePayloadWithDefaults

`func NewSecretCreatePayloadWithDefaults() *SecretCreatePayload`

NewSecretCreatePayloadWithDefaults instantiates a new SecretCreatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistry

`func (o *SecretCreatePayload) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *SecretCreatePayload) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *SecretCreatePayload) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetRegistryValue

`func (o *SecretCreatePayload) GetRegistryValue() string`

GetRegistryValue returns the RegistryValue field if non-nil, zero value otherwise.

### GetRegistryValueOk

`func (o *SecretCreatePayload) GetRegistryValueOk() (*string, bool)`

GetRegistryValueOk returns a tuple with the RegistryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryValue

`func (o *SecretCreatePayload) SetRegistryValue(v string)`

SetRegistryValue sets RegistryValue field to given value.


### GetEnvironmentValue

`func (o *SecretCreatePayload) GetEnvironmentValue() string`

GetEnvironmentValue returns the EnvironmentValue field if non-nil, zero value otherwise.

### GetEnvironmentValueOk

`func (o *SecretCreatePayload) GetEnvironmentValueOk() (*string, bool)`

GetEnvironmentValueOk returns a tuple with the EnvironmentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentValue

`func (o *SecretCreatePayload) SetEnvironmentValue(v string)`

SetEnvironmentValue sets EnvironmentValue field to given value.


### GetCertificate

`func (o *SecretCreatePayload) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SecretCreatePayload) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SecretCreatePayload) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetCertificateValue

`func (o *SecretCreatePayload) GetCertificateValue() string`

GetCertificateValue returns the CertificateValue field if non-nil, zero value otherwise.

### GetCertificateValueOk

`func (o *SecretCreatePayload) GetCertificateValueOk() (*string, bool)`

GetCertificateValueOk returns a tuple with the CertificateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValue

`func (o *SecretCreatePayload) SetCertificateValue(v string)`

SetCertificateValue sets CertificateValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


