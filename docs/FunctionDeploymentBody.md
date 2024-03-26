# FunctionDeploymentBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | [**Location**](Location.md) |  | 
**Replicas** | [**FunctionReplicas**](FunctionReplicas.md) |  | 

## Methods

### NewFunctionDeploymentBody

`func NewFunctionDeploymentBody(location Location, replicas FunctionReplicas, ) *FunctionDeploymentBody`

NewFunctionDeploymentBody instantiates a new FunctionDeploymentBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionDeploymentBodyWithDefaults

`func NewFunctionDeploymentBodyWithDefaults() *FunctionDeploymentBody`

NewFunctionDeploymentBodyWithDefaults instantiates a new FunctionDeploymentBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *FunctionDeploymentBody) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FunctionDeploymentBody) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FunctionDeploymentBody) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetReplicas

`func (o *FunctionDeploymentBody) GetReplicas() FunctionReplicas`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *FunctionDeploymentBody) GetReplicasOk() (*FunctionReplicas, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *FunctionDeploymentBody) SetReplicas(v FunctionReplicas)`

SetReplicas sets Replicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


