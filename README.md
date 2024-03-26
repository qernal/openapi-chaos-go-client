# Go API client for openapi_chaos_client

Central Management API - publicly exposed set of APIs for managing deployments

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.4.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi_chaos_client "github.com/qernal/openapi-chaos-go-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi_chaos_client.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi_chaos_client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi_chaos_client.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi_chaos_client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi_chaos_client.ContextOperationServerIndices` and `openapi_chaos_client.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi_chaos_client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_chaos_client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://chaos.qernal.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FunctionsAPI* | [**FunctionsDelete**](docs/FunctionsAPI.md#functionsdelete) | **Delete** /functions/{function_id} | Delete function
*FunctionsAPI* | [**FunctionsGet**](docs/FunctionsAPI.md#functionsget) | **Get** /functions/{function_id} | Get function (latest revision)
*FunctionsAPI* | [**FunctionsRevisionsGet**](docs/FunctionsAPI.md#functionsrevisionsget) | **Get** /functions/{function_id}/revisions/{function_revision_id} | Get a specific revision of a function
*FunctionsAPI* | [**FunctionsRevisionsList**](docs/FunctionsAPI.md#functionsrevisionslist) | **Get** /functions/{function_id}/revisions | List all revisions for a function
*FunctionsAPI* | [**FunctionsUpdate**](docs/FunctionsAPI.md#functionsupdate) | **Put** /functions/{function_id} | Update function
*FunctionsAPI* | [**ProjectsFunctionsCreate**](docs/FunctionsAPI.md#projectsfunctionscreate) | **Post** /projects/{project_id}/functions | Create function
*FunctionsAPI* | [**ProjectsFunctionsList**](docs/FunctionsAPI.md#projectsfunctionslist) | **Get** /projects/{project_id}/functions | List all functions within a project
*HostsAPI* | [**ProjectsHostsCreate**](docs/HostsAPI.md#projectshostscreate) | **Post** /projects/{project_id}/hosts | Create host for project
*HostsAPI* | [**ProjectsHostsDelete**](docs/HostsAPI.md#projectshostsdelete) | **Delete** /projects/{project_id}/hosts/{hostname} | Delete specific host by hostname
*HostsAPI* | [**ProjectsHostsGet**](docs/HostsAPI.md#projectshostsget) | **Get** /projects/{project_id}/hosts/{hostname} | Get specific host by hostname
*HostsAPI* | [**ProjectsHostsList**](docs/HostsAPI.md#projectshostslist) | **Get** /projects/{project_id}/hosts | List hosts for project
*HostsAPI* | [**ProjectsHostsUpdate**](docs/HostsAPI.md#projectshostsupdate) | **Put** /projects/{project_id}/hosts/{hostname} | Update specific host by hostname
*HostsAPI* | [**ProjectsHostsVerifyCreate**](docs/HostsAPI.md#projectshostsverifycreate) | **Post** /projects/{project_id}/hosts/{hostname}/verify | Schedule host verification task
*OrganisationsAPI* | [**OrganisationsCreate**](docs/OrganisationsAPI.md#organisationscreate) | **Post** /organisations | Create organisations
*OrganisationsAPI* | [**OrganisationsDelete**](docs/OrganisationsAPI.md#organisationsdelete) | **Delete** /organisations/{organisation_id} | Delete an organisation
*OrganisationsAPI* | [**OrganisationsGet**](docs/OrganisationsAPI.md#organisationsget) | **Get** /organisations/{organisation_id} | Get an organisation
*OrganisationsAPI* | [**OrganisationsList**](docs/OrganisationsAPI.md#organisationslist) | **Get** /organisations | List organisations
*OrganisationsAPI* | [**OrganisationsUpdate**](docs/OrganisationsAPI.md#organisationsupdate) | **Put** /organisations/{organisation_id} | Update an organisation
*ProjectsAPI* | [**OrganisationsProjectsList**](docs/ProjectsAPI.md#organisationsprojectslist) | **Get** /organisations/{organisation_id}/projects | Get all projects within an organisation
*ProjectsAPI* | [**ProjectsCreate**](docs/ProjectsAPI.md#projectscreate) | **Post** /projects | Create project
*ProjectsAPI* | [**ProjectsDelete**](docs/ProjectsAPI.md#projectsdelete) | **Delete** /projects/{project_id} | Delete project
*ProjectsAPI* | [**ProjectsGet**](docs/ProjectsAPI.md#projectsget) | **Get** /projects/{project_id} | Get project
*ProjectsAPI* | [**ProjectsList**](docs/ProjectsAPI.md#projectslist) | **Get** /projects | List projects
*ProjectsAPI* | [**ProjectsUpdate**](docs/ProjectsAPI.md#projectsupdate) | **Put** /projects/{project_id} | Update project
*ProvidersAPI* | [**ProvidersGet**](docs/ProvidersAPI.md#providersget) | **Get** /providers | Get available providers
*SecretsAPI* | [**ProjectsSecretsCreate**](docs/SecretsAPI.md#projectssecretscreate) | **Post** /projects/{project_id}/secrets | Create project secret
*SecretsAPI* | [**ProjectsSecretsDelete**](docs/SecretsAPI.md#projectssecretsdelete) | **Delete** /projects/{project_id}/secrets/{secret_name} | Delete project secret
*SecretsAPI* | [**ProjectsSecretsGet**](docs/SecretsAPI.md#projectssecretsget) | **Get** /projects/{project_id}/secrets/{secret_name} | Get project secret
*SecretsAPI* | [**ProjectsSecretsList**](docs/SecretsAPI.md#projectssecretslist) | **Get** /projects/{project_id}/secrets | List project secrets of a specific type
*SecretsAPI* | [**ProjectsSecretsUpdate**](docs/SecretsAPI.md#projectssecretsupdate) | **Put** /projects/{project_id}/secrets/{secret_name} | Update project secret
*TokensAPI* | [**AuthTokensCreate**](docs/TokensAPI.md#authtokenscreate) | **Post** /auth/tokens | Create new auth token
*TokensAPI* | [**AuthTokensDelete**](docs/TokensAPI.md#authtokensdelete) | **Delete** /auth/tokens/{token_id} | Delete token
*TokensAPI* | [**AuthTokensGet**](docs/TokensAPI.md#authtokensget) | **Get** /auth/tokens/{token_id} | Get token information
*TokensAPI* | [**AuthTokensList**](docs/TokensAPI.md#authtokenslist) | **Get** /auth/tokens | List all user auth tokens
*TokensAPI* | [**AuthTokensUpdate**](docs/TokensAPI.md#authtokensupdate) | **Put** /auth/tokens/{token_id} | Update token


## Documentation For Models

 - [AuthToken](docs/AuthToken.md)
 - [AuthTokenBody](docs/AuthTokenBody.md)
 - [AuthTokenMeta](docs/AuthTokenMeta.md)
 - [AuthTokenPatch](docs/AuthTokenPatch.md)
 - [BadRequestResponse](docs/BadRequestResponse.md)
 - [BadRequestResponseFields](docs/BadRequestResponseFields.md)
 - [ConflictResponse](docs/ConflictResponse.md)
 - [Date](docs/Date.md)
 - [DeletedResponse](docs/DeletedResponse.md)
 - [Function](docs/Function.md)
 - [FunctionBody](docs/FunctionBody.md)
 - [FunctionCompliance](docs/FunctionCompliance.md)
 - [FunctionDeployment](docs/FunctionDeployment.md)
 - [FunctionDeploymentBody](docs/FunctionDeploymentBody.md)
 - [FunctionEnv](docs/FunctionEnv.md)
 - [FunctionReplicas](docs/FunctionReplicas.md)
 - [FunctionReplicasAffinity](docs/FunctionReplicasAffinity.md)
 - [FunctionRoute](docs/FunctionRoute.md)
 - [FunctionScaling](docs/FunctionScaling.md)
 - [FunctionSize](docs/FunctionSize.md)
 - [FunctionType](docs/FunctionType.md)
 - [Host](docs/Host.md)
 - [HostBody](docs/HostBody.md)
 - [HostBodyPatch](docs/HostBodyPatch.md)
 - [HostVerificationStatus](docs/HostVerificationStatus.md)
 - [ListAuthTokens](docs/ListAuthTokens.md)
 - [ListFunction](docs/ListFunction.md)
 - [ListHosts](docs/ListHosts.md)
 - [ListOrganisationResponse](docs/ListOrganisationResponse.md)
 - [ListProjectResponse](docs/ListProjectResponse.md)
 - [ListSecretResponse](docs/ListSecretResponse.md)
 - [Location](docs/Location.md)
 - [NotFoundResponse](docs/NotFoundResponse.md)
 - [OrganisationBody](docs/OrganisationBody.md)
 - [OrganisationResponse](docs/OrganisationResponse.md)
 - [OrganisationsListPageParameter](docs/OrganisationsListPageParameter.md)
 - [PaginationLinks](docs/PaginationLinks.md)
 - [PaginationMeta](docs/PaginationMeta.md)
 - [ProjectBody](docs/ProjectBody.md)
 - [ProjectBodyPatch](docs/ProjectBodyPatch.md)
 - [ProjectResponse](docs/ProjectResponse.md)
 - [ProviderInner](docs/ProviderInner.md)
 - [ProviderInnerLocations](docs/ProviderInnerLocations.md)
 - [SecretBody](docs/SecretBody.md)
 - [SecretBodyPatch](docs/SecretBodyPatch.md)
 - [SecretCertificate](docs/SecretCertificate.md)
 - [SecretCreatePayload](docs/SecretCreatePayload.md)
 - [SecretCreateType](docs/SecretCreateType.md)
 - [SecretEnvironment](docs/SecretEnvironment.md)
 - [SecretMetaResponse](docs/SecretMetaResponse.md)
 - [SecretMetaResponseCertificatePayload](docs/SecretMetaResponseCertificatePayload.md)
 - [SecretMetaResponseDek](docs/SecretMetaResponseDek.md)
 - [SecretMetaResponsePayload](docs/SecretMetaResponsePayload.md)
 - [SecretMetaResponseRegistryPayload](docs/SecretMetaResponseRegistryPayload.md)
 - [SecretMetaType](docs/SecretMetaType.md)
 - [SecretRegistry](docs/SecretRegistry.md)
 - [SecretResponse](docs/SecretResponse.md)
 - [SecretResponsePayload](docs/SecretResponsePayload.md)
 - [UnauthorisedResponse](docs/UnauthorisedResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### token

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), openapi_chaos_client.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### cookie

- **Type**: API key
- **API key parameter name**: qernal_kratos_session
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: qernal_kratos_session and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi_chaos_client.ContextAPIKeys,
		map[string]openapi_chaos_client.APIKey{
			"qernal_kratos_session": {Key: "API_KEY_STRING"},
		},
	)
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

