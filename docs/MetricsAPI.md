# \MetricsAPI

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricsAggregationsList**](MetricsAPI.md#MetricsAggregationsList) | **Get** /metrics/aggregations/{metric_type} | Get metrics



## MetricsAggregationsList

> MetricsAggregationsList200Response MetricsAggregationsList(ctx, metricType).Page(page).FProject(fProject).FFunction(fFunction).FTimestamps(fTimestamps).FHistogramInterval(fHistogramInterval).Execute()

Get metrics



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
	metricType := "httprequests" // string | Metric aggregation type, types can be used with either a project or a function filter.  - httprequests: Aggregated HTTP requests - resourcestats: Aggregated resource stats (such as CPU, Memory and Network)  > Note: aggregations cannot return more than 300 data points 
	page := *openapiclient.NewOrganisationsListPageParameter() // OrganisationsListPageParameter | Query parameters for pagination (optional)
	fProject := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Project uuid reference (optional)
	fFunction := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Function uuid reference (optional)
	fTimestamps := *openapiclient.NewLogsListFTimestampsParameter() // LogsListFTimestampsParameter | Timestamp restriction for query (optional)
	fHistogramInterval := int32(600) // int32 | Histogram interval (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.MetricsAggregationsList(context.Background(), metricType).Page(page).FProject(fProject).FFunction(fFunction).FTimestamps(fTimestamps).FHistogramInterval(fHistogramInterval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.MetricsAggregationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricsAggregationsList`: MetricsAggregationsList200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.MetricsAggregationsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricType** | **string** | Metric aggregation type, types can be used with either a project or a function filter.  - httprequests: Aggregated HTTP requests - resourcestats: Aggregated resource stats (such as CPU, Memory and Network)  &gt; Note: aggregations cannot return more than 300 data points  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsAggregationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | [**OrganisationsListPageParameter**](OrganisationsListPageParameter.md) | Query parameters for pagination | 
 **fProject** | **string** | Project uuid reference | 
 **fFunction** | **string** | Function uuid reference | 
 **fTimestamps** | [**LogsListFTimestampsParameter**](LogsListFTimestampsParameter.md) | Timestamp restriction for query | 
 **fHistogramInterval** | **int32** | Histogram interval | 

### Return type

[**MetricsAggregationsList200Response**](MetricsAggregationsList200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

