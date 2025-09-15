# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Secret name | 
**Type** | [**SecretCreateType**](SecretCreateType.md) |  | 
**Payload** | Pointer to [**SecretPayload**](SecretPayload.md) |  | [optional] 
**Revision** | **int32** | Secret revision | 
**Date** | [**Date**](Date.md) |  | 

## Methods

### NewSecret

`func NewSecret(name string, type_ SecretCreateType, revision int32, date Date, ) *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Secret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Secret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Secret) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Secret) GetType() SecretCreateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Secret) GetTypeOk() (*SecretCreateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Secret) SetType(v SecretCreateType)`

SetType sets Type field to given value.


### GetPayload

`func (o *Secret) GetPayload() SecretPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Secret) GetPayloadOk() (*SecretPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Secret) SetPayload(v SecretPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Secret) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRevision

`func (o *Secret) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Secret) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Secret) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetDate

`func (o *Secret) GetDate() Date`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Secret) GetDateOk() (*Date, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Secret) SetDate(v Date)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


