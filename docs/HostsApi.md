# \HostsApi

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsHostsCreate**](HostsApi.md#ProjectsHostsCreate) | **Post** /projects/{project_id}/hosts | Create host for project
[**ProjectsHostsDelete**](HostsApi.md#ProjectsHostsDelete) | **Delete** /projects/{project_id}/hosts/{hostname} | Delete specific host by hostname
[**ProjectsHostsGet**](HostsApi.md#ProjectsHostsGet) | **Get** /projects/{project_id}/hosts/{hostname} | Get specific host by hostname
[**ProjectsHostsList**](HostsApi.md#ProjectsHostsList) | **Get** /projects/{project_id}/hosts | List hosts for project
[**ProjectsHostsUpdate**](HostsApi.md#ProjectsHostsUpdate) | **Put** /projects/{project_id}/hosts/{hostname} | Update specific host by hostname



## ProjectsHostsCreate

> Host ProjectsHostsCreate(ctx, projectId).HostBody(hostBody).Execute()

Create host for project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference
    hostBody := *openapiclient.NewHostBody("Host_example", "projects:secrets/MY_CERT", false) // HostBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HostsApi.ProjectsHostsCreate(context.Background(), projectId).HostBody(hostBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ProjectsHostsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsHostsCreate`: Host
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ProjectsHostsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsHostsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostBody** | [**HostBody**](HostBody.md) |  | 

### Return type

[**Host**](Host.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsHostsDelete

> DeletedResponse ProjectsHostsDelete(ctx, projectId, hostname).Execute()

Delete specific host by hostname

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference
    hostname := "example-domain.com" // string | Hostname

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HostsApi.ProjectsHostsDelete(context.Background(), projectId, hostname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ProjectsHostsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsHostsDelete`: DeletedResponse
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ProjectsHostsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 
**hostname** | **string** | Hostname | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsHostsDeleteRequest struct via the builder pattern


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


## ProjectsHostsGet

> Host ProjectsHostsGet(ctx, projectId, hostname).Execute()

Get specific host by hostname

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference
    hostname := "example-domain.com" // string | Hostname

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HostsApi.ProjectsHostsGet(context.Background(), projectId, hostname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ProjectsHostsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsHostsGet`: Host
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ProjectsHostsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 
**hostname** | **string** | Hostname | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsHostsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Host**](Host.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsHostsList

> ListHosts ProjectsHostsList(ctx, projectId).Page(page).Execute()

List hosts for project

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference
    page := map[string][]openapiclient.OrganisationsListPageParameter{"key": map[string]interface{}{ ... }} // OrganisationsListPageParameter | Query parameters for pagination (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HostsApi.ProjectsHostsList(context.Background(), projectId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ProjectsHostsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsHostsList`: ListHosts
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ProjectsHostsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsHostsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | [**OrganisationsListPageParameter**](OrganisationsListPageParameter.md) | Query parameters for pagination | 

### Return type

[**ListHosts**](ListHosts.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsHostsUpdate

> Host ProjectsHostsUpdate(ctx, projectId, hostname).HostBodyPatch(hostBodyPatch).Execute()

Update specific host by hostname

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference
    hostname := "example-domain.com" // string | Hostname
    hostBodyPatch := *openapiclient.NewHostBodyPatch() // HostBodyPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HostsApi.ProjectsHostsUpdate(context.Background(), projectId, hostname).HostBodyPatch(hostBodyPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ProjectsHostsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsHostsUpdate`: Host
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ProjectsHostsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 
**hostname** | **string** | Hostname | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsHostsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hostBodyPatch** | [**HostBodyPatch**](HostBodyPatch.md) |  | 

### Return type

[**Host**](Host.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

