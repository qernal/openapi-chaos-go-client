# SecretBodyRegistryPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registry** | **string** | Private registry domain/location, when using the private docker hub registry sepcify &#x60;docker.io&#x60; &gt; Without http scheme  | 
**RegistryValue** | **string** | Token used for auth to the registry | 

## Methods

### NewSecretBodyRegistryPayload

`func NewSecretBodyRegistryPayload(registry string, registryValue string, ) *SecretBodyRegistryPayload`

NewSecretBodyRegistryPayload instantiates a new SecretBodyRegistryPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretBodyRegistryPayloadWithDefaults

`func NewSecretBodyRegistryPayloadWithDefaults() *SecretBodyRegistryPayload`

NewSecretBodyRegistryPayloadWithDefaults instantiates a new SecretBodyRegistryPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistry

`func (o *SecretBodyRegistryPayload) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *SecretBodyRegistryPayload) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *SecretBodyRegistryPayload) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetRegistryValue

`func (o *SecretBodyRegistryPayload) GetRegistryValue() string`

GetRegistryValue returns the RegistryValue field if non-nil, zero value otherwise.

### GetRegistryValueOk

`func (o *SecretBodyRegistryPayload) GetRegistryValueOk() (*string, bool)`

GetRegistryValueOk returns a tuple with the RegistryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryValue

`func (o *SecretBodyRegistryPayload) SetRegistryValue(v string)`

SetRegistryValue sets RegistryValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


