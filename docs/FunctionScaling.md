# FunctionScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | CPU or Memory supported | 
**Low** | **int32** | For type to drop below before scale down | 
**High** | **int32** | For type to go above before scale up | 

## Methods

### NewFunctionScaling

`func NewFunctionScaling(type_ string, low int32, high int32, ) *FunctionScaling`

NewFunctionScaling instantiates a new FunctionScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionScalingWithDefaults

`func NewFunctionScalingWithDefaults() *FunctionScaling`

NewFunctionScalingWithDefaults instantiates a new FunctionScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FunctionScaling) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionScaling) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionScaling) SetType(v string)`

SetType sets Type field to given value.


### GetLow

`func (o *FunctionScaling) GetLow() int32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *FunctionScaling) GetLowOk() (*int32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *FunctionScaling) SetLow(v int32)`

SetLow sets Low field to given value.


### GetHigh

`func (o *FunctionScaling) GetHigh() int32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *FunctionScaling) GetHighOk() (*int32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *FunctionScaling) SetHigh(v int32)`

SetHigh sets High field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


