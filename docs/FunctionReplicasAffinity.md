# FunctionReplicasAffinity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **bool** | If there are &gt; 1 replica, make sure they&#39;re on different clusters | 
**Cloud** | **bool** | If there are &gt; 1 replica, make sure they&#39;re on different clouds | 

## Methods

### NewFunctionReplicasAffinity

`func NewFunctionReplicasAffinity(cluster bool, cloud bool, ) *FunctionReplicasAffinity`

NewFunctionReplicasAffinity instantiates a new FunctionReplicasAffinity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionReplicasAffinityWithDefaults

`func NewFunctionReplicasAffinityWithDefaults() *FunctionReplicasAffinity`

NewFunctionReplicasAffinityWithDefaults instantiates a new FunctionReplicasAffinity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *FunctionReplicasAffinity) GetCluster() bool`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *FunctionReplicasAffinity) GetClusterOk() (*bool, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *FunctionReplicasAffinity) SetCluster(v bool)`

SetCluster sets Cluster field to given value.


### GetCloud

`func (o *FunctionReplicasAffinity) GetCloud() bool`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *FunctionReplicasAffinity) GetCloudOk() (*bool, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *FunctionReplicasAffinity) SetCloud(v bool)`

SetCloud sets Cloud field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


