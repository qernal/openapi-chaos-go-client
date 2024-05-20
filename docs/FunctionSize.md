# FunctionSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **int32** | CPU in 0.1 vCPU increments, for a whole vCPU specify 1024 Must be in multiples of 128, with the same multiplier as memory from the base  | 
**Memory** | **int32** | Memory in 128 MB increments, values are integer always in MB Must be in multiples of 128, with the same multiplier as CPU from the base  | 

## Methods

### NewFunctionSize

`func NewFunctionSize(cpu int32, memory int32, ) *FunctionSize`

NewFunctionSize instantiates a new FunctionSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionSizeWithDefaults

`func NewFunctionSizeWithDefaults() *FunctionSize`

NewFunctionSizeWithDefaults instantiates a new FunctionSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *FunctionSize) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *FunctionSize) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *FunctionSize) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *FunctionSize) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *FunctionSize) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *FunctionSize) SetMemory(v int32)`

SetMemory sets Memory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


