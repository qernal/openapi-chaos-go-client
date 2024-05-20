# \ProjectsAPI

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganisationsProjectsList**](ProjectsAPI.md#OrganisationsProjectsList) | **Get** /organisations/{organisation_id}/projects | Get all projects within an organisation
[**ProjectsCreate**](ProjectsAPI.md#ProjectsCreate) | **Post** /projects | Create project
[**ProjectsDelete**](ProjectsAPI.md#ProjectsDelete) | **Delete** /projects/{project_id} | Delete project
[**ProjectsGet**](ProjectsAPI.md#ProjectsGet) | **Get** /projects/{project_id} | Get project
[**ProjectsList**](ProjectsAPI.md#ProjectsList) | **Get** /projects | List projects
[**ProjectsUpdate**](ProjectsAPI.md#ProjectsUpdate) | **Put** /projects/{project_id} | Update project



## OrganisationsProjectsList

> ListProjectResponse OrganisationsProjectsList(ctx, organisationId).Page(page).Execute()

Get all projects within an organisation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qernal/openapi-chaos-go-client"
)

func main() {
	organisationId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Organisation ID reference
	page := *openapiclient.NewOrganisationsListPageParameter() // OrganisationsListPageParameter | Query parameters for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.OrganisationsProjectsList(context.Background(), organisationId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.OrganisationsProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganisationsProjectsList`: ListProjectResponse
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.OrganisationsProjectsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganisationsProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | [**OrganisationsListPageParameter**](OrganisationsListPageParameter.md) | Query parameters for pagination | 

### Return type

[**ListProjectResponse**](ListProjectResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCreate

> ProjectResponse ProjectsCreate(ctx).ProjectBody(projectBody).Execute()

Create project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qernal/openapi-chaos-go-client"
)

func main() {
	projectBody := *openapiclient.NewProjectBody("3069614e-adc8-47cb-a69c-decf9c5f90fc", "first-api") // ProjectBody | Create/Update any field (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsCreate(context.Background()).ProjectBody(projectBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsCreate`: ProjectResponse
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectBody** | [**ProjectBody**](ProjectBody.md) | Create/Update any field | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDelete

> DeletedResponse ProjectsDelete(ctx, projectId).Execute()

Delete project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qernal/openapi-chaos-go-client"
)

func main() {
	projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsDelete(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsDelete`: DeletedResponse
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletedResponse**](DeletedResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsGet

> ProjectResponse ProjectsGet(ctx, projectId).Execute()

Get project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qernal/openapi-chaos-go-client"
)

func main() {
	projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsGet(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsGet`: ProjectResponse
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsList

> ListProjectResponse ProjectsList(ctx).Page(page).FName(fName).Execute()

List projects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qernal/openapi-chaos-go-client"
)

func main() {
	page := *openapiclient.NewOrganisationsListPageParameter() // OrganisationsListPageParameter | Query parameters for pagination (optional)
	fName := "my-proj*" // string | Filter resource on name, if the value ends in an asterix it will be treated as a partial search otherwise, it'll be an exact match  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsList(context.Background()).Page(page).FName(fName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsList`: ListProjectResponse
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | [**OrganisationsListPageParameter**](OrganisationsListPageParameter.md) | Query parameters for pagination | 
 **fName** | **string** | Filter resource on name, if the value ends in an asterix it will be treated as a partial search otherwise, it&#39;ll be an exact match  | 

### Return type

[**ListProjectResponse**](ListProjectResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsUpdate

> ProjectResponse ProjectsUpdate(ctx, projectId).ProjectBodyPatch(projectBodyPatch).Execute()

Update project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/qernal/openapi-chaos-go-client"
)

func main() {
	projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference
	projectBodyPatch := *openapiclient.NewProjectBodyPatch() // ProjectBodyPatch | Update any field (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsUpdate(context.Background(), projectId).ProjectBodyPatch(projectBodyPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsUpdate`: ProjectResponse
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectBodyPatch** | [**ProjectBodyPatch**](ProjectBodyPatch.md) | Update any field | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

