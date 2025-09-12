# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the payment method | [optional] 
**BillingAccount** | Pointer to **string** | Unique identifier for the billing account | [optional] 
**Name** | Pointer to **string** | Name on the payment method | [optional] 
**State** | Pointer to **string** | Current state of the payment method | [optional] 
**Address** | Pointer to [**PaymentMethodAddress**](PaymentMethodAddress.md) |  | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBillingAccount

`func (o *PaymentMethod) GetBillingAccount() string`

GetBillingAccount returns the BillingAccount field if non-nil, zero value otherwise.

### GetBillingAccountOk

`func (o *PaymentMethod) GetBillingAccountOk() (*string, bool)`

GetBillingAccountOk returns a tuple with the BillingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccount

`func (o *PaymentMethod) SetBillingAccount(v string)`

SetBillingAccount sets BillingAccount field to given value.

### HasBillingAccount

`func (o *PaymentMethod) HasBillingAccount() bool`

HasBillingAccount returns a boolean if a field has been set.

### GetName

`func (o *PaymentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *PaymentMethod) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PaymentMethod) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PaymentMethod) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PaymentMethod) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAddress

`func (o *PaymentMethod) GetAddress() PaymentMethodAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PaymentMethod) GetAddressOk() (*PaymentMethodAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PaymentMethod) SetAddress(v PaymentMethodAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PaymentMethod) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


