# \QuotasAPI

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganisationsQuotasGet**](QuotasAPI.md#OrganisationsQuotasGet) | **Get** /organisations/{organisation_id}/quotas/{quota_entity_quota} | Get specific organisation quota
[**OrganisationsQuotasList**](QuotasAPI.md#OrganisationsQuotasList) | **Get** /organisations/{organisation_id}/quotas | List organisation quotas
[**ProjectsQuotasGet**](QuotasAPI.md#ProjectsQuotasGet) | **Get** /projects/{project_id}/quotas/{quota_entity_quota} | Get specific project quota
[**ProjectsQuotasList**](QuotasAPI.md#ProjectsQuotasList) | **Get** /projects/{project_id}/quotas | List project quotas
[**UsersQuotasGet**](QuotasAPI.md#UsersQuotasGet) | **Get** /users/{user_id}/quotas/{quota_entity_quota} | Get specific user quota
[**UsersQuotasList**](QuotasAPI.md#UsersQuotasList) | **Get** /users/{user_id}/quotas | List user quotas



## OrganisationsQuotasGet

> []Quota OrganisationsQuotasGet(ctx, organisationId, quotaEntityQuota).Execute()

Get specific organisation quota



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
	quotaEntityQuota := "projects_secrets" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotasAPI.OrganisationsQuotasGet(context.Background(), organisationId, quotaEntityQuota).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotasAPI.OrganisationsQuotasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganisationsQuotasGet`: []Quota
	fmt.Fprintf(os.Stdout, "Response from `QuotasAPI.OrganisationsQuotasGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation ID reference | 
**quotaEntityQuota** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganisationsQuotasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Quota**](Quota.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganisationsQuotasList

> []Quota OrganisationsQuotasList(ctx, organisationId).Execute()

List organisation quotas



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotasAPI.OrganisationsQuotasList(context.Background(), organisationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotasAPI.OrganisationsQuotasList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganisationsQuotasList`: []Quota
	fmt.Fprintf(os.Stdout, "Response from `QuotasAPI.OrganisationsQuotasList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganisationsQuotasListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Quota**](Quota.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsQuotasGet

> []Quota ProjectsQuotasGet(ctx, projectId, quotaEntityQuota).Execute()

Get specific project quota



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
	quotaEntityQuota := "projects_secrets" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotasAPI.ProjectsQuotasGet(context.Background(), projectId, quotaEntityQuota).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotasAPI.ProjectsQuotasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsQuotasGet`: []Quota
	fmt.Fprintf(os.Stdout, "Response from `QuotasAPI.ProjectsQuotasGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 
**quotaEntityQuota** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsQuotasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Quota**](Quota.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsQuotasList

> []Quota ProjectsQuotasList(ctx, projectId).Execute()

List project quotas



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
	resp, r, err := apiClient.QuotasAPI.ProjectsQuotasList(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotasAPI.ProjectsQuotasList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsQuotasList`: []Quota
	fmt.Fprintf(os.Stdout, "Response from `QuotasAPI.ProjectsQuotasList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsQuotasListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Quota**](Quota.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersQuotasGet

> []Quota UsersQuotasGet(ctx, userId, quotaEntityQuota).Execute()

Get specific user quota



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
	userId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | User ID reference
	quotaEntityQuota := "projects_secrets" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotasAPI.UsersQuotasGet(context.Background(), userId, quotaEntityQuota).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotasAPI.UsersQuotasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersQuotasGet`: []Quota
	fmt.Fprintf(os.Stdout, "Response from `QuotasAPI.UsersQuotasGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID reference | 
**quotaEntityQuota** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersQuotasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Quota**](Quota.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersQuotasList

> []Quota UsersQuotasList(ctx, userId).Execute()

List user quotas



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
	userId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | User ID reference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotasAPI.UsersQuotasList(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotasAPI.UsersQuotasList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersQuotasList`: []Quota
	fmt.Fprintf(os.Stdout, "Response from `QuotasAPI.UsersQuotasList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersQuotasListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Quota**](Quota.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

