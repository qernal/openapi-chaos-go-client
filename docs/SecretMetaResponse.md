# SecretMetaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Project name | 
**Type** | [**SecretMetaType**](SecretMetaType.md) |  | 
**Payload** | Pointer to [**SecretMetaResponsePayload**](SecretMetaResponsePayload.md) |  | [optional] 
**Revision** | **int32** | Secret revision | 
**Date** | [**Date**](Date.md) |  | 

## Methods

### NewSecretMetaResponse

`func NewSecretMetaResponse(name string, type_ SecretMetaType, revision int32, date Date, ) *SecretMetaResponse`

NewSecretMetaResponse instantiates a new SecretMetaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretMetaResponseWithDefaults

`func NewSecretMetaResponseWithDefaults() *SecretMetaResponse`

NewSecretMetaResponseWithDefaults instantiates a new SecretMetaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecretMetaResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretMetaResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretMetaResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SecretMetaResponse) GetType() SecretMetaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretMetaResponse) GetTypeOk() (*SecretMetaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretMetaResponse) SetType(v SecretMetaType)`

SetType sets Type field to given value.


### GetPayload

`func (o *SecretMetaResponse) GetPayload() SecretMetaResponsePayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SecretMetaResponse) GetPayloadOk() (*SecretMetaResponsePayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SecretMetaResponse) SetPayload(v SecretMetaResponsePayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SecretMetaResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRevision

`func (o *SecretMetaResponse) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SecretMetaResponse) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SecretMetaResponse) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetDate

`func (o *SecretMetaResponse) GetDate() Date`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SecretMetaResponse) GetDateOk() (*Date, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SecretMetaResponse) SetDate(v Date)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


