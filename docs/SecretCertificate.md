# SecretCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | Public certificate | 
**CertificateValue** | **string** | Encrypted certificate private key | 

## Methods

### NewSecretCertificate

`func NewSecretCertificate(certificate string, certificateValue string, ) *SecretCertificate`

NewSecretCertificate instantiates a new SecretCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretCertificateWithDefaults

`func NewSecretCertificateWithDefaults() *SecretCertificate`

NewSecretCertificateWithDefaults instantiates a new SecretCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *SecretCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SecretCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SecretCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetCertificateValue

`func (o *SecretCertificate) GetCertificateValue() string`

GetCertificateValue returns the CertificateValue field if non-nil, zero value otherwise.

### GetCertificateValueOk

`func (o *SecretCertificate) GetCertificateValueOk() (*string, bool)`

GetCertificateValueOk returns a tuple with the CertificateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValue

`func (o *SecretCertificate) SetCertificateValue(v string)`

SetCertificateValue sets CertificateValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


