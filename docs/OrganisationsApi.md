# \OrganisationsApi

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrganisationsOrgId**](OrganisationsApi.md#DeleteOrganisationsOrgId) | **Delete** /organisations/{organisation_id} | Delete an organisation
[**GetOrganisations**](OrganisationsApi.md#GetOrganisations) | **Get** /organisations | List organisations
[**GetOrganisationsOrgId**](OrganisationsApi.md#GetOrganisationsOrgId) | **Get** /organisations/{organisation_id} | Get an organisation
[**PostOrganisations**](OrganisationsApi.md#PostOrganisations) | **Post** /organisations | Create organisations
[**PutOrganisationsOrgId**](OrganisationsApi.md#PutOrganisationsOrgId) | **Put** /organisations/{organisation_id} | Update an organisation



## DeleteOrganisationsOrgId

> DeletedResponse DeleteOrganisationsOrgId(ctx, organisationId).Execute()

Delete an organisation



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
    organisationId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Organisation ID reference

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationsApi.DeleteOrganisationsOrgId(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationsApi.DeleteOrganisationsOrgId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrganisationsOrgId`: DeletedResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganisationsApi.DeleteOrganisationsOrgId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganisationsOrgIdRequest struct via the builder pattern


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


## GetOrganisations

> ListOrganisationResponse GetOrganisations(ctx).Page(page).Execute()

List organisations



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
    page := map[string][]openapiclient.GetOrganisationsPageParameter{"key": map[string]interface{}{ ... }} // GetOrganisationsPageParameter | Query parameters for pagination (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationsApi.GetOrganisations(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationsApi.GetOrganisations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganisations`: ListOrganisationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganisationsApi.GetOrganisations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | [**GetOrganisationsPageParameter**](GetOrganisationsPageParameter.md) | Query parameters for pagination | 

### Return type

[**ListOrganisationResponse**](ListOrganisationResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisationsOrgId

> OrganisationResponse GetOrganisationsOrgId(ctx, organisationId).Execute()

Get an organisation



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
    organisationId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Organisation ID reference

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationsApi.GetOrganisationsOrgId(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationsApi.GetOrganisationsOrgId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganisationsOrgId`: OrganisationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganisationsApi.GetOrganisationsOrgId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganisationsOrgIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganisationResponse**](OrganisationResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganisations

> OrganisationResponse PostOrganisations(ctx).OrganisationBody(organisationBody).Execute()

Create organisations



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
    organisationBody := *openapiclient.NewOrganisationBody("Name_example") // OrganisationBody | Create/Update any field (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationsApi.PostOrganisations(context.Background()).OrganisationBody(organisationBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationsApi.PostOrganisations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostOrganisations`: OrganisationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganisationsApi.PostOrganisations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganisationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organisationBody** | [**OrganisationBody**](OrganisationBody.md) | Create/Update any field | 

### Return type

[**OrganisationResponse**](OrganisationResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrganisationsOrgId

> OrganisationResponse PutOrganisationsOrgId(ctx, organisationId).OrganisationBody(organisationBody).Execute()

Update an organisation



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
    organisationId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Organisation ID reference
    organisationBody := *openapiclient.NewOrganisationBody("Name_example") // OrganisationBody | Create/Update any field (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationsApi.PutOrganisationsOrgId(context.Background(), organisationId).OrganisationBody(organisationBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationsApi.PutOrganisationsOrgId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutOrganisationsOrgId`: OrganisationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganisationsApi.PutOrganisationsOrgId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrganisationsOrgIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organisationBody** | [**OrganisationBody**](OrganisationBody.md) | Create/Update any field | 

### Return type

[**OrganisationResponse**](OrganisationResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

