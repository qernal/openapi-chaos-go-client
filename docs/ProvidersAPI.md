# \ProvidersAPI

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvidersList**](ProvidersAPI.md#ProvidersList) | **Get** /providers | Get available providers



## ProvidersList

> ListProviderResponse ProvidersList(ctx).Page(page).Execute()

Get available providers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvidersAPI.ProvidersList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.ProvidersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersList`: ListProviderResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.ProvidersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvidersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | [**OrganisationsListPageParameter**](OrganisationsListPageParameter.md) | Query parameters for pagination | 

### Return type

[**ListProviderResponse**](ListProviderResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

