# PaymentMethodCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodId** | Pointer to **string** | Unique identifier for the payment method | [optional] 
**BillingAccountId** | Pointer to **string** | Unique identifier for the billing account | [optional] 
**Name** | Pointer to **string** | Name on the payment method | [optional] 
**Description** | Pointer to **string** | Description of this order (e.g. payment authorisation) | [optional] 
**State** | Pointer to **string** | Current state of the payment method | [optional] 
**Links** | Pointer to [**PaymentMethodCreateLinks**](PaymentMethodCreateLinks.md) |  | [optional] 

## Methods

### NewPaymentMethodCreate

`func NewPaymentMethodCreate() *PaymentMethodCreate`

NewPaymentMethodCreate instantiates a new PaymentMethodCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCreateWithDefaults

`func NewPaymentMethodCreateWithDefaults() *PaymentMethodCreate`

NewPaymentMethodCreateWithDefaults instantiates a new PaymentMethodCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodId

`func (o *PaymentMethodCreate) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *PaymentMethodCreate) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *PaymentMethodCreate) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *PaymentMethodCreate) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetBillingAccountId

`func (o *PaymentMethodCreate) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *PaymentMethodCreate) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *PaymentMethodCreate) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.

### HasBillingAccountId

`func (o *PaymentMethodCreate) HasBillingAccountId() bool`

HasBillingAccountId returns a boolean if a field has been set.

### GetName

`func (o *PaymentMethodCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethodCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethodCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentMethodCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentMethodCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentMethodCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentMethodCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentMethodCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *PaymentMethodCreate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PaymentMethodCreate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PaymentMethodCreate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PaymentMethodCreate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLinks

`func (o *PaymentMethodCreate) GetLinks() PaymentMethodCreateLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaymentMethodCreate) GetLinksOk() (*PaymentMethodCreateLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaymentMethodCreate) SetLinks(v PaymentMethodCreateLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaymentMethodCreate) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


