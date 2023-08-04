# Go API client for openapi-chaos-client

Central Management API - publicly exposed set of APIs for managing deployments

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi-chaos-client "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi-chaos-client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi-chaos-client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi-chaos-client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi-chaos-client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://chaos.qernal.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OrganisationsApi* | [**DeleteOrganisationsOrgId**](docs/OrganisationsApi.md#deleteorganisationsorgid) | **Delete** /organisations/{organisation_id} | Delete an organisation
*OrganisationsApi* | [**GetOrganisations**](docs/OrganisationsApi.md#getorganisations) | **Get** /organisations | List organisations
*OrganisationsApi* | [**GetOrganisationsOrgId**](docs/OrganisationsApi.md#getorganisationsorgid) | **Get** /organisations/{organisation_id} | Get an organisation
*OrganisationsApi* | [**PostOrganisations**](docs/OrganisationsApi.md#postorganisations) | **Post** /organisations | Create organisations
*OrganisationsApi* | [**PutOrganisationsOrgId**](docs/OrganisationsApi.md#putorganisationsorgid) | **Put** /organisations/{organisation_id} | Update an organisation
*ProjectsApi* | [**DeleteProjectsProjectId**](docs/ProjectsApi.md#deleteprojectsprojectid) | **Delete** /projects/{project_id} | Delete project
*ProjectsApi* | [**GetOrganisationsOrgIdProjects**](docs/ProjectsApi.md#getorganisationsorgidprojects) | **Get** /organisations/{organisation_id}/projects | Get all projects within an organisation
*ProjectsApi* | [**GetProjects**](docs/ProjectsApi.md#getprojects) | **Get** /projects | List projects
*ProjectsApi* | [**GetProjectsProjectId**](docs/ProjectsApi.md#getprojectsprojectid) | **Get** /projects/{project_id} | Get project
*ProjectsApi* | [**PostProjects**](docs/ProjectsApi.md#postprojects) | **Post** /projects | Create project
*ProjectsApi* | [**PutProjectsProjectId**](docs/ProjectsApi.md#putprojectsprojectid) | **Put** /projects/{project_id} | Update project


## Documentation For Models

 - [BadRequestResponse](docs/BadRequestResponse.md)
 - [BadRequestResponseFields](docs/BadRequestResponseFields.md)
 - [ConflictResponse](docs/ConflictResponse.md)
 - [DeletedResponse](docs/DeletedResponse.md)
 - [GetOrganisationsPageParameter](docs/GetOrganisationsPageParameter.md)
 - [ListOrganisationResponse](docs/ListOrganisationResponse.md)
 - [ListProjectResponse](docs/ListProjectResponse.md)
 - [NotFoundResponse](docs/NotFoundResponse.md)
 - [OrganisationBody](docs/OrganisationBody.md)
 - [OrganisationResponse](docs/OrganisationResponse.md)
 - [PaginationLinks](docs/PaginationLinks.md)
 - [PaginationMeta](docs/PaginationMeta.md)
 - [ProjectBody](docs/ProjectBody.md)
 - [ProjectResponse](docs/ProjectResponse.md)
 - [UnauthorisedResponse](docs/UnauthorisedResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### JWT

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@qernal.com

