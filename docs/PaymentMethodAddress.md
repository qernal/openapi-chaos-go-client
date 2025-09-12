# PaymentMethodAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the cardholder | [optional] 
**FirstName** | Pointer to **string** | First name of the cardholder | [optional] 
**LastName** | Pointer to **string** | Last name of the cardholder | [optional] 
**Organisation** | Pointer to **string** | Organisation name, if applicable | [optional] 
**AddressLine1** | Pointer to **string** | First line of the address | [optional] 
**AddressLine2** | Pointer to **string** | Second line of the address, if applicable | [optional] 
**City** | Pointer to **string** | City of the address | [optional] 
**County** | Pointer to **string** | County or state of the address | [optional] 
**PostalCode** | Pointer to **string** | Postal or ZIP code of the address | [optional] 
**Country** | Pointer to **string** | Country of the address (ISO 3166-1 alpha-2 code) | [optional] 
**PhoneNumber** | Pointer to **string** | Contact phone number associated with the payment method | [optional] 

## Methods

### NewPaymentMethodAddress

`func NewPaymentMethodAddress() *PaymentMethodAddress`

NewPaymentMethodAddress instantiates a new PaymentMethodAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodAddressWithDefaults

`func NewPaymentMethodAddressWithDefaults() *PaymentMethodAddress`

NewPaymentMethodAddressWithDefaults instantiates a new PaymentMethodAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PaymentMethodAddress) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PaymentMethodAddress) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PaymentMethodAddress) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PaymentMethodAddress) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstName

`func (o *PaymentMethodAddress) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PaymentMethodAddress) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PaymentMethodAddress) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PaymentMethodAddress) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PaymentMethodAddress) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PaymentMethodAddress) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PaymentMethodAddress) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PaymentMethodAddress) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetOrganisation

`func (o *PaymentMethodAddress) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *PaymentMethodAddress) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *PaymentMethodAddress) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *PaymentMethodAddress) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### GetAddressLine1

`func (o *PaymentMethodAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *PaymentMethodAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *PaymentMethodAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *PaymentMethodAddress) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *PaymentMethodAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *PaymentMethodAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *PaymentMethodAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *PaymentMethodAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *PaymentMethodAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PaymentMethodAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PaymentMethodAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PaymentMethodAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCounty

`func (o *PaymentMethodAddress) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *PaymentMethodAddress) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *PaymentMethodAddress) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *PaymentMethodAddress) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetPostalCode

`func (o *PaymentMethodAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PaymentMethodAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PaymentMethodAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PaymentMethodAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *PaymentMethodAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentMethodAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentMethodAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentMethodAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PaymentMethodAddress) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PaymentMethodAddress) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PaymentMethodAddress) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PaymentMethodAddress) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


