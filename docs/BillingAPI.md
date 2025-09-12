# \BillingAPI

All URIs are relative to *https://chaos.qernal.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsPaymentMethodsCreate**](BillingAPI.md#AccountsPaymentMethodsCreate) | **Post** /billing/accounts/{billing_account_id}/payment-methods | Create a new payment method for a billing account
[**AccountsPaymentMethodsList**](BillingAPI.md#AccountsPaymentMethodsList) | **Get** /billing/accounts/{billing_account_id}/payment-methods | List payment methods for a billing account
[**BillingAccountsCreate**](BillingAPI.md#BillingAccountsCreate) | **Post** /billing/accounts | Create billing account
[**BillingAccountsDelete**](BillingAPI.md#BillingAccountsDelete) | **Delete** /billing/accounts/{billing_account_id} | Delete billing account
[**BillingAccountsGet**](BillingAPI.md#BillingAccountsGet) | **Get** /billing/accounts/{billing_account_id} | Get billing account
[**BillingAccountsList**](BillingAPI.md#BillingAccountsList) | **Get** /billing/accounts | List billing accounts
[**BillingAccountsUpdate**](BillingAPI.md#BillingAccountsUpdate) | **Put** /billing/accounts/{billing_account_id} | Update billing account
[**PaymentMethodsDelete**](BillingAPI.md#PaymentMethodsDelete) | **Delete** /billing/payment-methods/{billing_payment_method_id} | Delete a specific payment method
[**PaymentMethodsGet**](BillingAPI.md#PaymentMethodsGet) | **Get** /billing/payment-methods/{billing_payment_method_id} | Retrieve metadata for a specific payment method
[**PaymentMethodsUpdate**](BillingAPI.md#PaymentMethodsUpdate) | **Put** /billing/payment-methods/{billing_payment_method_id} | Update a specific payment method



## AccountsPaymentMethodsCreate

> PaymentMethodCreate AccountsPaymentMethodsCreate(ctx, billingAccountId).PaymentMethodBody(paymentMethodBody).Execute()

Create a new payment method for a billing account



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
	billingAccountId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Billing account ID reference
	paymentMethodBody := *openapiclient.NewPaymentMethodBody() // PaymentMethodBody | Create/Update any field

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AccountsPaymentMethodsCreate(context.Background(), billingAccountId).PaymentMethodBody(paymentMethodBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AccountsPaymentMethodsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsPaymentMethodsCreate`: PaymentMethodCreate
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AccountsPaymentMethodsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingAccountId** | **string** | Billing account ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsPaymentMethodsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentMethodBody** | [**PaymentMethodBody**](PaymentMethodBody.md) | Create/Update any field | 

### Return type

[**PaymentMethodCreate**](PaymentMethodCreate.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsPaymentMethodsList

> []PaymentMethod AccountsPaymentMethodsList(ctx, billingAccountId).Execute()

List payment methods for a billing account



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
	billingAccountId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Billing account ID reference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.AccountsPaymentMethodsList(context.Background(), billingAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.AccountsPaymentMethodsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountsPaymentMethodsList`: []PaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.AccountsPaymentMethodsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingAccountId** | **string** | Billing account ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountsPaymentMethodsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PaymentMethod**](PaymentMethod.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingAccountsCreate

> BillingAccount BillingAccountsCreate(ctx).BillingAccountBody(billingAccountBody).Execute()

Create billing account



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
	billingAccountBody := *openapiclient.NewBillingAccountBody("Business Account") // BillingAccountBody | Create/Update any field

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingAccountsCreate(context.Background()).BillingAccountBody(billingAccountBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingAccountsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingAccountsCreate`: BillingAccount
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingAccountsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingAccountsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingAccountBody** | [**BillingAccountBody**](BillingAccountBody.md) | Create/Update any field | 

### Return type

[**BillingAccount**](BillingAccount.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingAccountsDelete

> DeletedResponse BillingAccountsDelete(ctx, billingAccountId).Execute()

Delete billing account



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
	billingAccountId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Billing account ID reference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingAccountsDelete(context.Background(), billingAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingAccountsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingAccountsDelete`: DeletedResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingAccountsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingAccountId** | **string** | Billing account ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingAccountsDeleteRequest struct via the builder pattern


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


## BillingAccountsGet

> BillingAccount BillingAccountsGet(ctx, billingAccountId).Execute()

Get billing account



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
	billingAccountId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Billing account ID reference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingAccountsGet(context.Background(), billingAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingAccountsGet`: BillingAccount
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingAccountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingAccountId** | **string** | Billing account ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingAccount**](BillingAccount.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingAccountsList

> []BillingAccount BillingAccountsList(ctx).Execute()

List billing accounts



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingAccountsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingAccountsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingAccountsList`: []BillingAccount
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingAccountsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillingAccountsListRequest struct via the builder pattern


### Return type

[**[]BillingAccount**](BillingAccount.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingAccountsUpdate

> BillingAccount BillingAccountsUpdate(ctx, billingAccountId).BillingAccountBody(billingAccountBody).Execute()

Update billing account



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
	billingAccountId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Billing account ID reference
	billingAccountBody := *openapiclient.NewBillingAccountBody("Business Account") // BillingAccountBody | Create/Update any field

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingAccountsUpdate(context.Background(), billingAccountId).BillingAccountBody(billingAccountBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingAccountsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingAccountsUpdate`: BillingAccount
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingAccountsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingAccountId** | **string** | Billing account ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingAccountsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingAccountBody** | [**BillingAccountBody**](BillingAccountBody.md) | Create/Update any field | 

### Return type

[**BillingAccount**](BillingAccount.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentMethodsDelete

> DeletedResponse PaymentMethodsDelete(ctx, billingPaymentMethodId).Execute()

Delete a specific payment method



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
	billingPaymentMethodId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Payment method ID reference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PaymentMethodsDelete(context.Background(), billingPaymentMethodId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PaymentMethodsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentMethodsDelete`: DeletedResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PaymentMethodsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingPaymentMethodId** | **string** | Payment method ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentMethodsDeleteRequest struct via the builder pattern


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


## PaymentMethodsGet

> PaymentMethod PaymentMethodsGet(ctx, billingPaymentMethodId).Execute()

Retrieve metadata for a specific payment method



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
	billingPaymentMethodId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Payment method ID reference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PaymentMethodsGet(context.Background(), billingPaymentMethodId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PaymentMethodsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentMethodsGet`: PaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PaymentMethodsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingPaymentMethodId** | **string** | Payment method ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentMethodsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentMethodsUpdate

> PaymentMethod PaymentMethodsUpdate(ctx, billingPaymentMethodId).PaymentMethodBody(paymentMethodBody).Execute()

Update a specific payment method



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
	billingPaymentMethodId := "3069614e-adc8-47cb-a69c-decf9c5f90fc" // string | Payment method ID reference
	paymentMethodBody := *openapiclient.NewPaymentMethodBody() // PaymentMethodBody | Create/Update any field

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.PaymentMethodsUpdate(context.Background(), billingPaymentMethodId).PaymentMethodBody(paymentMethodBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.PaymentMethodsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentMethodsUpdate`: PaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.PaymentMethodsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingPaymentMethodId** | **string** | Payment method ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentMethodsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentMethodBody** | [**PaymentMethodBody**](PaymentMethodBody.md) | Create/Update any field | 

### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

