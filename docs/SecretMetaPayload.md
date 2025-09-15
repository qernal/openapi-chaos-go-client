# SecretMetaPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | Public SSL certificate | 
**Registry** | **string** | Private registry domain/location, when using the private docker hub registry sepcify &#x60;docker.io&#x60; &gt; Without http scheme  | 
**Public** | **string** | Base64 encoded DEK public key | 

## Methods

### NewSecretMetaPayload

`func NewSecretMetaPayload(certificate string, registry string, public string, ) *SecretMetaPayload`

NewSecretMetaPayload instantiates a new SecretMetaPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretMetaPayloadWithDefaults

`func NewSecretMetaPayloadWithDefaults() *SecretMetaPayload`

NewSecretMetaPayloadWithDefaults instantiates a new SecretMetaPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *SecretMetaPayload) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SecretMetaPayload) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SecretMetaPayload) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetRegistry

`func (o *SecretMetaPayload) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *SecretMetaPayload) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *SecretMetaPayload) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetPublic

`func (o *SecretMetaPayload) GetPublic() string`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *SecretMetaPayload) GetPublicOk() (*string, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *SecretMetaPayload) SetPublic(v string)`

SetPublic sets Public field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


