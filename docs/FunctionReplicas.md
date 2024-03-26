# FunctionReplicas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Min** | **int32** | Minimum number of replicas to have | 
**Max** | **int32** | Maximum number of replicas to have | 
**Affinity** | [**FunctionReplicasAffinity**](FunctionReplicasAffinity.md) |  | 

## Methods

### NewFunctionReplicas

`func NewFunctionReplicas(min int32, max int32, affinity FunctionReplicasAffinity, ) *FunctionReplicas`

NewFunctionReplicas instantiates a new FunctionReplicas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionReplicasWithDefaults

`func NewFunctionReplicasWithDefaults() *FunctionReplicas`

NewFunctionReplicasWithDefaults instantiates a new FunctionReplicas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMin

`func (o *FunctionReplicas) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *FunctionReplicas) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *FunctionReplicas) SetMin(v int32)`

SetMin sets Min field to given value.


### GetMax

`func (o *FunctionReplicas) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *FunctionReplicas) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *FunctionReplicas) SetMax(v int32)`

SetMax sets Max field to given value.


### GetAffinity

`func (o *FunctionReplicas) GetAffinity() FunctionReplicasAffinity`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *FunctionReplicas) GetAffinityOk() (*FunctionReplicasAffinity, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *FunctionReplicas) SetAffinity(v FunctionReplicasAffinity)`

SetAffinity sets Affinity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


