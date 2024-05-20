# SecretRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registry** | **string** | Url to private container repository (for docker registry use docker.io) | 
**RegistryValue** | **string** | Token used for auth to the registry | 

## Methods

### NewSecretRegistry

`func NewSecretRegistry(registry string, registryValue string, ) *SecretRegistry`

NewSecretRegistry instantiates a new SecretRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretRegistryWithDefaults

`func NewSecretRegistryWithDefaults() *SecretRegistry`

NewSecretRegistryWithDefaults instantiates a new SecretRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistry

`func (o *SecretRegistry) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *SecretRegistry) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *SecretRegistry) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetRegistryValue

`func (o *SecretRegistry) GetRegistryValue() string`

GetRegistryValue returns the RegistryValue field if non-nil, zero value otherwise.

### GetRegistryValueOk

`func (o *SecretRegistry) GetRegistryValueOk() (*string, bool)`

GetRegistryValueOk returns a tuple with the RegistryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryValue

`func (o *SecretRegistry) SetRegistryValue(v string)`

SetRegistryValue sets RegistryValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


