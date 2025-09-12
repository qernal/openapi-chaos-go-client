# PaymentMethodCreateLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checkout** | Pointer to **string** | URL for the payment method checkout | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the checkout link was created | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Timestamp when the checkout link expires | [optional] 

## Methods

### NewPaymentMethodCreateLinks

`func NewPaymentMethodCreateLinks() *PaymentMethodCreateLinks`

NewPaymentMethodCreateLinks instantiates a new PaymentMethodCreateLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCreateLinksWithDefaults

`func NewPaymentMethodCreateLinksWithDefaults() *PaymentMethodCreateLinks`

NewPaymentMethodCreateLinksWithDefaults instantiates a new PaymentMethodCreateLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckout

`func (o *PaymentMethodCreateLinks) GetCheckout() string`

GetCheckout returns the Checkout field if non-nil, zero value otherwise.

### GetCheckoutOk

`func (o *PaymentMethodCreateLinks) GetCheckoutOk() (*string, bool)`

GetCheckoutOk returns a tuple with the Checkout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckout

`func (o *PaymentMethodCreateLinks) SetCheckout(v string)`

SetCheckout sets Checkout field to given value.

### HasCheckout

`func (o *PaymentMethodCreateLinks) HasCheckout() bool`

HasCheckout returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentMethodCreateLinks) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethodCreateLinks) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethodCreateLinks) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentMethodCreateLinks) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentMethodCreateLinks) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentMethodCreateLinks) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentMethodCreateLinks) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentMethodCreateLinks) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


