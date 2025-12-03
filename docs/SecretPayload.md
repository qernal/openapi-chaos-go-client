# SecretPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | Public TLS certificate | 
**Registry** | **string** | Private registry domain/location, when using the private docker hub registry sepcify &#x60;docker.io&#x60; &gt; Without http scheme  | 
**Dek** | **string** | Base64 encoded Data Encryption Key (DEK) | 

## Methods

### NewSecretPayload

`func NewSecretPayload(certificate string, registry string, dek string, ) *SecretPayload`

NewSecretPayload instantiates a new SecretPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretPayloadWithDefaults

`func NewSecretPayloadWithDefaults() *SecretPayload`

NewSecretPayloadWithDefaults instantiates a new SecretPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *SecretPayload) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SecretPayload) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SecretPayload) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetRegistry

`func (o *SecretPayload) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *SecretPayload) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *SecretPayload) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetDek

`func (o *SecretPayload) GetDek() string`

GetDek returns the Dek field if non-nil, zero value otherwise.

### GetDekOk

`func (o *SecretPayload) GetDekOk() (*string, bool)`

GetDekOk returns a tuple with the Dek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDek

`func (o *SecretPayload) SetDek(v string)`

SetDek sets Dek field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


