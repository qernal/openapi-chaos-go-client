# SecretBodyPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SecretCreateType**](SecretCreateType.md) |  | 
**Payload** | [**SecretCreatePayload**](SecretCreatePayload.md) |  | 
**Encryption** | **string** | Encryption entity | 

## Methods

### NewSecretBodyPatch

`func NewSecretBodyPatch(type_ SecretCreateType, payload SecretCreatePayload, encryption string, ) *SecretBodyPatch`

NewSecretBodyPatch instantiates a new SecretBodyPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretBodyPatchWithDefaults

`func NewSecretBodyPatchWithDefaults() *SecretBodyPatch`

NewSecretBodyPatchWithDefaults instantiates a new SecretBodyPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SecretBodyPatch) GetType() SecretCreateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretBodyPatch) GetTypeOk() (*SecretCreateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretBodyPatch) SetType(v SecretCreateType)`

SetType sets Type field to given value.


### GetPayload

`func (o *SecretBodyPatch) GetPayload() SecretCreatePayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SecretBodyPatch) GetPayloadOk() (*SecretCreatePayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SecretBodyPatch) SetPayload(v SecretCreatePayload)`

SetPayload sets Payload field to given value.


### GetEncryption

`func (o *SecretBodyPatch) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *SecretBodyPatch) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *SecretBodyPatch) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


