# \OrganisationsApi

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganisationsCreate**](OrganisationsApi.md#OrganisationsCreate) | **Post** /organisations | Create organisations
[**OrganisationsDelete**](OrganisationsApi.md#OrganisationsDelete) | **Delete** /organisations/{organisation_id} | Delete an organisation
[**OrganisationsGet**](OrganisationsApi.md#OrganisationsGet) | **Get** /organisations/{organisation_id} | Get an organisation
[**OrganisationsList**](OrganisationsApi.md#OrganisationsList) | **Get** /organisations | List organisations
[**OrganisationsUpdate**](OrganisationsApi.md#OrganisationsUpdate) | **Put** /organisations/{organisation_id} | Update an organisation



## OrganisationsCreate

> OrganisationResponse OrganisationsCreate(ctx).OrganisationBody(organisationBody).Execute()

Create organisations



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
    organisationBody := *openapiclient.NewOrganisationBody("my org") // OrganisationBody | Create/Update any field (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationsApi.OrganisationsCreate(context.Background()).OrganisationBody(organisationBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationsApi.OrganisationsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganisationsCreate`: OrganisationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganisationsApi.OrganisationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganisationsCreateRequest struct via the builder pattern


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


## OrganisationsDelete

> DeletedResponse OrganisationsDelete(ctx, organisationId).Execute()

Delete an organisation



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
    resp, r, err := apiClient.OrganisationsApi.OrganisationsDelete(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationsApi.OrganisationsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganisationsDelete`: DeletedResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganisationsApi.OrganisationsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganisationsDeleteRequest struct via the builder pattern


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


## OrganisationsGet

> OrganisationResponse OrganisationsGet(ctx, organisationId).Execute()

Get an organisation



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
    resp, r, err := apiClient.OrganisationsApi.OrganisationsGet(context.Background(), organisationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationsApi.OrganisationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganisationsGet`: OrganisationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganisationsApi.OrganisationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganisationsGetRequest struct via the builder pattern


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


## OrganisationsList

> ListOrganisationResponse OrganisationsList(ctx).Page(page).Execute()

List organisations



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
    page := map[string][]openapiclient.OrganisationsListPageParameter{"key": map[string]interface{}{ ... }} // OrganisationsListPageParameter | Query parameters for pagination (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationsApi.OrganisationsList(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationsApi.OrganisationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganisationsList`: ListOrganisationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganisationsApi.OrganisationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganisationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | [**OrganisationsListPageParameter**](OrganisationsListPageParameter.md) | Query parameters for pagination | 

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


## OrganisationsUpdate

> OrganisationResponse OrganisationsUpdate(ctx, organisationId).OrganisationBody(organisationBody).Execute()

Update an organisation



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
    organisationBody := *openapiclient.NewOrganisationBody("my org") // OrganisationBody | Create/Update any field (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganisationsApi.OrganisationsUpdate(context.Background(), organisationId).OrganisationBody(organisationBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganisationsApi.OrganisationsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganisationsUpdate`: OrganisationResponse
    fmt.Fprintf(os.Stdout, "Response from `OrganisationsApi.OrganisationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organisationId** | **string** | Organisation ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganisationsUpdateRequest struct via the builder pattern


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

