# \LogsAPI

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LogsList**](LogsAPI.md#LogsList) | **Get** /logs | Get logs



## LogsList

> ListLogResponse LogsList(ctx).Page(page).FProject(fProject).FFunction(fFunction).FTimestamps(fTimestamps).FQuery(fQuery).FLogType(fLogType).Execute()

Get logs



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
	fProject := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project uuid reference (optional)
	fFunction := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Function uuid reference (optional)
	fTimestamps := *openapiclient.NewLogsListFTimestampsParameter() // LogsListFTimestampsParameter | Timestamp restriction for query (optional)
	fQuery := "bootstrap" // string | Text query string (optional)
	fLogType := "error" // string | Type of log (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsAPI.LogsList(context.Background()).Page(page).FProject(fProject).FFunction(fFunction).FTimestamps(fTimestamps).FQuery(fQuery).FLogType(fLogType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.LogsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogsList`: ListLogResponse
	fmt.Fprintf(os.Stdout, "Response from `LogsAPI.LogsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | [**OrganisationsListPageParameter**](OrganisationsListPageParameter.md) | Query parameters for pagination | 
 **fProject** | **string** | Project uuid reference | 
 **fFunction** | **string** | Function uuid reference | 
 **fTimestamps** | [**LogsListFTimestampsParameter**](LogsListFTimestampsParameter.md) | Timestamp restriction for query | 
 **fQuery** | **string** | Text query string | 
 **fLogType** | **string** | Type of log | 

### Return type

[**ListLogResponse**](ListLogResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

