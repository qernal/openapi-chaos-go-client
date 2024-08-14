# FunctionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | ID of the project this function belongs to | 
**Version** | **string** | Function spec version | 
**Name** | **string** | Name of the function | 
**Description** | **string** | Description of what the function does | 
**Image** | **string** | Path to container image | 
**Type** | [**FunctionType**](FunctionType.md) |  | 
**Size** | [**FunctionSize**](FunctionSize.md) |  | 
**Port** | **int32** | Port the application runs on | 
**Routes** | Pointer to [**[]FunctionRoute**](FunctionRoute.md) | The public route/path to this function, only applicable to http type functions | [optional] 
**Scaling** | [**FunctionScaling**](FunctionScaling.md) |  | 
**Deployments** | [**[]FunctionDeploymentBody**](FunctionDeploymentBody.md) | List of deployments for this function | 
**Secrets** | [**[]FunctionEnv**](FunctionEnv.md) | List of environment variables for secrets | 
**Compliance** | [**[]FunctionCompliance**](FunctionCompliance.md) | Tags to limit deployment | 

## Methods

### NewFunctionBody

`func NewFunctionBody(projectId string, version string, name string, description string, image string, type_ FunctionType, size FunctionSize, port int32, scaling FunctionScaling, deployments []FunctionDeploymentBody, secrets []FunctionEnv, compliance []FunctionCompliance, ) *FunctionBody`

NewFunctionBody instantiates a new FunctionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionBodyWithDefaults

`func NewFunctionBodyWithDefaults() *FunctionBody`

NewFunctionBodyWithDefaults instantiates a new FunctionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *FunctionBody) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FunctionBody) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FunctionBody) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersion

`func (o *FunctionBody) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FunctionBody) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FunctionBody) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetName

`func (o *FunctionBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionBody) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FunctionBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FunctionBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImage

`func (o *FunctionBody) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FunctionBody) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FunctionBody) SetImage(v string)`

SetImage sets Image field to given value.


### GetType

`func (o *FunctionBody) GetType() FunctionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionBody) GetTypeOk() (*FunctionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionBody) SetType(v FunctionType)`

SetType sets Type field to given value.


### GetSize

`func (o *FunctionBody) GetSize() FunctionSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FunctionBody) GetSizeOk() (*FunctionSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FunctionBody) SetSize(v FunctionSize)`

SetSize sets Size field to given value.


### GetPort

`func (o *FunctionBody) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FunctionBody) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FunctionBody) SetPort(v int32)`

SetPort sets Port field to given value.


### GetRoutes

`func (o *FunctionBody) GetRoutes() []FunctionRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *FunctionBody) GetRoutesOk() (*[]FunctionRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *FunctionBody) SetRoutes(v []FunctionRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *FunctionBody) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetScaling

`func (o *FunctionBody) GetScaling() FunctionScaling`

GetScaling returns the Scaling field if non-nil, zero value otherwise.

### GetScalingOk

`func (o *FunctionBody) GetScalingOk() (*FunctionScaling, bool)`

GetScalingOk returns a tuple with the Scaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaling

`func (o *FunctionBody) SetScaling(v FunctionScaling)`

SetScaling sets Scaling field to given value.


### GetDeployments

`func (o *FunctionBody) GetDeployments() []FunctionDeploymentBody`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *FunctionBody) GetDeploymentsOk() (*[]FunctionDeploymentBody, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *FunctionBody) SetDeployments(v []FunctionDeploymentBody)`

SetDeployments sets Deployments field to given value.


### GetSecrets

`func (o *FunctionBody) GetSecrets() []FunctionEnv`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *FunctionBody) GetSecretsOk() (*[]FunctionEnv, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *FunctionBody) SetSecrets(v []FunctionEnv)`

SetSecrets sets Secrets field to given value.


### GetCompliance

`func (o *FunctionBody) GetCompliance() []FunctionCompliance`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *FunctionBody) GetComplianceOk() (*[]FunctionCompliance, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *FunctionBody) SetCompliance(v []FunctionCompliance)`

SetCompliance sets Compliance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


