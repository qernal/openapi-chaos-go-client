# FunctionRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Can be a regular expression | 
**Methods** | **[]string** | HTTP Verb(s) for this function | 
**Weight** | **int32** | The route weight for consideration | 

## Methods

### NewFunctionRoute

`func NewFunctionRoute(path string, methods []string, weight int32, ) *FunctionRoute`

NewFunctionRoute instantiates a new FunctionRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionRouteWithDefaults

`func NewFunctionRouteWithDefaults() *FunctionRoute`

NewFunctionRouteWithDefaults instantiates a new FunctionRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FunctionRoute) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FunctionRoute) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FunctionRoute) SetPath(v string)`

SetPath sets Path field to given value.


### GetMethods

`func (o *FunctionRoute) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *FunctionRoute) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *FunctionRoute) SetMethods(v []string)`

SetMethods sets Methods field to given value.


### GetWeight

`func (o *FunctionRoute) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *FunctionRoute) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *FunctionRoute) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


