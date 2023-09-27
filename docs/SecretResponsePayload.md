# SecretResponsePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | Public SSL certificate | 
**Registry** | **string** | Registry domain | 

## Methods

### NewSecretResponsePayload

`func NewSecretResponsePayload(certificate string, registry string, ) *SecretResponsePayload`

NewSecretResponsePayload instantiates a new SecretResponsePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretResponsePayloadWithDefaults

`func NewSecretResponsePayloadWithDefaults() *SecretResponsePayload`

NewSecretResponsePayloadWithDefaults instantiates a new SecretResponsePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *SecretResponsePayload) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SecretResponsePayload) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SecretResponsePayload) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetRegistry

`func (o *SecretResponsePayload) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *SecretResponsePayload) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *SecretResponsePayload) SetRegistry(v string)`

SetRegistry sets Registry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


