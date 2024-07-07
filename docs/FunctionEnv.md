# FunctionEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Key name | 
**Reference** | **string** | Reference for the secret, it&#39;s split up into 4 main parts; - [1] \&quot;projects:\&quot; - This references which resource type the secret is within - [2] \&quot;0a6b9ff3-6807-4820-b94b-5e1d7efcdd93\&quot; - The project UUID the secret is within - [3] \&quot;MY_SECRET\&quot; - The name of the secret - [4] \&quot;0\&quot; - The revision of the secret to use  | 

## Methods

### NewFunctionEnv

`func NewFunctionEnv(name string, reference string, ) *FunctionEnv`

NewFunctionEnv instantiates a new FunctionEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionEnvWithDefaults

`func NewFunctionEnvWithDefaults() *FunctionEnv`

NewFunctionEnvWithDefaults instantiates a new FunctionEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionEnv) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionEnv) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionEnv) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *FunctionEnv) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *FunctionEnv) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *FunctionEnv) SetReference(v string)`

SetReference sets Reference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


