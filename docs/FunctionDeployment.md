# FunctionDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the deployment | [optional] 
**Location** | [**Location**](Location.md) |  | 
**Replicas** | [**FunctionReplicas**](FunctionReplicas.md) |  | 

## Methods

### NewFunctionDeployment

`func NewFunctionDeployment(location Location, replicas FunctionReplicas, ) *FunctionDeployment`

NewFunctionDeployment instantiates a new FunctionDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionDeploymentWithDefaults

`func NewFunctionDeploymentWithDefaults() *FunctionDeployment`

NewFunctionDeploymentWithDefaults instantiates a new FunctionDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FunctionDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FunctionDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FunctionDeployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FunctionDeployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *FunctionDeployment) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FunctionDeployment) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FunctionDeployment) SetLocation(v Location)`

SetLocation sets Location field to given value.


### GetReplicas

`func (o *FunctionDeployment) GetReplicas() FunctionReplicas`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *FunctionDeployment) GetReplicasOk() (*FunctionReplicas, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *FunctionDeployment) SetReplicas(v FunctionReplicas)`

SetReplicas sets Replicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


