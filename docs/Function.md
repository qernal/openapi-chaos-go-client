# Function

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the function | 
**ProjectId** | **string** | ID of the project this function belongs to | 
**Version** | **string** | Function spec version | 
**Name** | **string** | Name of the function | 
**Description** | **string** | Description of what the function does | 
**Image** | **string** | Path to container image | 
**Revision** | **string** | Function revision | 
**Type** | [**FunctionType**](FunctionType.md) |  | 
**Size** | [**FunctionSize**](FunctionSize.md) |  | 
**Port** | **int32** | Port the application runs on | 
**Routes** | Pointer to [**[]FunctionRoute**](FunctionRoute.md) | The public route/path to this function, only applicable to http type functions | [optional] 
**Scaling** | [**FunctionScaling**](FunctionScaling.md) |  | 
**Deployments** | [**[]FunctionDeployment**](FunctionDeployment.md) | List of deployments for this function | 
**Secrets** | [**[]FunctionEnv**](FunctionEnv.md) | List of environment variables for secrets | 
**Compliance** | Pointer to [**[]FunctionCompliance**](FunctionCompliance.md) | Tags to limit deployment | [optional] 

## Methods

### NewFunction

`func NewFunction(id string, projectId string, version string, name string, description string, image string, revision string, type_ FunctionType, size FunctionSize, port int32, scaling FunctionScaling, deployments []FunctionDeployment, secrets []FunctionEnv, ) *Function`

NewFunction instantiates a new Function object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionWithDefaults

`func NewFunctionWithDefaults() *Function`

NewFunctionWithDefaults instantiates a new Function object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Function) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Function) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Function) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *Function) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Function) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Function) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersion

`func (o *Function) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Function) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Function) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetName

`func (o *Function) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Function) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Function) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Function) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Function) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Function) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImage

`func (o *Function) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Function) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Function) SetImage(v string)`

SetImage sets Image field to given value.


### GetRevision

`func (o *Function) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Function) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Function) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetType

`func (o *Function) GetType() FunctionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Function) GetTypeOk() (*FunctionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Function) SetType(v FunctionType)`

SetType sets Type field to given value.


### GetSize

`func (o *Function) GetSize() FunctionSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Function) GetSizeOk() (*FunctionSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Function) SetSize(v FunctionSize)`

SetSize sets Size field to given value.


### GetPort

`func (o *Function) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Function) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Function) SetPort(v int32)`

SetPort sets Port field to given value.


### GetRoutes

`func (o *Function) GetRoutes() []FunctionRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *Function) GetRoutesOk() (*[]FunctionRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *Function) SetRoutes(v []FunctionRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *Function) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetScaling

`func (o *Function) GetScaling() FunctionScaling`

GetScaling returns the Scaling field if non-nil, zero value otherwise.

### GetScalingOk

`func (o *Function) GetScalingOk() (*FunctionScaling, bool)`

GetScalingOk returns a tuple with the Scaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaling

`func (o *Function) SetScaling(v FunctionScaling)`

SetScaling sets Scaling field to given value.


### GetDeployments

`func (o *Function) GetDeployments() []FunctionDeployment`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *Function) GetDeploymentsOk() (*[]FunctionDeployment, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *Function) SetDeployments(v []FunctionDeployment)`

SetDeployments sets Deployments field to given value.


### GetSecrets

`func (o *Function) GetSecrets() []FunctionEnv`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *Function) GetSecretsOk() (*[]FunctionEnv, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *Function) SetSecrets(v []FunctionEnv)`

SetSecrets sets Secrets field to given value.


### GetCompliance

`func (o *Function) GetCompliance() []FunctionCompliance`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *Function) GetComplianceOk() (*[]FunctionCompliance, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *Function) SetCompliance(v []FunctionCompliance)`

SetCompliance sets Compliance field to given value.

### HasCompliance

`func (o *Function) HasCompliance() bool`

HasCompliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


