# \TokensApi

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTokensCreate**](TokensApi.md#AuthTokensCreate) | **Post** /auth/tokens | Create new auth token
[**AuthTokensDelete**](TokensApi.md#AuthTokensDelete) | **Delete** /auth/tokens/{token_id} | Delete token
[**AuthTokensGet**](TokensApi.md#AuthTokensGet) | **Get** /auth/tokens/{token_id} | Get token information
[**AuthTokensList**](TokensApi.md#AuthTokensList) | **Get** /auth/tokens | List all user auth tokens
[**AuthTokensUpdate**](TokensApi.md#AuthTokensUpdate) | **Put** /auth/tokens/{token_id} | Update token



## AuthTokensCreate

> AuthToken AuthTokensCreate(ctx).AuthTokenBody(authTokenBody).Execute()

Create new auth token



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
    authTokenBody := *openapiclient.NewAuthTokenBody("TF token", int32(90)) // AuthTokenBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.AuthTokensCreate(context.Background()).AuthTokenBody(authTokenBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.AuthTokensCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokensCreate`: AuthToken
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.AuthTokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authTokenBody** | [**AuthTokenBody**](AuthTokenBody.md) |  | 

### Return type

[**AuthToken**](AuthToken.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokensDelete

> DeletedResponse AuthTokensDelete(ctx, tokenId).Execute()

Delete token

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
    tokenId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Token ID reference

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.AuthTokensDelete(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.AuthTokensDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokensDelete`: DeletedResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.AuthTokensDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Token ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensDeleteRequest struct via the builder pattern


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


## AuthTokensGet

> AuthTokenMeta AuthTokensGet(ctx, tokenId).Execute()

Get token information

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
    tokenId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Token ID reference

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.AuthTokensGet(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.AuthTokensGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokensGet`: AuthTokenMeta
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.AuthTokensGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Token ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthTokenMeta**](AuthTokenMeta.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokensList

> ListAuthTokens AuthTokensList(ctx).Page(page).Execute()

List all user auth tokens

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
    page := map[string][]openapiclient.OrganisationsListPageParameter{"key": map[string]interface{}{ ... }} // OrganisationsListPageParameter | Query parameters for pagination (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.AuthTokensList(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.AuthTokensList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokensList`: ListAuthTokens
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.AuthTokensList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | [**OrganisationsListPageParameter**](OrganisationsListPageParameter.md) | Query parameters for pagination | 

### Return type

[**ListAuthTokens**](ListAuthTokens.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokensUpdate

> AuthToken AuthTokensUpdate(ctx, tokenId).AuthTokenPatch(authTokenPatch).Execute()

Update token

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
    tokenId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Token ID reference
    authTokenPatch := *openapiclient.NewAuthTokenPatch() // AuthTokenPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.AuthTokensUpdate(context.Background(), tokenId).AuthTokenPatch(authTokenPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.AuthTokensUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokensUpdate`: AuthToken
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.AuthTokensUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Token ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokensUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authTokenPatch** | [**AuthTokenPatch**](AuthTokenPatch.md) |  | 

### Return type

[**AuthToken**](AuthToken.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

