# SecretBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Secret name | 
**Type** | [**SecretCreateType**](SecretCreateType.md) |  | 
**Payload** | [**SecretCreatePayload**](SecretCreatePayload.md) |  | 
**Encryption** | **string** | Encryption entity | 

## Methods

### NewSecretBody

`func NewSecretBody(name string, type_ SecretCreateType, payload SecretCreatePayload, encryption string, ) *SecretBody`

NewSecretBody instantiates a new SecretBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretBodyWithDefaults

`func NewSecretBodyWithDefaults() *SecretBody`

NewSecretBodyWithDefaults instantiates a new SecretBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecretBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretBody) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SecretBody) GetType() SecretCreateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretBody) GetTypeOk() (*SecretCreateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretBody) SetType(v SecretCreateType)`

SetType sets Type field to given value.


### GetPayload

`func (o *SecretBody) GetPayload() SecretCreatePayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SecretBody) GetPayloadOk() (*SecretCreatePayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SecretBody) SetPayload(v SecretCreatePayload)`

SetPayload sets Payload field to given value.


### GetEncryption

`func (o *SecretBody) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *SecretBody) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *SecretBody) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


