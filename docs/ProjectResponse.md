# ProjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Project id | 
**OrgId** | **string** | Organisation id | 
**Name** | **string** | Project name | 
**Date** | [**Date**](Date.md) |  | 

## Methods

### NewProjectResponse

`func NewProjectResponse(id string, orgId string, name string, date Date, ) *ProjectResponse`

NewProjectResponse instantiates a new ProjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectResponseWithDefaults

`func NewProjectResponseWithDefaults() *ProjectResponse`

NewProjectResponseWithDefaults instantiates a new ProjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectResponse) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *ProjectResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ProjectResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ProjectResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetName

`func (o *ProjectResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDate

`func (o *ProjectResponse) GetDate() Date`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ProjectResponse) GetDateOk() (*Date, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ProjectResponse) SetDate(v Date)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


