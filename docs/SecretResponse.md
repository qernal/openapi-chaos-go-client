# SecretResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Secret name | 
**Type** | [**SecretCreateType**](SecretCreateType.md) |  | 
**Payload** | Pointer to [**SecretResponsePayload**](SecretResponsePayload.md) |  | [optional] 
**Revision** | **int32** | Secret revision | 
**Date** | [**Date**](Date.md) |  | 

## Methods

### NewSecretResponse

`func NewSecretResponse(name string, type_ SecretCreateType, revision int32, date Date, ) *SecretResponse`

NewSecretResponse instantiates a new SecretResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretResponseWithDefaults

`func NewSecretResponseWithDefaults() *SecretResponse`

NewSecretResponseWithDefaults instantiates a new SecretResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecretResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SecretResponse) GetType() SecretCreateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretResponse) GetTypeOk() (*SecretCreateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretResponse) SetType(v SecretCreateType)`

SetType sets Type field to given value.


### GetPayload

`func (o *SecretResponse) GetPayload() SecretResponsePayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SecretResponse) GetPayloadOk() (*SecretResponsePayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SecretResponse) SetPayload(v SecretResponsePayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SecretResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRevision

`func (o *SecretResponse) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SecretResponse) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SecretResponse) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetDate

`func (o *SecretResponse) GetDate() Date`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SecretResponse) GetDateOk() (*Date, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SecretResponse) SetDate(v Date)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


