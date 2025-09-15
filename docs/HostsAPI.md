# \HostsAPI

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsHostsCreate**](HostsAPI.md#ProjectsHostsCreate) | **Post** /projects/{project_id}/hosts | Create host for project
[**ProjectsHostsDelete**](HostsAPI.md#ProjectsHostsDelete) | **Delete** /projects/{project_id}/hosts/{hostname} | Delete specific host by hostname
[**ProjectsHostsGet**](HostsAPI.md#ProjectsHostsGet) | **Get** /projects/{project_id}/hosts/{hostname} | Get specific host by hostname
[**ProjectsHostsList**](HostsAPI.md#ProjectsHostsList) | **Get** /projects/{project_id}/hosts | List hosts for project
[**ProjectsHostsUpdate**](HostsAPI.md#ProjectsHostsUpdate) | **Put** /projects/{project_id}/hosts/{hostname} | Update specific host by hostname
[**ProjectsHostsVerifyCreate**](HostsAPI.md#ProjectsHostsVerifyCreate) | **Post** /projects/{project_id}/hosts/{hostname}/verify | Schedule host verification task



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
	openapiclient "github.com/qernal/openapi-chaos-go-client"
)

func main() {
	projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference
	hostBody := *openapiclient.NewHostBody("example-domain.com", "projects:d8f6ede4-e3a4-414d-b6c3-b8a972773e56/MY_CERT", false) // HostBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ProjectsHostsCreate(context.Background(), projectId).HostBody(hostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ProjectsHostsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsHostsCreate`: Host
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ProjectsHostsCreate`: %v\n", resp)
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
	openapiclient "github.com/qernal/openapi-chaos-go-client"
)

func main() {
	projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference
	hostname := "example-domain.com" // string | Hostname

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ProjectsHostsDelete(context.Background(), projectId, hostname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ProjectsHostsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsHostsDelete`: DeletedResponse
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ProjectsHostsDelete`: %v\n", resp)
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
	openapiclient "github.com/qernal/openapi-chaos-go-client"
)

func main() {
	projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference
	hostname := "example-domain.com" // string | Hostname

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ProjectsHostsGet(context.Background(), projectId, hostname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ProjectsHostsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsHostsGet`: Host
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ProjectsHostsGet`: %v\n", resp)
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

> []Host ProjectsHostsList(ctx, projectId).Page(page).Execute()

List hosts for project

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
	page := *openapiclient.NewOrganisationsListPageParameter() // OrganisationsListPageParameter | Query parameters for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ProjectsHostsList(context.Background(), projectId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ProjectsHostsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsHostsList`: []Host
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ProjectsHostsList`: %v\n", resp)
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

[**[]Host**](Host.md)

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
	openapiclient "github.com/qernal/openapi-chaos-go-client"
)

func main() {
	projectId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project ID reference
	hostname := "example-domain.com" // string | Hostname
	hostBodyPatch := *openapiclient.NewHostBodyPatch() // HostBodyPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ProjectsHostsUpdate(context.Background(), projectId, hostname).HostBodyPatch(hostBodyPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ProjectsHostsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsHostsUpdate`: Host
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ProjectsHostsUpdate`: %v\n", resp)
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


## ProjectsHostsVerifyCreate

> Host ProjectsHostsVerifyCreate(ctx, projectId, hostname).Execute()

Schedule host verification task

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
	hostname := "example-domain.com" // string | Hostname

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.ProjectsHostsVerifyCreate(context.Background(), projectId, hostname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.ProjectsHostsVerifyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsHostsVerifyCreate`: Host
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.ProjectsHostsVerifyCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 
**hostname** | **string** | Hostname | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsHostsVerifyCreateRequest struct via the builder pattern


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

