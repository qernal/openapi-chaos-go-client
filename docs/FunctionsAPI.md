# \FunctionsAPI

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FunctionsCreate**](FunctionsAPI.md#FunctionsCreate) | **Post** /functions | Create function
[**FunctionsDelete**](FunctionsAPI.md#FunctionsDelete) | **Delete** /functions/{function_id} | Delete function
[**FunctionsGet**](FunctionsAPI.md#FunctionsGet) | **Get** /functions/{function_id} | Get function (latest revision)
[**FunctionsRevisionsGet**](FunctionsAPI.md#FunctionsRevisionsGet) | **Get** /functions/{function_id}/revisions/{function_revision_id} | Get a specific revision of a function
[**FunctionsRevisionsList**](FunctionsAPI.md#FunctionsRevisionsList) | **Get** /functions/{function_id}/revisions | List all revisions for a function
[**FunctionsUpdate**](FunctionsAPI.md#FunctionsUpdate) | **Put** /functions/{function_id} | Update function
[**ProjectsFunctionsList**](FunctionsAPI.md#ProjectsFunctionsList) | **Get** /projects/{project_id}/functions | List all functions within a project



## FunctionsCreate

> Function FunctionsCreate(ctx).FunctionBody(functionBody).Execute()

Create function



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
	functionBody := *openapiclient.NewFunctionBody("51687d2f-07b0-4260-8ecb-f5098305fdd4", "1.0.0", "my-function", "My function does this", "docker.io/my-image:latest", openapiclient.FunctionType("http"), *openapiclient.NewFunctionSize(int32(128), int32(128)), int32(8080), *openapiclient.NewFunctionScaling("cpu", int32(30), int32(60)), []openapiclient.FunctionDeploymentBody{*openapiclient.NewFunctionDeploymentBody(*openapiclient.NewLocation("51687d2f-07b0-4260-8ecb-f5098305fdd4"), *openapiclient.NewFunctionReplicas(int32(1), int32(5), *openapiclient.NewFunctionReplicasAffinity(true, true)))}, []openapiclient.FunctionEnv{*openapiclient.NewFunctionEnv("MY_ENV_VAR", "projects:0a6b9ff3-6807-4820-b94b-5e1d7efcdd93/MY_SECRET@0")}, []openapiclient.FunctionCompliance{openapiclient.FunctionCompliance("soc2")}) // FunctionBody | Create/Update any field

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsCreate(context.Background()).FunctionBody(functionBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsCreate`: Function
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionBody** | [**FunctionBody**](FunctionBody.md) | Create/Update any field | 

### Return type

[**Function**](Function.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsDelete

> DeletedResponse FunctionsDelete(ctx, functionId).Execute()

Delete function



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
	functionId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Function ID reference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsDelete(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsDelete`: DeletedResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** | Function ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsDeleteRequest struct via the builder pattern


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


## FunctionsGet

> Function FunctionsGet(ctx, functionId).Execute()

Get function (latest revision)



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
	functionId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Function ID reference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsGet(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsGet`: Function
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** | Function ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Function**](Function.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsRevisionsGet

> Function FunctionsRevisionsGet(ctx, functionId, functionRevisionId).Execute()

Get a specific revision of a function



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
	functionId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Function ID reference
	functionRevisionId := "1069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Function revision ID reference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsRevisionsGet(context.Background(), functionId, functionRevisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsRevisionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsRevisionsGet`: Function
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsRevisionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** | Function ID reference | 
**functionRevisionId** | **string** | Function revision ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsRevisionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Function**](Function.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsRevisionsList

> ListFunction FunctionsRevisionsList(ctx, functionId).Page(page).Execute()

List all revisions for a function



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
	functionId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Function ID reference
	page := *openapiclient.NewOrganisationsListPageParameter() // OrganisationsListPageParameter | Query parameters for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsRevisionsList(context.Background(), functionId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsRevisionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsRevisionsList`: ListFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsRevisionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** | Function ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsRevisionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | [**OrganisationsListPageParameter**](OrganisationsListPageParameter.md) | Query parameters for pagination | 

### Return type

[**ListFunction**](ListFunction.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsUpdate

> Function FunctionsUpdate(ctx, functionId).Function(function).Execute()

Update function



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
	functionId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Function ID reference
	function := *openapiclient.NewFunction("51687d2f-07b0-4260-8ecb-f5098305fdd4", "51687d2f-07b0-4260-8ecb-f5098305fdd4", "1.0.0", "my-function", "My function does this", "docker.io/my-image:latest", "51687d2f-07b0-4260-8ecb-f5098305fdd4", openapiclient.FunctionType("http"), *openapiclient.NewFunctionSize(int32(128), int32(128)), int32(8080), *openapiclient.NewFunctionScaling("cpu", int32(30), int32(60)), []openapiclient.FunctionDeployment{*openapiclient.NewFunctionDeployment(*openapiclient.NewLocation("51687d2f-07b0-4260-8ecb-f5098305fdd4"), *openapiclient.NewFunctionReplicas(int32(1), int32(5), *openapiclient.NewFunctionReplicasAffinity(true, true)))}, []openapiclient.FunctionEnv{*openapiclient.NewFunctionEnv("MY_ENV_VAR", "projects:0a6b9ff3-6807-4820-b94b-5e1d7efcdd93/MY_SECRET@0")}, []openapiclient.FunctionCompliance{openapiclient.FunctionCompliance("soc2")}) // Function | Update any field

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsUpdate(context.Background(), functionId).Function(function).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsUpdate`: Function
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string** | Function ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **function** | [**Function**](Function.md) | Update any field | 

### Return type

[**Function**](Function.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsFunctionsList

> ListFunction ProjectsFunctionsList(ctx, projectId).Page(page).Execute()

List all functions within a project



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
	resp, r, err := apiClient.FunctionsAPI.ProjectsFunctionsList(context.Background(), projectId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.ProjectsFunctionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsFunctionsList`: ListFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.ProjectsFunctionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsFunctionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | [**OrganisationsListPageParameter**](OrganisationsListPageParameter.md) | Query parameters for pagination | 

### Return type

[**ListFunction**](ListFunction.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

