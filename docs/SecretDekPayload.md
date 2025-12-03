# SecretDekPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dek** | **string** | Base64 encoded Data Encryption Key (DEK) | 

## Methods

### NewSecretDekPayload

`func NewSecretDekPayload(dek string, ) *SecretDekPayload`

NewSecretDekPayload instantiates a new SecretDekPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretDekPayloadWithDefaults

`func NewSecretDekPayloadWithDefaults() *SecretDekPayload`

NewSecretDekPayloadWithDefaults instantiates a new SecretDekPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDek

`func (o *SecretDekPayload) GetDek() string`

GetDek returns the Dek field if non-nil, zero value otherwise.

### GetDekOk

`func (o *SecretDekPayload) GetDekOk() (*string, bool)`

GetDekOk returns a tuple with the Dek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDek

`func (o *SecretDekPayload) SetDek(v string)`

SetDek sets Dek field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


