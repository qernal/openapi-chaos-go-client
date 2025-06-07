# \UsersAPI

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersQuotasGet**](UsersAPI.md#UsersQuotasGet) | **Get** /users/{user_id}/quotas/{quota_entity_quota} | Get specific user quota
[**UsersQuotasList**](UsersAPI.md#UsersQuotasList) | **Get** /users/{user_id}/quotas | List user quotas



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
	resp, r, err := apiClient.UsersAPI.UsersQuotasGet(context.Background(), userId, quotaEntityQuota).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersQuotasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersQuotasGet`: []Quota
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersQuotasGet`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.UsersQuotasList(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersQuotasList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersQuotasList`: []Quota
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersQuotasList`: %v\n", resp)
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

