# \SecretsApi

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsSecretsCreate**](SecretsApi.md#ProjectsSecretsCreate) | **Post** /projects/{project_id}/secrets | Create project secret
[**ProjectsSecretsDelete**](SecretsApi.md#ProjectsSecretsDelete) | **Delete** /projects/{project_id}/secrets/{secret_name} | Delete project secret
[**ProjectsSecretsGet**](SecretsApi.md#ProjectsSecretsGet) | **Get** /projects/{project_id}/secrets/{secret_name} | Get project secret
[**ProjectsSecretsList**](SecretsApi.md#ProjectsSecretsList) | **Get** /projects/{project_id}/secrets | List project secrets of a specific type
[**ProjectsSecretsUpdate**](SecretsApi.md#ProjectsSecretsUpdate) | **Put** /projects/{project_id}/secrets/{secret_name} | Update project secret



## ProjectsSecretsCreate

> SecretResponse ProjectsSecretsCreate(ctx, projectId).SecretBody(secretBody).Execute()

Create project secret



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
    secretBody := *openapiclient.NewSecretBody("Name_example", openapiclient.SecretCreateType("registry"), openapiclient.SecretCreatePayload{SecretCertificate: openapiclient.NewSecretCertificate("<x509 certificate pem format>", "<base64 encrypted pkcs8 or pkcs1 pem format>")}, "keys/dek/123") // SecretBody | Create/Update any field

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.ProjectsSecretsCreate(context.Background(), projectId).SecretBody(secretBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.ProjectsSecretsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsSecretsCreate`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.ProjectsSecretsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsSecretsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretBody** | [**SecretBody**](SecretBody.md) | Create/Update any field | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsSecretsDelete

> DeletedResponse ProjectsSecretsDelete(ctx, projectId, secretName).Execute()

Delete project secret



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
    secretName := "MY_SECRET" // string | Unique secret name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.ProjectsSecretsDelete(context.Background(), projectId, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.ProjectsSecretsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsSecretsDelete`: DeletedResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.ProjectsSecretsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 
**secretName** | **string** | Unique secret name | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsSecretsDeleteRequest struct via the builder pattern


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


## ProjectsSecretsGet

> SecretMetaResponse ProjectsSecretsGet(ctx, projectId, secretName).Execute()

Get project secret



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
    secretName := "MY_SECRET" // string | Unique secret name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.ProjectsSecretsGet(context.Background(), projectId, secretName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.ProjectsSecretsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsSecretsGet`: SecretMetaResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.ProjectsSecretsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 
**secretName** | **string** | Unique secret name | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsSecretsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecretMetaResponse**](SecretMetaResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsSecretsList

> ListSecretResponse ProjectsSecretsList(ctx, projectId).Page(page).SecretType(secretType).Execute()

List project secrets of a specific type



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
    secretType := openapiclient.SecretMetaType("registry") // SecretMetaType | Type of secret to filter on (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.ProjectsSecretsList(context.Background(), projectId).Page(page).SecretType(secretType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.ProjectsSecretsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsSecretsList`: ListSecretResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.ProjectsSecretsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsSecretsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | [**OrganisationsListPageParameter**](OrganisationsListPageParameter.md) | Query parameters for pagination | 
 **secretType** | [**SecretMetaType**](SecretMetaType.md) | Type of secret to filter on | 

### Return type

[**ListSecretResponse**](ListSecretResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsSecretsUpdate

> SecretResponse ProjectsSecretsUpdate(ctx, projectId, secretName).SecretBodyPatch(secretBodyPatch).Execute()

Update project secret



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
    secretName := "MY_SECRET" // string | Unique secret name
    secretBodyPatch := *openapiclient.NewSecretBodyPatch(openapiclient.SecretCreateType("registry"), openapiclient.SecretCreatePayload{SecretCertificate: openapiclient.NewSecretCertificate("<x509 certificate pem format>", "<base64 encrypted pkcs8 or pkcs1 pem format>")}, "keys/dek/123") // SecretBodyPatch | Update any field

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretsApi.ProjectsSecretsUpdate(context.Background(), projectId, secretName).SecretBodyPatch(secretBodyPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretsApi.ProjectsSecretsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsSecretsUpdate`: SecretResponse
    fmt.Fprintf(os.Stdout, "Response from `SecretsApi.ProjectsSecretsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 
**secretName** | **string** | Unique secret name | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsSecretsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **secretBodyPatch** | [**SecretBodyPatch**](SecretBodyPatch.md) | Update any field | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

