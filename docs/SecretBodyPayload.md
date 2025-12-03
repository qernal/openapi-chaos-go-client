# SecretBodyPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registry** | **string** | Private registry domain/location, when using the private docker hub registry sepcify &#x60;docker.io&#x60; &gt; Without http scheme  | 
**RegistryValue** | **string** | Token used for auth to the registry | 
**EnvironmentValue** | **string** | Encrypted environment value | 
**Certificate** | **string** | Public TLS certificate | 
**CertificateValue** | **string** | Encrypted certificate private key | 

## Methods

### NewSecretBodyPayload

`func NewSecretBodyPayload(registry string, registryValue string, environmentValue string, certificate string, certificateValue string, ) *SecretBodyPayload`

NewSecretBodyPayload instantiates a new SecretBodyPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretBodyPayloadWithDefaults

`func NewSecretBodyPayloadWithDefaults() *SecretBodyPayload`

NewSecretBodyPayloadWithDefaults instantiates a new SecretBodyPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistry

`func (o *SecretBodyPayload) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *SecretBodyPayload) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *SecretBodyPayload) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetRegistryValue

`func (o *SecretBodyPayload) GetRegistryValue() string`

GetRegistryValue returns the RegistryValue field if non-nil, zero value otherwise.

### GetRegistryValueOk

`func (o *SecretBodyPayload) GetRegistryValueOk() (*string, bool)`

GetRegistryValueOk returns a tuple with the RegistryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryValue

`func (o *SecretBodyPayload) SetRegistryValue(v string)`

SetRegistryValue sets RegistryValue field to given value.


### GetEnvironmentValue

`func (o *SecretBodyPayload) GetEnvironmentValue() string`

GetEnvironmentValue returns the EnvironmentValue field if non-nil, zero value otherwise.

### GetEnvironmentValueOk

`func (o *SecretBodyPayload) GetEnvironmentValueOk() (*string, bool)`

GetEnvironmentValueOk returns a tuple with the EnvironmentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentValue

`func (o *SecretBodyPayload) SetEnvironmentValue(v string)`

SetEnvironmentValue sets EnvironmentValue field to given value.


### GetCertificate

`func (o *SecretBodyPayload) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SecretBodyPayload) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SecretBodyPayload) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetCertificateValue

`func (o *SecretBodyPayload) GetCertificateValue() string`

GetCertificateValue returns the CertificateValue field if non-nil, zero value otherwise.

### GetCertificateValueOk

`func (o *SecretBodyPayload) GetCertificateValueOk() (*string, bool)`

GetCertificateValueOk returns a tuple with the CertificateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValue

`func (o *SecretBodyPayload) SetCertificateValue(v string)`

SetCertificateValue sets CertificateValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


